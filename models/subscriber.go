package models

import (
	"pub-sub/exception"
	"pub-sub/helpers"
)

type Subscriber struct {
	messageChan   chan Message
	readToConsume chan bool
}

func NewSubscriber(route string) (*Subscriber, error) {
	for _, t := range GetBroker().topics {
		if helpers.CheckIfElementExistInSliceThatsMatchs(route, GetBroker().topics, func(element *Topic) string {
			return element.Name
		}) {
			readToConsume := make(chan bool)
			sub := &Subscriber{readToConsume: readToConsume}
			t.Subscribers = append(t.Subscribers, sub)
			return sub, nil
		}
	}
	return nil, exception.ErrorTopicNotFound
}

func (s *Subscriber) Consume() Message {
	if s.messageChan == nil {
		s.messageChan = make(chan Message)
	}
	for {
		s.readToConsume <- true
		messsage := <-s.messageChan
		if !messsage.Flag {
			return messsage
		}
	}
}

func (s *Subscriber) StopConsume() {
	s.readToConsume <- false
	s.messageChan = nil
}
