**English** | [日本語](README.ja.md)

<img src="assets/logo.png" alt="gDeck">

![gDeck demo](assets/demo.gif)

# gdeck

**gdeck** is a lightweight API testing tool that works from both the CLI and an interactive TUI.

- Send HTTP requests from the terminal
- Save and reuse request definitions
- Browse, run, edit, and delete saved requests in the TUI
- Substitute `{{VAR}}` placeholders at run time via env files

---

## Installation

```bash
go build -o gdeck
mv gdeck /usr/local/bin/
```

To try it locally without installing globally:

```bash
go build -o gdeck
./gdeck --help
```

> On macOS, you can also symlink: `ln -s $(pwd)/gdeck /usr/local/bin/gdeck`

---

## Quick Start

```bash
# 1. Save a request with a placeholder
gdeck save getUser GET https://api.example.com/users/{{USER_ID}} \
  -H 'Authorization: Bearer {{TOKEN}}'

# 2. Set env values
gdeck env set USER_ID 42
gdeck env set TOKEN abc123

# 3. Run the saved request
gdeck run getUser
```

---

## CLI Commands

### `gdeck get [url]`

Send a one-off GET request.

```bash
gdeck get https://example.com/api/status
gdeck get https://example.com/api/status -v -o response.json -t 30
```

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--verbose` | `-v` | `false` | Include response headers in output |
| `--output` | `-o` | `""` | Write response body to a file |
| `--timeout` | `-t` | `10` | Timeout in seconds |

---

### `gdeck post [url]`

Send a one-off POST request.

```bash
gdeck post https://example.com/api/items \
  -d '{"name":"test"}' \
  -H 'Content-Type: application/json'
```

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--data` | `-d` | `""` | Request body |
| `--header` | `-H` | `[]` | Request header (repeatable, `Key: Value`) |
| `--verbose` | `-v` | `false` | Include response headers in output |
| `--output` | `-o` | `""` | Write response body to a file |
| `--timeout` | `-t` | `10` | Timeout in seconds |

---

### `gdeck save [({path/})name] [method] [url]`

Save a request definition as JSON for later reuse via `run` or the TUI.

```bash
gdeck save SampleCmd POST https://example.com/api/items \
  -d '{"key":"{{TOKEN}}"}' \
  -H 'Content-Type: application/json'

# Nested path support
gdeck save api/users getUser GET https://api.example.com/users/{{USER_ID}}
```

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--data` | `-d` | `""` | Request body |
| `--header` | `-H` | `[]` | Request header (repeatable) |

- Any HTTP method is supported (GET, POST, PUT, PATCH, DELETE, etc.)
- Saved to `~/.gdeck/requests/{path/}name.json`
- Names containing `..` are rejected (path traversal guard)
- Overwrites an existing file and prints `Updated:` instead of `Saved:`

---

### `gdeck run [({path/})name]`

Execute one or more saved requests.

```bash
gdeck run SampleCmd
gdeck run SampleCmd --env dev -v -t 30
gdeck run "saved_commands/*"
```

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--verbose` | `-v` | `false` | Verbose response output |
| `--data` | `-d` | `""` | Override saved request body |
| `--header` | `-H` | `[]` | Override/merge headers |
| `--timeout` | `-t` | `10` | Timeout in seconds |
| `--env` | | `""` | Named environment file |

- Uses the HTTP method stored in the saved request
- Replaces `{{KEY}}` placeholders in URL, body, and headers using env files
- Wildcards match multiple saved requests (e.g. `"folder/*"`)
- Header overrides merge with saved headers (override wins by key, case-insensitive)

---

### `gdeck list`

List all saved request names (one per line, extension stripped).

```bash
gdeck list
```

---

### `gdeck show [({path/})name]`

Display saved request detail as indented JSON.

```bash
gdeck show SampleCmd
gdeck show "folder/*"
```

---

### `gdeck delete [({path/})name]`

Delete saved request file(s).

```bash
gdeck delete SampleCmd
gdeck delete "folder/*"
```

Prints `Status: 204 No-Content` on success.

---

## TUI Mode

Launch the interactive UI:

```bash
gdeck tui
```

The TUI is not launched automatically — use the `tui` subcommand explicitly.

