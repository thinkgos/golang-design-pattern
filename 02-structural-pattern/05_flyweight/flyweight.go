package flyweight

import "fmt"

/*********** 重复不变的数据 ******/

// 共享不变的数据,用于复用
type ImageFlyweight struct {
	data string
}

func NewImageFlyweight(filename string) *ImageFlyweight {
	return &ImageFlyweight{
		data: fmt.Sprintf("image data %s", filename), // Load image file
	}
}

func (i *ImageFlyweight) Data() string {
	return i.data
}

/***** 单例共享 *****/

// 单例管理共享不变的数据
var imageFactory = &ImageFlyweightFactory{
	make(map[string]*ImageFlyweight),
}

type ImageFlyweightFactory struct {
	maps map[string]*ImageFlyweight
}

func (f *ImageFlyweightFactory) Get(filename string) *ImageFlyweight {
	image := f.maps[filename]
	if image == nil {
		image = NewImageFlyweight(filename)
		f.maps[filename] = image
	}
	return image
}

// 调用实例,调用共享数据
type ImageViewer struct {
	*ImageFlyweight
}

func NewImageViewer(filename string) *ImageViewer {
	return &ImageViewer{
		imageFactory.Get(filename),
	}
}

func (i *ImageViewer) Display() {
	fmt.Printf("Display: %s\n", i.Data())
}
