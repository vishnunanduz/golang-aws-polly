package main

import (
	"github.com/vishnunanduz/go-aws-polly/service"
)

var (
	joey service.PollyService = service.NewJoeyPollyService()
)

func main() {
	err := joey.SynthesizeText("Hi, I am vishnujith", "test.mp3")
	if err != nil {
		panic(err)
	}
}
