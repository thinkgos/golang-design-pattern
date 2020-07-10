package facade

import "fmt"

//API is facade interface of facade package
type API interface {
	Test() string
}

// AModuleAPI 细节1
type AModuleAPI interface {
	TestA() string
}

// BModuleAPI 细节2
type BModuleAPI interface {
	TestB() string
}

//facade implement
type apiImpl struct {
	a AModuleAPI
	b BModuleAPI
}

func NewAPI() API {
	return &apiImpl{
		a: NewAModuleAPI(),
		b: NewBModuleAPI(),
	}
}

func (a *apiImpl) Test() string {
	aRet := a.a.TestA()
	bRet := a.b.TestB()
	return fmt.Sprintf("%s\n%s", aRet, bRet)
}

type aModuleImpl struct{}

// NewAModuleAPI return new AModuleAPI
func NewAModuleAPI() AModuleAPI {
	return &aModuleImpl{}
}

func (*aModuleImpl) TestA() string {
	return "A module running"
}

type bModuleImpl struct{}

//NewBModuleAPI return new BModuleAPI
func NewBModuleAPI() BModuleAPI {
	return &bModuleImpl{}
}

func (*bModuleImpl) TestB() string {
	return "B module running"
}
