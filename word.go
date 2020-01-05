// テキストからの単語の切り出し、ストップワードなど。

package main

import (
	"strings"

	"github.com/bunji2/mecab"
)

var stopWordPrefixes = []string{
	"接頭辞",
	"接尾辞",
	"記号",
	"助詞",
	"助動詞",
	"特殊",
	"判定詞",
}
var stopWords = []string{
	"動詞_し",
}

var mp *mecab.Proc

// InitializeWordExtraction : 初期化
func InitializeWordExtraction() (err error) {
	if mp == nil {
		mecab.Init(mecab.Config{})
		mp, err = mecab.NewProc()
	}
	return
}

// MakeDocFromText : テキストから文書データ作成
func MakeDocFromText(text string) (r TypeDoc) {
	r = TypeDoc{}
	words := mp.Write(text)
	for _, word := range words {
		if !isStopWord(word) {
			tword := TypeWord(word)
			r[tword] = r[tword] + 1
		}
	}
	return
}

// isStopWord : 与えられた単語がストップワードなら true
func isStopWord(word string) bool {
	for _, prefix := range stopWordPrefixes {
		if strings.HasPrefix(word, prefix) {
			return true
		}
	}
	for _, stopWord := range stopWords {
		if word == stopWord {
			return true
		}
	}
	return false
}
