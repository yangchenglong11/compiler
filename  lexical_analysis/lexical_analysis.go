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

func GetContent(path string) (*Buffer, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	c, _ := ioutil.ReadAll(f)
	return NewBuffer(c), nil
}

func LexicalAnalysis(path string) (*Tokens, *Symbles, error) {
	var (
		token  Tokens
		symble Symbles
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
			if t == MachineCode[Integer] {
				s := Symble{
					Number: len(symble.S) + 1,
					Kind:   t,
					Name:   fmt.Sprintf("%d", int(num)),
				}
				symble.S = append(symble.S, s)
				to := Token{
					Label: len(token.T),
					Name:  fmt.Sprintf("%d", int(num)),
					Code:  t,
					Addr:  len(symble.S),
				}
				token.T = append(token.T, to)
				continue
			}
			s := Symble{
				Number: len(symble.S) + 1,
				Kind:   t,
				Name:   fmt.Sprintf("%f", num),
			}
			symble.S = append(symble.S, s)
			to := Token{
				Label: len(token.T),
				Name:  fmt.Sprintf("%f", num),
				Code:  t,
				Addr:  len(symble.S),
			}
			token.T = append(token.T, to)
			continue
		case Letter:
			t, s, err := handleLetter(b, c)
			if err != nil {
				return nil, nil, err
			}
			sy := Symble{
				Number: len(symble.S) + 1,
				Kind:   t,
				Name:   s,
			}
			symble.S = append(symble.S, sy)
			to := Token{
				Label: len(token.T),
				Name:  s,
				Code:  t,
				Addr:  len(symble.S),
			}
			token.T = append(token.T, to)
			continue
		case Other:
			o, s, err := handlerOther(b, c)
			if err != nil {
				return nil, nil, err
			}
			if o == 39 {
				continue
			}
			sy := Symble{
				Number: len(symble.S) + 1,
				Kind:   o,
				Name:   s,
			}
			symble.S = append(symble.S, sy)
			to := Token{
				Label: len(token.T),
				Name:  s,
				Code:  o,
				Addr:  len(symble.S),
			}
			token.T = append(token.T, to)
			continue
		}
	}

	return nil, nil, err
}

