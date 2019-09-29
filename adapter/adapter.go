package main

import "fmt"

/*
 *Target- 适配的目标抽象类 对应为 MediaPlayer
 *request():void  对应为  pay(audioType string,fileName string)
 */
type MediaPlayer interface {
	Pay(audioType string,fileName string)
}

/*
 *Adaptee -适配者接口 对应 AdvancedMediaPlayer -高级媒体播放器接口
 *SpecificRequest 对应为 PlayVlc(fileName string)
 *SpecificRequest 对应为 PayMp4(fileName string)
 */
type AdvancedMediaPlayer interface {
	PlayVlc(fileName string)
	PayMp4(fileName string)
}

//适配者具体实现类 对应 VlcPlayer-Vlc播放器
type VlcPlayer struct{}

//Adaptee SpecificRequest 对应 PlayVlc
func (v *VlcPlayer)PlayVlc(fileName string){
	fmt.Println("Playing vlc file. Name: "+fileName)
}

//Adaptee SpecificRequest 对应 PayMp4
func (v *VlcPlayer)PayMp4(fileName string){
	//Do not thing
}

//适配者具体实现类 对应 Mp4Player-Mp4播放器
type Mp4Player  struct{}

//Adaptee SpecificRequest 对应 PlayVlc
func (m *Mp4Player)PlayVlc(fileName string){
	//Do not thing
}

//Adaptee SpecificRequest 对应 PayMp4
func (m *Mp4Player)PayMp4(fileName string){
	fmt.Println("Playing mp4 file. Name: "+fileName)
}


//Adapter- 适配器类 对应为 MediaAdapter
type MediaAdapter struct{
	AdvancedMediaPlayer AdvancedMediaPlayer
}

func NewMediaAdapter(audioType string)*MediaAdapter{
	if audioType == "vlc"{
		return &MediaAdapter{AdvancedMediaPlayer:&VlcPlayer{}}
	}else if audioType == "mp4"{
		return &MediaAdapter{AdvancedMediaPlayer:&Mp4Player{}}
	}
	return nil
}

//Adapter- Request实现Target抽象类接口 对应 MediaPlayer Pay(audioType string,fileName string)
func (a *MediaAdapter)Pay(audioType string,fileName string){
	//执行Adaptee的SpecificRequest 方法
	if audioType == "vlc"{
		a.AdvancedMediaPlayer.PlayVlc(fileName)
	} else if audioType == "mp4"{
		a.AdvancedMediaPlayer.PayMp4(fileName)
	}
}

//AudioPlayer -实现了MediaPlayer
type AudioPlayer struct{
}

func (a *AudioPlayer)Pay(audioType string,fileName string){
	if audioType == "mp3"{
		fmt.Println("Playing mp3 file. Name: "+fileName)
	}else if audioType == "vlc" || audioType == "mp4"{
		mediaAdapter := NewMediaAdapter(audioType)
		mediaAdapter.Pay(audioType,fileName)
	}else{
		fmt.Println("Invalid media. ",audioType ," format not supported ")
	}
}


func main(){
	audioPlayer := AudioPlayer{}

	audioPlayer.Pay("mp3", "beyond the horizon.mp3")
	audioPlayer.Pay("mp4", "alone.mp4")
	audioPlayer.Pay("vlc", "far far away.vlc")
	audioPlayer.Pay("avi", "mind me.avi")
}

