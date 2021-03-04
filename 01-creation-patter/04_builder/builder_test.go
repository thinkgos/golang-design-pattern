package builder

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBuilder1(t *testing.T) {
	builder := &Builder1{}
	director := NewDirector(builder)
	director.Construct()
	res := builder.GetResult()
	require.Equal(t, "123", res)
}

func TestBuilder2(t *testing.T) {
	builder := &Builder2{}
	director := NewDirector(builder)
	director.Construct()
	res := builder.GetResult()
	require.Equal(t, 6, res)
}
