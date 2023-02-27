package gohttp

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsAcceptable(t *testing.T) {
	require.True(t, isAcceptable(`application\/(hal[+\-])?json`, "application/json; charset utf-8"))
	require.True(t, isAcceptable(`application\/(hal[+\-])?json`, "application/json; charset utf-8"))
	require.False(t, isAcceptable(`application\/(hal[+\-])?json`, "text/html; charset utf-8"))
}
