# apictl(api-tester-cli)
CLIからAPI接続テストを行うツール

## 使用ツール
### ◆cobra

`go get github.com/spf13/cobra@latest`

### ◆cobra-cli

root.goなどを簡単に作成してくれる便利CLI。
いつか導入する。

`go install github.com/spf13/cobra-cli@latest`

## ツールの使い方(Local)
### ◆単発実行

`go run main.go apictl get <URL>`

### ◆バイナリを作成しツール名称で実行

`go build -o apictl`

`mv apictl /usr/local/bin/`

`apictl get <URL>`

シンボリックリンクを作成しておけばビルドだけで済む

`ln -s $(pwd)/apictl /usr/local/bin/apictl`

## 実装時ルール
* コミット前には必ず`go fmt`を実行する