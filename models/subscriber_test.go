package models

import (
	"pub-sub/exception"
	"testing"
)

func TestSubcriber(t *testing.T) {
	testCases := []struct {
		name          string
		route         string
		wantedMessage Message
		wantedError   error
	}{
		{
			name:          "success",
			route:         "test",
			wantedMessage: Message{Body: []byte("test")},
			wantedError:   nil,
		},
		{
			name:          "topic_not_found",
			route:         "test_notaaa",
			wantedMessage: Message{Body: []byte("test")},
			wantedError:   exception.ErrorTopicNotFound,
		},
	}
	LoadBrokerForTest()
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			PublishOnTestTopic()
			subscriber, err := NewSubscriber(tc.route)
			if err != tc.wantedError {
				t.Errorf("error on publish want (%v), got (%v)", tc.wantedError, err)

			}
			if err == nil {
				defer subscriber.StopConsume()
				message := subscriber.Consume()
				if string(message.Body) != string(tc.wantedMessage.Body) {
					t.Errorf("error on publish want (%v), got (%v)", string(message.Body), string(tc.wantedMessage.Body))
				}
			}

		})
	}
}

func PublishOnTestTopic() {
	publisher := NewPublisher()
	publisher.PublishOnTopic("test", Message{Body: []byte("test")})
}
