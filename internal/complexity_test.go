package internal

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_complexity(t *testing.T) {
	model := setupModel(t)
	file := readFile(t, "testdata/main.go")
	got, err := complexity(context.TODO(), model, file)
	require.NoError(t, err)
	require.NotEmpty(t, got.Time)
	require.NotEmpty(t, got.Cyclomatic)
}
