package linkparser

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func argClosure(fileName string) func() {
	return func() {
		b, err := os.ReadFile("../../" + fileName)
		html := string(b)
		r := strings.NewReader(html)
		var parser Parser
		err = parser.ReadHtml(r)
		handleErr(err)
		err = parser.ParseTree()
		handleErr(err)
		for _, l := range parser.Result {
			fmt.Println(l)
		}
	}
}

func TestSampleOne(t *testing.T) {
	argClosure("sample1.html")()
}
func TestSampleTwo(t *testing.T) {
	argClosure("sample2.html")()
}
func TestSampleThree(t *testing.T) {
	argClosure("sample3.html")()
}
func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
