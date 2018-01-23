/*
 * Revision History:
 *     Initial: 2018/01/17        Wang RiYu
 */

package syntax_analysis

import (
  "fmt"
  "errors"
  "strings"

  lexical "github.com/yangchenglong11/compiler/lexical_analysis"
)

const (
  Ab = ">" // Above
  Be = "<" // Below
  Eq = "=" // Equal
)

type Equ struct {
  op     int //四元式操作码
  op1    int // 操作数在符号表中的入口地址
  op2    int // 操作数在符号表中的入口地址
  result int // 结果变量在符号表中的入口地址
}

var Equs []Equ

type Parser struct {
  grammar       map[string]string // 文法
  Vt            []string          // 终结符集
  Vn            []string          // 非终结符集
  relationTable [][]string        // 优先关系表
}

func (parser *Parser) Init(grammar map[string]string, Vt, Vn []string, relationTable [][]string) {
  parser.Vt = Vt
  parser.Vn = Vn
  parser.grammar = grammar
  parser.relationTable = relationTable
}

func (parser Parser) vtContrainsAny(str string) bool { // Vt 是否包含 str
  for i := range parser.Vt {
    if parser.Vt[i] == str {
      return true
    }
  }
  return false
}

func (parser Parser) getRelation(a, b string) (string, error) { // 获取 a 与 b 的关系
  var (
    indexA = -1
    indexB = -1
  )
  for i, v := range parser.Vt {
    if a == v {
      indexA = i
    }
    if b == v {
      indexB = i
    }
  }
  //fmt.Println("index", indexA, indexB)
  if indexA < 0 {
    return "", errors.New("no this letter")
  }
  if indexB < 0 {
    return Eq, nil
  }

  return parser.relationTable[indexA][indexB], nil
}

func (parser Parser) generateCode(input Stack) {
  split := strings.Split(input.ToString(), " ")

  if len(split) > 0 {
    if split[0] == "while" {
      newEqu := Equ{op: lexical.MachineCode["jmp"], op1: lexical.MachineCode[""], op2: lexical.MachineCode[""], result: 0}
      Equs = append(Equs, newEqu)
      //fmt.Printf("--- while %s %+v\n", split, newEqu)
    }

    if split[0] == "if" {
      //fmt.Printf("--- if %s %+v\n", split, Equs)
    }
  }

  if len(split) > 1 {
    if split[1] == "<" || split[1] == ">" {
      newEqu := Equ{op: lexical.MachineCode[fmt.Sprintf("j%s", split[1])], op1: input[0].Addr, op2: input[2].Addr, result: len(Equs)}
      Equs = append(Equs, newEqu)
      Equs = append(Equs, Equ{op: lexical.MachineCode["jmp"], op1: lexical.MachineCode[""], op2: lexical.MachineCode[""], result: 0})
      //fmt.Printf("--- relop %s %+v\n", split, newEqu)
    }

    if split[1] == ":=" {
      newEqu := Equ{op: lexical.MachineCode[split[1]], op1: len(Equs)-1, op2: lexical.MachineCode[""], result: input[0].Addr}
      Equs = append(Equs, newEqu)
      //fmt.Printf("--- := %s %+v\n", split, newEqu)
    }

    if split[1] == "-" || split[1] == "+" {
      newEqu := Equ{op: lexical.MachineCode[split[1]], op1: input[0].Addr, op2: input[2].Addr, result: len(Equs)-1}
      Equs = append(Equs, newEqu)
      //fmt.Printf("--- operator %s %+v\n", split, newEqu)
    }
  }
}

