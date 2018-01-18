/*
 * MIT License
 *
 * Copyright (c) 2017 Yang Chenglong
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

/*
 * Revision History:
 *     Initial: 2017/12/25        Yang Chenglong
 */

package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"os"
)

func what(a byte) int {
	if (a >= 48) && (a <= 57) {
		//０-９数字
		return Number
	} else if ((a >= 'a') && (a <= 'z')) || ((a >= 'A') && (a <= 'Z')) {
		//a-z的字母
		return Letter
	} else {
		//其他的标点符号
		return Other
	}
}

func GetContent(path string) (*bytes.Buffer, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	c, _ := ioutil.ReadAll(f)
	return bytes.NewBuffer(c), nil
}

func LexicalAnalysis(path string) (*[]Token, *[]Symble, error) {
	var (
		token  []Token
		symble []Symble
	)
	c, err := GetContent(path)
	if err != nil {
		return nil, nil, err
	}
	for {
		b, err := CheckByte(c)
		if err != nil {
			if err == io.EOF {
				return &token, &symble, nil
			}
		}
		switch what(b) {
		case Number:
			t, num, err := handleNumber(b, c)
			if err != nil {
				return nil, nil, err
			}
			s := Symble{
				Number: len(symble) + 1,
				Kind:   t,
				Name:   fmt.Sprintf("%s", num),
			}
			symble = append(symble, s)
			to := Token{
				Label: len(token),
				Name:  fmt.Sprintf("%s", num),
				Code:  t,
				Addr:  len(symble),
			}
			token = append(token, to)
			continue
		case Letter:
			t,s,err := handleLetter(b, c)
			if err != nil {
				return nil, nil, err
			}
			if s == "" {
				continue
			}
			sy := Symble{
				Number: len(symble) + 1,
				Kind:   t,
				Name:   s,
			}
			symble = append(symble, sy)
			to := Token{
				Label: len(token),
				Name:  s,
				Code:  t,
				Addr:  len(symble),
			}
			token = append(token, to)
			continue
		case Other:
			o,s,err := handlerOther(b, c)
			if err != nil {
				return nil, nil, err
			}
			sy := Symble{
				Number: len(symble) + 1,
				Kind:   o,
				Name:   s,
			}
			symble = append(symble, sy)
			to := Token{
				Label: len(token),
				Name:  s,
				Code:  o,
				Addr:  len(symble),
			}
			token = append(token, to)
			continue
		}
	}

	return nil, nil, err
}

func CheckByte(f *bytes.Buffer) (byte, error) {
	cSubTemp, err := f.ReadByte()
	if err != nil {
		if err == io.EOF {
			return Space, err
		}
		return 0x0, err
	}

	t := what(cSubTemp)
	if t <= Letter {
		return cSubTemp, nil
	}

	if cSubTemp == NewLine {
		rows = rows + 1
	}

	//判断是否为空格,合并多个空格为一个
	if isSpace(cSubTemp) {
		for isSpace(cSubTemp) {
			cSubTemp, err = f.ReadByte()
			if err != nil {
				if err == io.EOF {
					return Space, err
				}
				return 0x0, err
			}
		}
		f.UnreadByte()
		return Space, nil
	}

	//判断注释
	if cSubTemp == '/' {
		//临时变量，用于检查下一个是否为'/'
		var cTemp byte

		//再读入一个字符
		cTemp, _ = f.ReadByte()

		//若为注释一直读入直到换行符，否则退回刚才读入的字符
		if cTemp == '/' {
			for cTemp != '\n' {
				cTemp, err = f.ReadByte()
			}
			//遇到注释，在注释结尾返回空格
			return NewLine, nil
		} else {
			//退回刚才读入的字符
			f.UnreadByte()
		}
	}
	return cSubTemp, nil
}

