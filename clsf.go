// Naive Bayes Classifier
package nbclsf

import (
	"math"
)

// TypeDoc : 文書データの型
type TypeDoc map[TypeWord]int

//type TypeDoc []TypeWord

// TypeWord : 単語データの型
type TypeWord string

// TypeCat : カテゴリデータの型
type TypeCat string

// TypeProbWordGivenCat : P(w|C)の関数の型
type TypeProbWordGivenCat func(TypeWord, TypeCat) float64

// TypeProbDocGivenCat : P(D|C)を求める関数の型
type TypeProbDocGivenCat func(TypeDoc, TypeCat, TypeProbWordGivenCat) float64

// TypeProbCatGivenDoc : P(C|D)を求める関数の型
type TypeProbCatGivenDoc func(TypeDoc, TypeCat, TypeProbWordGivenCat) float64

// InitializeClsf : 初期化
func InitializeClsf() (err error) {
	numDocsCat = map[TypeCat]int{}
	catList = map[TypeCat]int{}
	wordList = map[TypeWord]int{}
	numWordInCat = map[TypeCat]map[TypeWord]int{}
	numAllWordsInCat = map[TypeCat]int{}
	return
}

var numAllDocs int             // すべての文書の数
var numDocsCat map[TypeCat]int // 各カテゴリごとの文書の数

//var catList []TypeCat // カテゴリのリスト
var catList map[TypeCat]int // カテゴリのリスト

func numCats() int {
	return len(catList)
}

func catID(cat TypeCat) int {
	return catList[cat]
}

func catName(catID int) TypeCat {
	for cat, cID := range catList {
		if catID == cID {
			return cat
		}
	}
	return TypeCat("unknown")
}

// numAllWords : すべての単語の個数（重複なし）
var numAllWords int

// wordList : 単語のリスト
var wordList map[TypeWord]int

// numWordInCat: ある単語があるカテゴリに含まれる個数
var numWordInCat map[TypeCat]map[TypeWord]int

// 例：numWordInCat[cat][word] = 単語 word がカテゴリ cat に含まれる個数

// numAllWordsInCat : カテゴリに含まれる個数
var numAllWordsInCat map[TypeCat]int

// 例：numAllWordsInCat[cat] = カテゴリ cat に含まれる単語の個数

// ProbCat : P(C) を求める関数。
//           カテゴリに含まれる文書の全文書数に対する割合
func ProbCat(cat TypeCat) float64 {
	//return float64(numDocsCat[cat]+1) / float64(numAllDocs+numCats())
	return float64(numDocsCat[cat]) / float64(numAllDocs)
}

/*
// ProbWordGivenCat : P(w|C) を求める関数。
//                    単語 word の単語がカテゴリ cat に含まれる確率
func ProbWordGivenCat(word TypeWord, cat TypeCat) float64 {
	return float64(numWordInCat[cat][word]) / float64(numAllWordsInCat[cat])
}
*/

// ProbWordGivenCat : P(w|C) を求める関数。
//                                 単語 word がカテゴリ cat に含まれる確率(スムージング拡張版)
func ProbWordGivenCat(word TypeWord, cat TypeCat) float64 {
	// ラプラススムージング
	// 　　　　カテゴリに含まれるwordの個数　＋　　　１
	// ＝　ーーーーーーーーーーーーーーーーーーーーーーーーーー
	// 　　　カテゴリに含まれる全単語の個数　＋　全単語の個数
	num := numWordInCat[cat][word] + 1
	//sum := numAllWordsInCat[cat] + len(numWordInCat[cat])
	sum := numAllWordsInCat[cat] + numAllWords
	return float64(num) / float64(sum)
}

/*
// ProbDocGivenCat : P(D|C) を求める関数。
// 文書 doc からなる文書がカテゴリ cat に含まれる確率
func ProbDocGivenCat(doc TypeDoc, cat TypeCat, probWordGivenCat TypeProbWordGivenCat) (r float64) {
	r = 1.0
	for _, word := range doc {
		r *= probWordGivenCat(word, cat)
	}
	return
}
*/

// LogProbDocGivenCat : log P(D|C) を求める関数。
//                      文書 doc がカテゴリ cat に含まれる確率の対数
func LogProbDocGivenCat(doc TypeDoc, cat TypeCat) (r float64) {
	r = 0.0
	for word, num := range doc {
		r += float64(num) * math.Log(ProbWordGivenCat(word, cat))
	}
	return
}

/*
// ProbCatGivenDoc : P(C|D) を求める関数。
// カテゴリ cat に属する文書群に文書 doc が含まれる確率
func ProbCatGivenDoc(doc TypeDoc, cat TypeCat, probWordGivenCat TypeProbWordGivenCat) (r float64) {
	//fmt.Println("#", "ProbCatGivenDoc", "doc =", doc, "cat =", cat)
	r = ProbCat(cat) * ProbDocGivenCat(doc, cat, probWordGivenCat)
	//fmt.Println("#", "ProbCatGivenDoc", "r =", r)
	return
}
*/

// LogProbCatGivenDoc : log P(C|D) を求める関数
// カテゴリ cat に属する文書群に文書 doc が含まれる確率の対数
func LogProbCatGivenDoc(doc TypeDoc, cat TypeCat) (r float64) {
	//fmt.Println("#", "LogProbCatGivenDoc", "doc =", doc, "cat =", cat)
	r = math.Log(ProbCat(cat)) + LogProbDocGivenCat(doc, cat)
	//fmt.Println("#", "LogProbCatGivenDoc", "r =", r)
	return
}

// PredictCat : 与えられた文書のカテゴリを推定する
func PredictCat(doc TypeDoc) (cat TypeCat) {
	cat = catName(0)
	maxValue := LogProbCatGivenDoc(doc, cat)
	//fmt.Println("#", "*PredictCat", "cat =", catList[0], "result =", maxValue)
	for i := 1; i < numCats(); i++ {
		result := LogProbCatGivenDoc(doc, catName(i))
		//fmt.Println("#", "PredictCat", "cat =", catList[i], "result =", result)
		if result > maxValue {
			maxValue = result
			cat = catName(i)
		}
	}
	//fmt.Println("#", "PredictCat", "cat =", cat)
	return
}

// AddCat : カテゴリの追加。既出の場合はなにもしない。
func AddCat(cat TypeCat) int {
	// カテゴリ cat が初出かどうか検査する
	_, ok := catList[cat]
	if !ok { // カテゴリ cat が初出の場合
		// カテゴリリストに追加
		catList[cat] = len(catList)
		// カテゴリの文書を初期化
		numWordInCat[cat] = map[TypeWord]int{}
	}
	return catList[cat]
}

// Train : 文書 doc をカテゴリ cat として学習する
func Train(doc TypeDoc, cat TypeCat) {

	// カテゴリの追加
	AddCat(cat)

	// すべての文書の数をインクリメント
	numAllDocs++

	// カテゴリ cat の文書の数をインクリメント
	numDocsCat[cat] = numDocsCat[cat] + 1

	for word, num := range doc {
		// 単語 word を単語リストに追加。
		wordList[word] = wordList[word] + num // 単語の出現回数
		// 単語 word がカテゴリ cat に含まれる個数をインクリメント
		numWordInCat[cat][word] = numWordInCat[cat][word] + num
		// カテゴリ cat に含まれる単語の個数をインクリメント
		numAllWordsInCat[cat] = numAllWordsInCat[cat] + num
	}

	// すべての単語の個数を計算
	numAllWords = len(wordList)
}
