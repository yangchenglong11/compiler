/*
 * MIT License
 *
 * Copyright (c) 2017 SmartestEE Co., Ltd..
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
 *     Initial: 2017/12/25        Yang Chenglong
 */

package main

import (
	"fmt"
)

const (
	SBUFSIZE   = 256 // 定义扫描缓冲区的大小
	Identifier = "TypeIdentifier"
	Integer    = "TypeInteger"
	Real       = "TypeReal"

	Number = 0
	Letter = 1
	Other  = 2

	Space   = 0x20
	NewLine = '\n'
)

var MachineCode = map[string]int{"and": 1, "begin": 2, "bool": 3, "do": 4, "else": 5, "end": 6, "false": 7,
	"if": 8, "integer": 9, "not": 10, "or": 11, "program": 12, "real": 13, "then": 14, "true": 15, "var": 16,
	"while": 17, Identifier: 18, Integer: 19, Real: 20,
	"(": 21, ")": 22, "+": 23, "-": 24, "*": 25, "/": 26, ".": 27, ",": 28, ":": 29, ";": 30, ":=": 31, "=": 32,
	"<=": 33, "<": 34, "<>": 35, ">": 36, ">=": 37, "err": 38, "": 39}

type LexicalError struct {
	Rows     int
	Kind     int
	Describe string
}

type Token struct {
	Label int    // 单词序号
	Name  string // 单词本身
	Code  int    // 单词的机内码
	Addr  int    // 地址，单词为保留字时为-1，为标识符或常数时为大于0的数值，即在符号表中的入口地址。

}

type Symble struct {
	Number int    //序号
	Kind   int    //类型
	Name   string //名字
}

type Tokens struct {
	T []Token
}

func (t Tokens) String() {
	for i := range t.T {
		fmt.Printf("%+v\n", t.T[i])
	}
}

type Symbles struct {
	S []Symble
}

func (s Symbles) String() {
	for i := range s.S {
		fmt.Printf("%+v\n", s.S[i])
	}
}

var rows = 0

var LexicalErrors []LexicalError
