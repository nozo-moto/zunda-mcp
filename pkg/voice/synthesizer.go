package voice

import (
	"context"
	"io"

	openapi "zundamon-mcp/pkg/generated"
)

// Synthesizer は音声合成エンジンのクエリ取得と合成処理を抽象化したインターフェースです。
type Synthesizer interface {
	// AudioQuery はテキストと話者IDから音声合成用のクエリを取得します。
	AudioQuery(ctx context.Context, text string, speakerID int32) (*openapi.AudioQuery, error)
	// Synthesize は話者IDとAudioQueryからWAVデータを生成して返します。
	Synthesize(ctx context.Context, speakerID int32, query *openapi.AudioQuery, enableInterrogativeUpspeak bool) ([]byte, error)
}

// ClientSynthesizer は生成済みのOpenAPIクライアントを使ってSynthesizerを実装します。
type ClientSynthesizer struct {
	api *openapi.DefaultApiService
}

// NewClientSynthesizer はVOICEVOXエンジンへの接続URLを指定して新しいSynthesizerを生成します。
func NewClientSynthesizer(baseURL string) Synthesizer {
	cfg := openapi.NewConfiguration()
	cfg.Servers[0].URL = baseURL
	client := openapi.NewAPIClient(cfg)
	return &ClientSynthesizer{api: client.DefaultApi}
}

// AudioQuery はテキストと話者IDから音声合成用のクエリを取得します。
func (cs *ClientSynthesizer) AudioQuery(ctx context.Context, text string, speakerID int32) (*openapi.AudioQuery, error) {
	req := cs.api.AudioQueryAudioQueryPost(ctx).Text(text).Speaker(speakerID)
	query, _, err := req.Execute()
	return query, err
}

// Synthesize は話者IDとAudioQueryを元にWAV形式の音声データを生成して返します。
func (cs *ClientSynthesizer) Synthesize(ctx context.Context, speakerID int32, query *openapi.AudioQuery, enableInterrogativeUpspeak bool) ([]byte, error) {
	builder := cs.api.SynthesisSynthesisPost(ctx).Speaker(speakerID).AudioQuery(*query)
	if !enableInterrogativeUpspeak {
		builder = builder.EnableInterrogativeUpspeak(enableInterrogativeUpspeak)
	}
	_, resp, err := builder.Execute()
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}
