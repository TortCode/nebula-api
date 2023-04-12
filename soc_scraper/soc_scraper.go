package main

import (
	"context"

	"github.com/chromedp/chromedp"
)

var socUrl string = `https://soc.utdallas.edu/`

func main() {
	// obtain context
	ctx, cancel := chromedp.NewContext(context.Background())
	// ensure cleanup occurs
	defer cancel()

	err := chromedp.Run(ctx,
		loginToSoc(),
	)

	if err != nil {
		panic(err)
	}
}

func loginToSoc() chromedp.Tasks {
	exploreQuery := `.main > .whitebox > .c3::nth-child(2) > .sidenav > ul > li:first-child`
	return chromedp.Tasks{
		chromedp.Navigate(socUrl),
		chromedp.Click(exploreQuery, chromedp.ByQuery),
	}
}