func (parser Parser) Analysis(stack, input *Stack) (bool, error) { // 算符优先分析过程
  inputLen := len(input.ToString()) + 10
  stackLen := len(stack.ToString()) + 10
  totalLen := inputLen + stackLen
  fmt.Printf(fmt.Sprintf("%%-%ds%%%ds%%%ds\n", stackLen, inputLen-4, 42), "栈", "输入流", "操作")
  fmt.Printf(fmt.Sprintf("%%-%ds%%%ds%%%ds\n", stackLen, inputLen, 45), stack.ToString(), input.ToString(), "initial")
  var k = 0
  for input.Left() != "#" || stack.ToString() != "# N" {
    newStr := input.Left()
    curStr := stack.Top()
    //fmt.Println("栈顶元素与输入元素", curStr, newStr)

    j := k
    if !parser.vtContrainsAny(curStr) {
      j = k - 1
    }
    //fmt.Println("j k 下标", j, k)

    curStr = stack.Index(j)
    relation, err := parser.getRelation(curStr, newStr)
    //fmt.Println("栈顶终结符比较", curStr, relation, newStr)
    if err != nil {
      return false, err
    } else {
      if relation == Be || relation == Eq {
        stack.Push(input.Shift())
        operation := fmt.Sprintf("%s < %s, push %s", curStr, newStr, newStr)
        stackLen := len(stack.ToString()) + 10
        fmt.Printf(fmt.Sprintf("%%-%ds%%%ds%%%ds\n", stackLen, totalLen-stackLen, 45), stack.ToString(), input.ToString(), operation)
        k++
      } else if relation == Ab {
        for {
          q := curStr
          if j > 0 && parser.vtContrainsAny(stack.Index(j-1)) {
            j--
          } else if j > 1 && !parser.vtContrainsAny(stack.Index(j - 1)) {
            j -= 2
          }
          p := stack.Index(j)
          //fmt.Println("当前元素p q", p, q)
          relation, err := parser.getRelation(p, q)
          if err != nil {
            return false, err
          } else {
            if relation == Be {
              //fmt.Println("下标p q j k", p, q, j, k)
              //fmt.Println("当前栈", stack.ToString(), j, k)
              replace := Stack(*stack)[j+1:k+1]

              token := TokenN
              if len(replace) == 1 && (replace[0].Output == "id" || replace[0].Output == "i") {
                token = Token{Output: "N", Name: replace[0].Name, Addr: replace[0].Addr, Label: replace[0].Label, Code: replace[0].Code}
              }
              operation := fmt.Sprintf("%s < %s > %s, replace %s", p, q, newStr, replace.ToString())
              stack.Replace(j+1, k+1, token)
              stackLen := len(stack.ToString()) + 10
              fmt.Printf(fmt.Sprintf("%%-%ds%%%ds%%%ds\n", stackLen, totalLen-stackLen, 45), stack.ToString(), input.ToString(), operation)
              parser.generateCode(replace) // 生成四元式
              k = j + 1
              break
            } else if relation == Eq {
              curStr = p
            } else {
              return false, err
            }
          }
        }
      } else {
        return false, err
      }
    }
  }

  //fmt.Println("end", stack.ToString(), input.ToString())
  return true, nil
}

func (parser Parser) DisplayGrammar() {
  fmt.Printf("----------\n   %s\n----------\n", "文法")

  result := make(map[string]string)
  for k, v := range parser.grammar {
    value, ok := result[v]
    if ok {
      result[v] = fmt.Sprintf("%s | %s", value, k)
    } else {
      result[v] = k
    }
  }

  for k, v := range result {
    fmt.Printf("%s -> %s\n", k, v)
  }
  fmt.Println()
}

func (parser Parser) DisplayRelationTable() {
  length := len(parser.Vt)
  fmt.Printf("--------------------\n   %s\n--------------------\n", "算符优先关系表")
  fmt.Print("       ")
  for i := 0; i < length; i++ {
    fmt.Printf("%8s", parser.Vt[i])
  }
  fmt.Println()
  for i := 0; i < length; i++ {
    fmt.Printf("%-7s %7s", parser.Vt[i], parser.relationTable[i][0])
    for j := 1; j < len(parser.relationTable[i]); j++ {
      fmt.Printf("%8s", parser.relationTable[i][j])
    }
    fmt.Println()
  }
  fmt.Println()
}
