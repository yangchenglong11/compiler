/*
 * Revision History:
 *     Initial: 2018/01/22        Yang Chenglong
 */

package main

import (
	"fmt"
	gen "github.com/yangchenglong11/compiler/general_code"
	lex "github.com/yangchenglong11/compiler/lexical_analysis"
)

func main() {
	/*
	(+,a,b,c)
	(-,b,c,d)
	(*,c,d,a)
	(/,b,4,a)
	(jne,d,c,a)
	(jl,c,b,d)
	*/
	e := []gen.Equ{
		{
			Op:     lex.MachineCode["+"],
			Op1:    1,
			Op2:    2,
			Result: 3,
		},
		{
			Op:     lex.MachineCode["-"],
			Op1:    2,
			Op2:    3,
			Result: 4,
		},
		{
			Op:     lex.MachineCode["*"],
			Op1:    3,
			Op2:    4,
			Result: 1,
		},
		{
			Op:     lex.MachineCode["/"],
			Op1:    2,
			Op2:    5,
			Result: 1,
		},
		{
			Op:     lex.MachineCode["jne"],
			Op1:    4,
			Op2:    3,
			Result: 1,
		},
		{
			Op:     lex.MachineCode["jl"],
			Op1:    3,
			Op2:    2,
			Result: 4,
		},
	}

	s := lex.Symbles{
		S: []lex.Symble{
			{Number: 1, Kind: 18, Name: "a"},
			{Number: 2, Kind: 18, Name: "b"},
			{Number: 3, Kind: 18, Name: "c"},
			{Number: 4, Kind: 18, Name: "d"},
			{Number: 5, Kind: 19, Name: "4"},
			{Number: 6, Kind: 19, Name: "5"},
		},
	}
	gen.InitSymble(s)
	g := gen.DivBasicBlock(e)
	st := gen.HandleBlocks(g)
	fmt.Println(st)
}
