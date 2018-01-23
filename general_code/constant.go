/*
 * Revision History:
 *     Initial: 2018/01/22        Yang Chenglong
 */

package general_code

import (
	lex "github.com/yangchenglong11/compiler/lexical_analysis"
)

var OpCode = map[int]string{
	lex.MachineCode["+"]: "ADD",
	lex.MachineCode["-"]: "SUB",
	lex.MachineCode["*"]: "MUL",
	lex.MachineCode["/"]: "DIV",
}

type REG struct {
	Name  string
	Value []int
}

var T Symbles

type Symble struct {
	Symble    lex.Symble
	IsActive int
	IsUsed   int
}

type Symbles struct {
	S []Symble
}

var M = REG{
	Name:  "Memory",
	Value: make([]int, 0),
}
var BX = REG{
	Name:  "BX",
	Value: make([]int, 0),
}

var DX = REG{
	Name:  "BX",
	Value: make([]int, 0),
}

var R = []REG{BX, DX, M}

var AVALUE = make(map[int][]REG)


var code string
// var RVALUE  = make(map[REG][]int)
