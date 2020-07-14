package flyweight

import "fmt"

type ImageFlyweightFactory struct {
	maps map[string]*ImageFlyweight
}

// 共享不变的数据,用于复用
type ImageFlyweight struct {
	data string
}

func NewImageFlyweight(filename string) *ImageFlyweight {
	// Load image file
	data := fmt.Sprintf("image data %s", filename)
	return &ImageFlyweight{
		data: data,
	}
}

func (i *ImageFlyweight) Data() string {
	return i.data
}

// 单例管理共享不变的数据
var imageFactory = &ImageFlyweightFactory{make(map[string]*ImageFlyweight)}

func GetImageFlyweightFactory() *ImageFlyweightFactory {
	return imageFactory
}

func (f *ImageFlyweightFactory) Get(filename string) *ImageFlyweight {
	image := f.maps[filename]
	if image == nil {
		image = NewImageFlyweight(filename)
		f.maps[filename] = image
	}
	return image
}

// 调用实例
type ImageViewer struct {
	*ImageFlyweight
}

func NewImageViewer(filename string) *ImageViewer {
	image := GetImageFlyweightFactory().Get(filename)
	return &ImageViewer{
		ImageFlyweight: image,
	}
}

func (i *ImageViewer) Display() {
	fmt.Printf("Display: %s\n", i.Data())
}
