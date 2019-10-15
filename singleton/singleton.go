package main

import (
	"fmt"
	"sync"
)

//Singleton 是单例模式类
type Singleton struct{}

var (
	singleton *Singleton
	once sync.Once
)

//GetInstance 用于获取单例模式对象
func GetInstance() *Singleton{
	once.Do(func(){
		singleton = &Singleton{}
	})
	return singleton
}

func main(){
	ins1 := GetInstance()
	ins2 := GetInstance()
	if ins1 != ins2{
		fmt.Println("instance is not equal")
	}
}
