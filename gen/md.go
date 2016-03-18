package main

type Md struct {
	Tags    []string
	Cate    string
	content []byte
}

func (self *Md) Html() []byte {
}

func NewMd(fileName string) (m *Md, err error) {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}
	cates := re
}
