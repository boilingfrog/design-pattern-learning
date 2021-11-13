package 适配器模式

import "testing"

func TestPlayer(t *testing.T) {
	ad := AudioPlayer{}
	ad.Play("mp4", "荷塘月色")
	ad.Play("vlc", "小苹果")
	ad.Play("mp3", "天空之城")
}
