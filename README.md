# 🛠️ tools

A bag of toys. A project to take ownership over the terminal tools I use every day.

## Projects

### grup (grep)

A replacement for (my usage of) `grep`. Faster than `grep`.

Options:

- `-n` show line numbers
- `-re` treat query as a regular expression
- `-w` (debug) set number of search workers

Quirks:

- Recursively searches by default (similar to other `grep` replacements)
- RegExp syntax: https://github.com/google/re2/wiki/Syntax

Example:

```bash
grup -n -re ^module .
# go.mod:1 module github.com/healeycodes/tools

grup -n "\"bufio\"" .
# utils/search.go:4 	"bufio"
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

A replacement for (my usage of) `cat`.

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
