/*
 * Revision History:
 *     Initial: 2018/01/19        Yang Chenglong
 */

package general_code

import (
	lex "github.com/yangchenglong11/compiler/lexical_analysis"
)

type Equ struct {
	Op     int // 四元式操作码
	Op1    int // 操作数在符号表中的入口地址
	Op2    int // 操作数在符号表中的入口地址
	Result int // 结果变量在符号表中的入口地址
}

type BasicBlock struct {
}

type GenStruct struct {
	Label    int // 语句序号
	Code     int // 语句的块内码
	Equ      Equ // 原四元式
	Out_port int // 记录该四元式是否为一个基本块的入口，是则为1，否则为0。
}

func DivBasicBlock(e []Equ) []GenStruct {
	g := make([]GenStruct, len(e))

	for i := range e {
		gen := GenStruct{Equ: e[i]}
		g = append(g, gen)
	}

	for i := range g {
		if i == 0 {
			g[i].Out_port = 1
		}

		if isJump(g[i].Equ.Op) {
			g[i].Equ.Result = 1
			if i < len(g)-1 {
				g[i+1].Out_port = 1
			}
		}
	}

	count := 1
	b := 1
	for i := range g {
		g[i].Label = count
		g[i].Code = b
		b = b + 1
		if g[i].Out_port == 1 {
			count = count + 1
			b = 1
		}
	}
	return g
}

func isJump(i int) bool {
	if (i >= lex.MachineCode["jmp"] && i <= lex.MachineCode["jnz"]) {
		return true
	}
	return false
}

func General(g []GenStruct) string {
	var l int
	if len(g) <= 0 {
		return ""
	}

	l = g[0].Label

	for i := range g {
		if g[i].Label <= l {

		}
	}
}
