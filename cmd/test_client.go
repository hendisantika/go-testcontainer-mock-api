package main

import (
	"fmt"
	http_client "go-testcontainer-mock-api/internal/http-client"
	"os"
)

func main() {
	client, err := http_client.NewClient("https://api.agify.io")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	res, err := client.GetAge("Yuji")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(res)
}
