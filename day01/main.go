package day01

import (
	"fmt"
	"sync"
)

// SimpleSingleton 简单单例模式
type SimpleSingleton struct {
	a string
}

var ss *SimpleSingleton

func GetInstance() *SimpleSingleton {
	return ss
}

//懒汉式单例模式

var (
	lazySingleton *SimpleSingleton
	once          = &sync.Once{}
)

func GetLazyInstance() *SimpleSingleton {
	once.Do(func() {
		lazySingleton = &SimpleSingleton{
			a: "hello world",
		}
	})
	return lazySingleton
}

// IApiController 工厂模式
type IApiController interface {
	GetFiles()
}
type ApiController struct{}

func (a *ApiController) GetFiles() {
	fmt.Println("api get files")
}

func NewApiController() IApiController {
	return &ApiController{}
}
