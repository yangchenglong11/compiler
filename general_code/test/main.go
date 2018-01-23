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
	e := []gen.Equ{
		{
			Op:     1,
			Op1:    3,
			Op2:    0,
			Result: 2,
		},
		{
			Op:     4,
			Op1:    4,
			Op2:    0,
			Result: 1,
		},
	}

	s := lex.Symbles{
		S: []lex.Symble{
			{Number: 1, Kind: 18, Name: "a"},
			{Number: 2, Kind: 18, Name: "b"},
			{Number: 3, Kind: 19, Name: "0"},
			{Number: 4, Kind: 19, Name: "1"},
		},
	}
	gen.InitSymble(s)
	g := gen.DivBasicBlock(e)
	fmt.Printf("%+v", g)
	st := gen.HandleBlocks(g)
	fmt.Println(st)
}
