package singleton

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSingleton(t *testing.T) {
	ins1 := GetInstance()
	ins2 := GetInstance()
	require.Equal(t, ins1, ins2)
}

func TestParallelSingleton(t *testing.T) {
	const parCount = 100

	wg := sync.WaitGroup{}
	instances := make([]*Singleton, parCount)

	wg.Add(parCount)
	for i := 0; i < parCount; i++ {
		go func(index int) {
			instances[index] = GetInstance()
			wg.Done()
		}(i)
	}
	wg.Wait()

	for i := 1; i < parCount; i++ {
		require.Equal(t, instances[i-1], instances[i])
	}
}
