package main

import (
	"fmt"
	"encoding/json"
	"log"
	"io/ioutil"
)

type Bot struct {
	id string `json:"bot_id"`
	groupID	string `json:"group_id"`
	host	string	`json:"host"`
	port	string	`json:"port"`
	logfile	string	`json:"logfile"`
	server	string
	Hooks	map[string]func(InMsg) string
}

type InMsg struct {
	avatar_url string `json:"avatar_url"`
	id	string	`json:"id"`
	name	string	`json:"name"`
	sender_id	string	`json:"sender_id"`
	sender_type	string	`json:"sender_type"`
	system	bool	`json:"system"`
	text	string	`json:"text"`
	user_id	string	`json:"user_id"`
}

type OutMsg struct {
	id	string	`json:"bot_id"`
	text	string	`json:"text"`
}

// Parses JSON and creates bot
func CreateJsonBot(filename string) (*Bot, error) {
	file, err := ioutil.ReadFile(filename)

	var bot Bot
	if err != nil {
		log.Fatal("Unable to parse bot config")
		return nil, err
	}

	json.Unmarshal(file, &bot)

	bot.server = bot.host + ":" + bot.port
	log.Printf("Creating bot at %s\nLogging at %s\n", bot.server, bot.logfile)
	bot.Hooks = make(map[string]func(InMsg) string)

	log.Println("Bot Created:"+bot.id)
	return &bot, err

}

func HelloTest(str string) {
	fmt.Println(str)
}

func main() {
        fmt.Println("hello world!")
	bot, err := CreateJsonBot("bot_config.json")
	if err != nil {
		log.Fatal("Could not create bot from JSON")
	}

	log.Println(bot.server)
	bot.AddHook("Hi!$", hello)
	bot.AddHook("Hello!$", hello2)
	bot.AddHook("test$", test)
	bot.AddHook("request", suggestion)
	log.Printf("Listening on %v...\n", bot.server)
	http.HandleFunc("/", bot.Handler())
	log.Fatal(http.ListenAndServe(bot.server, nil))
}

/*
 Test hook functions
 Each hook should match a certain string, and if it matches
 it should return a string of text
 Hooks will be traversed until match occurs
*/
func hello(msg groupmebot.InboundMessage) (string) {
        resp := fmt.Sprintf("Hi, %v.", msg.Name)
        return resp
}

func hello2(msg groupmebot.InboundMessage) (string) {
        resp := fmt.Sprintf("Hello, %v.", msg.Name)
        return resp
}

func test(msg groupmebot.InboundMessage) (string) {
        resp := fmt.Sprintf("This bot is a work in prorgress, %v.", msg.Name)
        return resp
}

func suggestion(msg groupmebot.InboundMessage) (string) {
        resp := fmt.Sprintf("Your suggestion has been saved, %v.", msg.Name)
        log.Printf("Suggestion from %s: msg %s\n", msg.Name, msg)
        return resp
}



