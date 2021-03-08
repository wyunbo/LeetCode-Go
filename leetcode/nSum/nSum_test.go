package leetcode

import (
	"fmt"
	"testing"
)

type question18 struct {
	para18
	ans18
}

// para 是参数
// one 代表第一个参数
type para18 struct {
	a []int
	n int
	t int
}

// ans 是答案
// one 代表第一个答案
type ans18 struct {
	one [][]int
}

func Test_Problem18(t *testing.T) {

	qs := []question18{

		{
			para18{[]int{1, 2, 3, 4, 5}, 4, 12},
			ans18{[][]int{{1, 2, 4, 5}}},
		},
	}

	fmt.Printf("------------------------n Sum------------------------\n")

	for _, q := range qs {
		_, p := q.ans18, q.para18
		fmt.Printf("【input】:%v       【output】:%v\n", p, nSum(p.a, p.n, p.t))
	}
	fmt.Printf("\n\n\n")
}
