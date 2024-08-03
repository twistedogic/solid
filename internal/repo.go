package internal

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

var langMap = map[string]string{
	".go":   "go",
	".rs":   "rust",
	".md":   "markdown",
	".sh":   "bash",
	".java": "java",
	".js":   "javascript",
	".ts":   "typescript",
	".py":   "python",
}

type File struct {
	Name    string
	Content string
}

func (f File) NumberOfLines() int {
	loc := 0
	for _, line := range strings.Split(f.Content, "\n") {
		if strings.TrimSpace(line) != "" {
			loc += 1
		}
	}
	return loc
}

func ReadFile(path string) (f File, err error) {
	f.Name = path
	file, err := os.Open(path)
	if err != nil {
		return
	}
	b, err := io.ReadAll(file)
	if err != nil {
		return
	}
	f.Content = string(b)
	return
}

type CodeSnippet struct {
	Name    string
	Lang    string
	Snippet string
}

func NewCodeSnippet(f File) CodeSnippet {
	lang := ""
	if e, ok := langMap[filepath.Ext(f.Name)]; ok {
		lang = e
	}
	return CodeSnippet{Name: f.Name, Lang: lang, Snippet: codeSnippet(f.Content, lang)}
}

func extractCode(snippet string) string {
	start := false
	lines := strings.Split(snippet, "\n")
	code := make([]string, 0, len(lines))
	for _, line := range lines {
		switch {
		case strings.HasPrefix(line, "```"):
			start = !start
		case start:
			code = append(code, line)
		}
	}
	if len(code) == 0 {
		return snippet
	}
	return strings.Join(code, "\n")
}
