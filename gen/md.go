package main

import (
	"regexp"
    "github.com/russross/blackfriday"
)



type Md struct {
	Tags    []string
	Cate    string
	Content []byte
}

var (
	reCate = regexp.MustCompile("@\\((\\w+)\\)")
	reTags = regexp.MustCompile("@\\(\\w+\\)\\s*(\\[\\w+\\])+")
)

func (self *Md) Html() []byte {
    return blackfriday.MarkdownCommon(self.Content)
}

func findCate(content string) (cate string) {
	f := reCate.FindStringSubmatch(content)
	if len(f) < 2 {
		return
	}
	return f[1]
}

func findTags(content string) (tags []string) {
	outs := reTags.FindStringSubmatch(content)
    if len(outs) < 2 {
        return nil
    }
    out := outs[1:]
	for _, v := range out {
		tags = append(tags, v[1:len(v)-1])
	}
	return
}
