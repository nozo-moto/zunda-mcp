package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"

	openapi "github.com/nozo-moto/zunda-mcp/pkg/generated"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	// コマンドライン引数を定義
	textPtr := flag.String("text", "", "読み上げるテキスト")
	speakerPtr := flag.Int("speaker", 1, "VOICEVOXの話者ID")

	flag.Parse() // コマンドライン引数をパース

	// text引数が指定されている場合はCLIモードとして実行
	if *textPtr != "" {
		fmt.Printf("CLIモードで実行: text='%s', speaker=%d\n", *textPtr, *speakerPtr)
		ctx := context.Background()
		err := speakText(ctx, *textPtr, int32(*speakerPtr))
		if err != nil {
			fmt.Fprintf(os.Stderr, "エラー: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("音声再生が完了しました。")
		return // CLIモードの場合はここで終了
	}

	// text引数が指定されていない場合はMCPサーバーモードとして実行
	// fmt.Println("MCPサーバーモードで起動します...")

	// Create MCP server
	s := server.NewMCPServer(
		"音声読み上げサーバー",
		"1.0.0",
	)

	// 音声読み上げツールを追加
	speakTool := mcp.NewTool("speak_text",
		mcp.WithDescription("テキストを音声で読み上げます"),
		mcp.WithString("text",
			mcp.Required(),
			mcp.Description("読み上げるテキスト"),
		),
		mcp.WithNumber("speaker",
			mcp.Description("VOICEVOXの話者ID (デフォルトは1)"),
		),
	)

	// ツールハンドラーを追加
	s.AddTool(speakTool, speakHandler)

	// STDIOサーバーを起動
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("サーバーエラー: %v\n", err)
	}
}

func speakHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	text, ok := request.Params.Arguments["text"].(string)
	if !ok || text == "" {
		return nil, errors.New("テキストは必須で文字列である必要があります")
	}

	// speakerIDを取得 (オプション、デフォルトは1)
	speakerIDParam, ok := request.Params.Arguments["speaker"].(float64) // MCPでは数値はfloat64になる可能性がある
	speakerID := int32(1)                                               // デフォルトの話者ID
	if ok {
		speakerID = int32(speakerIDParam)
	}

	// speakText関数を呼び出す
	if err := speakText(ctx, text, speakerID); err != nil {
		// エラーハンドリング: エラーメッセージを返すか、より詳細な情報を提供
		return nil, fmt.Errorf("音声合成・再生エラー: %w", err)
	}

	// 成功メッセージを返す
	return mcp.NewToolResultText(fmt.Sprintf("「%s」を読み上げました", text)), nil
}

// speakText はテキストを音声合成し、afplayで再生する
func speakText(ctx context.Context, text string, speakerID int32) error {
	// VOICEVOX APIクライアントの設定
	cfg := openapi.NewConfiguration()
	cfg.Host = "localhost:50021" // VOICEVOX APIのエンドポイント
	cfg.Scheme = "http"          // Schemeを設定
	client := openapi.NewAPIClient(cfg)

	// 1. /audio_query でクエリを取得
	audioQueryRequest := client.DefaultApi.AudioQueryAudioQueryPost(ctx).Speaker(speakerID).Text(text)
	audioQuery, resp, err := audioQueryRequest.Execute()
	if err != nil {
		var openapiErr openapi.GenericOpenAPIError
		if errors.As(err, &openapiErr) {
			fmt.Printf("OpenAPI Error Body: %s\n", string(openapiErr.Body()))
		}
		return fmt.Errorf("audio_queryリクエストエラー: %w", err)
	}
	if resp.StatusCode != 200 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		return fmt.Errorf("audio_queryエラー: status=%d, body=%s", resp.StatusCode, string(bodyBytes))
	}
	resp.Body.Close()

	// 2. /synthesis で音声データを生成
	synthesisRequest := client.DefaultApi.SynthesisSynthesisPost(ctx).Speaker(speakerID)
	synthesisRequest = synthesisRequest.AudioQuery(*audioQuery)

	wavData, synthResp, err := synthesisRequest.Execute()
	if err != nil {
		var openapiErr openapi.GenericOpenAPIError
		if errors.As(err, &openapiErr) {
			fmt.Printf("OpenAPI Error Body: %s\n", string(openapiErr.Body()))
		}
		return fmt.Errorf("synthesisリクエストエラー: %w", err)
	}
	if synthResp.StatusCode != 200 {
		bodyBytes, _ := io.ReadAll(synthResp.Body)
		synthResp.Body.Close()
		return fmt.Errorf("synthesisエラー: status=%d, body=%s", synthResp.StatusCode, string(bodyBytes))
	}
	defer synthResp.Body.Close()
	defer wavData.Close() // 自動的に一時ファイルが削除される

	// 3. afplayで再生
	cmd := exec.CommandContext(ctx, "afplay", wavData.Name())
	if err := cmd.Run(); err != nil {
		if errors.Is(err, exec.ErrNotFound) {
			return fmt.Errorf("`afplay`コマンドが見つかりません。macOS環境で実行してください。")
		}
		return fmt.Errorf("afplay実行エラー: %w", err)
	}

	return nil // 成功
}
