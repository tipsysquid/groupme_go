package groupmebot

import (
	"fmt"
	"encoding/json"
	"net/http"
	"os"
	"strings"
)

type Bot struct {
	id string `json:"bot_id"`
	groupID	string `json:"group_id"`
	host	string	`json:"host"`
	port	string	`json:"port"`
	logfile	string	`json:"logfile"`
	server	string
	Hooks	map[string]funct(InMsg) string
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
