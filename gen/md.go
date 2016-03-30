package main

import (
    "regexp"
    "github.com/russross/blackfriday"
    "io/ioutil"
    "github.com/satori/go.uuid"
    "path"
    "fmt"
    "strings"
)

type HtmlGener struct {
    mds     []Md
    tagMap  map[string][]Md
    cateMap map[string][]Md
    rootDir string
}

func NewHtmlGener(mds []Md, rootDir string) *HtmlGener {
    return &HtmlGener{
        mds: mds,
        rootDir: rootDir,
    }
}

func (self *HtmlGener) Gen() error {
    if err := self.genBase(); err != nil {
        return err
    }
    return nil
}

func (self *HtmlGener) genBase() error {
    indexLinks := make([]string, len(self.mds))
    for i, m := range self.mds {
        htmlName := uuid.NewV4().String()
        htmlPath := path.Join(self.rootDir, htmlName + ".html")
        linkContent := htmlName
        linkHref := fmt.Sprintf(`./html/%s.html`, htmlName)
        indexLinks[i] = fmt.Sprintf(`<a href="%s">%s</a>`, linkHref, linkContent)
        err := ioutil.WriteFile(htmlPath, m.Html(), 0644)
        if err != nil {
            return err
        }
    }
    if err := ioutil.WriteFile(path.Join(self.rootDir, "..", "index.html"),
        []byte(strings.Join(indexLinks, "")), 0644); err != nil {
        return err
    }
    return nil
}

func (self *HtmlGener) genTags() error {
    //for k, v := range self.tagMap {
    //    htmlName := path.Join(self.rootDir, fmt.Sprintf(`tag_%s.html`, k))
    //
    //}
    return nil
}

func (self *HtmlGener) genCate() error {
    return nil
}

func (self *HtmlGener) genIndex() error {
    return nil
}

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
        tags = append(tags, v[1:len(v) - 1])
    }
    return
}
