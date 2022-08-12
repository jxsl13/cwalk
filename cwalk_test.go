package cwalk_test

import (
	"bytes"
	"io/fs"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/jxsl13/cwalk"
	"github.com/stretchr/testify/require"
)

func TestFullPath(t *testing.T) {
	require := require.New(t)
	testDir, err := os.Getwd()
	if err != nil {
		t.Error(err)
		return
	}

	err = cwalk.Walk(testDir, func(path string, info fs.FileInfo, err error) error {
		require.True(strings.HasPrefix(path, testDir), "prefix is not: ", testDir)
		return nil
	})
	if err != nil {
		t.Error(err)
		return
	}
}

func TestFewFiles(t *testing.T) {
	require := require.New(t)
	cmd := exec.Command("go", "env", "GOPATH")
	buf := bytes.Buffer{}
	cmd.Stdout = &buf

	err := cmd.Run()
	if err != nil {
		t.Error(err)
		return
	}

	goRoot := strings.TrimSpace(buf.String())

	cwalk.Walk(goRoot, func(path string, info fs.FileInfo, err error) error {
		require.True(strings.HasPrefix(path, goRoot), "prefix is not: ", goRoot)
		return nil
	})
}

func TestManyFiles(t *testing.T) {
	require := require.New(t)

	root := "/"

	cwalk.Walk(root, func(path string, info fs.FileInfo, err error) error {
		require.True(strings.HasPrefix(path, root), "prefix is not: ", root)
		return nil
	})
}

func TestSingleFile(t *testing.T) {
	require := require.New(t)

	issue := "/etc/issue"

	err := cwalk.Walk(issue, func(path string, info fs.FileInfo, err error) error {
		require.NoError(err)
		require.True(strings.HasPrefix(path, issue), "prefix is not: ", issue)
		return nil
	})
	require.NoError(err)
}
