package main

import (
    "github.com/satori/go.uuid"
	"io/ioutil"
	"log"
    "path"
    "os"
)

const md_DIR = "./md"

func main() {
    //(md files in md dir) -> md structs
    mds, err := ParseMd(md_DIR)
    if err != nil {
        log.Fatal(`parse all md struct:`, err)
    }
    //generate html files of `content html`
    for _, m := range mds {
        html := m.Html()
        htmlName := uuid.V4()
    }
}

func ParseMd(mdDir string) (mds []Md, err error) {
	files, err := ioutil.ReadDir(mdDir)
	if err != nil {
        return
	}
    var mds []Md
	for _, f := range files {
        var cate string //dir > @(cate)
		if f.IsDir() {
            mds = append(mds, ParseDirInMdDir(path.Join(mdDir, f.Name()))...)
            continue
        }
        regularFile, err := os.Open(f)
        if err != nil {
            return
        }
        content, err := ioutil.ReadAll(regularFile)
        cate = findCate(content)
        tags := findTags(content)
        mds = append(mds, Md {
            Cate: cate,
            Tags: tags,
            Content: content,
        })
        regularFile.Close()
	}
}

func ParseDirInMdDir(p string) (mds []Md, err error) {
    mdFiles, err := getFileInDir(p)
    if err != nil {
        return
    }
    for _, mf := range mdFiles {
        f, err := os.Open(mf)
        if err != nil {
            return
        }
        content, err := ioutil.ReadAll(f)
        if err != nil {
            return
        }
        f.Close()
        tags := findTags(content)
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
            fsTmp, err := getFileInDir(pathNow)
            if err != nil {
                return
            }
            files = append(files, fsTmp...)
            continue
        }
        files = append(files, pathNow)
    }
    return
}
