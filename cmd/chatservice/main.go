package main

func main() {
	configs, err := LoadConfig(".")
	if err != nil {
		panic(err)
	}
}
