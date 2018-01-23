/*
 * Revision History:
 *     Initial: 2018/01/19        Yang Chenglong
 */

package general_code

import (
	"fmt"
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
	Label        int // 语句序号
	Code         int // 语句的块内码
	Equ          Equ // 原四元式
	Out_port     int // 记录该四元式是否为一个基本块的入口，是则为1，否则为0。
	Op1IsActive  int
	Op1IsUsed    int
	Op2IsActive  int
	Op2IsUsed    int
	ResuIsActive int
	ResuIsUsed   int
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
	if i >= lex.MachineCode["jmp"] && i <= lex.MachineCode["jnz"] {
		return true
	}
	return false
}

func HandleBlocks(g []GenStruct) string {
	var l int
	//var code string
	var temp []GenStruct
	if len(g) <= 0 {
		return ""
	}

	l = g[0].Label

	for i := range g {
		for j := i; ; j++ {
			if g[i].Label <= l {
				temp = append(temp, g[i])
			}
		}
		//code = fmt.Sprintf("%s\n%s", code, GeneralCode(temp))
	}

	return code
}

func isRegContain(r []REG, t REG) bool {
	for i := range r {
		if r[i].Name == t.Name {
			return true
		}
	}
	return false
}

func isValueContain(r []int, v int) bool {
	for i := range r {
		if r[i] == v {
			return true
		}
	}
	return false
}

func HandleVariableInfo(g []GenStruct) []GenStruct{
	for i := len(g) - 1; i > 0; i-- {
		g[i].ResuIsActive = GetActive(g[i].Equ.Result)
		g[i].ResuIsUsed = GetUsed(g[i].Equ.Result)
		T.S[GetIndex(g[i].Equ.Result)].IsUsed = 0
		T.S[GetIndex(g[i].Equ.Result)].IsActive = 0
		g[i].Op1IsActive = GetActive(g[i].Equ.Op1)
		g[i].Op1IsUsed = GetUsed(g[i].Equ.Op1)
		g[i].Op2IsActive = GetActive(g[i].Equ.Op2)
		g[i].Op2IsUsed = GetUsed(g[i].Equ.Op2)
		T.S[GetIndex(g[i].Equ.Op1)].IsUsed = i
		T.S[GetIndex(g[i].Equ.Op1)].IsActive = 1
		T.S[GetIndex(g[i].Equ.Op2)].IsUsed = i
		T.S[GetIndex(g[i].Equ.Op2)].IsActive = 1
	}

	return g
}

func GETREG(g GenStruct) REG {
	var re REG
	for i := range R {
		s := AVALUE[g.Equ.Op1]
		if (isRegContain(s, R[i]) && isValueContain(R[i].Value, g.Equ.Op1)) || (g.Equ.Op1 == g.Equ.Result) || (g.Op1IsUsed==0&&g.Op1IsActive==0) {
			re = R[i]
		} else if len(R[i].Value) == 0 {
			re = R[i]
		} else {
			re = R[i]
			for j := range R[i].Value {
				if !isRegContain(AVALUE[R[i].Value[j]],M) {
					str := fmt.Sprintf("MOV M %s",R[i].Name)
					code = fmt.Sprintf("%s\n%s",code,str)
					R[i].Value = DeleteValue(R[i].Value,R[i].Value[j])
					AVALUE[R[i].Value[j]] = []REG{M}
				}
			}
		}
	}

	return re
}

func GeneralCode(g []GenStruct) {
	g = HandleVariableInfo(g)
	for i := range g {

		if g[i].Equ.Result > 0 && g[i].Equ.Op2 > 0 && g[i].Equ.Op1 > 0 {
			code = fmt.Sprintf("MOV AX, %d\n%s AX, %s\nMOV %s, AX", GetName(g[i].Equ.Op1), OpCode[g[i].Equ.Op], GetName(g[i].Equ.Op2), GetName(g[i].Equ.Result))
		}

		if g[i].Equ.Result > 0 && g[i].Equ.Op2 > 0 && g[i].Equ.Op1 > 0 {
			code = fmt.Sprintf("MOV AX, %d\n%s AX, %s\nMOV %s, AX", GetName(g[i].Equ.Op1), OpCode[g[i].Equ.Op], GetName(g[i].Equ.Op2), GetName(g[i].Equ.Result))
		}

	}
}

func DeleteValue(s []int, d int)[]int {
	for i := range s{
		if s[i] == d {
			for j := i ;j<len(s)-1;j++ {
				s[j] = s[j+1]
			}
		}
	}
	return s
}

func GetName(in int) string {
	for i := range T.S {
		if T.S[i].Symble.Number == in {
			return T.S[i].Symble.Name
		}
	}
	return ""
}

func GetIndex(in int) int {
	for i := range T.S {
		if T.S[i].Symble.Number == in {
			return i
		}
	}
	return 0
}

func GetActive(in int) int {
	for i := range T.S {
		if T.S[i].Symble.Number == in {
			return T.S[i].IsActive
		}
	}
	return 0
}

func GetUsed(in int) int {
	for i := range T.S {
		if T.S[i].Symble.Number == in {
			return T.S[i].IsUsed
		}
	}
	return 0
}
