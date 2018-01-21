/*
 * Revision History:
 *     Initial: 2018/01/17        Wang RiYu
 */

package main

import (
  "fmt"
  "errors"
  "strings"
  "os"
  "bufio"
)

var (
  Above = ">"
  Below = "<"
  Equal = "="
  /* 文法
    E -> E + T | T
    T -> T * F | F
    F -> P ↑ F | P
    P -> (E) | i
  */
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
  l := input.Len() + 2
  width := fmt.Sprintf("%%-%ds%%%ds%%20s\n", l, l)
  fmt.Printf(fmt.Sprintf("%%-%ds%%%ds%%16s\n", l-2, l-2), "栈", "输入流", "操作")
  fmt.Printf(width, stack.ToString(), input.ToString(), "initial")
  var k = 0
  for input.Left() != "#" || stack.ToString() != "#N" {
    newStr := input.Left()
    curStr := stack.Top()
    //fmt.Println("栈顶元素与输入元素", curStr, newStr)

    j := k
    if !isContrainAny(Vt, curStr) {
      j = k - 1
    }
    //fmt.Println("j k 下标", j, k)

    curStr = stack.Index(j)
    relation, err := getRelation(curStr, newStr)
    //fmt.Println("栈顶终结符比较", curStr, relation, newStr)
    if err != nil {
      return false, err
    } else {
      if relation == Below || relation == Equal {
        stack.Push(input.Shift())
        operation := fmt.Sprintf("%s<%s,push %s", curStr, newStr, newStr)
        fmt.Printf(width, stack.ToString(), input.ToString(), operation)
        k++
      } else if relation == Above {
        for {
          q := curStr
          if j > 0 && isContrainAny(Vt, stack.Index(j-1)) {
            j--
          } else if j > 1 && !isContrainAny(Vt, stack.Index(j-1)) {
            j -= 2
          }
          p := stack.Index(j)
          //fmt.Println("当前元素p q", p, q)
          relation, err := getRelation(p, q)
          if err != nil {
            return false, err
          } else {
            if relation == Below {
              //fmt.Println("下标p q j k", p, q, j, k)
              //fmt.Println("当前栈", stack.ToString(), j, k)
              operation := fmt.Sprintf("%s<%s>%s,replace %s", p, q, newStr, Stack(*stack)[j+1:k+1].ToString())
              stack.Replace(j+1, k+1, "N")
              fmt.Printf(width, stack.ToString(), input.ToString(), operation)
              k = j + 1
              break
            } else if relation == Equal {
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

func checkInput(valid, inputSlice []string) bool {
  for i := range inputSlice {
    if !isContrainAny(valid, inputSlice[i]) {
      return false
    }
  }
  return true
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
  fmt.Println("\n输入语句, 以#结束, 每个终结符以空格隔开, 例如 i + i * i #")
  //fmt.Scanln(&input)
  reader := bufio.NewReader(os.Stdin)
  strBytes, _, err := reader.ReadLine()
  if err != nil {
    fmt.Println(err)
  } else {
    input = string(strBytes)
    split := strings.Split(input, " ")
    if input[len(input)-1:] != "#" || !checkInput(Vt, split) {
      fmt.Println("unvalid input")
    } else {
      for i := range split {
        inputStack.Push(split[i])
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
}
