/*
 * Revision History:
 *     Initial: 2017/12/25        Yang Chenglong
 */

package lexical_analysis

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

	ErrIdentifier = 1
	DesIdentifier = "A number at the beginning of a identifier."

	ErrManyPoint = 2
	DesManyPoint = "There is more than one point in the real."

	ErrReal = 3
	DesReal = "The decimal part of the real number appears in letters"
)

var MachineCode = map[string]int{"and": 1, "begin": 2, "bool": 3, "do": 4, "else": 5, "end": 6, "false": 7,
	"if": 8, "integer": 9, "not": 10, "or": 11, "program": 12, "real": 13, "then": 14, "true": 15, "var": 16,
	"while": 17, Identifier: 18, Integer: 19, Real: 20,
	"(": 21, ")": 22, "+": 23, "-": 24, "*": 25, "/": 26, ".": 27, ",": 28, ":": 29, ";": 30, ":=": 31, "=": 32,
	"<=": 33, "<": 34, "<>": 35, ">": 36, ">=": 37, "err": 38, "": 39,"jmp":40,"jl":41,"jg":42,"jne":43}

type LexicalError struct {
	Rows     int
	Kind     int
	Describe string
}

type Token struct {
	Label int    // 单词序号
	Name  string // 单词本身
	Code  int    // 单词的机内码
	Addr  int    // 地址，单词为保留字时为-1，为运算符时为-2，为标识符或常数时为大于0的数值，即在符号表中的入口地址。

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

var rows = 1

var LexicalErrors []LexicalError
