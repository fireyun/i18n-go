// Code generated by running "go generate" in golang.org/x/text. DO NOT EDIT.

package translations

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
)

type dictionary struct {
	index []uint32
	data  string
}

func (d *dictionary) Lookup(key string) (data string, ok bool) {
	p, ok := messageKeyToIndex[key]
	if !ok {
		return "", false
	}
	start, end := d.index[p], d.index[p+1]
	if start == end {
		return "", false
	}
	return d.data[start:end], true
}

func init() {
	dict := map[string]catalog.Dictionary{
		"en": &dictionary{index: enIndex, data: enData},
		"zh": &dictionary{index: zhIndex, data: zhData},
	}
	fallback := language.MustParse("en")
	cat, err := catalog.NewFromMap(dict, catalog.Fallback(fallback))
	if err != nil {
		panic(err)
	}
	message.DefaultCatalog = cat
}

var messageKeyToIndex = map[string]int{
	"My name is %s, I'm %d years old.": 2,
	"Welcome!":                         0,
	"Who are you? How old are you?":    1,
}

var enIndex = []uint32{ // 4 elements
	0x00000000, 0x00000009, 0x00000027, 0x0000004e,
} // Size: 40 bytes

const enData string = "" + // Size: 78 bytes
	"\x02Welcome!\x02Who are you? How old are you?\x02My name is %[1]s, I'm %" +
	"[2]d years old."

var zhIndex = []uint32{ // 4 elements
	0x00000000, 0x0000000a, 0x00000026, 0x0000004f,
} // Size: 40 bytes

const zhData string = "" + // Size: 79 bytes
	"\x02欢迎！\x02你是谁？你多大了？\x02我的名字叫%[1]s，我%[2]d岁了。"

	// Total table size 237 bytes (0KiB); checksum: 9054E346
