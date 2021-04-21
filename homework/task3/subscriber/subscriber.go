package subscriber

import (
	"github.com/nsqio/go-nsq"
)

type Subscriber struct {
	consumer *nsq.Consumer
}

func NewSubscriber(channel string, handler nsq.Handler, addr string) (*Subscriber, error) {
	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("topic", channel, config)
	if err != nil {
		return nil, err
	}

	consumer.AddHandler(handler)

	err = consumer.ConnectToNSQLookupd(addr)
	if err != nil {
		return nil, err
	}

	return &Subscriber{consumer}, nil
}

func (s *Subscriber) Stop() {
	s.consumer.Stop()
}
