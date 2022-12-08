package main

type Application struct {
	Notifier Notifier
}

func (a *Application) SetNotifier(notifier Notifier) {
	a.Notifier = notifier
}

func (a *Application) SendMessage(message string) {
	a.Notifier.Send(message)
}
