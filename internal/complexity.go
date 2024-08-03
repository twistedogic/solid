package internal

import (
	"context"
)

type Complexity struct {
	Time       string
	Cyclomatic uint
}

const complexityTemplate = `Analyze the {{.Lang}} code snippet and return the time complexity and cyclomatic complexity in JSON ONLY.
Input:
{{.Snippet}}
`

func complexity(ctx context.Context, m Model, f File) (metrics Complexity, err error) {
	code := NewCodeSnippet(f)
	err = m.GenerateJSON(ctx, complexityTemplate, code, &metrics)
	return
}
