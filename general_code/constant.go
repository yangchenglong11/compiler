/*
 * Revision History:
 *     Initial: 2018/01/22        Yang Chenglong
 */

package general_code

var (
	Jump = map[int]bool{J:true,Jl:true,Jg:true,Jne:true}
)

const (
	J = iota+1
	Jl
	Jg
	Jne

)