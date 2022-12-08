package main

import "fmt"

type FacebookNotifier struct {
	Notifier Notifier
}

func (f *FacebookNotifier) Send(message string) {
	fmt.Println("Facebook message: " + message)
	if f.Notifier != nil {
		f.Notifier.Send(message)
	}
}
