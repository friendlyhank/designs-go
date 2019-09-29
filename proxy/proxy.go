package main

import "fmt"

//Image 对应 Subject 抽象主题角色
type Image interface{
	//Display() 对应request()
	Display()
}

//RealImage interface Image 对应 RealSubject 真实主题角色
type RealImage struct{
	fileName string
}

func NewRealImage(fileName string)*RealImage{
	realImage := &RealImage{}
	realImage.loadFromDisk(fileName)
	return realImage
}

//Display- 对应 RealSubject request()
func (r *RealImage)Display(){
	fmt.Println("Displaying "+r.fileName)
}

func (r *RealImage)loadFromDisk(fileName string){
	fmt.Println("Loading "+fileName)
}

//ProxyImage interface Image 对应Proxy
type ProxyImage struct{
	realImage *RealImage
	fileName string
}

func NewProxyImage(fileName string)*ProxyImage{
	return &ProxyImage{}
}

func (p *ProxyImage)Display(){
	//preRequest()
	if p.realImage == nil{
		p.realImage = NewRealImage(p.fileName)
	}
	//Display
	p.realImage.Display()
}

func main(){
	image := NewProxyImage("test_10mb.jpg")
	image.Display()
	fmt.Println("")
	image.Display()
}

