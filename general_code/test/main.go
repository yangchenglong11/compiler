/*
 * Revision History:
 *     Initial: 2018/01/22        Yang Chenglong
 */

package main

import (
	gen "github.com/yangchenglong11/compiler/general_code"
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

	gen.DivBasicBlock(e)
}
