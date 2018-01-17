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
	"io"
	"io/ioutil"
	"os"
)

func what(a byte) int {
	if (a >= 48) && (a <= 57) {
		//０-９数字
		return 0
	} else if (a >= 97) && (a <= 122) {
		//a-z的字母
		return 1
	} else {
		//其他的标点符号
		return 2
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
	return nil, nil, nil
}

func CheckByte(f *bytes.Buffer) (byte, error) {

}

func isspace(b byte) bool {
	return b == 0x20
}
