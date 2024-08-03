package internal

import (
	"context"
	"math"
)

func unique(vals []string) []string {
	m := make(map[string]struct{}, len(vals))
	for _, v := range vals {
		m[v] = struct{}{}
	}
	uniqueVals := make([]string, 0, len(m))
	for k := range m {
		uniqueVals = append(uniqueVals, k)
	}
	return uniqueVals
}

type HalsteadMetrics struct {
	Operators []string
	Operands  []string
}

func (h HalsteadMetrics) operators() float64         { return float64(len(h.Operators)) }
func (h HalsteadMetrics) operands() float64          { return float64(len(h.Operands)) }
func (h HalsteadMetrics) distinctOperators() float64 { return float64(len(unique(h.Operators))) }
func (h HalsteadMetrics) distinctOperands() float64  { return float64(len(unique(h.Operands))) }
func (h HalsteadMetrics) length() float64            { return h.operators() + h.operands() }

func (h HalsteadMetrics) Volume() float64 {
	return h.length() * math.Log2(h.distinctOperators()+h.distinctOperands())
}

func (h HalsteadMetrics) Difficulty() float64 {
	return h.distinctOperators() / 2.0 * h.operands() / h.distinctOperands()
}

func (h HalsteadMetrics) Effort() float64 { return h.Difficulty() * h.Volume() }

const halsteadTemplate = `Analyze the {{.Lang}} code snippet and return all the {{.Lang}} keywords as operators and the rest as operands in JSON ONLY.
Input:
{{.Snippet}}
`

func halstead(ctx context.Context, m Model, f File) (metrics HalsteadMetrics, err error) {
	code := NewCodeSnippet(f)
	err = m.GenerateJSON(ctx, halsteadTemplate, code, &metrics)
	return
}
