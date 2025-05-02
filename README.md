# テキスト読み上げMCPサーバー

このプロジェクトは、[Model Context Protocol (MCP)](https://github.com/mark3labs/mcp-go)を使用して、テキストを音声に変換して読み上げるサーバーを実装したものです。

## 機能

- テキストを音声に変換して読み上げる
- 複数の言語に対応（日本語、英語、フランス語など）

## インストール

```
go mod tidy
go build -o tts-server
```

## 使い方

サーバーを起動します：

```
./tts-server
```

このサーバーはMCP互換のクライアントからの接続を待ち受けます。

### ツール

#### speak_text

テキストを音声で読み上げます。

**パラメータ：**

- `text` (必須): 読み上げるテキスト
- `language` (オプション): 読み上げ言語コード（デフォルト: "ja"）
  - 対応言語: "ja" (日本語), "en" (英語), "fr" (フランス語), など

**使用例：**

```json
{
  "jsonrpc": "2.0",
  "method": "mcp.call_tool",
  "params": {
    "name": "speak_text",
    "arguments": {
      "text": "こんにちは、世界",
      "language": "ja"
    }
  },
  "id": 1
}
```

## 注意点

- 音声の再生には、スピーカーが必要です
- 一部の環境では、オーディオ出力の設定が必要な場合があります

## 依存ライブラリ



## openapi clientの作り方

VOICE VOXを起動すると

[http://localhost:50021/docs](http://localhost:50021/docs)

にopenapiの定義があります

```
$ wget http://localhost:50021/openapi.json
$ openapi-generator generate \
  -i openapi.json \
  -g go \
  -o ./pkg//generated
```

