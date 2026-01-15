# ghcat

`ghcat` is a small CLI tool that fetches a **single file from GitHub** and prints it to `stdout`.

It is designed to work like `cat`, but for GitHub repositories.

> Fetch smart, display with your favorite tools.

---

## Why

If you frequently:

- open files on GitHub just to copy them
- download ZIPs for a single file
- manually build `raw.githubusercontent.com` URLs

`ghcat` removes that friction.

It resolves branches, handles auth when needed, and stays pipe-friendly.

---

## Installation

### Using Go

```bash
go install github.com/eos175/ghcat@latest
```

---

## Usage

### Short form (recommended)

```bash
ghcat owner/repo/path/to/file.go
```

Example:

```bash
ghcat eos175/tachyon-boilerplate/pkg/utils/utils.go
```

`ghcat` will automatically discover the repository default branch (`main`, `master`, etc).

---

### GitHub UI URLs

```bash
ghcat https://github.com/owner/repo/blob/branch/path/to/file.go
```

Example:

```bash
ghcat https://github.com/golang/go/blob/master/src/fmt/print.go
```

---

### Raw GitHub URLs

```bash
ghcat https://raw.githubusercontent.com/owner/repo/branch/path/to/file.go
```

> ⚠️ Raw URLs only work for **public repositories**.

---

### GitHub Gist URLs

```bash
ghcat https://gist.github.com/owner/gist_id
```

Example:

```bash
ghcat https://gist.github.com/eos175/cd4265e9d050c90dbfd11fa60f024732
```

---

## Piping (recommended)

`ghcat` always writes plain text to `stdout`, making it ideal for pipes:

```bash
ghcat owner/repo/file.go | bat
ghcat owner/repo/file.go | less
ghcat owner/repo/file.go | sed -n '1,50p'
```

---

## Authentication

`ghcat` automatically uses authentication if available.

Resolution order:

1. `GITHUB_TOKEN` environment variable
2. `~/.github_token` file
3. No authentication (public repositories)

For private repositories, authentication is required.

Recommended permissions for `~/.github_token`:

```bash
chmod 600 ~/.github_token
```

---

## What ghcat does NOT do

- No browsing
- No pagination
- No formatting
- No syntax highlighting
- No flags for line numbers or display

`ghcat` does one thing: **fetch a file and print it**.

Use existing tools (`bat`, `less`, `sed`, `tail`) for everything else.

---

## Philosophy

- One job
- Explicit inputs
- No heuristics
- No magic
- Unix-style composition

---

## License

MIT
