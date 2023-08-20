package models

import (
	"pub-sub/exception"
)

type MessageBody interface {
	[]byte
}

type broker struct {
	topics []*Topic
}

var singleBroker *broker

func GetBroker() *broker {
	if singleBroker == nil {
		singleBroker = &broker{}
	}
	return singleBroker
}

func (b *broker) NewTopic(route string) {
	for _, v := range b.topics {
		if v.Name == route {
			return
		}
	}
	topic := &Topic{Name: route}
	b.topics = append(b.topics, topic)
	go topic.SendMessagesToSubscribers()
}

func (b *broker) addMessageToTopic(route string, message Message) error {
	for _, topic := range b.topics {
		if topic.Name == route {
			topic.Messages = append(topic.Messages, message)
			return nil
		}
	}
	return exception.ErrorTopicNotFound
}
