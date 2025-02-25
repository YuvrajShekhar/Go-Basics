package main

import "fmt"

func main() {
	website_map := map[string]string{
		"Google":  "https://google.com",
		"Amazone": "https://amazon.com",
	}

	fmt.Println(website_map)
	website_map["LinkedIN"] = "https://www.linkedin.com"
	fmt.Println(website_map)

	delete(website_map, "Google")
	fmt.Println(website_map)
}
