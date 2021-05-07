package consumer

import (
	. "LWA/pubsub"
	"encoding/json"
	"fmt"
	"time"
)

func Consumer(ps *PubSub){
	for{
		time.Sleep(1*time.Second)
		for k,v := range ps.TopicData{
			v.Lock.Lock()
			if len(v.Data) > 0{
				subs := ps.Subs[k]
				for _,users := range subs{
					users.Data[k] = v.Data
					for _,data := range v.Data{
						response := map[string]string{
							"topic" : k,
							"message" : data,
							"sentTo" : users.UserName,
						}

						res,_ := json.Marshal(response)
						fmt.Println(string(res))
					}
				}
				ps.TopicData[k] = TopicStruct{
					Lock: ps.TopicData[k].Lock,
					Data: []string{},
				}
			}
			v.Lock.Unlock()
		}
	}
}
