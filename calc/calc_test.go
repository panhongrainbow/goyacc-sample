package calc

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Calc(t *testing.T) {
	CalcParse(&CalcLex{S: "aa=1\n"})
	require.Equal(t, 1, regs["aa"])
}
