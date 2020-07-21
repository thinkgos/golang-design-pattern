package decorator

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDecorator(t *testing.T) {
	var c Component = new(BaseComponent)

	c = WarpAddDecorator(c, 10)
	c = WarpMulDecorator(c, 8)

	require.Equal(t, 80, c.Calc())
}
