package main

import "fmt"

// 中介者模式：是用来降低多个对象和类之间的通信复杂性。
// 这种模式提供了一个中介类，该类通常处理不同类之间的通信，并支持松耦合，使代码易于维护。中介者模式属于行为型模式

type Cpu struct {
}

func (c *Cpu) DecodeMedia() {
	fmt.Println("音频文件解码...")
}

type HardDisk struct {
}

func (hd *HardDisk) Read() {
	fmt.Println("正在读取硬盘...")
}
func (hd *HardDisk) Write() {
	fmt.Println("正在写入硬盘...")
}

type SoundCard struct {
}

func (sc *SoundCard) Out(song string) {
	fmt.Println(song, "播放中...")
}

type Ethernet struct {
}

func (e *Ethernet) Download(song string) {
	fmt.Println(song, "下载中...")

}

type MainBoard struct {
	cpu       Cpu
	soundCard SoundCard
	hd        HardDisk
	eth       Ethernet
}

func (mb *MainBoard) PlayMusic(song string) {
	mb.eth.Download(song)

	mb.hd.Write()

	mb.hd.Read()

	mb.cpu.DecodeMedia()

	mb.soundCard.Out(song)
}

func main() {
	main := MainBoard{}
	main.PlayMusic("123.mp3")
}
