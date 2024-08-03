package internal

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func readFile(t *testing.T, path string) File {
	t.Helper()
	f, err := ReadFile(path)
	require.NoError(t, err)
	return f
}
