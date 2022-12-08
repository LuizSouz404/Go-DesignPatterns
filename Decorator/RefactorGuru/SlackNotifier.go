package main

import "fmt"

type SlackNotifier struct {
	Notifier Notifier
}

func (s *SlackNotifier) Send(message string) {
	fmt.Println("Slack message: " + message)
	if s.Notifier != nil {
		s.Notifier.Send(message)
	}
}
