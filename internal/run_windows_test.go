package internal_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tufin/oasdiff/internal"
)

func Test_InvalidFileWindows(t *testing.T) {
	var stderr bytes.Buffer
	require.Equal(t, 102, internal.Run(cmdToArgs("oasdiff diff no-file ../data/openapi-test3.yaml"), io.Discard, &stderr))
	require.Equal(t, "Error: failed to load base spec from \"no-file\": open no-file: The system cannot find the file specified.\n", stderr.String())
}

func Test_ComposedModeInvalidFileWindows(t *testing.T) {
	var stderr bytes.Buffer
	require.Equal(t, 103, internal.Run(cmdToArgs("oasdiff diff ../data/allof/* ../data/allof/* --composed --flatten"), io.Discard, &stderr))
	require.Equal(t, "Error: failed to load base specs from glob \"../data/allof/*\": failed to flatten allOf in \"..\\data\\allof\\invalid.yaml\": unable to resolve Type conflict: all Type values must be identical\n", stderr.String())
}