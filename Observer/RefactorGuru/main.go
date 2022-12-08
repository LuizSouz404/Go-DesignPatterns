package main

func main() {
	publisher := &Publisher{}

	subscriber := &ConcreteSubscriber{
		NameSubscriber: "Luiz",
	}
	publisher.Subscriber(subscriber)

	subscriberOne := &ConcreteSubscriber{
		NameSubscriber: "Brunna",
	}
	publisher.Subscriber(subscriberOne)

	subscriberTwo := &ConcreteSubscriber{
		NameSubscriber: "João",
	}
	publisher.Subscriber(subscriberTwo)

	publisherThree := &ConcreteSubscriber{
		NameSubscriber: "Pedro",
	}
	publisher.Subscriber(publisherThree)

	publisher.NotifySubscribers()
	publisher.Unsubscriber(publisherThree)
	publisher.NotifySubscribers()
}
