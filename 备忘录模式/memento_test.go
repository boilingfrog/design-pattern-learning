package 备忘录模式

import "testing"

func TestVideo_Watch(t *testing.T) {
	video := &Video{
		time: 10,
	}

	t.Log(video.Status())
	progress := video.Save()

	video.Watch(30)
	t.Log(video.Status())

	video.Load(progress)
	t.Log(video.Status())
}
