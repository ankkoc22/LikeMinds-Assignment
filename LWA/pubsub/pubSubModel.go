package PubSub

import (
	. "LWA/users"
	"log"
	"sync"
)

type PubSub struct {
	TopicData map[string]TopicStruct
	Subs map[string][]User
}

type TopicStruct struct{
	Lock *sync.Mutex
	Data []string
}

func NewPubSub() *PubSub {
	ps := &PubSub{}
	ps.TopicData = make(map[string]TopicStruct)
	ps.Subs = make(map[string][]User)
	return ps
}

func (ps *PubSub) Subscribe(topic string, user User) {
	if _,ok:=ps.TopicData[topic];!ok{
		log.Println("Topic does not exist")
		return
	}
	ps.Subs[topic] = append(ps.Subs[topic], user)
}

func (ps *PubSub) Publish(topic string, msg string) {
	if _,ok:=ps.TopicData[topic];!ok{
		log.Println("Topic does not exist")
		return
	}
	lock := ps.TopicData[topic].Lock
	lock.Lock()
	defer lock.Unlock()
	ps.TopicData[topic] = TopicStruct{
		Lock: lock,
		Data: append(ps.TopicData[topic].Data,msg),
	}
}

func (ps *PubSub) AddTopic(topic string){
	ps.TopicData[topic] = TopicStruct{
		Lock: &sync.Mutex{},
	}
}

func (ps *PubSub) RemoveTopic(topic string){
	if _,ok:=ps.TopicData[topic];!ok{
		log.Println("Topic does not exist")
		return
	}
	delete(ps.TopicData,topic)
}