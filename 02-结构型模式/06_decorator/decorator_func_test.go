package decorator

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDecoratorFunc(t *testing.T) {
	c := NewEmptyBase()

	c = WarpAddDecoratorFunc(c, 10)
	c = WarpMulDecoratorFunc(c, 8)

	require.Equal(t, 80, c())
}
