package main

import "fmt"

/*
 *TODO 实现最好用异步转同步的过程
 */

//Command -命令抽象类
type Command interface {
	execute()
}

//FlipUpCommand 具体命令实现类-开灯
type FlipUpCommand struct{
	theLight *Light
}

//execute- 具体命令实现类的执行方法
func (uc *FlipUpCommand)execute(){
	uc.theLight.turnOn()
}

//FlipDownCommand -具体命令实现类-关灯
type FlipDownCommand  struct{
	theLight *Light
}

func (dc *FlipDownCommand)execute(){
	dc.theLight.turnOff()
}

//Invoker - 调用者
type Switch struct{
	command Command
}

func (s *Switch)setCommand(command Command){
	s.command = command
}

func (s *Switch)action(){
	s.command.execute()
}

//Receiver- 接收者
type Light struct{
}

func (l *Light)turnOn(){
	fmt.Println("The light is on")
}

func (l *Light)turnOff(){
	fmt.Println("The light is off")
}

func main(){
	//new Receiver and FlipDownCommand
	receiver := &Light{}
	flipDowncommand := &FlipDownCommand{theLight:receiver}

	//客户端通过调用者来执行命令
	invoker := &Switch{}
	//FlipDown
	invoker.setCommand(flipDowncommand)
	invoker.action()

	flipUpCommand := &FlipUpCommand{theLight:receiver}
	//FlipDown
	invoker.setCommand(flipUpCommand)
	invoker.action()

}
