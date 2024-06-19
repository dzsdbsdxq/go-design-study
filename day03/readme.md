# 代理模式

代理模式是一种常见的设计模式，它为其他对象提供了一个代理或占位符，以便在需要时控制对这些对象的访问。

在 Go 语言中，代理模式通常使用接口实现，以下是一个简单的例子，演示如何使用代理模式来控制对某个对象的访问：
```golang
// Subject 接口定义了需要被代理的对象的方法
type Subject interface {
    Request() string
}

// RealSubject 是需要被代理的对象
type RealSubject struct{}

func (r *RealSubject) Request() string {
    return "RealSubject: 处理请求"
}

// Proxy 是代理对象，它包含了一个指向 RealSubject 的引用
type Proxy struct {
    realSubject *RealSubject
}

func (p *Proxy) Request() string {
    // 在这里可以进行一些额外的操作，例如鉴权、日志等
    result := "Proxy: 转发请求到 RealSubject\n"
    if p.realSubject == nil {
        p.realSubject = &RealSubject{}
    }
    result += " " + p.realSubject.Request()
    return result
}

// 使用代理模式来使用 RealSubject 对象
func main() {
    proxy := &Proxy{}
    fmt.Println(proxy.Request())
}
```