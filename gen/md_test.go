package main

import (
	"fmt"
	"testing"
)

func TestFindCate(t *testing.T) {
	sub := `abc`
	content := fmt.Sprintf(`@(%s) [f] fjja;df`, sub)
	c := findCate(content)
	if sub != c {
		t.Errorf(`%s not equal %s`, sub, c)
	}
	content = fmt.Sprintf(`@(jjj [fla]`)
	c = findCate(content)
	if c != "" {
		t.Errorf(`c should equal ""`)
	}
}

func TestFindTags(t *testing.T) {
	tag := "def"
	content := fmt.Sprintf(`@(fff) [%s]`, tag)
	tags := findTags(content)
	var isFound bool
	for _, t := range tags {
		if t == tag {
			isFound = true
		}
	}
	if !isFound {
		t.Errorf(`%s should contain %s`, tags, tag)
	}
}
