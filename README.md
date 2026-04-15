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
### ◆バイナリを使用しない実行

`go run main.go get <URL>`

`go run main.go post <URL>`

### ◆バイナリ使用した実行

`apictl get <URL>`

`apictl post <URL>`

### ◆オプション

`-v`: 詳細出力

`-o`: ファイルエクスポート

`-d`: postデータを設定

`-H`: reqeust headerを設定

### ◆例
```
go run main.go  post https://jsonplaceholder.typicode.com/posts \
  -d '{"id":"1"}' \
  -o '/Users/nunokawa/projects/apictl/post.json' \
  -v
```


### ◆バイナリを作成方法

`go build -o apictl`

`mv apictl /usr/local/bin/`

`apictl get <URL>`

シンボリックリンクを作成しておけばビルドだけで済む

`ln -s $(pwd)/apictl /usr/local/bin/apictl`

## 実装時ルール
* コミット前には必ず`go fmt`を実行する