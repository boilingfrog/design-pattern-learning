package 迭代器模式

import (
	"testing"
)

func TestIterator(t *testing.T) {
	names := names{
		"小明", "小豆", "小龙",
	}
	nameRepository := names.NewIterator()
	for nameRepository.HasNext() {
		t.Log(nameRepository.Next())
	}
}
