package embedhtml

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

const extension = ".embed.html"

type EmbedHTML struct {
	Options
}

func (c *EmbedHTML) Run() (err error) {
	if c.About {
		fmt.Println("Visit https://github.com/gonejack/embede-html")
		return
	}
	if len(c.HTML) == 0 {
		return errors.New("no .html files given")
	}
	return c.run()
}
func (c *EmbedHTML) run() (err error) {
	for _, html := range c.HTML {
		err = c.process(html)
		if err != nil {
			return
		}
	}
	return
}
func (c *EmbedHTML) process(html string) (err error) {
	if strings.HasSuffix(html, extension) {
		return
	}

	opt := chromedp.WithBrowserOption(
		chromedp.WithDialTimeout(time.Minute),
	)
	ctx, cancel := chromedp.NewContext(context.TODO(), opt)
	defer cancel()

	content := ""
	//output := strings.TrimSuffix(html, ".html") + extension
	requests := make(map[network.RequestID]*network.Request)
	response := make(map[network.RequestID]*network.Response)
	respBody := make(map[network.RequestID][]byte)

	chromedp.ListenTarget(ctx, func(v interface{}) {
		switch ev := v.(type) {
		case *network.EventRequestWillBeSent:
			log.Printf("req: [reqId=%s, url=%s]", ev.RequestID, ev.Request.URL)
			requests[ev.RequestID] = ev.Request
		case *network.EventLoadingFinished:
			log.Printf("rsp: [reqId=%s]", ev.RequestID)
		case *network.EventResponseReceived:
			response[ev.RequestID] = ev.Response
		case *network.EventResponseReceivedExtraInfo:
		case *network.EventDataReceived:
		default:
		}
	})

	err = chromedp.Run(ctx,
		chromedp.Navigate(pathToURI(html)),
		chromedp.ActionFunc(func(ctx context.Context) error {
			for reqId := range requests {
				body, err := network.GetResponseBody(reqId).Do(ctx)
				if err == nil {
					respBody[reqId] = body
				}
			}
			return nil
		}),
		chromedp.ActionFunc(func(ctx context.Context) error {
			return chromedp.OuterHTML("html", &content).Do(ctx)
		}),
	)
	if err != nil {
		return fmt.Errorf("cannot render %s: %s", html, err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))
	if err != nil {
		return
	}

	doc.Find("img,link,script").Each(func(i int, e *goquery.Selection) {

	})

	return
}
