package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

func Test(question string, answer string) (string, bool) {
	in := bytes.NewBufferString(question)
	out := new(bytes.Buffer)
	err := Calc(in, out)
	if err != nil {
		return err.Error(), false
	}
	result := out.String()
	if result != answer {
		return "results not match\nGot: " + result + "\nExpected: " + answer, false
	}
	return "", true
}

func NewStack() *Stack {
	return &Stack{}
}

type Stack struct {
	nodes []int
	count int
}

func (s *Stack) Push(n int) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

func (s *Stack) Pop() int {
	if s.count == 0 {
		panic(errors.New("stack is empty"))
	}
	s.count--
	return s.nodes[s.count]
}

func Calc(in io.Reader, out io.Writer) (resultErr error) {
	defer func() {
		if err := recover(); err != nil {
			resultErr = fmt.Errorf("calc: %v", err)
		}
	}()

	s := NewStack()
	for {
		var input string
		_, err := fmt.Fscan(in, &input)
		if err != nil {
			if err != io.EOF {
				panic(err)
			}
			return
		}

		switch input {
		case "+":
			s.Push(s.Pop() + s.Pop())
		case "-":
			s.Push(-s.Pop() + s.Pop())
		case "*":
			s.Push(s.Pop() * s.Pop())
		case "=":
			_, err = fmt.Fprint(out, s.Pop())
			if err != nil {
				panic(err)
			}
			return nil
		default:
			n, err := strconv.Atoi(input)
			if err != nil {
				panic(err)
			}
			s.Push(n)
		}
	}
}

func main() {
	if err := Calc(os.Stdin, os.Stdout); err != nil {
		fmt.Print("error: " + err.Error())
	}
}
