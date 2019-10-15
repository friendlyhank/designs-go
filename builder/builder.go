package main

import "fmt"

//Builder 是生成器接口
type Builder interface {
	Part1()
	Part2()
	Part3()
}

type Director struct {
	builder Builder
}

// NewDirector ...
func NewDirector(builder Builder) *Director {
	return &Director{
		builder: builder,
	}
}

//Construct Product
func (d *Director) Construct() {
	d.builder.Part1()
	d.builder.Part2()
	d.builder.Part3()
}

type Builder1 struct {
	result string
}

func (b *Builder1) Part1() {
	b.result += "1"
}

func (b *Builder1) Part2() {
	b.result += "2"
}

func (b *Builder1) Part3() {
	b.result += "3"
}

func (b *Builder1) GetResult() string {
	return b.result
}

type Builder2 struct {
	result int
}

func (b *Builder2) Part1() {
	b.result += 1
}

func (b *Builder2) Part2() {
	b.result += 2
}

func (b *Builder2) Part3() {
	b.result += 3
}

func (b *Builder2) GetResult() int {
	return b.result
}

func main(){
	builder1 := &Builder1{}
	director1 := NewDirector(builder1)
	director1.Construct()
	res1 := builder1.GetResult()
	if res1 != "123" {
		fmt.Sprintf("Builder1 fail expect 123 acture %s",res1)
	}

	builder2 := &Builder2{}
	director2 := NewDirector(builder2)
	director2.Construct()
	res2 := builder2.GetResult()
	if res2 != 6 {
		fmt.Sprintf("Builder2 fail expect 6 acture %d", res2)
	}

}
