package main

import "fmt"

func main() {
	notifier := &FacebookNotifier{}

	notifier.Send("Nova promoção")

	fmt.Println("------------------------------------")

	notifierSMSAndFacebook := &SMSNotifier{
		Notifier: notifier,
	}

	notifierSMSAndFacebook.Send("Refactoring Guru")

	fmt.Println("------------------------------------")

	notifierSMSFacebookSlack := &SlackNotifier{
		Notifier: notifierSMSAndFacebook,
	}

	notifierSMSFacebookSlack.Send("Final Decorator")
}
