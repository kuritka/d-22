package main

import (
	"fmt"
)

type Button struct {
	eventListeners map[string][]chan string
}

func NewButton() *Button {
	b := new(Button)
	b.eventListeners = make(map[string][]chan string)
	return b
}

func (b *Button) AddEventListener(event string, responseChannel chan string) {
	if _, found := b.eventListeners[event]; found {
		b.eventListeners[event] = append(b.eventListeners[event], responseChannel)
		return
	}
	b.eventListeners[event] = []chan string{responseChannel}
}

func (b *Button) RemoveEventListener(event string, listenerChannel chan string) {
	if _, found := b.eventListeners[event]; found {
		without := []chan string{}
		for _, v := range b.eventListeners[event] {
			if v != listenerChannel {
				without = append(without, v)
			}
		}
		b.eventListeners[event] = without
	}
}

func (b *Button) TriggerEvent(event string, response string) {
	if _, found := b.eventListeners[event]; found {
		for _, v := range b.eventListeners[event] {
			// useful, because we don't want to have unresponsive UI.
			// we want to inform handler and continue
			go func(handler chan string) {
				handler <- response
			}(v)
		}
	}
}

func main() {
	btn := NewButton()

	handlerOne := make(chan string)
	handlerTwo := make(chan string)

	btn.AddEventListener("click", handlerOne)
	btn.AddEventListener("click", handlerTwo)

	go func() {
		for {
			msg := <-handlerOne
			fmt.Println("HandlerOne ", msg)
		}
	}()

	go func() {
		for {
			msg := <-handlerTwo
			fmt.Println("HandlerTwo ", msg)
		}
	}()

	btn.TriggerEvent("click", "Button clicked")

	btn.RemoveEventListener("click", handlerTwo)

	btn.TriggerEvent("click", "Button clicked AGAIN!!")

	fmt.Scanln()
}
