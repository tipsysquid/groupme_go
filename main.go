package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/adammohammed/groupmebot"
)

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

func help(msg groupmebot.InboundMessage) (string) {
	resp := fmt.Sprintf("Available commands: request\nType 'request' followed by an idea, and it will be saved for later")
	return resp
}

func main() {

	bot, err := groupmebot.NewBotFromJson("mybot_cfg.json")
	if err != nil {
		log.Fatal("Could not create bot structure")
	}

	// Make a list of functions
	bot.AddHook("Hi!$", hello)
	bot.AddHook("Hello!$", hello2)
	bot.AddHook("test$",test)
	bot.AddHook("request",suggestion)
	bot.AddHook("Request",suggestion)
	bot.AddHook("/help",help)
	// Create Server to listen for incoming POST from GroupMe
	log.Printf("Listening on %v...\n", bot.Server)
	http.HandleFunc("/", bot.Handler())
	initMsg(bot)
	log.Fatal(http.ListenAndServe(bot.Server, nil))
}

func initMsg (bot *groupmebot.GroupMeBot){
	bot.SendMessage("System warming up...\nType /help for available commands")
	bot.SendMessage("Now supporting lazy-patch for uppercase R!")
}
