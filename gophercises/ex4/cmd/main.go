package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tomschdev/gotime/gophercises/ex4/pkg/linkparser"
)

func main() {
	exampleHtml := `
<html>
<body>
  <h1>Hello!</h1>
  <a href="/other-page">A link to another page</a>
</body>
</html>
	`
	r := strings.NewReader(exampleHtml)
	var parser linkparser.Parser
	err := parser.ReadHtml(r)
	handleErr(err)

	err = parser.ParseTree()
	handleErr(err)

	for _, l := range parser.Result {
		fmt.Println(l)
	}

}
func handleErr(err error) {
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
}
