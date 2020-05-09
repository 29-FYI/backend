package main

import (
	"net/url"
	"unicode/utf8"

	"github.com/29-FYI/twentynine"
)

type linkring struct {
	links [twentynine.TwentyNine]twentynine.Link
	i     int
}

func (lr linkring) LinkLink(link twentynine.Link) (linkring, bool) {
	if len(link.Headline) < 8 || len(link.URL) < 8 || len(link.Headline) > 128 || len(link.URL) > 128 {
		return lr, false
	}
	if !utf8.ValidString(link.Headline) || !utf8.ValidString(link.Headline) {
		return lr, false
	}
	u, err := url.Parse(link.URL)
	if err != nil {
		return lr, false
	}
	if u.Scheme != "https" {
		return lr, false
	}

	lr.links[lr.i] = link
	lr.i = (lr.i + 1) % twentynine.TwentyNine
	return lr, true
}

func (lr linkring) Link(i int) (link twentynine.Link) {
	if i > twentynine.TwentyNine {
		return
	}
	return lr.links[(twentynine.TwentyNine-i-1+lr.i)%twentynine.TwentyNine]
}

func (lr linkring) Links() (links []twentynine.Link) {
	links = make([]twentynine.Link, 0)
	for i := 0; i < twentynine.TwentyNine; i++ {
		link := lr.Link(i)
		if link.Headline == "" {
			break
		}
		links = append(links, link)
	}
	return
}
