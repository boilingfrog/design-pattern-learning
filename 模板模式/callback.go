package 模板模式

import "fmt"

type TestPaper1 struct {
}
type Callback func(num int, answer string)

func (t *TestPaper1) testQuestion1() {
	fmt.Println("问题：中国有多少个民族")
}

func (t *TestPaper1) testQuestion2() {
	fmt.Println("问题：中国有多大")
}

func (t *TestPaper1) subCallback(num int, an string, callback Callback) {
	switch num {
	case 1:
		t.testQuestion1()
	case 2:
		t.testQuestion1()
	default:
		panic("不支持的题目")
	}

	callback(num, an)
}

// 回调函数的具体实现
func (t *TestPaper1) answer(num int, an string) {
	fmt.Println(num, an)
}

type student3 struct {
	*TestPaper1
}

func (s *student3) answer() {
	s.subCallback(1, "56", s.TestPaper1.answer)
	s.subCallback(2, "不知道", s.TestPaper1.answer)
}
