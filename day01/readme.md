# 心得
## 单例模式
> 1、该类在整个运行周期中仅能够被实例化一次
> 2、该类的实例化对象对外是不可见的，且必须自行提供一个公共的访问点供客户端去使用
> 3、该实例应被自行创建

而单例模式的实现又分为了两种，分别是饿汉模式与懒汉模式。

### 懒汉模式 
懒汉模式并不会在一开始就去实例化该单例，而是在第一次使用到它的时候，才会将其初始化并返回。。这就引伸出了一个问题，我们如何让这个单例只在第一次被调用的时候而初始化？换言之怎么让该实例被初始化的业务代码只能被全局调用一次？
```golang
package main

import (
	"fmt"
	"sync"
)

var once sync.Once

type singelton struct {}

var instance *singelton

func GetInstance() *singelton {

	once.Do(func(){
		instance = new(singelton)
	})

	return instance
}

func (s *singelton) DoPrint() {
	fmt.Println("666")
}

func main() {
	s := GetInstance()
	s.SomeThing()
}
```

## 工厂模式

### 简单工厂
只要请求工厂类，传入不同的参数，进而获取到对应的产品实例，不用关心其实现细节，只管获取实例使用即可.
