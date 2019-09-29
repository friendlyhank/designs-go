package main

import(
	"fmt"
)

type CPU struct {
}

func (c *CPU)start(){
	fmt.Println("启动CPU...")
}

type Memory struct{
}

func (m *Memory)start(){
	fmt.Println("启动内存管理...")
}

type Disk struct{
}

func (d *Disk)start(){
	fmt.Println("启动硬盘...")
}

type Computer struct{}

//Façade
func (c *Computer)startComputer(){
	//启动CPU
	cpu := &CPU{}
	cpu.start()

	//内存
	memory:= &Memory{}
	memory.start()

	//硬盘
	disk:= &Disk{}
	disk.start()
}

//Client
func main(){
	computer := &Computer{}
	computer.startComputer()
}

