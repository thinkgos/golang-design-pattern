package simplefactory

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHiAPI_Say(t *testing.T) {
	api := NewAPI(hi)
	s := api.Say("Tom")

	require.Equal(t, "Hi, Tom", s)
}

func TestHelloAPI_Say(t *testing.T) {
	api := NewAPI(hello)
	s := api.Say("Tom")

	require.Equal(t, "Hello, Tom", s)
}
