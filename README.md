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

`-o`: ファイルエクスポート

`-t`: タイムアウトを設定

### ◆POST

`gdeck post {URL}`

---オプション---

`-v`: 詳細出力

`-o`: ファイルエクスポート

`-d`: postデータを設定

`-H`: reqeust headerを設定

`-t`: タイムアウトを設定

### ◆SAVE

`gdeck save {NAME} {METHOD} {URL}`

---オプション---

`-d`: postデータを設定

`-H`: reqeust headerを設定

---環境変数---

`{{HOGE}}`: この形式にすることで環境変数を使用できる

runコマンド実行時に明示的に値をセット。

### ◆RUN

`gdeck run {NAME or PATH}`

`gdeck run "{NAME or PATH}/*"` で一括実行（※ダブルクォーテーションで囲まないとzshエラー起こります）

---オプション---

`-v`: 詳細出力

`-d`: postデータを設定

`-H`: reqeust headerを設定

`-t`: タイムアウトを設定

---環境変数---

saveコマンドで登録した際に`{{HOGE}}`のような形式があれば、実行時に値をセットできる

`HOGE=batsumaru \ gdeck run {SAVED COMMAND NAME}`

### ◆SHOW

`gdeck show {NAME or path}`

`gdeck run "{NAME or PATH}/*"` で一括実行（※ダブルクォーテーションで囲まないとzshエラー起こります）

---オプション---

### ◆LIST

`gdeck list`

---オプション---

### ◆DELETE

`gdeck delete {NAME or path}`

`gdeck run "{NAME or PATH}/*"` で一括実行（※ダブルクォーテーションで囲まないとzshエラー起こります）

---オプション---

### ◆ENV

- set <br>`gdeck env set KEY VALUE`
- show: <br>`gdeck env show KEY`
- list: <br>`gdeck env list`
- delete: <br>`gdeck env delete KEY`

## バイナリの作成方法

`go build -o gdeck`

`mv gdeck /usr/local/bin/`

`gdeck get {URL}`

シンボリックリンクを作成しておけばビルドだけで済む

`ln -s $(pwd)/gdeck /usr/local/bin/gdeck`

## 実装時ルール

- コミット前には必ず`go fmt`を実行する
