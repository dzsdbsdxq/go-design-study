# 建造者模型
在 Go 语言中，建造者模式可以用于创建复杂的对象，将对象的构建过程和表示分离开来。
```golang
// PersonInterface 接口表示一个人的信息和行为
type PersonInterface interface {
    Speak(message string)
    Sleep()
}

// Person 结构体实现了 PersonInterface 接口
type Person struct {
    name    string
    age     int
    gender  string
    address string
}

// Speak 方法实现了 PersonInterface 接口中的说话方法
func (p *Person) Speak(message string) {
    fmt.Printf("%s says: %s\n", p.name, message)
}

// Sleep 方法实现了 PersonInterface 接口中的睡觉方法
func (p *Person) Sleep() {
    fmt.Printf("%s is sleeping now...\n", p.name)
}

// PersonBuilderInterface 接口表示一个人的信息构造器
type PersonBuilderInterface interface {
    SetName(name string) PersonBuilderInterface
    SetAge(age int) PersonBuilderInterface
    SetGender(gender string) PersonBuilderInterface
    SetAddress(address string) PersonBuilderInterface
    Build() PersonInterface
}

// PersonBuilder 结构体表示一个人的建造器
type PersonBuilder struct {
    name    string
    age     int
    gender  string
    address string
}

// SetName 方法设置人的姓名
func (b *PersonBuilder) SetName(name string) PersonBuilderInterface {
    b.name = name
    return b
}

// SetAge 方法设置人的年龄
func (b *PersonBuilder) SetAge(age int) PersonBuilderInterface {
    b.age = age
    return b
}

// SetGender 方法设置人的性别
func (b *PersonBuilder) SetGender(gender string) PersonBuilderInterface {
    b.gender = gender
    return b
}

// SetAddress 方法设置人的地址
func (b *PersonBuilder) SetAddress(address string) PersonBuilderInterface {
    b.address = address
    return b
}

// Build 方法创建一个新的 Person 实例
func (b *PersonBuilder) Build() PersonInterface {
    return &Person{
        name:    b.name,
        age:     b.age,
        gender:  b.gender,
        address: b.address,
    }
}

```
```golang
func main() {
    // 创建一个 PersonBuilder 实例
    builder := &PersonBuilder{}

    // 设置人的信息
    person := builder.SetName("Tom").
        SetAge(18).
        SetGender("Male").
        SetAddress("123 Main St").
        Build()

    // 调用人的说话和睡觉方法
    person.Speak("Hello, World!")
    person.Sleep()
}

```

# 原型模式
原型模式（Prototype Pattern）是一种创建型设计模式，其核心思想是通过复制现有对象来创建新的对象，而不是使用传统的构造函数创建。这种方式可以避免大量的构造函数调用，从而提高代码的性能和可读性。
