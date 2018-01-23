/*
 * Revision History:
 *     Initial: 2018/01/22        Wang RiYu
 */

package main

import (
	"fmt"

	lexical "github.com/yangchenglong11/compiler/lexical_analysis"
	syntax "github.com/yangchenglong11/compiler/syntax_analysis"
)

func main() {
	var (
		Ab     = syntax.Ab
		Be     = syntax.Be
		Eq     = syntax.Eq
		input  syntax.Stack
		stack  syntax.Stack
		parser syntax.Parser

		/*
		   P’ -> #P#
		   P -> program L
		   L -> S | id, L | id : K | var L; G
		   K -> integer | bool | real
		   G -> begin S end
		   S -> id := E | if B then S else S | while B do S
		   B -> id < I | id > I
		   E -> id + I | id - I
		   I -> i | id | (E) | E
		*/
		grammar = map[string]string{
			"#P#":                "P’",
			"program L":          "P",
			"S":                  "L",
			"id, L":              "L",
			"id : K":             "L",
			"var L; G":           "L",
			"brgin S end":        "G",
			"id := E":            "S",
			"if B then S else S": "S",
			"while B do S":       "S",
			"integer":            "K",
			"bool":               "K",
			"real":               "K",
			"id > I":             "B",
			"id < I":             "B",
			"i":                  "I",
			"id":                 "I",
			"(E)":                "I",
			"E":                  "I",
			"id + I":             "E",
			"id - I":             "E",
		}
		relation = [][]string{
			{Eq, Be, "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""},
			{Ab, "", Be, "", "", Be, "", "", "", "", "", "", "", Be, "", "", Be, "", "", "", "", "", "", "", ""},
			{Ab, "", "", Eq, Eq, "", Ab, "", "", "", "", Ab, Eq, "", Ab, Ab, "", Ab, Eq, Eq, Eq, Eq, "", "", Ab},
			{Ab, "", Be, "", "", Be, Ab, "", "", "", "", "", "", Be, "", "", Be, "", "", "", "", "", "", "", ""},
			{Ab, "", "", "", "", "", Ab, Be, Be, Be, "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""},
			{"", "", Be, "", "", Be, Eq, "", "", "", "", "", "", Be, "", "", Be, "", "", "", "", "", "", "", ""},
			{Ab, "", "", "", "", "", Ab, "", "", "", Be, "", "", "", "", "", "", "", "", "", "", "", "", "", ""},
			{Ab, "", "", "", "", "", Ab, "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""},
			{Ab, "", "", "", "", "", Ab, "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""},
			{Ab, "", "", "", "", "", Ab, "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""},
			{"", "", Be, "", "", "", "", "", "", "", "", Eq, "", Be, "", "", Be, "", "", "", "", "", "", "", ""},
			{Ab, "", "", "", "", "", Ab, "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""},
			{Ab, "", Be, "", "", "", Ab, "", "", "", "", Ab, "", "", "", Ab, "", "", "", "", "", "", "", "", ""},
			{"", "", Be, "", "", "", "", "", "", "", "", "", "", "", Eq, "", "", "", "", "", "", "", "", "", ""},
			{"", "", Be, "", "", "", "", "", "", "", "", "", "", Be, "", Eq, Be, "", "", "", "", "", "", "", ""},
			{Ab, "", Be, "", "", "", Ab, "", "", "", "", Ab, "", Be, "", Ab, Be, "", "", "", "", "", "", "", ""},
			{"", "", Be, "", "", "", "", "", "", "", "", "", "", "", "", "", "", Eq, "", "", "", "", "", "", ""},
			{Ab, "", Be, "", "", "", Ab, "", "", "", "", Ab, "", Be, "", Ab, Be, "", "", "", "", "", "", "", ""},
			{"", "", Be, "", "", "", "", "", "", "", "", "", "", "", Ab, "", "", Ab, "", "", "", "", Be, Be, ""},
			{"", "", Be, "", "", "", "", "", "", "", "", "", "", "", Ab, "", "", Ab, "", "", "", "", Be, Be, ""},
			{Ab, "", Be, "", "", "", Ab, "", "", "", "", Ab, "", "", Ab, Ab, "", Ab, "", "", "", "", Be, Be, Ab},
			{Ab, "", Be, "", "", "", Ab, "", "", "", "", Ab, "", "", Ab, Ab, "", Ab, "", "", "", "", Be, Be, Ab},
			{Ab, "", "", "", "", "", Ab, "", "", "", "", Ab, "", "", Ab, Ab, "", Ab, "", "", "", "", "", "", Ab},
			{"", "", Be, "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", Eq},
			{Ab, "", "", "", "", "", Ab, "", "", "", "", Ab, "", "", Ab, Ab, "", Ab, "", "", "", "", "", "", Ab},
		}
		Vt = []string{"#", "program", "id", ",", ":", "var", ";", "integer", "bool", "real", "begin", "end", ":=", "if", "then", "else", "while", "do", "<", ">", "+", "-", "i", "(", ")"}
		Vn = []string{"P’", "P", "L", "S", "G", "K", "B", "E", "I"}
	)
	parser.Init(grammar, Vt, Vn, relation)
	parser.DisplayGrammar()
	parser.DisplayRelationTable()

	tokens, symbles, err := lexical.LexicalAnalysis("test.lu")
	if err != nil {
		fmt.Println(err)
		return
	}
	if tokens != nil {
		fmt.Printf("----------------\n     %s\n----------------\n", "token表")
		tokens.String()
		for _, v := range tokens.T {
			if v.Addr < 0 {
				t := syntax.Token{Label: v.Label, Name: v.Name, Code: v.Code, Addr: v.Addr, Output: v.Name}
				input.Push(t)
			} else {
				if v.Code == lexical.MachineCode[lexical.Identifier] {
					t := syntax.Token{Label: v.Label, Name: v.Name, Code: v.Code, Addr: v.Addr, Output: "id"}
					input.Push(t)
				} else {
					t := syntax.Token{Label: v.Label, Name: v.Name, Code: v.Code, Addr: v.Addr, Output: "i"}
					input.Push(t)
				}
			}
		}
		if symbles != nil {
			fmt.Printf("\n----------------\n     %s\n----------------\n", "符号表")
			symbles.String()
		}

		fmt.Printf("\n----------------\n    %s\n----------------\n", "语法分析")
		stack.Push(syntax.Token{Output: "#"})
		input.Push(syntax.Token{Output: "#"})
		result, err := parser.Analysis(&stack, &input)
		if err != nil {
			fmt.Println(err)
			return
		} else {
			if result {
				fmt.Println("归约分析成功")

				fmt.Printf("\n----------------\n   %s\n----------------\n", "输出四元式")
				for _, v := range syntax.Equs {
					fmt.Printf("%+v\n", v)
				}
			} else {
				fmt.Println("归约分析失败")
			}
		}
	}

	if len(lexical.LexicalErrors) > 0 {
		fmt.Printf("\n----------------\n    %s\n----------------\n", "词法错误")
		for i := range lexical.LexicalErrors {
			fmt.Printf("%+v", lexical.LexicalErrors[i])
			fmt.Println()
		}
	}
}
