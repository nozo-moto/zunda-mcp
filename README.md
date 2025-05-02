# テキスト読み上げMCPサーバー

MCP経由でずんだもんに喋らせます


## 機能

- VOICE BOXを利用して、テキストからずんだもんの読み上げ音声をWAVで作成し、macOSの `afplay` で再生する

## インストール

```
go install github.com/nozo-moto/zunda-mcp@latest
```

## 使い方

### MCP 

mcp jsonに

```
    "zundamon-mcp": {
      "command": "$HOME/go/bin/zundamon-mcp"
    },
```

[VOICE BOX](https://voicevox.hiroshiba.jp/) をインストールして起動しておいてください

Mac以外ではafplayが存在しないので動かないです

### CLI

``` bash
$ go run main.go -text こんにちはなのだ
CLIモードで実行: text='こんにちはなのだ', speaker=1
音声再生が完了しました。
```


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

