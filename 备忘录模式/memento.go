package 备忘录模式

import "fmt"

type Memento interface{}

type Video struct {
	time int
}

type videoMemento struct {
	time int
}

func (g *Video) Watch(time int) {
	g.time += time
}

func (g *Video) Save() Memento {
	return &videoMemento{
		time: g.time,
	}
}

func (g *Video) Load(m Memento) {
	gm := m.(*videoMemento)
	g.time = gm.time
}

func (g *Video) Status() string {
	return fmt.Sprintf("video watch time:%d", g.time)
}
