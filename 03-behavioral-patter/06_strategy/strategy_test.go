package strategy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStrategy(t *testing.T) {
	cpu := NewComputer(100, 200, &Add{})

	require.Equal(t, 300, cpu.Do())
	cpu.SetStrategy(&Mul{})
	require.Equal(t, 20000, cpu.Do())
	cpu.SetStrategy(&Sub{})
	require.Equal(t, -100, cpu.Do())
}
