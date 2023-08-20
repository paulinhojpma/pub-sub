package models

import (
	"sync"
)

type Topic struct {
	Name        string
	Messages    []Message
	Subscribers []*Subscriber
}

func (t *Topic) SendMessagesToSubscribers() {
	for {
		var wg sync.WaitGroup
		messageToSend := Message{Flag: true}
		if len(t.Messages) > 0 {
			messageToSend, t.Messages = t.Messages[0], t.Messages[1:]
		}

		tam := len(t.Subscribers)

		wg.Add(tam)
		for _, sub := range t.Subscribers {
			go func(sub *Subscriber) {
				defer wg.Done()
				select {
				case readToConsume := <-sub.readToConsume:
					if readToConsume {
						sub.messageChan <- messageToSend
					}
				default:
					if sub.messageChan != nil {
						sub.messageChan <- messageToSend

					}
				}
				if !messageToSend.Flag {
					messageToSend.Consumed = true
				}
			}(sub)

		}
		wg.Wait()
		if !messageToSend.Consumed && !messageToSend.Flag {
			t.Messages = append(t.Messages, messageToSend)
		}
	}
}
