package utils

import (
	"html/template"
	"os"
	"path/filepath"
	"testing"

	"github.com/barrett370/sally/templates"
	"github.com/stretchr/testify/require"
)

func GetCombinedTemplates(dir string) (*template.Template, error) {
	// Clones default templates to then merge with the user defined templates.
	// This allows for the user to only override certain templates, but not all
	// if they don't want.
	templates, err := templates.Templates.Clone()
	if err != nil {
		return nil, err
	}
	return templates.ParseGlob(filepath.Join(dir, "*.html"))
}

// TempFile persists contents and returns the path and a clean func
func TempFile(t *testing.T, contents string) (path string) {
	content := []byte(contents)
	tmpfile, err := os.CreateTemp("", "sally-tmp")
	require.NoError(t, err, "unable to create tmpfile")

	_, err = tmpfile.Write(content)
	require.NoError(t, err, "unable to write tmpfile")

	err = tmpfile.Close()
	require.NoError(t, err, "unable to close tmpfile")

	t.Cleanup(func() {
		_ = os.Remove(tmpfile.Name())
	})

	return tmpfile.Name()
}
