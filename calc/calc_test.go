package calc

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Calc(t *testing.T) {
	l := &CalcLex{S: "aa=1\n"}
	CalcParse(l)
	fmt.Println(l.result)

	l = &CalcLex{S: "?a\n"}
	CalcParse(l)
	fmt.Println(l.result)
	require.Equal(t, l.result.(Test).node, "variable")
	require.Equal(t, l.result.(Test).value[0], "a")

	l = &CalcLex{S: "?aa\n"}
	CalcParse(l)
	fmt.Println(l.result)
	require.Equal(t, l.result.(Test).node, "variable")
	require.Equal(t, l.result.(Test).value[0], "aa")

	l = &CalcLex{S: "?[1\n"}
	CalcParse(l)
	fmt.Println(l.result)
	require.Equal(t, l.result.(Test).node, "array2")
	require.Equal(t, l.result.(Test).value[0], []int{1})

	l = &CalcLex{S: "?[1,2\n"}
	CalcParse(l)
	fmt.Println(l.result)
	require.Equal(t, l.result.(Test).node, "array2")
	require.Equal(t, l.result.(Test).value[0], []int{1, 2})
}
