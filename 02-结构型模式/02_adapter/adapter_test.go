package adapter

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAdapter(t *testing.T) {
	provider := NewAdaptee()
	target := NewAdapter(provider)

	require.Equal(t, "Adaptee method", target.Request())
}
