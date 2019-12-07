package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	AddCode       = 1
	MultiplyCode  = 2
	StoreCode     = 3
	OutputCode    = 4
	JumpTrueCode  = 5
	JumpFalseCode = 6
	LessThanCode  = 7
	EqualsCode    = 8
	HaltCode      = 99
)

type Program struct {
	Tape            []int
	CurrentPosition int
}

func NewProgram(t []int) *Program {
	return &Program{Tape: t, CurrentPosition: 0}
}

func (p *Program) write(index, value int) {
	p.Tape[index] = value
}

func (p *Program) getParam(param int) int {
	var ret int
	switch param {
	case 1:
		if (p.Tape[p.CurrentPosition]/100)%10 == 0 {
			ret = p.Tape[p.Tape[p.CurrentPosition+1]]
		} else {
			ret = p.Tape[p.CurrentPosition+1]
		}
	case 2:
		if (p.Tape[p.CurrentPosition]/1000)%10 == 0 {
			ret = p.Tape[p.Tape[p.CurrentPosition+2]]
		} else {
			ret = p.Tape[p.CurrentPosition+2]
		}
	case 3:
		if (p.Tape[p.CurrentPosition]/10000)%10 == 0 {
			ret = p.Tape[p.Tape[p.CurrentPosition+3]]
		} else {
			ret = p.Tape[p.CurrentPosition+3]
		}
	}
	return ret
}
func (p *Program) Execute() {
	for {
		intcode := p.Tape[p.CurrentPosition] % 100
		switch intcode {
		case AddCode:
			p.write(p.Tape[p.CurrentPosition+3], p.getParam(1)+p.getParam(2))
			p.CurrentPosition += 4
		case MultiplyCode:
			p.write(p.Tape[p.CurrentPosition+3], p.getParam(1)*p.getParam(2))
			p.CurrentPosition += 4
		case StoreCode:
			buf := bufio.NewReader(os.Stdin)
			fmt.Print("> ")
			s, err := buf.ReadString('\n')
			if err != nil {
				panic(err)
			}
			val, err := strconv.Atoi(strings.Trim(s, "\n"))
			if err != nil {
				panic(err)
			}
			p.write(p.Tape[p.CurrentPosition+1], val)
			p.CurrentPosition += 2
		case OutputCode:
			os.Stdout.WriteString(strconv.Itoa(p.getParam(1)) + "\n")
			p.CurrentPosition += 2
		case JumpTrueCode:
			if p.getParam(1) != 0 {
				p.CurrentPosition = p.getParam(2)
			} else {
				p.CurrentPosition += 3
			}
		case JumpFalseCode:
			if p.getParam(1) == 0 {
				p.CurrentPosition = p.getParam(2)
			} else {
				p.CurrentPosition += 3
			}
		case LessThanCode:
			if p.getParam(1) < p.getParam(2) {
				p.write(p.Tape[p.CurrentPosition+3], 1)
			} else {
				p.write(p.Tape[p.CurrentPosition+3], 0)
			}
			p.CurrentPosition += 4
		case EqualsCode:
			if p.getParam(1) == p.getParam(2) {
				p.write(p.Tape[p.CurrentPosition+3], 1)
			} else {
				p.write(p.Tape[p.CurrentPosition+3], 0)
			}
			p.CurrentPosition += 4
		case HaltCode:
			return
		default:
			panic("unknown intcode")
		}
		p.Execute()
	}
}