### Layout

```
┌─────────────────────────────────────────┐
│  gdeck TUI                              │
├──────────────────┬──────────────────────┤
│  Requests (35%)  │  Preview / Response  │
│  scrollable list │  or Save/Edit form   │
├──────────────────┴──────────────────────┤
│  context-sensitive shortcut bar         │
└─────────────────────────────────────────┘
```

- **Left pane:** saved request list with color-coded HTTP methods
- **Right pane:** request preview, loading spinner, response, or save/edit form

### Keybindings

#### Normal mode

| Key | Action |
|-----|--------|
| `q`, `Ctrl+c` | Quit |
| `←` / `→` | Switch focus between list and right pane |
| `↑` / `↓` | Move cursor, scroll list, load preview (list focus) |
| `Enter` | Run selected request (list focus) |
| `/` | Enter search mode (filter by name) |
| `s` | Open new save form |
| `e` | Edit selected request |
| `d` | Delete confirmation (list focus) |
| `↑` / `↓` | Scroll right pane (response focus) |

#### Search mode

| Key | Action |
|-----|--------|
| *typing* | Filter requests by name (case-insensitive) |
| `↑` / `↓` | Navigate filtered list |
| `Enter` | Run filtered request |
| `Esc` | Exit search and restore previous selection |

After a run completes in search mode, the filter is cleared and the cursor returns to the executed request.

#### Save / Edit mode

| Key | Action |
|-----|--------|
| `Tab` | Next field |
| `Shift+Tab` | Previous field |
| `Ctrl+s` | Save |
| `Esc` | Cancel |

Form fields: Name, Method, URL, Headers (one per line), Body.

- Editing with a renamed name saves the new file and deletes the old one
- `{{KEY}}` placeholders are supported in headers and body

#### Delete confirmation

| Key | Action |
|-----|--------|
| `y` | Confirm delete |
| `n` | Cancel |
| `Esc` | Cancel (when an error message is shown) |

### TUI limitations

The TUI uses the same store and runner as the CLI, but does not expose these CLI-only options:

- `gdeck get` / `gdeck post` (ad-hoc requests without saving)
- `gdeck env` (env file management)
- `run --env`, `-d`, `-H`, `-t`, `-v` (uses default `~/.gdeck/.env` only)

---

## Environment Variables

Manage substitution variables with the `env` subcommand:

```bash
gdeck env set KEY VALUE [--env NAME]
gdeck env show KEY [--env NAME]
gdeck env list [--env NAME]
gdeck env delete KEY [--env NAME]
```

| `--env` value | File path |
|---------------|-----------|
| *(empty / omitted)* | `~/.gdeck/.env` |
| `dev` | `~/.gdeck/envs/dev.env` |

File format: `KEY=VALUE` per line. Lines starting with `#` and blank lines are ignored.

Use placeholders in saved requests:

```bash
gdeck save myAPI GET https://api.example.com/{{PATH}}
gdeck env set PATH users/42
gdeck run myAPI
```

- Placeholder pattern: `{{WORD}}` (alphanumeric and underscore only)
- Applied to URL, body, and headers at run time
- Run fails with an error if a placeholder has no matching env value

---

## Data Storage

All data is stored under `~/.gdeck/`:

```
~/.gdeck/
├── .env                 # default env file
├── envs/
│   └── dev.env          # named env files
└── requests/
    ├── SampleCmd.json
    └── api/
        └── users/
            └── getUser.json
```

Saved request JSON schema:

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

## Tips

- Use nested paths to organize requests: `gdeck save api/users getUser GET ...`
- Run, show, or delete multiple requests at once with wildcards: `gdeck run "api/*"`
- Switch environments per run: `gdeck run myAPI --env prod`
- Use `gdeck tui` for browsing, previewing, and editing saved requests interactively

---

## Tech Stack

| Component | Library |
|-----------|---------|
| CLI | [Cobra](https://github.com/spf13/cobra) |
| TUI | [Bubble Tea](https://github.com/charmbracelet/bubbletea), [bubbles](https://github.com/charmbracelet/bubbles), [Lipgloss](https://github.com/charmbracelet/lipgloss) |
| Language | Go 1.24 |

When contributing, run `go fmt ./...` before committing.
