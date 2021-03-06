package utils

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

var (
	DOT       byte = 46
	SEPARATOR      = "  "
	LINEBREAK      = "\n"
)

type file struct {
	dir  bool
	name string
	date time.Time
	size int
}

func ListFiles(path string, hideDotFiles bool, details bool) string {
	dir, err := os.ReadDir(path)
	if err != nil {
		log.Fatalf("Couldn't read %s: %s", dir, err)
	}

	files := make([]file, 0)
	for _, item := range dir {
		info, err := item.Info()
		if err != nil {
			log.Fatalf("Couldn't stat %s: %s", item, err)
		}

		name := info.Name()
		if !hideDotFiles || len(name) > 0 && name[0] != DOT {
			files = append(files, file{
				info.IsDir(),
				name,
				info.ModTime(),
				int(info.Size()),
			})
		}
	}

	if details {
		// Most recent first
		sort.Slice(files, func(i, j int) bool {
			return files[i].date.After(files[j].date)
		})
	} else {
		// Alphabetically
		sort.Slice(files, func(i, j int) bool {
			return strings.ToLower(files[i].name) < strings.ToLower(files[j].name)
		})

		// Directories first
		sort.Slice(files, func(i, j int) bool {
			return files[i].dir
		})
	}

	output := make([]string, 0)
	for _, file := range files {
		output = append(output, formatFile(file, details))
	}

	if details {
		return strings.Join(output, LINEBREAK)
	} else {
		return strings.Join(output, SEPARATOR)
	}
}

func formatFile(file file, details bool) string {
	if details {
		return fmt.Sprintf("%s %6s %s%s", file.date.Format("2006-01-02 15:04"), byteCountSI(file.size), file.name, dirChar(file.dir))
	}
	return fmt.Sprintf("%s%s", file.name, dirChar(file.dir))
}

func dirChar(dir bool) string {
	if dir {
		return "/"
	}
	return ""
}

// https://yourbasic.org/golang/formatting-byte-size-to-human-readable-format/
func byteCountSI(b int) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB",
		float64(b)/float64(div), "kMGTPE"[exp])
}
