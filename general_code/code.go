/*
 * MIT License
 *
 * Copyright (c) 2018 SmartestEE Co., Ltd..
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

/*
 * Revision History:
 *     Initial: 2018/01/19        Yang Chenglong
 */

package general_code

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
	if _, ok := Jump[i]; ok == true {
		return true
	}
	return false
}

func main() {
	e := []Equ{
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
		{
			Op:     2,
			Op1:    1,
			Op2:    1,
			Result: 3,
		},
		{
			Op:     1,
			Op1:    1,
			Op2:    1,
			Result: 4,
		},
		{
			Op:     1,
			Op1:    1,
			Op2:    1,
			Result: 5,
		},
		{
			Op:     1,
			Op1:    1,
			Op2:    1,
			Result: 6,
		},
		{
			Op:     1,
			Op1:    1,
			Op2:    1,
			Result: 7,
		},
		{
			Op:     1,
			Op1:    1,
			Op2:    1,
			Result: 8,
		},
		{
			Op:     1,
			Op1:    1,
			Op2:    1,
			Result: 9,
		},
		{
			Op:     1,
			Op1:    1,
			Op2:    1,
			Result: 10,
		},
		{
			Op:     1,
			Op1:    1,
			Op2:    1,
			Result: 11,
		},
		{
			Op:     1,
			Op1:    1,
			Op2:    1,
			Result: 12,
		},
	}
}
