package main

type Player struct {
	name string
}

func (p * Player) setName(n string) {
	(*p).name = n
}
