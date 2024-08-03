package internal

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func setupModel(t *testing.T) Model {
	t.Helper()
	model, err := DefaultModel()
	if err != nil {
		t.Skip("skipping test for local ollama.")
	}
	return model
}

type input struct{ A, B int }
type output struct{ Result int }

func Test_ollama_GenerateJSON(t *testing.T) {
	tmpl := `Return the sum of {{.A}} and {{.B}} in JSON ONLY.`
	req := input{A: 1, B: 10}
	got := output{}
	want := output{Result: 11}
	model := setupModel(t)
	require.NoError(t, model.GenerateJSON(context.TODO(), tmpl, req, &got))
	require.Equal(t, want, got)
}
