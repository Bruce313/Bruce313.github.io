package main

import (
	"io/ioutil"
	"log"
    "path"
    "os"
)

const (
    md_DIR = "./md"
    html_DIR = "./html"
)

func main() {
    //(md files in md dir) -> md structs
    mds, err := ParseMd(md_DIR)
    if err != nil {
        log.Fatal(`parse all md struct:`, err)
    }
    htmlGener := NewHtmlGener(mds, html_DIR)
    if err := htmlGener.Gen(); err != nil {
        log.Fatal(`gen html:`, err)
    }
    //generate html files of `content html`
    //for _, m := range mds {
    //    html := m.Html()
    //    htmlName := uuid.NewV4().String()
    //    htmlPath := path.Join(html_DIR, htmlName + ".html")
    //    err = ioutil.WriteFile(htmlPath, html, 0644)
    //    if err != nil {
    //        log.Fatal(`write html file:`, err)
    //    }
    //}
    //gen nav by tags
    //gen nav html by cate
    //gen index.html
}

func ParseMd(mdDir string) (mds []Md, err error) {
	files, err := ioutil.ReadDir(mdDir)
	if err != nil {
        return
	}
	for _, f := range files {
        var cate string //dir > @(cate)
		if f.IsDir() {
            mdsInDir, errParse := ParseDirInMdDir(path.Join(mdDir, f.Name()))
            if errParse != nil {
                return nil, errParse
            }
            mds = append(mds, mdsInDir...)
            continue
        }
        regularFile, errOpen := os.Open(path.Join(mdDir, f.Name()))
        if errOpen != nil {
            return nil, errOpen
        }
        content, errReadAll := ioutil.ReadAll(regularFile)
        if errReadAll != nil {
            return nil, errReadAll
        }
        cate = findCate(string(content))
        tags := findTags(string(content))
        mds = append(mds, Md {
            Cate: cate,
            Tags: tags,
            Content: content,
        })
        regularFile.Close()
	}
    return
}

func ParseDirInMdDir(p string) (mds []Md, err error) {
    mdFiles, err := getFileInDir(p)
    if err != nil {
        return
    }
    for _, mf := range mdFiles {
        f, errOpen := os.Open(mf)
        if errOpen != nil {
            return nil, errOpen
        }
        content, errReadAll := ioutil.ReadAll(f)
        if errReadAll != nil {
            return nil, errReadAll
        }
        f.Close()
        tags := findTags(string(content))
        mds = append(mds, Md {
            Cate: path.Base(p),
            Tags: tags,
            Content: content,
        })
    }
    return
}

func getFileInDir(dirName string) (files []string, err error) {
    fs, err := ioutil.ReadDir(dirName)
    if err != nil {
        return
    }
    for _, f := range fs {
        pathNow := path.Join(dirName, f.Name())
        if f.IsDir() {
            fsTmp, errGetFile := getFileInDir(pathNow)
            if errGetFile != nil {
                return nil, errGetFile
            }
            files = append(files, fsTmp...)
            continue
        }
        files = append(files, pathNow)
    }
    return
}
