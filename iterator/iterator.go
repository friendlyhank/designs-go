package main

import (
	"container/list"
	"fmt"
)

/*
 *实际上container容器已经实现了堆、链、环,而且slice,map也已经非常成熟好用
 *
 */

//Iterator -抽象迭代器类
type Iterator interface{
	Next() *list.Element
	hasNext()bool
}

//ConcreteIterator -具体迭代器类
type ConcreteIterator struct{
	cursor  int
	list *list.List
	obj *list.Element
}

func NewConcreteIterator(list *list.List)*ConcreteIterator{
	return &ConcreteIterator{list:list}
}

func (ci *ConcreteIterator)hasNext()bool{
	if ci.cursor == ci.list.Len(){
		return false
	}
	return true
}

func (ci *ConcreteIterator)Next()*list.Element{
	if ci.hasNext(){
		if ci.cursor == 0{
			ci.obj = ci.list.Front()
		}else {
			ci.obj = ci.obj.Next()
		}
		ci.cursor ++
	}
	return ci.obj
}

//Aggregate - 抽象容器类
type Aggregate interface {
	Add(obj interface{})
	Remove(obj interface{})
}

//ConcreteAggregate 容器具体类
type ConcreteAggregate struct{
	list *list.List
}

func (ag *ConcreteAggregate)iterator()Iterator {
	return NewConcreteIterator(ag.list)
}

func (ag *ConcreteAggregate)Add(value interface{}){
	ag.list.PushBack(value)
}

func (ag *ConcreteAggregate)Remove(value interface{}){
	//for i :=0;i<3;i++ do you know ?
	for obj := ag.list.Front();obj != nil;obj = obj.Next(){
		if obj.Value.(string) == value{
			ag.list.Remove(obj)
			break
		}
	}
}

func main(){
	ag := &ConcreteAggregate{list:list.New()}
	ag.Add("小明")
	ag.Add("小红")
	ag.Add("小刚")

	it := ag.iterator()
	for it.hasNext(){
		str := it.Next().Value.(string)
		fmt.Println(str)
	}
}
