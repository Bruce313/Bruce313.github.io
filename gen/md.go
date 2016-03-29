package main

import (
	"regexp"
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
	return nil
}

func findCate(content string) (cate string) {
	f := reCate.FindStringSubmatch(content)
	if len(f) < 2 {
		return
	}
	return f[1]
}

func findTags(content string) (tags []string) {
	out := reTags.FindStringSubmatch(content)[1:]
	for _, v := range out {
		tags = append(tags, v[1:len(v)-1])
	}
	return
}
