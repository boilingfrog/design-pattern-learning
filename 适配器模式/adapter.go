package 适配器模式

import "fmt"

// 基础的播放功能
type MediaPlayer interface {
	play(audioType string, fileName string)
}

// 不同的播放器平台
type AdvancedMediaPlayer interface {
	playVlc(fileName string)
	playMp4(fileName string)
}

// VlcPlayers
type VlcPlayers struct {
}

func (v *VlcPlayers) playVlc(fileName string) {
	fmt.Println("正在播放" + fileName)
}

func (v *VlcPlayers) playMp4(fileName string) {
	fmt.Println("格式不支持")
}

// Mp4Player
type Mp4Player struct {
}

func (m *Mp4Player) playVlc(fileName string) {
	fmt.Println("格式不支持")
}

func (m *Mp4Player) playMp4(fileName string) {
	fmt.Println("正在播放" + fileName)
}

// 适配器
type MediaAdapter struct {
	MusicPlayer AdvancedMediaPlayer
}

func NewMediaAdapter(audioType string) *MediaAdapter {
	var mediaAdapter MediaAdapter
	switch audioType {
	case "vlc":
		mediaAdapter.MusicPlayer = &VlcPlayers{}
	case "mp4":
		mediaAdapter.MusicPlayer = &Mp4Player{}
	default:
		panic("不支持的类型")
	}
	return &mediaAdapter
}
func (m *MediaAdapter) play(audioType string, fileName string) {
	switch audioType {
	case "vlc":
		m.MusicPlayer.playVlc(fileName)
	case "mp4":
		m.MusicPlayer.playMp4(fileName)

	}
}

// AudioPlayer 音频播放器类
type AudioPlayer struct {
	mediaAdapter *MediaAdapter
}

// Play 播放音频
func (auPlayer *AudioPlayer) Play(audioType, fileName string) {
	if audioType == "mp3" {
		fmt.Println("正在播放" + fileName)
		return
	}
	auPlayer.mediaAdapter = NewMediaAdapter(audioType)
	auPlayer.mediaAdapter.play(audioType, fileName)
}
