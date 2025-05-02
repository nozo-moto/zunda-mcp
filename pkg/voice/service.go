package voice

import (
	"context"
	"fmt"
	"os"
	"os/exec"
)

// Service は音声合成のファサードとして機能し、合成とファイル処理を統合します。
type Service struct {
	synthesizer Synthesizer
	processor   TextProcessor
}

// NewService は新しい音声合成サービスを作成します。
func NewService(synthesizer Synthesizer, processor TextProcessor) *Service {
	return &Service{
		synthesizer: synthesizer,
		processor:   processor,
	}
}

// ProcessText はテキストから音声を生成し、各ステップでファイルに保存します。
func (s *Service) ProcessText(ctx context.Context, text string, speakerID int32,
	textFile, queryFile, audioFile string) (string, error) {

	// テキストを保存
	if err := s.processor.SaveText(text, textFile); err != nil {
		return "", fmt.Errorf("テキストの保存に失敗しました: %w", err)
	}

	// 音声合成クエリを取得
	query, err := s.synthesizer.AudioQuery(ctx, text, speakerID)
	if err != nil {
		return "", fmt.Errorf("音声合成クエリの取得に失敗しました: %w", err)
	}

	// クエリを保存
	if err := s.processor.SaveAudioQuery(query, queryFile); err != nil {
		return "", fmt.Errorf("クエリの保存に失敗しました: %w", err)
	}

	// 音声を合成
	audioData, err := s.synthesizer.Synthesize(ctx, speakerID, query, true)
	if err != nil {
		return "", fmt.Errorf("音声合成に失敗しました: %w", err)
	}

	// 音声ファイルを保存
	if err := s.processor.SaveAudio(audioData, audioFile); err != nil {
		return "", fmt.Errorf("音声ファイルの保存に失敗しました: %w", err)
	}

	// 絶対パスを取得
	audioPath, err := s.processor.GetAbsolutePath(audioFile)
	if err != nil {
		return "", fmt.Errorf("パスの取得に失敗しました: %w", err)
	}

	return audioPath, nil
}

// ProcessTextAndPlay はテキストから音声を生成し再生します。一時ファイルを使用します。
func (s *Service) ProcessTextAndPlay(ctx context.Context, text string, speakerID int32) error {
	// テキストを一時ファイルに保存
	textTmpFile, err := s.processor.CreateTempFile("text-", ".txt", []byte(text))
	if err != nil {
		return fmt.Errorf("一時テキストファイルの作成に失敗しました: %w", err)
	}
	defer os.Remove(textTmpFile) // 後始末

	// 音声合成クエリを取得
	query, err := s.synthesizer.AudioQuery(ctx, text, speakerID)
	if err != nil {
		return fmt.Errorf("音声合成クエリの取得に失敗しました: %w", err)
	}

	// クエリをJSON形式に変換して一時ファイルに保存
	queryJSON, err := s.processor.ConvertAudioQueryToJSON(query)
	if err != nil {
		return fmt.Errorf("クエリのJSON変換に失敗しました: %w", err)
	}

	queryTmpFile, err := s.processor.CreateTempFile("query-", ".json", queryJSON)
	if err != nil {
		return fmt.Errorf("一時クエリファイルの作成に失敗しました: %w", err)
	}
	defer os.Remove(queryTmpFile) // 後始末

	// 音声を合成
	audioData, err := s.synthesizer.Synthesize(ctx, speakerID, query, true)
	if err != nil {
		return fmt.Errorf("音声合成に失敗しました: %w", err)
	}

	// 音声を一時ファイルに保存
	audioTmpFile, err := s.processor.CreateTempFile("audio-", ".wav", audioData)
	if err != nil {
		return fmt.Errorf("一時音声ファイルの作成に失敗しました: %w", err)
	}
	defer os.Remove(audioTmpFile) // 後始末

	// afplayで再生
	cmd := exec.Command("afplay", audioTmpFile)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("音声の再生に失敗しました: %w", err)
	}

	return nil
}
