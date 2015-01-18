package main

import (
	"os"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) < 2 {
		os.Stderr.WriteString("Input a search word.\n")
		os.Exit(1)
	}

	doc, err := goquery.NewDocument("http://eow.alc.co.jp/search?q=" + os.Args[1])
	if err != nil {
		panic(err)
	}

	r := doc.Find("#resultsList ul li").First()

	os.Stdout.WriteString(r.Find(".midashi").First().Text())

	r.Find(".wordclass").Each(func(_ int, w *goquery.Selection) {
		os.Stdout.WriteString("\n" + w.Text())

		w.Next().Find("li").Each(func(i int, l *goquery.Selection) {
			os.Stdout.WriteString("\n" + strconv.Itoa(i+1) + ". ")
			for c := l.Nodes[0].FirstChild; c != nil; c = c.NextSibling {
				if c.Type == html.ElementNode {
					if c.Data == "br" {
						os.Stdout.WriteString("\n")
					}
					continue
				}

				os.Stdout.WriteString(c.Data)
			}
		})
	})

	os.Stdout.WriteString("\n発音: " + r.Find(".pron").First().Text() + "\n")
}
