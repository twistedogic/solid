package internal

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_halstead(t *testing.T) {
	model := setupModel(t)
	file := readFile(t, "testdata/main.go")
	got, err := halstead(context.TODO(), model, file)
	require.NoError(t, err)
	require.NotEmpty(t, got.Operators)
	require.NotEmpty(t, got.Operands)
}
