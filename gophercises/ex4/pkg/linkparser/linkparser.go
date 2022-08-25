package linkparser

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Parser struct {
	Tree   *html.Node
	Result []Link
}

func (p *Parser) ReadHtml(r io.Reader) error {
	doc, err := html.Parse(r)
	if err != nil {
		return err
	}
	p.Tree = doc
	return nil
}

func (p *Parser) ParseTree() error {
	linkNodes := recordLinkNodes(p.Tree)
	for _, n := range linkNodes {
		p.Result = append(p.Result, newLink(n))
	}
	return nil
}

func newLink(n *html.Node) (link Link) {
	for _, att := range n.Attr {
		if att.Key == "href" {
			link.Href = att.Val
			break
		}
	}
	link.Text = strings.Join(strings.Fields(extractText(n)), " ")
	return
}

func extractText(n *html.Node) (text string) {
	if n.Type == html.TextNode {
		text = n.Data
		return
	}
	if n.Type != html.ElementNode {
		text = ""
		return
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		text += extractText(c) + ""
	}
	return
}

func recordLinkNodes(n *html.Node) (linkNodes []*html.Node) {
	if n.Type == html.ElementNode && n.Data == "a" {
		linkNodes = []*html.Node{n}
		return
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		linkNodes = append(linkNodes, recordLinkNodes(c)...)
	}
	return
}
