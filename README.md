# gdeck(api-tester-cli)

CLIからAPI接続テストを行うツール

## 使用ツール

### ◆cobra

`go get github.com/spf13/cobra@latest`

### ◆cobra-cli

root.goなどを簡単に作成してくれる便利CLI。
いつか導入する。

`go install github.com/spf13/cobra-cli@latest`

## 使い方

### ◆GET

`gdeck get {URL}`

---オプション---

`-v`: 詳細出力

`-o`: ファイル出力パス

`-t`: タイムアウト（秒）

### ◆POST

`gdeck post {URL}`

---オプション---

`-v`: 詳細出力

`-o`: ファイル出力パス

`-d`: リクエストボディ

`-H`: リクエストヘッダー

`-t`: タイムアウト（秒）

### ◆SAVE

`gdeck save {NAME} {METHOD} {URL}`

---オプション---

`-d`: リクエストボディ

`-H`: リクエストヘッダー

---環境変数---

保存されたリクエスト内に `{{HOGE}}` のような形式があれば、`gdeck run` 実行時に環境変数置換されます。

### ◆RUN

`gdeck run {NAME or PATH}`

`gdeck run "{NAME or PATH}/*"` で複数ファイルを一括実行できます（※シェルのワイルドカード展開を防ぐため、引用符で囲んでください）。

---オプション---

`-v`: 詳細出力

`-d`: リクエストボディ上書き

`-H`: リクエストヘッダー上書き

`-t`: タイムアウト（秒）

`--env`: 環境名を指定して環境変数ファイルを切り替え

---例---

`gdeck run SampleCmd --env dev`

### ◆SHOW

`gdeck show {NAME or PATH}`

保存済みコマンドの詳細を表示します。

### ◆LIST

`gdeck list`

保存済みコマンドの一覧を表示します。

### ◆DELETE

`gdeck delete {NAME or PATH}`

保存済みコマンドを削除します。

### ◆ENV

`gdeck env` は環境変数管理のサブコマンドです。

- `gdeck env set KEY VALUE [--env NAME]`
- `gdeck env show KEY [--env NAME]`
- `gdeck env list [--env NAME]`
- `gdeck env delete KEY [--env NAME]`

名前付き環境を使う場合は `--env` を指定します。

## バイナリの作成方法

`go build -o gdeck`

`mv gdeck /usr/local/bin/`

`gdeck get {URL}`

シンボリックリンクを作成しておけばビルドだけで済む

`ln -s $(pwd)/gdeck /usr/local/bin/gdeck`

## 実装時ルール

- コミット前には必ず`go fmt`を実行する
