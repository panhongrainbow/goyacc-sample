package calc

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Calc(t *testing.T) {
	/*l2 := &CalcLex2{input: bytes.NewBuffer([]byte("aa=1\n"))}
	CalcParse2(l2)
	fmt.Println(l2.result)*/

	// >>>>> >>>>> test converting letters to variable
	/*
	 with "\n"
	 rule: '?' variable '\n'
	 input: "?aa@\n"
	 output: "syntax error"
	 explanation: aa@ is not a variable
	*/
	/*
	 without "\n"
	 rule: '?' variable
	 input: "?aa@"
	 output: "aa"
	 explanation: aa is not a variable and ignore the @ because it doesn't know the end of the line
	*/

	// one letter
	l := &CalcLex{S: "ASK a ?"}
	CalcParse(l)
	require.Equal(t, l.result.(Test).node, "variable")
	require.Equal(t, l.result.(Test).value[0], "a")
	require.Equal(t, l.err, nil)

	// two LETTERs
	l = &CalcLex{S: "ASK ab ?"}
	CalcParse(l)
	require.Equal(t, l.result.(Test).node, "variable")
	require.Equal(t, l.result.(Test).value[0], "ab")
	require.Equal(t, l.err, nil)

	// three LETTERs
	l = &CalcLex{S: "ASK abc ?"}
	CalcParse(l)
	require.Equal(t, l.result.(Test).node, "variable")
	require.Equal(t, l.result.(Test).value[0], "abc")
	require.Equal(t, l.err, nil)

	// contain not LETTERs
	l = &CalcLex{S: "ASK ab@ ?"}
	CalcParse(l)
	require.Equal(t, l.result, nil)
	require.Equal(t, l.err.Error(), "syntax error")

	// >>>>> >>>>> test converting DIGITs to a number

	// one DIGIT (decimal)
	l = &CalcLex{S: "ASK 1 ?"}
	CalcParse(l)
	require.Equal(t, l.result.(Test).node, "number")
	require.Equal(t, l.result.(Test).value[0], 1)
	require.Equal(t, l.err, nil)

	// two DIGITs (decimal)
	l = &CalcLex{S: "ASK 12 ?"}
	CalcParse(l)
	require.Equal(t, l.result.(Test).node, "number")
	require.Equal(t, l.result.(Test).value[0], 12)
	require.Equal(t, l.err, nil)

	// three DIGITs (decimal)
	l = &CalcLex{S: "ASK 123 ?"}
	CalcParse(l)
	require.Equal(t, l.result.(Test).node, "number")
	require.Equal(t, l.result.(Test).value[0], 123)
	require.Equal(t, l.err, nil)

	// octal
	l = &CalcLex{S: "ASK 012 ?"}
	CalcParse(l)
	require.Equal(t, l.result.(Test).node, "number")
	require.Equal(t, l.result.(Test).value[0], 10)
	require.Equal(t, l.err, nil)

	// >>>>> >>>>> test converting numbers to a sub array (array2)

	l = &CalcLex{S: "ASK [1 ?"}
	CalcParse(l)
	require.Equal(t, l.result.(Test).node, "array")
	require.Equal(t, l.result.(Test).value[0], []int{1})
	require.Equal(t, l.err, nil)

	l = &CalcLex{S: "ASK [1,2 ?"}
	CalcParse(l)
	require.Equal(t, l.result.(Test).node, "array")
	require.Equal(t, l.result.(Test).value[0], []int{1, 2})
	require.Equal(t, l.err, nil)

	// >>>>> >>>>> test converting numbers to an array

	l = &CalcLex{S: "ASK [] ?"}
	CalcParse(l)
	/*require.Equal(t, l.result.(Test).node, "array")
	require.Equal(t, l.result.(Test).value[0], []int{})*/
	require.Equal(t, l.err, nil)

	return
}
