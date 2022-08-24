package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/tomschdev/gotime/gophercises/ex3/pkg/cyoa"
)

func main() {
	port := flag.Int("port", 3030, "select port to start cyoa server")
	storyFile := flag.String("story", "cyoa.json", "file containing chapters for cyoa game")
	flag.Parse()
	f, err := os.Open(*storyFile)
	handleErr(err)

	var story cyoa.Story
	err = story.ReadStoryFile(f)
	handleErr(err)

	handler := story.MakeHandler()
	fmt.Printf("Listening on port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), handler))

}

func handleErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
