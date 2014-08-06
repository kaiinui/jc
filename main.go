package main

import (
	"github.com/nsf/termbox-go"
	"log"
)

func main() {
	_ = termbox.Init()
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	termbox.Flush()
loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyCtrlS {
				termbox.Sync()
			}
			if ev.Key == termbox.KeyCtrlQ {
				break loop
			}
			if ev.Key == termbox.KeyCtrlC {
				break loop
			}

			write(&ev)

			termbox.Flush()
		}
	}

}

func write(ev *termbox.Event) {
	log.Println(ev.Ch)
}
