package proxy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProxy(t *testing.T) {
	realSub := new(RealSubject)
	proxy := NewProxy(realSub)

	require.Equal(t, "pre:real:after", proxy.Do())
}
