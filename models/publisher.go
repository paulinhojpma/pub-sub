package models

type Publisher struct {
	broker *broker
}

func NewPublisher() *Publisher {
	return &Publisher{
		broker: GetBroker(),
	}
}

func (p *Publisher) PublishOnTopic(route string, message Message) error {
	return p.broker.addMessageToTopic(route, message)
}
