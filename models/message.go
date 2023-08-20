package models

type Message struct {
	Body     []byte
	Consumed bool
	Flag     bool
}
