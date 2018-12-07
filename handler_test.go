package handler

import (
	"testing"
)
const checkMark  = "\u2713"
const ballotX = "\u2717"

func TestAdd(t *testing.T){
	t.Log("测试Add方法")
	var vars = []struct{
		A int
		B int
		R int
	}{
		{3,5,8},
		{4,6,10},
	}
	for k,v := range vars {
		sum := Add(v.A,v.B)
		if sum == v.R{
			t.Logf("测试用例%d验证通过%s",k+1,checkMark)
		}else{
			t.Errorf("测试用例%d验证不通过%s",k+1,ballotX)
		}
	}
}
