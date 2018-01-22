package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"strings"
	"github.com/adammohammed/groupmebot"
	"github.com/abrander/coinmarketcap"
)

/*
 Test hook functions
 Each hook should match a certain string, and if it matches
 it should return a string of text
 Hooks will be traversed until match occurs
*/

func test(msg groupmebot.InboundMessage) (string) {
	resp := fmt.Sprintf("This bot is a work in prorgress, %v.", msg.Name)
	return resp
}

func request(msg groupmebot.InboundMessage) (string) {
	resp := fmt.Sprintf("Your request has been saved, %v.", msg.Name)
        log.Printf("Suggestion from %s: msg %s\n", msg.Name, msg)
	return resp
}

func help(msg groupmebot.InboundMessage) (string) {
	resp := fmt.Sprintf("Available commands: request\nType '/request' followed by an idea, and it will be saved for later\n Type 'marketcap' for total crypto marketcap\n Type '/eth for current ethereum price\n Type '/price %' where % is the symbol name of the token you are looking for")
	return resp
}

func ethprice(msg groupmebot.InboundMessage) (string) {
	client, _ := coinmarketcap.NewClient()

	ticker, _ := client.Ticker(
		coinmarketcap.Currency("ethereum"),
	)

	coininfo := ticker.CoinBySymbol("eth")
	price, _ := coininfo.Price("USD")

	resp := fmt.Sprintf("Current Ethereum price in USD: %.0f (Updated %s ago)", 
		price, time.Since(coininfo.LastUpdated),
	)
	return resp
}

func symbolprice(msg groupmebot.InboundMessage) (string) {
	client, _ := coinmarketcap.NewClient()
	
	//symbol := splitAfter(msg.Text, '/price', "'")
	strArray := strings.Fields(msg.Text)

	log.Printf(msg.Text)
	log.Printf(strArray[0])

	symbol := strArray[1]
	log.Printf(symbol)
	ticker, _ :=client.Ticker(
		coinmarketcap.Limit(300),
	)

	coininfo := ticker.CoinBySymbol(symbol)
	price, _:= coininfo.Price("USD")

	resp := fmt.Sprintf("Current price of %s (%s) in USD: %.2f", 
		coininfo.Name, coininfo.Symbol, price,)

	return resp

}

func marketcap(msg groupmebot.InboundMessage) (string) {
	client, _ := coinmarketcap.NewClient()

	globaldata, _ := client.GlobalData(
		coinmarketcap.Convert("USD"),
	)

	cap, _ := globaldata.MarketCap("USD")

	resp :=	fmt.Sprintf("Global market cap in USD: %.0f (Updated %s ago)\n",
		cap,
		time.Since(globaldata.LastUpdated),
	)
	return resp
}


func main() {

	bot, err := groupmebot.NewBotFromJson("mybot_cfg.json")
	if err != nil {
		log.Fatal("Could not create bot structure")
	}

	// Make a list of functions
	bot.AddHook("/test$",test)
	bot.AddHook("/request",request)
	bot.AddHook("/help",help)
	bot.AddHook("/marketcap",marketcap)
	bot.AddHook("/eth", ethprice)
	bot.AddHook("/price", symbolprice)
	// Create Server to listen for incoming POST from GroupMe
	log.Printf("Listening on %v...\n", bot.Server)
	http.HandleFunc("/", bot.Handler())
	initMsg(bot)
	log.Fatal(http.ListenAndServe(bot.Server, nil))
}

func initMsg (bot *groupmebot.GroupMeBot){
	//bot.SendMessage("System warming up...\nType /help for available commands")
	bot.SendMessage("Updated Limit query to include top 300 tokens in search")
}

