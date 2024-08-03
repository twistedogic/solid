package internal

import (
	"context"
	"go/parser"
	"go/token"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_review(t *testing.T) {
	model := setupModel(t)
	file := readFile(t, "testdata/main.go")
	got, err := review(context.TODO(), model, file)
	require.NoError(t, err)
	fset := token.NewFileSet()
	_, err = parser.ParseFile(fset, "", got, 0)
	require.NoError(t, err)
}
