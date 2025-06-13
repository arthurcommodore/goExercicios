package main

import (
	"github.com/nsf/termbox-go"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(10, 5, '★', termbox.ColorRed, termbox.ColorBlack)
	termbox.Flush()

	termbox.PollEvent() // Espera por uma tecla para sair
}
