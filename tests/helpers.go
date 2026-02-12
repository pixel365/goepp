package tests

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/pixel365/goepp"

	"github.com/pixel365/goepp/command"
)

func mustRead(t *testing.T, path string) []byte {
	t.Helper()

	path = filepath.Clean(path)
	b, err := os.ReadFile(path)

	require.NoError(t, err)

	return b
}

func mustParse(t *testing.T, parser goepp.Parser, path string) command.Commander {
	t.Helper()

	path = filepath.Clean(path)
	cmd, err := parser.Parse(mustRead(t, path))

	require.NoError(t, err)

	return cmd
}

func mustParseAndValidate(t *testing.T, parser goepp.Parser, path string) {
	t.Helper()

	path = filepath.Clean(path)
	cmd := mustParse(t, parser, path)

	require.NoError(t, cmd.Validate())
}

func mustFailParse(t *testing.T, parser goepp.Parser, path string) {
	t.Helper()

	path = filepath.Clean(path)
	_, err := parser.Parse(mustRead(t, path))

	require.Error(t, err)
}
