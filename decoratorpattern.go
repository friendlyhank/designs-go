package main

import "fmt"

/**
 *设计模式-结构型
 *装饰器模式
 */

// step1: 编写基础功能，刚开始不需要定义接口
type Base struct {
}
func (b *Base) Call() string {
	return "base is called"
}

// step2: 将上面的方法声明为接口类型，基础功能中的 Call() 调用自动满足下面的接口
type DecoratorI interface {
	Call() string
}

// step3: 编写新增功能，结构中保存接口类型的参数
type Decorator struct {
	derorator DecoratorI
}

func (d *Decorator) Call() string {
	return "decorator: " + d.derorator.Call()
}

func main() {
	base := &Base{}
	fmt.Println(base.Call())

	decorator := Decorator{base}
	fmt.Println(decorator.Call())
}