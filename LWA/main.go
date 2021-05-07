package main

import (
	"LWA/consumer"
	PubSub "LWA/pubsub"
	"LWA/users"
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	ps := PubSub.NewPubSub()
	go consumer.Consumer(ps)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		scannerText := scanner.Text()
		tokens := strings.Split(scannerText, " ")
		switch tokens[0] {
		case "addUser":
			if len(tokens) < 3 {
				log.Println("Enter proper command")
				break
			}
			users.AddUser(tokens[1], tokens[2])
		case "addTopic":
			if len(tokens) < 3 {
				log.Println("Enter proper command")
				break
			}
			user, err := users.GetUser(tokens[2])
			if err != nil {
				log.Println(err.Error())
				break
			}
			if user.Role != "admin" {
				log.Println("User should be an admin")
				break
			}
			ps.AddTopic(tokens[1])
		case "subscribeTopic":
			if len(tokens) < 3 {
				log.Println("Enter proper command")
				break
			}
			user, err := users.GetUser(tokens[2])
			if err != nil {
				log.Println(err.Error())
				break
			}
			ps.Subscribe(tokens[1], user)
		case "postEvent":
			if len(tokens) < 2 {
				log.Println("Enter proper command")
				break
			}
			event := make(map[string]string)
			jsonData := scannerText[strings.IndexByte(scannerText, ' ')+1:]
			err := json.Unmarshal([]byte(jsonData), &event)
			if err != nil {
				log.Println("Unable to read json message :- ", err.Error())
				break
			}
			ps.Publish(event["topicName"], event["text"])
		case "viewSubscribedTopics":
			if len(tokens) < 2 {
				log.Println("Enter proper command")
				break
			}
			user, err := users.GetUser(tokens[1])
			if err != nil {
				log.Println(err.Error())
				break
			}

			if len(user.Data) == 0 {
				fmt.Println("No Topics Subscribed")
				break
			}

			for k, _ := range user.Data {
				fmt.Print(k, " ")
			}

			fmt.Println()
		case "removeUser":
			if len(tokens) < 3 {
				log.Println("Enter proper command")
				break
			}
			user, err := users.GetUser(tokens[2])
			if err != nil {
				log.Println(err.Error())
				break
			}
			if user.Role != "admin" {
				log.Println("User should be an admin")
				break
			}
			if tokens[1] == tokens[2] {
				log.Println("User cannot remove themselves")
				break
			}

			users.RemoveUser(tokens[1])
		case "removeTopic":
			if len(tokens) < 3 {
				log.Println("Enter proper command")
				break
			}
			user, err := users.GetUser(tokens[2])
			if err != nil {
				log.Println(err.Error())
				break
			}
			if user.Role != "admin" {
				log.Println("User should be an admin")
				break
			}
			ps.RemoveTopic(tokens[1])
		default:
			fmt.Println("Invalid Command")
		}

	}
}
