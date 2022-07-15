package calc

import (
	"fmt"
	"testing"
)

func Test_Calc(t *testing.T) {
	l := &CalcLex{S: "aa=1\n"}
	CalcParse(l)
	fmt.Println(l.result)
}
