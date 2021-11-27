package 备忘录模式

import "fmt"

type Memento interface{}

type Video struct {
	time int
}

type videoMemento struct {
	time int
}

func (v *Video) Watch(time int) {
	v.time += time
}

func (v *Video) Save() Memento {
	return &videoMemento{
		time: v.time,
	}
}

func (v *Video) Load(m Memento) {
	gm := m.(*videoMemento)
	v.time = gm.time
}

func (v *Video) Status() string {
	return fmt.Sprintf("video watch time:%d", v.time)
}
