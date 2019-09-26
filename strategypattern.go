package main

import "fmt"

/*
 *抽象策略类
 */
type Strategy interface{
	doOperation(num1 int,num2 int)int
}

/*
 *具体策略类-相加算法
 */
type OperationAdd struct{
}

func (operationAdd *OperationAdd)doOperation(num1 int,num2 int)int{
	return num1 + num2
}

/*
 *具体策略类-相减算法
 */
type OperationSubstract struct{
}

func (operationSubstract *OperationSubstract)doOperation(num1 int,num2 int)int{
	return num1 - num2
}

/*
 *具体策略类-相乘算法
 */
type OperationMultiply struct{
}

func (operationMultiply *OperationMultiply)doOperation(num1 int,num2 int)int{
	return num1 * num2
}

/*
 *环境类
 */
type Context struct{
	strategy Strategy
}

func (context *Context)executeStrategy(num1 int,num2 int)int{
	return context.strategy.doOperation(num1,num2)
}

func main(){
	context := &Context{&OperationAdd{}}
	fmt.Println(context.executeStrategy(10,5))

	context = &Context{&OperationSubstract{}}
	fmt.Println(context.executeStrategy(10,5))

	context = &Context{&OperationMultiply{}}
	fmt.Println(context.executeStrategy(10,5))
}


