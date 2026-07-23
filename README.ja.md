[English](README.md) | **日本語**

<img src="assets/logo.png" alt="gDeck">

![gDeck demo](assets/demo.gif)

# gdeck

**gdeck** は、CLI と TUI の両方で API テストを実行できる軽量なツールです。

- ターミナルから HTTP リクエストを送信
- リクエスト定義の保存と再利用
- TUI で保存済みリクエストの閲覧・実行・編集・削除
- 実行時に `{{VAR}}` プレースホルダーを env ファイルで置換

---

## インストール

### Homebrew（macOS）

macOS 専用。[Releases](https://github.com/Nunokawa-Lab/gdeck/releases) のビルド済みバイナリをインストールします（Go は不要）。

**方法A — 1コマンド（Homebrew 6.0+。この Cask のみ trust）:**

```bash
brew install --cask Nunokawa-Lab/gdeck/gdeck
```

**方法B — tap してからインストール:**

[Homebrew 6.0+](https://docs.brew.sh/Tap-Trust) では、サードパーティ tap は trust しないと Cask が読み込まれません。tap 全体を trust するか、この Cask だけ trust してください:

```bash
brew tap Nunokawa-Lab/gdeck
brew trust Nunokawa-Lab/gdeck              # tap 全体を trust
brew install --cask gdeck
```

または gdeck の Cask だけ trust する（推奨）:

```bash
brew tap Nunokawa-Lab/gdeck
brew trust --cask Nunokawa-Lab/gdeck/gdeck
brew install --cask gdeck
```

Linux ユーザーは下記の **go install** または **バイナリをダウンロード** を利用してください。

アップグレード:

```bash
brew update
brew upgrade --cask gdeck
```

### go install

[Go](https://go.dev/dl/) 1.24 以降が必要です。

```bash
go install github.com/Nunokawa-Lab/gdeck@latest
```

`$HOME/go/bin`（または `$GOPATH/bin`）が `PATH` に通っていることを確認してください。

特定のバージョンをインストールする場合:

```bash
go install github.com/Nunokawa-Lab/gdeck@v0.1.1
```

### バイナリをダウンロード

ビルド済みバイナリは [Releases](https://github.com/Nunokawa-Lab/gdeck/releases) から入手できます。

**macOS (Apple Silicon)**

```bash
curl -LO https://github.com/Nunokawa-Lab/gdeck/releases/latest/download/gdeck-darwin-arm64
chmod +x gdeck-darwin-arm64
mv gdeck-darwin-arm64 /usr/local/bin/gdeck
```

**macOS (Intel)**

```bash
curl -LO https://github.com/Nunokawa-Lab/gdeck/releases/latest/download/gdeck-darwin-amd64
chmod +x gdeck-darwin-amd64
mv gdeck-darwin-amd64 /usr/local/bin/gdeck
```

**Linux**

```bash
curl -LO https://github.com/Nunokawa-Lab/gdeck/releases/latest/download/gdeck-linux-amd64
chmod +x gdeck-linux-amd64
sudo mv gdeck-linux-amd64 /usr/local/bin/gdeck
```

Windows の場合は [Releases](https://github.com/Nunokawa-Lab/gdeck/releases) から `gdeck-windows-amd64.exe` をダウンロードし、`PATH` に追加してください。

> **macOS:** 初回起動時に Gatekeeper によりブロックされる場合があります。そのときは以下を実行してください:
> ```bash
> xattr -d com.apple.quarantine gdeck-darwin-arm64   # または gdeck-darwin-amd64
> ```

---

## クイックスタート

```bash
# 1. プレースホルダー付きでリクエストを保存
gdeck save getUser GET https://api.example.com/users/{{USER_ID}} \
  -H 'Authorization: Bearer {{TOKEN}}'

# 2. 環境変数を設定
gdeck env set USER_ID 42
gdeck env set TOKEN abc123

# 3. 保存済みリクエストを実行
gdeck run getUser
```

---

## CLI コマンド

### `gdeck get [url]`

ワンショット GET リクエストを送信します。

```bash
gdeck get https://example.com/api/status
gdeck get https://example.com/api/status -v -o response.json -t 30
```

| フラグ | 短縮 | デフォルト | 説明 |
|--------|------|-----------|------|
| `--verbose` | `-v` | `false` | レスポンスヘッダーを出力に含める |
| `--output` | `-o` | `""` | レスポンスボディをファイルに書き出す |
| `--timeout` | `-t` | `10` | タイムアウト（秒） |

---

### `gdeck post [url]`

ワンショット POST リクエストを送信します。

```bash
gdeck post https://example.com/api/items \
  -d '{"name":"test"}' \
  -H 'Content-Type: application/json'
```

| フラグ | 短縮 | デフォルト | 説明 |
|--------|------|-----------|------|
| `--data` | `-d` | `""` | リクエストボディ |
| `--header` | `-H` | `[]` | リクエストヘッダー（繰り返し指定可、`Key: Value` 形式） |
| `--verbose` | `-v` | `false` | レスポンスヘッダーを出力に含める |
| `--output` | `-o` | `""` | レスポンスボディをファイルに書き出す |
| `--timeout` | `-t` | `10` | タイムアウト（秒） |

---

### `gdeck save [({path/})name] [method] [url]`

リクエスト定義を JSON として保存し、`run` や TUI から再利用できるようにします。

```bash
gdeck save SampleCmd POST https://example.com/api/items \
  -d '{"key":"{{TOKEN}}"}' \
  -H 'Content-Type: application/json'

# ネストパス対応
gdeck save api/users getUser GET https://api.example.com/users/{{USER_ID}}
```

| フラグ | 短縮 | デフォルト | 説明 |
|--------|------|-----------|------|
| `--data` | `-d` | `""` | リクエストボディ |
| `--header` | `-H` | `[]` | リクエストヘッダー（繰り返し指定可） |

- 任意の HTTP メソッドに対応（GET, POST, PUT, PATCH, DELETE など）
- `~/.gdeck/requests/{path/}name.json` に保存
- 名前に `..` を含む場合は拒否（パストラバーサル防止）
- 既存ファイルがある場合は上書きし、`Saved:` の代わりに `Updated:` を表示

---

### `gdeck run [({path/})name]`

保存済みリクエストを 1 件以上実行します。

```bash
gdeck run SampleCmd
gdeck run SampleCmd --env dev -v -t 30
gdeck run "saved_commands/*"
```

| フラグ | 短縮 | デフォルト | 説明 |
|--------|------|-----------|------|
| `--verbose` | `-v` | `false` | 詳細なレスポンス出力 |
| `--data` | `-d` | `""` | 保存済みボディの上書き |
| `--header` | `-H` | `[]` | ヘッダーの上書き / マージ |
| `--timeout` | `-t` | `10` | タイムアウト（秒） |
| `--env` | | `""` | 名前付き環境ファイル |

- 保存済みリクエストに記録された HTTP メソッドで実行
- URL / ボディ / ヘッダー内の `{{KEY}}` を env ファイルの値で置換
- ワイルドカードで複数リクエストを一括実行（例: `"folder/*"`）
- ヘッダー上書きは保存済みヘッダーとマージ（キーが重複した場合は上書き側が優先、大文字小文字は区別しない）

---

### `gdeck list`

保存済みリクエスト名を一覧表示します（1 行 1 件、拡張子なし）。

```bash
gdeck list
```

---

### `gdeck show [({path/})name]`

保存済みリクエストの詳細をインデント付き JSON で表示します。

```bash
gdeck show SampleCmd
gdeck show "folder/*"
```

---

### `gdeck delete [({path/})name]`

保存済みリクエストファイルを削除します。

```bash
gdeck delete SampleCmd
gdeck delete "folder/*"
```

成功時に `Status: 204 No-Content` を表示します。

---

## TUI モード

インタラクティブ UI を起動:

```bash
gdeck tui
```

TUI は自動起動しません — `tui` サブコマンドを明示的に実行してください。

### 画面構成

```
┌─────────────────────────────────────────┐
│  gdeck TUI                              │
├──────────────────┬──────────────────────┤
│  リクエスト (35%) │  プレビュー / レスポンス │
│  スクロール一覧   │  または保存 / 編集フォーム │
├──────────────────┴──────────────────────┤
│  コンテキスト別ショートカットバー         │
└─────────────────────────────────────────┘
```

- **左ペイン:** HTTP メソッドが色分けされた保存リクエスト一覧
- **右ペイン:** リクエストプレビュー、ローディング、レスポンス、または保存 / 編集フォーム

### キーバインド

#### 通常モード

| キー | 操作 |
|------|------|
| `q`, `Ctrl+c` | 終了 |
| `←` / `→` | 一覧と右ペインのフォーカス切り替え |
| `↑` / `↓` | カーソル移動、一覧スクロール、プレビュー更新（一覧フォーカス時） |
| `Enter` | 選択中リクエストを実行（一覧フォーカス時） |
| `/` | 検索モード（名前でフィルタ） |
| `s` | 新規保存フォームを開く |
| `e` | 選択中リクエストを編集 |
| `d` | 削除確認（一覧フォーカス時） |
| `↑` / `↓` | 右ペインをスクロール（レスポンスフォーカス時） |

#### 検索モード

| キー | 操作 |
|------|------|
| *入力* | 名前でフィルタ（大文字小文字を区別しない） |
| `↑` / `↓` | フィルタ結果を移動 |
| `Enter` | フィルタ結果のリクエストを実行 |
| `Esc` | 検索を終了し、以前の選択位置に戻る |

検索モードで実行が完了すると、フィルタが解除され、実行したリクエストの位置にカーソルが戻ります。

#### 保存 / 編集モード

| キー | 操作 |
|------|------|
| `Tab` | 次のフィールド |
| `Shift+Tab` | 前のフィールド |
| `Ctrl+s` | 保存 |
| `Esc` | キャンセル |

フォーム項目: Name, Method, URL, Headers（1 行 1 ヘッダー）, Body

- 編集時に名前を変更すると、新ファイルを保存して旧ファイルを削除
- ヘッダーとボディで `{{KEY}}` プレースホルダーに対応

#### 削除確認

| キー | 操作 |
|------|------|
| `y` | 削除を確定 |
| `n` | キャンセル |
| `Esc` | キャンセル（エラーメッセージ表示時） |

### TUI の制限

TUI は CLI と同じ store / runner を使いますが、以下の CLI 専用機能は提供しません:

- `gdeck get` / `gdeck post`（保存せずにワンショット実行）
- `gdeck env`（env ファイル管理）
- `run --env`, `-d`, `-H`, `-t`, `-v`（デフォルトの `~/.gdeck/.env` のみ使用）

---

## 環境変数

`env` サブコマンドで置換用の変数を管理します:

```bash
gdeck env set KEY VALUE [--env NAME]
gdeck env show KEY [--env NAME]
gdeck env list [--env NAME]
gdeck env delete KEY [--env NAME]
```

| `--env` の値 | ファイルパス |
|-------------|-------------|
| *(空 / 省略)* | `~/.gdeck/.env` |
| `dev` | `~/.gdeck/envs/dev.env` |

ファイル形式: 1 行 1 件の `KEY=VALUE`。`#` で始まる行と空行は無視されます。

保存済みリクエストでのプレースホルダー利用例:

```bash
gdeck save myAPI GET https://api.example.com/{{PATH}}
gdeck env set PATH users/42
gdeck run myAPI
```

- プレースホルダー形式: `{{WORD}}`（英数字とアンダースコアのみ）
- 実行時に URL / ボディ / ヘッダーへ適用
- 対応する env 値がない場合はエラーで停止

---

## データ保存先

すべてのデータは `~/.gdeck/` 配下に保存されます:

```
~/.gdeck/
├── .env                 # デフォルト env ファイル
├── envs/
│   └── dev.env          # 名前付き env ファイル
└── requests/
    ├── SampleCmd.json
    └── api/
        └── users/
            └── getUser.json
```

保存済みリクエストの JSON スキーマ:

```json
{
  "name": "SampleCmd",
  "method": "POST",
  "url": "https://api.example.com/{{PATH}}",
  "headers": ["Content-Type: application/json", "Authorization: Bearer {{TOKEN}}"],
  "body": "{\"key\":\"value\"}"
}
```

---

## 便利なヒント

- ネストパスでリクエストを整理: `gdeck save api/users getUser GET ...`
- ワイルドカードで複数リクエストを一括操作: `gdeck run "api/*"`
- 実行ごとに環境を切り替え: `gdeck run myAPI --env prod`
- `gdeck tui` で保存済みリクエストを対話的に閲覧・プレビュー・編集

---

## 技術スタック

| コンポーネント | ライブラリ |
|---------------|-----------|
| CLI | [Cobra](https://github.com/spf13/cobra) |
| TUI | [Bubble Tea](https://github.com/charmbracelet/bubbletea), [bubbles](https://github.com/charmbracelet/bubbles), [Lipgloss](https://github.com/charmbracelet/lipgloss) |
| 言語 | Go 1.24 |

コントリビュート時は、コミット前に `go fmt ./...` を実行してください。
