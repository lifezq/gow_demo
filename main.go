package main

import (
	"log"
	"net/http"
	"time"

	"./websrv/v1"

	"github.com/lifezq/gow"
)

func main() {

	s := gow.New()

	s.SetConfig(&gow.Config{
		BaseUrl:      "/web",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	})

	s.RegisterController("say", &v1.Demo{})

	log.Printf("Server is running %s", ":16000")

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	s.Run(":16000")
}
