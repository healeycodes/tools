# 🛠️ tools

This is a learning project and a chance to take ownership over the common terminal tools I rely on every day.

I hope to add to this monorepo over time.

## el (ls)

A replacement for (my usage of) `ls`. Turns out I rely on about 1% of the available [flags](https://man7.org/linux/man-pages/man1/ls.1.html).

Options:

- `-h` ignore entries starting with `.`
- `-d` display, and sort by, date

Quirks:

- Display dot files by default
- Sorted alphabetically (with a lowercase compare)
- Directories are put first (unless `-d` is passed)

Example:

```
el
# utils/  .gitignore  el  go.mod  main.go  README.md
```

## Tests

Test all projects.

```bash
go test ./...
```
