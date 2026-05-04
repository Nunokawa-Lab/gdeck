# apictl(api-tester-cli)
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

`apictl get <URL>`

---オプション---

`-v`: 詳細出力

`-o`: ファイルエクスポート

`-t`: タイムアウトを設定


### ◆POST

`apictl post <URL>`

---オプション---

`-v`: 詳細出力

`-o`: ファイルエクスポート

`-d`: postデータを設定

`-H`: reqeust headerを設定

`-t`: タイムアウトを設定


### ◆SAVE

`apictl save <NAME> <METHOD> <URL>`

---オプション---

`-d`: postデータを設定

`-H`: reqeust headerを設定

---環境変数---

`{{HOGE}}`: この形式にすることで環境変数を使用できる

runコマンド実行時に明示的に値をセット。


### ◆RUN

`apictl run <SAVED COMMAND NAME>`

---オプション---

`-v`: 詳細出力

`-d`: postデータを設定

`-H`: reqeust headerを設定

`-t`: タイムアウトを設定

---環境変数---

saveコマンドで登録した際に`{{HOGE}}`のような形式があれば、実行時に値をセットできる

`HOGE=batsumaru \ apictl run <SAVED COMMAND NAME>`


### ◆LIST

`apictl list`

---オプション---


### ◆DELETE

`apictl delete <SAVED COMMAND NAME>`

---オプション---


## バイナリの作成方法

`go build -o apictl`

`mv apictl /usr/local/bin/`

`apictl get <URL>`

シンボリックリンクを作成しておけばビルドだけで済む

`ln -s $(pwd)/apictl /usr/local/bin/apictl`

## 実装時ルール
* コミット前には必ず`go fmt`を実行する