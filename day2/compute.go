package main

import "fmt"

const (
	AddCode      = 1
	MultiplyCode = 2
	HaltCode     = 99
)

type Program struct {
	Tape            []int
	CurrentPosition int
}

func NewProgram(t []int) *Program {
	return &Program{Tape: t, CurrentPosition: 0}
}

func (p *Program) movePosition(i int) {
	p.CurrentPosition = i
}

func (p *Program) moveLeft(num int) {
	p.CurrentPosition -= num
}

func (p *Program) moveRight(num int) {
	p.CurrentPosition += num
}

func (p *Program) readCurrentValue() int {
	return p.Tape[p.CurrentPosition]
}

func (p *Program) readValue(i int) int {
	return p.Tape[i]
}
func (p *Program) writeValue(i, v int) {
	p.Tape[i] = v
}

func (p *Program) Execute() {
	for {
		switch pos := p.CurrentPosition; p.Tape[pos] {
		case AddCode:
			first := p.readValue(p.CurrentPosition + 1)
			second := p.readValue(p.CurrentPosition + 2)
			dest := p.readValue(p.CurrentPosition + 3)
			p.writeValue(dest, p.readValue(first)+p.readValue(second))
		case MultiplyCode:
			first := p.readValue(p.CurrentPosition + 1)
			second := p.readValue(p.CurrentPosition + 2)
			dest := p.readValue(p.CurrentPosition + 3)
			p.writeValue(dest, p.readValue(first)*p.readValue(second))
		case HaltCode:
			return
		default:
			fmt.Printf("%+v", p)
			panic("unknown opcode")
		}
		p.moveRight(4)
		p.Execute()
	}
}
