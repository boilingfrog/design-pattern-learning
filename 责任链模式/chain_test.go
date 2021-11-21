package 责任链模式

import (
	"testing"

	"github.com/tj/assert"
)

func TestApproveChain(t *testing.T) {
	c1 := NewHeadTeacherChain()
	c2 := NewDepManagerChain()
	c3 := NewDeanTeacherChain()

	c1.SetApprover(c2)
	c2.SetApprover(c3)

	var c Teacher = c1
	assert.Equal(t, true, c.HandleApproveRequest("小明", 3))
	assert.Equal(t, true, c.HandleApproveRequest("小红", 2))
	assert.Equal(t, false, c.HandleApproveRequest("小龙", 30))
}
