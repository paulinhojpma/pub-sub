package models

import (
	"pub-sub/exception"
	"testing"
)

func TestPublish(t *testing.T) {
	testCases := []struct {
		name        string
		route       string
		wantedError error
	}{
		{
			name:        "success",
			route:       "test",
			wantedError: nil,
		},
		{
			name:        "error_topic_dont_exists",
			route:       "test_not",
			wantedError: exception.ErrorTopicNotFound,
		},
	}
	LoadBrokerForTest()
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			publisher := NewPublisher()
			err := publisher.PublishOnTopic(tc.route, Message{Body: []byte("test")})
			if err != tc.wantedError {
				t.Errorf("error on publish want (%v), got (%v)", tc.wantedError, err)
			}

		})
	}
}

func LoadBrokerForTest() {
	broker := GetBroker()
	broker.NewTopic("test")

}
