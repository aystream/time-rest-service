package main

import "github.com/aystream/time-rest-service/src/app"

func main() {
	newApp := &app.App{}
	newApp.Initialize()
	newApp.Run(":8080")
}
