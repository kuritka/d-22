package main

import "fmt"

type Message struct {
	To      []string
	From    []string
	Content string
}

type FailedMessage struct {
	ErrorMessage    string
	OriginalMessage Message
}

func main() {
	msgCh := make(chan Message, 1)
	errCh := make(chan FailedMessage, 1)
	select {
	case receivedMsg := <-msgCh:
		fmt.Println(receivedMsg)
	case failedMsg := <-errCh:
		fmt.Println(failedMsg)
	default:
		fmt.Println("No messages ")
	}
}

//func main() {
//	msgCh := make(chan Message, 1)
//	errCh := make(chan FailedMessage, 1)
//	msg := Message{
//		To:      []string{"michalK@seznam.cz"},
//		From:    []string{"robot@roboto.io"},
//		Content: "Keep it secret, keep it safe.",
//	}
//	err := FailedMessage{
//		ErrorMessage:    "Message intercepted by someone",
//		OriginalMessage: Message{},
//	}
//	msgCh <- msg
//	errCh <- err
//	fmt.Println(<-msgCh)
//	fmt.Println(<-errCh)
//
//}
