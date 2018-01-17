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

package main

import (
  "fmt"
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
  ids = [7]string{"+", "*", "↑", "i", "(", ")", "#"}
)

func reverse(array Stack) Stack {
  var result Stack
  newArray := array
  for i := newArray.Len(); i > 0 ; i-- {
    if r, err := newArray.Pop(); err == nil {
      result.Push(r)
    }
  }

  return result
}

func main() {
  fmt.Printf("%18s\n", "算符优先关系表")
  for i := 0; i < len(ids); i++ {
    fmt.Printf("%5s", ids[i])
  }
  fmt.Println()
  for i := 0; i < len(ids); i++ {
    fmt.Printf("%1s %3s", ids[i], relation[i][0])
    for j := 1; j < len(relation[i]); j++ {
      fmt.Printf("%5s", relation[i][j])
    }
    fmt.Println()
  }

  /* 输入语句 */
  var (
    input string
    inputStack Stack
    stack Stack
  )
  fmt.Println("\n输入语句, 以#结束:")
  fmt.Scanln(&input)
  for i := len(input) - 1; i >= 0; i-- {
    inputStack.Push(string(input[i]))
  }
  stack.Push("#")
  fmt.Printf("%10s%10s\n", "栈", "输入栈")
  fmt.Printf("%10s%10s\n", stack, reverse(inputStack))
}
