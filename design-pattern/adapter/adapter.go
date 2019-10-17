package main

import "fmt"

// 适配器模式： 作为两个不兼容的接口之间的桥梁。这种类型的设计模式属于结构型模式，它结合了两个独立接口的功能。
//这种模式涉及到一个单一的类，该类负责加入独立的或不兼容的接口功能

type Player interface {
	Play(format, file string)
}

type AudioPlayer struct {
}

func NewPlayer() Player {
	return &AudioPlayer{}
}

type PlayerAdapter struct {
}

func (pa *PlayerAdapter) PlayMp4(format, file string) {
	fmt.Println("PlayerAdapter", format, file)
}

func NewPlayerAdapter() *PlayerAdapter {
	return &PlayerAdapter{}
}

func (*AudioPlayer) Play(format, file string) {
	switch format {
	case "mp3":
		fmt.Println(format, file)
	case "mp4":
		NewPlayerAdapter().PlayMp4(format, file)
	}
}

func main() {
	player := NewPlayer()
	player.Play("mp3", "booty music.mp3")
	player.Play("mp4", "apologize.mp4")
}
