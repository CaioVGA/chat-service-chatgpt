package main

import "github.com/CaioVGA/chat-service-chatgpt/configs"

func main() {
	configs, err := LoadConfig(".")
	if err != nil {
		panic(err)
	}
}