func CheckByte(f *Buffer) (byte, error) {
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

	if cSubTemp == '\n' {
		rows = rows + 1
		return '$', nil
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
				if err != nil {
					if err == io.EOF {
						return Space, err
					}
					return 0x0, err
				}
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

func handleNumber(b byte, buf *Buffer) (int, float64, error) {
	var (
		temp byte
		num  = float64(int(b) - 48)
		err  error
	)

	// continue read
	temp, err = CheckByte(buf)
	if err != nil {
		if err == io.EOF {
			goto finish
		}
		return 0, 0, err
	}

	if temp == Space || temp == '$' {
		goto finish
	}

	// is a Integer
	for what(temp) == 0 {
		num = num*10 + float64(int(temp)-48)
		temp, err = CheckByte(buf)
		if err != nil {
			if err == io.EOF {
				goto finish
			}
			return 0, 0, err
		}
	}

	// solve the bug as 3g34
	if what(temp) == Letter {
		t, err := CheckByte(buf)
		if err != nil {
			if err == io.EOF {
				goto finish
			}
			return 0, 0, err
		}
		if what(t) == Number {
			for i := 0; what(t) == 0; i++ {
				t, err = CheckByte(buf)
				if err != nil {
					if err == io.EOF {
						goto finish
					}
					return 0, 0, err
				}
			}

			num_err := LexicalError{
				Rows:     rows,
				Kind:     ErrIdentifier,
				Describe: DesIdentifier,
			}
			LexicalErrors = append(LexicalErrors, num_err)
		} else {
			buf.UnreadByte()
		}

		goto finish
	}

	// is a real
	if temp == '.' {
		// continue read
		temp, err = CheckByte(buf)
		if err != nil {
			if err == io.EOF {
				goto finish
			}
			return 0, 0, err
		}

		// calculate num
		for i := 0; what(temp) == 0; i++ {
			num = num + float64(int(temp)-48)/(math.Pow10(i+1))
			temp, err = CheckByte(buf)
			if err != nil {
				if err == io.EOF {
					goto finish
				}
				return 0, 0, err
			}
		}

		// solve the bug as 23.32.32
		if temp == '.' {
			num_err := LexicalError{
				Rows:     rows,
				Kind:     ErrManyPoint,
				Describe: DesManyPoint,
			}
			LexicalErrors = append(LexicalErrors, num_err)
			temp, err = CheckByte(buf)
			for i := 0; what(temp) == 0; i++ {
				temp, err = CheckByte(buf)
				if err != nil {
					if err == io.EOF {
						goto finish
					}
					return 0, 0, err
				}
			}
		}

		if what(temp) == Letter {
			temp, err = CheckByte(buf)
			if err != nil {
				if err == io.EOF {
					goto finish
				}
				return 0, 0, err
			}
			if what(temp) == 0 {
				for i := 0; what(temp) == 0; i++ {
					temp, err = CheckByte(buf)
					if err != nil {
						if err == io.EOF {
							goto finish
						}
						return 0, 0, err
					}
				}

				nu_err := LexicalError{
					Rows:     rows,
					Kind:     ErrReal,
					Describe: DesReal,
				}
				LexicalErrors = append(LexicalErrors, nu_err)
			} else {
				buf.UnreadByte()
			}
		}

		buf.UnreadByte()
		goto real
	}

real:
	return MachineCode[Real], num, nil
finish:
	return MachineCode[Integer], num, nil
}

func handleLetter(b byte, buf *Buffer) (int, string, error) {
	var (
		temp []byte
		err  error
	)
	temp = append(temp, b)

	if buf.Len() >= 1 {
		b, _ = CheckByte(buf)
		if what(b) != Letter {
			buf.UnreadByte()
			goto finish
		}
		temp = append(temp, b)
		switch string(temp) {
		case "do", "if", "or":
			return MachineCode[string(temp)], string(temp), nil
		default:
			buf.UnreadByte()
			temp = temp[:len(temp)-1]
		}
	}

	if buf.Len() >= 2 {
		for i := 0; i < 2; i++ {
			b, _ := CheckByte(buf)
			if what(b) != Letter {
				buf.UnreadByte()
				goto finish
			}
			temp = append(temp, b)
		}
		switch string(temp) {
		case "var", "not", "and", "end":
			return MachineCode[string(temp)], string(temp), nil
		default:
			for i := 0; i < 2; i++ {
				buf.UnreadByte()
			}
			temp = temp[:len(temp)-2]
		}
	}

	if buf.Len() >= 3 {
		for i := 0; i < 3; i++ {
			b, _ := CheckByte(buf)
			if what(b) != Letter {
				buf.UnreadByte()
				goto finish
			}
			temp = append(temp, b)
		}
		switch string(temp) {
		case "true", "bool", "else", "real", "then":
			return MachineCode[string(temp)], string(temp), nil
		default:
			temp = temp[:len(temp)-3]
			for i := 0; i < 3; i++ {
				buf.UnreadByte()
			}
		}
	}

	if buf.Len() >= 4 {
		for i := 0; i < 4; i++ {
			b, _ := CheckByte(buf)
			if what(b) != Letter {
				buf.UnreadByte()
				goto finish
			}
			temp = append(temp, b)
		}
		switch string(temp) {
		case "false", "begin", "while":
			return MachineCode[string(temp)], string(temp), nil
		default:
			for i := 0; i < 4; i++ {
				buf.UnreadByte()
			}
			temp = temp[:len(temp)-4]
		}
	}

	if buf.Len() >= 6 {
		for i := 0; i < 6; i++ {
			b, _ := CheckByte(buf)
			if what(b) != Letter {
				buf.UnreadByte()
				goto finish
			}
			temp = append(temp, b)
		}
		switch string(temp) {
		case "program", "integer":
			return MachineCode[string(temp)], string(temp), nil
		default:
			for i := 0; i < 6; i++ {
				buf.UnreadByte()
			}
			temp = temp[:len(temp)-6]
		}
	}

	b, err = CheckByte(buf)
	if err != nil {
		if err == io.EOF {
			return MachineCode[Identifier], string(temp), nil
		}
		return 0, "", err
	}
	if b == '$' {
		return MachineCode[Identifier], string(temp), nil
	}
	temp = append(temp, b)

	for what(b) == Letter {
		temp = append(temp, b)
		b, err = CheckByte(buf)
		if err != nil {
			if err == io.EOF {
				return MachineCode[Identifier], string(temp), nil
			}
			return 0, "", err
		}
	}
	buf.UnreadByte()
	temp = temp[:len(temp)-1]
finish:
	return MachineCode[Identifier], string(temp), nil
}

func handlerOther(b byte, buf *Buffer) (int, string, error) {
	var (
		temp []byte
	)
	temp = append(temp, b)
	if isSpace(b) || b == '$' {
		return MachineCode[""], "", nil
	}

	if buf.Len() >= 0 {
		switch string(temp) {
		case "<", "+", ">", "=", ";", ":", ")", "(", ",", ".", "/", "*", "-":
			if buf.Len() >= 1 {
				b, _ = CheckByte(buf)
				temp = append(temp, b)
				switch string(temp) {
				case ":=", ">=", "<>", "<=":
					return MachineCode[string(temp)], string(temp), nil
				default:
					buf.UnreadByte()
					temp = temp[:len(temp)-1]
				}
			}

			return MachineCode[string(temp)], string(temp), nil
		default:
			return MachineCode["err"], "#", nil
		}
	}

	return MachineCode["err"], "#", nil
}

func isSpace(b byte) bool {
	return b == Space
}

func main() {
	t, s, err := LexicalAnalysis("./test.lu")
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

	for i := range LexicalErrors {
		fmt.Printf("%+v", LexicalErrors[i])
		fmt.Println()
	}
}
