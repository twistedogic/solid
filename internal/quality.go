package internal

import (
	"context"
	"math"
)

func maintainabilityIndex(f File, h HalsteadMetrics, c Complexity) float64 {
	loc := f.NumberOfLines()
	index := (171.0 - 5.2*math.Log(h.Volume()) - 0.23*float64(c.Cyclomatic) - 16.2*math.Log(float64(loc))) * 100 / 171
	if index < 0.0 {
		return 0
	}
	return index
}

func MI(ctx context.Context, m Model, f File) (float64, error) {
	h, err := halstead(ctx, m, f)
	if err != nil {
		return 0, err
	}
	c, err := complexity(ctx, m, f)
	return maintainabilityIndex(f, h, c), err
}

const miTemplate = `You are senior code analyst.
Analyze the {{.Lang}} code snippet and return Maintainability Index.
Input:
{{.Snippet}}
`

func MaintainabilityIndex(ctx context.Context, m Model, f File) (string, error) {
	code := NewCodeSnippet(f)
	return m.Generate(ctx, miTemplate, code)
}
