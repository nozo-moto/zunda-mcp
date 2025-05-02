# テキスト読み上げMCPサーバー

MCP経由でずんだもんに喋らせま

## 機能

- テキストを音声に変換して読み上げる

## インストール

```
go install github.com/nozo-moto/zunda-mcp@latest
```

## 使い方

mcp jsonに

```
    "zundamon-mcp": {
      "command": "$HOME/go/bin/zundamon-mcp"
    },
```

[VOICE BOX](https://voicevox.hiroshiba.jp/) をインストールして起動しておいてください

Mac以外で動くかは検証してないです


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

