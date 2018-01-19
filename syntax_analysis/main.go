/*
 * Revision History:
 *     Initial: 2018/01/17        Wang RiYu
 */

/* L语言文法定义
程序定义:
<main> -> program <ID> <Body>
<Body> -> <VarIntro> <Begin>

变量定义:
<VarIntro> -> var <VarDef> | ε
<VarDef> -> <IDTable>: <Type> | <IDTable>: <Type>; <VarDef>
<IDTable> -> <ID>, <IDTable> | <ID>

语句定义：
<Begin> -> begin <Sentence> end
<Sentence> -> <Execute>; <Sentence> | <Execute>
<Execute> -> <SimpleSt> | <StructSt>
<SimpleSt> -> <Assignment>
<Assignment> -> <Variable>:=<Expression>
<Variable> -> <ID>
<StructSt> -> <Begin> | <IfSt> | <WhileSt>
<IfSt> -> if <BoolExpress> then <Execute> | if <BoolExpress> then <Execute> else <Execute>
<WhileSt> -> while <BoolExpress> do <Execute>

表达式定义:
<Expression> -> <ArithmeticExp> | <BoolExpress>
<ArithmeticExp> -> <ArithmeticExp> + <Item> | <ArithmeticExp> - <Item> | <Item>
<Item> -> <Item> * <Factor> | <Item> / <Factor> | <Factor>
<Factor> -> <ArithmeticNum> | (<ArithmeticExp>)
<ArithmeticNum> -> <ID> | <Integer> | <Real>
<BoolExpress> -> <BoolExpress> or <BoolItem> | <BoolItem>
<BoolItem> -> <BoolItem> and <BoolFactor> | <BoolFactor>
<BoolFactor> -> not <BoolFactor> | <BoolValue>
<BoolValue> -> <BoolConstant> | <ID> | (<BoolExpress>) | <RelationExpress>
<RalationExpress> -> <ID> <Rop> <ID>
<Rop> -> < | <= | = | > | >= | <>

类型定义:
<Type> -> integer | bool | real

单词定义:
<ID> -> <Letter> | <ID> <Letter> | <ID> <Number>
<Integer> -> <Number> | <Integer> <Number>
<Real> -> <Integer> | <Real> <Number>
<BoolValue> -> true | false

字符定义:
<Letter> -> A│B│C│D│E│F│G│H│I│J│K│L│M│N│O│P│Q│R│S│T│U│V│W│X│Y│Z│a│b│c│d│e│f│g│h│i│j│k│l│m│n│o│p│q│r│s│t│u│v│w│x│y│z
<Number> -> 0│1│2│3│4│5│6│7│8│9
*/

/*
E -> E + T | T
T -> T * F | F
F -> P ↑ F | P
P -> (E) | i
*/
package main

import (
  "fmt"
  "errors"
)

var (
  Above    = ">"
  Below    = "<"
  Equal    = "="
  relation = [7][7]string{// 算符优先关系，-1代表小于，1代表大于，0代表等于
    {Above, Below, Below, Below, Below, Above, Above},
    {Above, Above, Below, Below, Below, Above, Above},
    {Above, Above, Below, Below, Below, Above, Above},
    {Above, Above, Above, "", "", Above, Above},
    {Below, Below, Below, Below, Below, Equal, ""},
    {Above, Above, Above, "", "", Above, Above},
    {Below, Below, Below, Below, Below, "", Equal},
  }
  Vt = []string{"+", "*", "↑", "i", "(", ")", "#"} // 终结符集
  Vn = []string{"E", "T", "F", "P"}                // 非终结符
)

func isContrainAny(slice []string, string string) bool {
  for i := range slice {
    if slice[i] == string {
      return true
    }
  }
  return false
}

func getRelation(a, b string) (string, error) { // 获取 a 与 b 的关系
  var (
    indexA = -1
    indexB = -1
  )
  for i, v := range Vt {
    if a == v {
      indexA = i
    }
    if b == v {
      indexB = i
    }
  }
  if indexA < 0 || indexB < 0 {
    return "", errors.New("no this letter")
  }

  return relation[indexA][indexB], nil
}

func analysis(stack, input *Stack) (bool, error) {
  l := len(*input) + 2
  width := fmt.Sprintf("%%-%ds%%%ds%%16s\n", l, l)
  fmt.Printf(fmt.Sprintf("%%-%ds%%%ds%%12s\n", l-2, l-2), "栈", "输入流", "操作")
  fmt.Printf(width, *stack, input.Reverse(), "initial")
  var k = 0
  for input.Top() != "#" {
    newStr := input.Top()
    curStr := stack.Top()
    fmt.Println("a1", newStr, curStr)

    j := k
    if !isContrainAny(Vt, curStr) {
      j = k - 1
    }

    for {
      if relation, err := getRelation(stack.Index(j), newStr); err == nil && relation == Above {
        q := stack.Index(j)
        if isContrainAny(Vt, stack.Index(j-1)) {
          j--
        } else {
          j -= 2
        }
        for {
          if relation, err := getRelation(stack.Index(j), q); err == nil && relation == Below {
            q = stack.Index(j)
            if isContrainAny(Vt, stack.Index(j-1)) {
              j--
            } else {
              j -= 2
            }
            operation := fmt.Sprintf("%s<%s>%s,replace %s", stack.Index(j), stack.Index(k), newStr, string(*stack)[j+1:k+1])
            stack.Replace(j+1, k+1, "N")
            fmt.Printf(width, *stack, input.Reverse(), operation)
            k++
          } else if err != nil {
            return false, err
          }
        }
      } else if err != nil {
        return false, err
      }
    }

    relation, err := getRelation(stack.Index(j), newStr)
    if err != nil {
      return false, err
    } else {
      if relation == Below || relation == Equal {
        stack.Push(input.Pop())
        operation := fmt.Sprintf("%s<%s,push %s", curStr, newStr, newStr)
        fmt.Printf(width, *stack, input.Reverse(), operation)
      }
    }
  }

  return true, nil
}

func main() {
  fmt.Printf("%18s\n", "算符优先关系表")
  for i := 0; i < len(Vt); i++ {
    fmt.Printf("%5s", Vt[i])
  }
  fmt.Println()
  for i := 0; i < len(Vt); i++ {
    fmt.Printf("%1s %3s", Vt[i], relation[i][0])
    for j := 1; j < len(relation[i]); j++ {
      fmt.Printf("%5s", relation[i][j])
    }
    fmt.Println()
  }

  /* 输入语句 */
  var (
    input      string
    inputStack Stack
    stack      Stack
  )
  fmt.Println("\n输入语句, 以#结束:")
  fmt.Scanln(&input)
  if input[len(input)-1:] != "#" {
    fmt.Println("unvalid input")
  } else {
    for i := len(input) - 1; i >= 0; i-- {
      inputStack.Push(string(input[i]))
    }
    stack.Push("#")

    result, err := analysis(&stack, &inputStack)
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
