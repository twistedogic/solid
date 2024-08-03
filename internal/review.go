package internal

import (
	"context"
)

const reviewTemplate = `You are a senior developer who add code review comment inline ONLY.
Analyze the {{.Lang}} code snippet that will be provided to you to review, focusing on potential bugs, code intent, naming and performance. Submit a version that is with code review comments.
- **no quotes**
- **plain {{.Lang}} code ONLY**
- **no explanation**
Input:
{{.Snippet}}
`

func review(ctx context.Context, m Model, f File) (string, error) {
	code := NewCodeSnippet(f)
	res, err := m.Generate(ctx, reviewTemplate, code)
	return extractCode(res), err
}

func ReviewCode(ctx context.Context, m Model, f File) (string, error) {
	code := NewCodeSnippet(f)
	res, err := m.Generate(ctx, reviewTemplate, code)
	return res, err
}
