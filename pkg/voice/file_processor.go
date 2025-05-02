package voice

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// TextProcessor はテキストと音声ファイルの入出力を処理します。
type TextProcessor interface {
	// SaveText はテキストをファイルに保存します。
	SaveText(text string, filePath string) error
	// SaveAudioQuery はAudioQueryをJSONとして保存します。
	SaveAudioQuery(query interface{}, filePath string) error
	// ConvertAudioQueryToJSON はAudioQueryをJSON形式のバイト配列に変換します。
	ConvertAudioQueryToJSON(query interface{}) ([]byte, error)
	// SaveAudio はバイナリデータを音声ファイルとして保存します。
	SaveAudio(audioData []byte, filePath string) error
	// GetAbsolutePath は相対パスから絶対パスを取得します。
	GetAbsolutePath(filePath string) (string, error)
	// CreateTempFile は一時ファイルを作成し、データを書き込みます。
	CreateTempFile(prefix string, suffix string, data []byte) (string, error)
}

// FileProcessor はファイルシステムを使用したTextProcessorの実装です。
type FileProcessor struct{}

// NewFileProcessor は新しいFileProcessorを作成します。
func NewFileProcessor() TextProcessor {
	return &FileProcessor{}
}

// SaveText はテキストをファイルに保存します。
func (fp *FileProcessor) SaveText(text string, filePath string) error {
	return ioutil.WriteFile(filePath, []byte(text), 0644)
}

// ConvertAudioQueryToJSON はAudioQueryをJSON形式のバイト配列に変換します。
func (fp *FileProcessor) ConvertAudioQueryToJSON(query interface{}) ([]byte, error) {
	return json.MarshalIndent(query, "", "  ")
}

// SaveAudioQuery はオブジェクトをJSONとして保存します。
func (fp *FileProcessor) SaveAudioQuery(query interface{}, filePath string) error {
	queryJSON, err := fp.ConvertAudioQueryToJSON(query)
	if err != nil {
		return fmt.Errorf("JSONへの変換に失敗しました: %w", err)
	}

	return ioutil.WriteFile(filePath, queryJSON, 0644)
}

// SaveAudio はバイナリデータを音声ファイルとして保存します。
func (fp *FileProcessor) SaveAudio(audioData []byte, filePath string) error {
	return ioutil.WriteFile(filePath, audioData, 0644)
}

// GetAbsolutePath は相対パスから絶対パスを取得します。
func (fp *FileProcessor) GetAbsolutePath(filePath string) (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("現在のディレクトリの取得に失敗しました: %w", err)
	}
	return filepath.Join(cwd, filePath), nil
}

// CreateTempFile は一時ファイルを作成し、データを書き込みます。
func (fp *FileProcessor) CreateTempFile(prefix string, suffix string, data []byte) (string, error) {
	// 一時ファイルを作成
	tmpFile, err := ioutil.TempFile("", prefix+"*"+suffix)
	if err != nil {
		return "", fmt.Errorf("一時ファイルの作成に失敗しました: %w", err)
	}
	defer tmpFile.Close()

	// データを書き込み
	if _, err := tmpFile.Write(data); err != nil {
		os.Remove(tmpFile.Name()) // エラー時はファイルを削除
		return "", fmt.Errorf("一時ファイルへの書き込みに失敗しました: %w", err)
	}

	return tmpFile.Name(), nil
}
