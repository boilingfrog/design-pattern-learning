package 模板模式

import "fmt"

type testPaperCallback struct {
}

func (t *testPaperCallback) testQuestion1() {
	fmt.Println("问题1：中国有多少个民族")
}

func (t *testPaperCallback) testQuestion2() {
	fmt.Println("问题2：中国有多大")
}

func (t *testPaperCallback) subCallback(callback CallbackImpl) {
	t.testQuestion1()
	t.testQuestion2()
	callback.callback()
}

type CallbackImpl interface {
	callback()
}

type student3 struct {
	*testPaperCallback
}

func (s *student3) callback() {
	fmt.Println("答案1：56")
	fmt.Println("答案2：测试")
}

func doPaperCallback(student *student3) {
	student.subCallback(&student3{})
}
