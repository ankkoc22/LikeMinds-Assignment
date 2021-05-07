package users

import (
	"errors"
	"log"
)

var UserMap map[string]User

type User struct{
	UserName string
	Role string
	Data map[string][]string
}

func init(){
	UserMap = make(map[string]User)
}

func AddUser(username,role string){
	if role != "admin" && role != "user"{
		log.Println("Invalid Role")
		return
	}
	UserMap[username] = User{
		UserName: username,
		Role:     role,
		Data:     make(map[string][]string),
	}
}

func GetUser(username string) (User,error){
	if _,ok := UserMap[username];!ok{
		return User{},errors.New("User does not exists")
	}

	return UserMap[username],nil
}

func RemoveUser(username string) {
	if _,ok := UserMap[username];!ok{
		log.Println("User does not exists")
		return
	}
	delete(UserMap,username)
}