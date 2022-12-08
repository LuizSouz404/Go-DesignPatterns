package main

import "fmt"

type SMSNotifier struct {
	Notifier Notifier
}

func (s *SMSNotifier) Send(message string) {
	fmt.Println("SMS message: " + message)
	if s.Notifier != nil {
		s.Notifier.Send(message)
	}
}
