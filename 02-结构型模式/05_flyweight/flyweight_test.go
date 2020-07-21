package flyweight

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func ExampleFlyweight() {
	viewer := NewImageViewer("image1.png")
	viewer.Display()
	// Output:
	// Display: image data image1.png
}

func TestFlyweight(t *testing.T) {
	viewer1 := NewImageViewer("image1.png")
	viewer2 := NewImageViewer("image1.png")

	require.Equal(t, viewer1.ImageFlyweight, viewer2.ImageFlyweight)
}
