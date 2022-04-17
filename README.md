# 🛠️ tools

This is a learning project and a chance to take ownership over the common terminal tools I rely on every day.

I hope to add to this monorepo over time.

## Tests

Test all projects.

```bash
go test ./...
```

## Projects

### grup (grep)

A replacement for (my usage of) `grep`.

Quirks:

- Search the pipe, or recursively search the given directory
- There's a space after the line number so, e.g. with VS Code, you can click through to the file
- RegExp syntax: https://github.com/google/re2/wiki/Syntax

Example:

```
grup ^module .
# go.mod:1 module github.com/healeycodes/tools

grup "\"bufio\"" .
# main.go:4       "bufio"

echo 1 | grup 1
# 1
```

### el (ls)

A replacement for (my usage of) `ls`.

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

### kat (cat)

A replacement for (my usage of) `cat`. I've only ever passed a file argument or piped to `cat`.

Example:

```
echo 1 | kat
# 1

echo 1 >> some_file
kat some_file
# 1
```
