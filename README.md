# gdeck

**gdeck** は、CLI と TUI の両方で API テストを実行できる軽量なツールです。

- 🚀 シンプルな HTTP リクエスト実行
- 📁 保存済みリクエスト管理
- 🎛️ インタラクティブ TUI モード
- ♻️ 環境変数置換対応

---

## はじめに

### インストール

```bash
go build -o gdeck
mv gdeck /usr/local/bin/
```

または、ローカルで試すだけなら:

```bash
go build -o gdeck
```

> macOS では `ln -s $(pwd)/gdeck /usr/local/bin/gdeck` でシンボリックリンクを作成できます。

---

## 使い方ガイド

### 1. GET リクエスト

```bash
gdeck get https://example.com/api/status
```

#### オプション
- `-v`: 詳細出力
- `-o`: ファイル出力パス
- `-t`: タイムアウト（秒）

---

### 2. POST リクエスト

```bash
gdeck post https://example.com/api/items -d '{"name":"test"}'
```

#### オプション
- `-v`: 詳細出力
- `-o`: ファイル出力パス
- `-d`: リクエストボディ
- `-H`: リクエストヘッダー
- `-t`: タイムアウト（秒）

---

### 3. 保存して再利用

```bash
gdeck save SampleCmd POST https://example.com/api/items
```

#### 保存時オプション
- `-d`: リクエストボディ
- `-H`: リクエストヘッダー

保存済みコマンド内の `{{HOGE}}` 形式は、`gdeck run` 実行時に環境変数として置換されます。

---

### 4. 保存済みコマンドの実行

```bash
gdeck run SampleCmd
```

複数ファイルを一括実行する場合:

```bash
gdeck run "saved_commands/*"
```

#### オプション
- `-v`: 詳細出力
- `-d`: リクエストボディ上書き
- `-H`: リクエストヘッダー上書き
- `-t`: タイムアウト（秒）
- `--env`: 環境名を指定して環境変数ファイルを切り替え

#### 例

```bash
gdeck run SampleCmd --env dev
```

---

### 5. 保存済みコマンドの確認

```bash
gdeck show SampleCmd
```

```bash
gdeck list
```

---

### 6. 削除

```bash
gdeck delete SampleCmd
```

---

## TUI モード

`gdeck tui`

gdeck の TUI は、保存済みリクエストを左右ペインで見ながら操作できるインタラクティブ画面です。

### TUI の特徴

- 左ペイン: 保存リクエスト一覧
- 右ペイン: 選択中リクエストのプレビュー / レスポンス
- `↑` / `↓`: リクエスト選択
- `Enter`: 選択中リクエスト実行
- `←` / `→`: ペイン切り替え
- `q` / `ctrl+c`: 終了

---

## 環境変数管理 (`env`)

`gdeck env` では、環境変数の登録・参照・削除が可能です。

```bash
gdeck env set KEY VALUE [--env NAME]
gdeck env show KEY [--env NAME]
gdeck env list [--env NAME]
gdeck env delete KEY [--env NAME]
```

`--env` を指定すると、名前付き環境を切り替えて利用できます。

---

## 便利なヒント

- `gdeck tui` で直感的に保存リクエストを操作できます。
- 保存済みリクエストでは `{{KEY}}` を使い、実行時に値を置換できます。
- 開発時は `go fmt ./...` を忘れずに実行してください。

---

## 開発メモ

- CLI は `cobra` を使用しています。
- TUI は `bubbletea` / `lipgloss` で構築しています。
- コード修正後は `go fmt` を実行してください。
