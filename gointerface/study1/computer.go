package main

type Card interface {
	Dis()
}
type CPU interface {
	Cal()
}
type Memory interface {
	Sto()
}

type Computer struct {
	cpu    CPU
	memory Memory
	card   Card
}

func NewComputer(card Card, memory Memory, cpu CPU) *Computer {
	return &Computer{
		cpu:    cpu,
		card:   card,
		memory: memory,
	}
}

func (this *Computer) Dowork() {
	this.card.Dis()
	this.cpu.Cal()
	this.memory.Sto()
}
