package 模板模式

import "fmt"

type TestPaperImpl interface {
	testQuestion1()
	testQuestion2()
	Answer1()
	Answer2()
}

type TestPaper struct {
}

func (t *TestPaper) testQuestion1() {
	fmt.Println("问题：中国有多少个民族")
}

func (t *TestPaper) testQuestion2() {
	fmt.Println("问题：中国有多大")
}

func (t *TestPaper) Answer1() {
}

func (t *TestPaper) Answer2() {
}

// 封装具体步骤
func doPaper(paper TestPaperImpl) {
	paper.testQuestion1()
	paper.Answer1()

	paper.testQuestion2()
	paper.Answer2()
}

type student1 struct {
	TestPaper
}

func (s *student1) Answer1() {
	fmt.Println("答案：56")
}

func (s *student1) Answer2() {
	fmt.Println("答案：很大")
}

type student2 struct {
	TestPaper
}

func (s *student2) Answer1() {
	fmt.Println("答案：51")
}

func (s *student2) Answer2() {
	fmt.Println("答案：不知道")
}