func handleNumber(b byte, buf *bytes.Buffer) (int, float64, error) {
	var (
		temp = b
		num  = float64(int(temp) - 48)
		err  error
	)

	// continue read
	temp, err = buf.ReadByte()
	if err != nil {
		if err == io.EOF {
			goto finish
		}
		return 0, 0, err
	}

	// not a number
	if what(temp) != 0 {
		// is a real
		if temp == '.' {
			// continue read
			temp, err = buf.ReadByte()
			if err != nil {
				if err == io.EOF {
					goto finish
				}
				return 0, 0, err
			}
			// calculate num
			for i := 0; what(temp) == 0; i++ {
				num = num + float64(int(temp)-48)/(math.Pow10(i))
				temp, err = buf.ReadByte()
				if err != nil {
					if err == io.EOF {
						goto finish
					}
					return 0, 0, err
				}
			}
			goto real

		}
	}

	// is a Integer
	for what(temp) == 0 {
		num = num*10 + float64(int(temp)-48)
		temp, err = buf.ReadByte()
		if err != nil {
			if err == io.EOF {
				goto finish
			}
			return 0, 0, err
		}
	}

real:
	fmt.Println("(", MachineCode[Real], ",", num, ")")
	return MachineCode[Real], num, nil
finish:

	fmt.Println("(", MachineCode[Integer], ",", num, ")")
	return MachineCode[Integer], num, nil

}

func handleLetter(b byte, buf *bytes.Buffer) (int,string,error) {
	var (
		t []byte
		temp []byte
		err error
	)
	temp = append(temp,b)
	switch  {
	case buf.Len()>=6:
		t = make([]byte,6)
		buf.Read(t)
		temp = append(temp,t...)
		switch string(temp) {
		case "integer":
		case "program":
			return MachineCode[string(temp)],string(temp),nil
		default:
			goto other
		}
	case buf.Len()>=4:
		t = make([]byte,4)
		buf.Read(t)
		temp = append(temp,t...)
		switch string(temp) {
		case "begin":
		case "false":
		case "while":
			return MachineCode[string(temp)],string(temp),nil
		default:
			goto other
		}
	case buf.Len()>=3:
		t = make([]byte,3)
		buf.Read(t)
		temp = append(temp,t...)
		switch string(temp) {
		case "bool":
		case "else":
		case "real":
		case "then":
		case "true":
			return MachineCode[string(temp)],string(temp),nil
		default:
			goto other
		}
	case buf.Len()>=2:
		t = make([]byte,2)
		buf.Read(t)
		temp = append(temp,t...)
		switch string(temp) {
		case "not":
		case "and":
		case "end":
		case "var":
			return MachineCode[string(temp)],string(temp),nil
		default:
			goto other
		}
	case buf.Len()>=1:
		b,_ = buf.ReadByte()
		temp = append(temp,b)
		switch string(temp) {
		case "if":
		case "or":
		case "do":
			return MachineCode[string(temp)],string(temp),nil
		default:
			goto other
		}
	case buf.Len()>=0:
		return MachineCode[Identifier], string(temp), nil

	default:
	}
other:
	b, err = buf.ReadByte()
	if err != nil {
		if err == io.EOF {
			return MachineCode[Identifier],string(temp),nil
		}
		return 0, "", err
	}
	for what(b) == Letter {
		temp = append(temp,b)
		b, err = buf.ReadByte()
		if err != nil {
			if err == io.EOF {
				return MachineCode[Identifier],string(temp),nil
			}
			return 0, "", err
		}
	}
	return MachineCode[Identifier],string(temp),nil
}


func handlerOther(b byte, buf *bytes.Buffer) (int,string,error){
	var (
		temp []byte
	)
	temp = append(temp,b)
	switch  {
	case buf.Len()>=1:
		b,_ = buf.ReadByte()
		temp = append(temp,b)
		switch string(temp) {
		case ">=":
		case "<=":
		case "<>":
		case ":=":
			return MachineCode[string(temp)],string(temp),nil
		default:
			return MachineCode["err"],"#",nil
		}
	case buf.Len()>=0:
		switch string(temp) {
		case "+":
		case "-":
		case "*":
		case "/":
		case ".":
		case ",":
		case "(":
		case ")":
		case ":":
		case ";":
		case "=":
		case ">":
		case "<":
			return MachineCode[string(temp)],string(temp),nil
		case "":
			return MachineCode[""],"",nil
		default:
			return MachineCode["err"],"#",nil
		}

	default:

	}
	return MachineCode["err"],"#",nil
}

func isSpace(b byte) bool {
	return b == Space
}

func main() {
	t,s,err := LexicalAnalysis("./test.lu")
	if err != nil {
		fmt.Println(err)
		return
	}
	if t!= nil {
		fmt.Println(*t)
	}
	if s != nil {
		fmt.Println(*s)
	}

}