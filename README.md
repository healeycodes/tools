# 🛠️ tools

This is a learning project and a chance to take ownership over the common terminal tools I rely on every day.

I hope to add to this monorepo over time.

## Projects

### grup (grep)

A replacement for (my usage of) `grep`.

Quirks:

- Search data from a pipe, or search the given files and/or directories
- There's a space after the line number so, e.g. with VS Code, you can click through to the file
- RegExp syntax: https://github.com/google/re2/wiki/Syntax

Example:

```bash
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
- `-d` sort by date and display metadata

Quirks:

- Display dot files by default
- Sorted alphabetically (with a lowercase compare)
- Directories are put first (unless `-d` is passed)

Example:

```bash
el
# utils/  .gitignore  el  go.mod  main.go  README.md

el -d utils/
# 2022-04-17 13:27  695 B listfiles_test.go
# 2022-04-17 13:18 2.1 kB listfiles.go
# 2022-04-15 13:49    0 B listfiles_test_files/
```

### kat (cat)

A replacement for (my usage of) `cat`. I've only ever passed a file argument or piped to `cat`.

Example:

```bash
echo 1 | kat
# 1

echo 1 >> some_file
kat some_file
# 1
```

## Tests

Test all projects.

```bash
./test.sh
```

## Build

Build all projects (for: windows, linux, darwin).

```bash
./build.sh
```
