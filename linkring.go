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

func (lr linkring) LinkLink(link twentynine.Link) (nlr linkring, ok bool) {
	if len(link.Headline) < 8 || len(link.URL) < 8 || len(link.Headline) > 128 || len(link.URL) > 128 {
		return
	}
	if !utf8.ValidString(link.Headline) || !utf8.ValidString(link.URL) {
		return
	}
	if _, err := url.Parse(link.URL); err != nil {
		return
	}
	lr.links[lr.i] = link
	lr.i = (lr.i + 1) % twentynine.TwentyNine
	return lr, true
}

func (lr linkring) Link(i int) (link twentynine.Link, ok bool) {
	if i > twentynine.TwentyNine {
		return
	}
	link = lr.links[(twentynine.TwentyNine-i-1+lr.i)%twentynine.TwentyNine]
	ok = true
	return
}

func (lr linkring) Links() (links []twentynine.Link) {
	links = make([]twentynine.Link, 0)
	for i := 0; i < twentynine.TwentyNine; i++ {
		link, _ := lr.Link(i)
		if link.Headline == "" {
			break
		}
		links = append(links, link)
	}
	return
}
