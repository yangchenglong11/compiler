/*
 * Revision History:
 *     Initial: 2018/01/17        Wang RiYu
 */

package main

import (
  "fmt"
  "errors"
  "strings"
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

    for {
      curStr = stack.Index(j)
      if relation, err := getRelation(curStr, newStr); err == nil && relation == Above {
        q := curStr
        //fmt.Println("当前元素q", q)
        if j > 0 && isContrainAny(Vt, stack.Index(j-1)) {
          j--
        } else if j > 1 && !isContrainAny(Vt, stack.Index(j-1)) {
          j -= 2
        }
        for {
          p := stack.Index(j)
          //fmt.Println("当前元素pq", p, q)
          if relation, err := getRelation(p, q); err == nil {
            if relation == Below {
              //fmt.Println("栈内终结符比较", p, relation, q)
              if index := strings.IndexAny(stack.ToString(), q) - strings.IndexAny(stack.ToString(), p); index == 1 || index == 0 {
                operation := fmt.Sprintf("%s<%s>%s,replace %s", p, q, newStr, string(Stack(*stack)[j+1:]))
                stack.Replace(j+1, k+1, "N")
                fmt.Printf(width, stack.ToString(), input.ToString(), operation)
                k = j + 1
                break
              }
              q = p
              if j > 0 && isContrainAny(Vt, stack.Index(j-1)) {
                if j-1 > 0 {
                  j--
                }
              } else if j > 1 && !isContrainAny(Vt, stack.Index(j-1)) {
                if j-2 > 0 {
                  j -= 2
                }
              }
              //fmt.Println("下标q j k", q, j, k)
              //fmt.Println("当前栈", stack.ToString(), j, k)
              operation := fmt.Sprintf("%s<%s>%s,replace %s", p, curStr, newStr, string(*stack)[j+1:])
              stack.Replace(j+1, k+1, "N")
              fmt.Printf(width, stack.ToString(), input.ToString(), operation)
              k = j + 1
            } else if relation == Equal {
              if stack.ToString() == "#N" {
                if input.Left() == "#" {
                  return true, nil
                } else {
                  break
                }
              }
              if strings.IndexAny(stack.ToString(), q) == strings.IndexAny(stack.ToString(), p) {
                break
              }
              q = p
              if j > 0 && isContrainAny(Vt, stack.Index(j-1)) {
                j--
              } else if j > 1 && !isContrainAny(Vt, stack.Index(j-1)) {
                j -= 2
              }
            } else {
              break
            }
          } else {
            return false, err
          }
        }
      } else if err != nil {
        return false, err
      } else if relation != Above {
        break
      }
    }

    if input.Left() == "#" && stack.ToString() == "#N" {
      return true, nil
    }

    relation, err := getRelation(stack.Index(j), newStr)
    //fmt.Println("栈顶终结符比较", stack.Index(j), relation, newStr)
    if err != nil {
      return false, err
    } else {
      if relation == Below || relation == Equal {
        stack.Push(input.Shift())
        operation := fmt.Sprintf("%s<%s,push %s", stack.Index(j), newStr, newStr)
        fmt.Printf(width, stack.ToString(), input.ToString(), operation)
        k++
      } else {
        return false, err
      }
    }
  }

  fmt.Println("end", stack.ToString(), input.ToString())
  return true, nil
}

func checkInput(valid []string, input string) bool {
  split := strings.Split(input, "")
  for i := range split {
    if !isContrainAny(valid, split[i]) {
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
  fmt.Println("\n输入语句, 以#结束, 例如 i+i*i# :")
  fmt.Scanln(&input)
  if input[len(input)-1:] != "#" || !checkInput(Vt, input) {
    fmt.Println("unvalid input")
  } else {
    inputStack.Push(input)
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
