package main

import "fmt"

//Shape- 对应Product abstract 抽象产品角色
type Shape interface {
	Draw()//use()
}

//Rectangle -长方形 对应 ConcreteProductA 具体产品角色
type Rectangle struct{}

//use()
func (r *Rectangle)Draw(){
	fmt.Println("Inside Rectangle::draw() method.")
}

//Square- 正方形 对应 ConcreteProductB 具体产品角色
type Square struct{}

//use()
func (s *Square)Draw(){
	fmt.Println("Inside Square::draw() method.")
}

//Circle- 圆形 对应 ConcreteProductC 具体产品角色
type Circle  struct{}

//use()
func (c *Circle)Draw(){
	fmt.Println("Inside Circle::draw() method.")
}

//ShapeFactory 对应 Factory 工厂角色
type ShapeFactory struct{
}

func (sf *ShapeFactory)GetShape(shapeType string)Shape{
	if shapeType == "CIRCLE"{
		return &Circle{}
	}else if shapeType == "RECTANGLE"{
		return &Rectangle{}
	}else if shapeType == "SQUARE"{
		return &Square{}
	}
	return nil
}

func main(){
	shapeFactory  := &ShapeFactory{}

	//获取 Circle 的对象，并调用它的 draw 方法
	shap1 := shapeFactory.GetShape("CIRCLE")
	shap1.Draw()

	////获取 Rectangle 的对象，并调用它的 draw 方法
	shap2 := shapeFactory.GetShape("RECTANGLE")
	shap2.Draw()

	//获取 Square 的对象，并调用它的 draw 方法
	shap3 := shapeFactory.GetShape("SQUARE")
	shap3.Draw()
}




