// ナイーブベイズによるマルチクラス分類器。
// Naive Bayes Classifier for multi classes

package nbclsf

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"math"
	"os"
)

// TypeClsf : 分類器の型
type TypeClsf struct {
	// すべての文書の数
	NumAllDocs int `json:"num_all_docs"`

	// 各カテゴリごとの文書の数
	NumDocsCat map[TypeCat]int `json:"num_docs_cat"`

	// カテゴリのリスト
	CatList map[TypeCat]int `json:"cat_list"`

	// NumAllWords : すべての単語の個数（重複なし）
	NumAllWords int `json:"num_all_words"`

	// WordList : 単語のリスト
	WordList map[TypeWord]int `json:"word_list"`

	// NumWordInCat: ある単語があるカテゴリに含まれる個数
	NumWordInCat map[TypeCat]map[TypeWord]int `json:"num_word_in_cat"`

	// 例：NumWordInCat[cat][word] = 単語 word がカテゴリ cat に含まれる個数

	// NumAllWordsInCat : カテゴリに含まれる個数
	NumAllWordsInCat map[TypeCat]int `json:"num_all_words_in_cat"`

	// 例：NumAllWordsInCat[cat] = カテゴリ cat に含まれる単語の個数

}

// NewClsf : 新しい分類器の作成。multi class 用
func NewClsf() (clsf *TypeClsf) {
	clsf = &TypeClsf{
		NumDocsCat:       map[TypeCat]int{},
		CatList:          map[TypeCat]int{},
		WordList:         map[TypeWord]int{},
		NumWordInCat:     map[TypeCat]map[TypeWord]int{},
		NumAllWordsInCat: map[TypeCat]int{},
	}
	return
}

// Load : 分類器データのファイルからの読み出し
func Load(inFile string) (c *TypeClsf, err error) {
	var bytes []byte
	bytes, err = ioutil.ReadFile(inFile)
	if err != nil {
		return
	}
	var d TypeClsf
	err = json.Unmarshal(bytes, &d)
	if err != nil {
		return
	}

	c = &d
	return
}

// Save : 分類器データのファイルへの保存
func (c *TypeClsf) Save(outFile string) (err error) {

	var w *os.File
	w, err = os.Create(outFile)
	if err != nil {
		return
	}
	defer w.Close()
	var b []byte
	b, err = json.Marshal(c)
	if err != nil {
		return
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "  ")
	if err != nil {
		return
	}
	_, err = out.WriteTo(w)
	return
}

// ProbCat : P(C) を求める関数。
// カテゴリに含まれる文書の全文書数に対する割合。
func (c *TypeClsf) ProbCat(cat TypeCat) float64 {
	return float64(c.NumDocsCat[cat]+1) / float64(c.NumAllDocs+len(c.CatList))
	//	return float64(c.NumDocsCat[cat]) / float64(c.NumAllDocs)
}

// ProbWordGivenCat : P(w|C) を求める関数。
// 単語 word の単語がカテゴリ cat に含まれる確率。
func (c *TypeClsf) ProbWordGivenCat(word TypeWord, cat TypeCat) float64 {
	// ラプラススムージング
	// 　　　　カテゴリに含まれるwordの個数　＋　　　１
	// ＝　ーーーーーーーーーーーーーーーーーーーーーーーーーー
	// 　　　カテゴリに含まれる全単語の個数　＋　全単語の個数
	num := c.NumWordInCat[cat][word] + 1
	sum := c.NumAllWordsInCat[cat] + c.NumAllWords
	return float64(num) / float64(sum)
}

// LogProbDocGivenCat : log P(D|C) を求める関数。
// 文書 doc がカテゴリ cat に含まれる確率の対数
func (c *TypeClsf) LogProbDocGivenCat(doc TypeDoc, cat TypeCat) (r float64) {
	r = 0.0
	for word, num := range doc {
		r += float64(num) * math.Log(c.ProbWordGivenCat(word, cat))
	}
	return
}

// LogProbCatGivenDoc : log P(C|D) を求める関数。
// カテゴリ cat に属する文書群に文書 doc が含まれる確率の対数
func (c *TypeClsf) LogProbCatGivenDoc(doc TypeDoc, cat TypeCat) (r float64) {
	//fmt.Println("#", "LogProbCatGivenDoc", "doc =", doc, "cat =", cat)
	r = math.Log(c.ProbCat(cat)) + c.LogProbDocGivenCat(doc, cat)
	//fmt.Println("#", "LogProbCatGivenDoc", "r =", r)
	return
}

// Predict : 与えられた文書のカテゴリを推定する
func (c *TypeClsf) Predict(doc TypeDoc) (cat TypeCat) {
	cats := []TypeCat{}
	for c := range c.CatList {
		cats = append(cats, c)
	}
	cat = cats[0]
	maxValue := c.LogProbCatGivenDoc(doc, cat)
	//fmt.Println("#", "*PredictCat", "cat =", CatList[0], "result =", maxValue)
	for i := 1; i < len(cats); i++ {
		result := c.LogProbCatGivenDoc(doc, cats[i])
		//fmt.Println("#", "PredictCat", "cat =", CatList[i], "result =", result)
		if result > maxValue {
			maxValue = result
			cat = cats[i]
		}
	}
	//fmt.Println("#", "PredictCat", "cat =", cat)
	return
}

// Train : 文書 doc をカテゴリ cat として学習する
func (c *TypeClsf) Train(doc TypeDoc, cat TypeCat) {

	// カテゴリを追加
	c.addCat(cat)
	// すべての文書の数をインクリメント
	c.NumAllDocs++

	// カテゴリ cat の文書の数をインクリメント
	c.NumDocsCat[cat] = c.NumDocsCat[cat] + 1

	for word, num := range doc {
		// 単語 word を単語リストに追加。
		c.WordList[word] = c.WordList[word] + num // 単語の出現回数となる
		// 単語 word がカテゴリ cat に含まれる個数をインクリメント
		c.NumWordInCat[cat][word] = c.NumWordInCat[cat][word] + num
		// カテゴリ cat に含まれる単語の個数をインクリメント
		c.NumAllWordsInCat[cat] = c.NumAllWordsInCat[cat] + num
	}

	// すべての単語の個数を計算
	c.NumAllWords = len(c.WordList)
}

// addCat : カテゴリの追加
func (c *TypeClsf) addCat(cat TypeCat) {

	// カテゴリ cat が既出の場合、なにもしない。
	if c.isKnownCat(cat) {
		return
	}

	// 以下、カテゴリ cat が初出の場合
	// カテゴリリストに追加
	c.CatList[cat] = len(c.CatList)
	// カテゴリの文書を初期化
	c.NumWordInCat[cat] = map[TypeWord]int{}
}

func (c TypeClsf) isKnownCat(cat TypeCat) bool {
	for knownCat := range c.CatList {
		if cat == knownCat {
			return true
		}
	}
	return false
}
