/*
 * Revision History:
 *     Initial: 2018/01/22        Yang Chenglong
 */

package main

import (
	"fmt"

	lex "github.com/yangchenglong11/compiler/lexical_analysis"
)

func main() {
	t, s, err := lex.LexicalAnalysis("../test.lu")
	if err != nil {
		fmt.Println(err)
		return
	}
	if t != nil {
		t.String()
	}
	fmt.Println()
	if s != nil {
		s.String()
	}

	for i := range lex.LexicalErrors {
		fmt.Printf("%+v", lex.LexicalErrors[i])
		fmt.Println()
	}
}
