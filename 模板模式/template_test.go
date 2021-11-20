package 模板模式

import (
	"fmt"
	"testing"
)

func TestTestPaper(t *testing.T) {
	st1 := &student1{}
	doPaper(st1)

	fmt.Println("++++++++++++++")
	st2 := &student2{}
	doPaper(st2)
}

func TestTestPaper_callback(t *testing.T) {
	st1 := &student3{}
	doPaperCallback(st1)
}
