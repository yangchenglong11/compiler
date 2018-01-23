/*
 * Revision History:
 *     Initial: 2018/01/22        Wang RiYu
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	syntax "github.com/yangchenglong11/compiler/syntax_analysis"
)

func main() {
	var (
		Ab       = syntax.Ab
		Be       = syntax.Be
		Eq       = syntax.Eq
		inputStr string
		input    syntax.Stack
		stack    syntax.Stack
		parser   syntax.Parser

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

	fmt.Println("\n输入语句, 以#结束, 每个终结符以空格隔开, 例如 program while id > 0 do if id > i then id := id - i else id := id + i #:")
	reader := bufio.NewReader(os.Stdin)
	strBytes, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println(err)
	} else {
		inputStr = string(strBytes)
		split := strings.Split(inputStr, " ")
		if split[len(split)-1] != "#" {
			fmt.Println("unvalid input")
		} else {
			for i := range split {
				input.Push(syntax.Token{Output: split[i]})
			}
			stack.Push(syntax.Token{Output: "#"})

			result, err := parser.Analysis(&stack, &input)
			if err != nil {
				fmt.Println(err)
			} else {
				if result {
					fmt.Println("归约分析成功")
				} else {
					fmt.Println("归约分析失败")
				}
			}
		}
	}
}
