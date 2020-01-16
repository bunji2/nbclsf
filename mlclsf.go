// ナイーブベイズによるマルチラベル分類器の実験的実装
// An experimental implementation of Naive Bayes Classifier for multi-label
// [TODO] 各ラベルごとの TypeClsf の配列として実装すると、それぞれ同じ WordList を保持してしまい無駄が多い。効率的な実装を今後行う。

package nbclsf

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// TypeMLClsf : マルチラベル分類器の型
type TypeMLClsf struct {
	// マルチラベルの次数
	DegreeMLabel int `json:"degree_mlabel"`
	// 各ラベルごとの分類器
	Clsfs []*TypeClsf `json:"clsfs"`
}

// NewMLClsf : 新しいマルチラベル分類器の作成。
func NewMLClsf(degreeMLabel int) (r *TypeMLClsf) {
	clsfs := make([]*TypeClsf, degreeMLabel)
	for i := 0; i < degreeMLabel; i++ {
		clsfs[i] = NewClsf()
	}
	r = &TypeMLClsf{
		DegreeMLabel: degreeMLabel,
		Clsfs:        clsfs,
	}
	return
}

// LoadMLClsf : マルチラベル分類器のファイルからの読み出し
func LoadMLClsf(inFile string) (c *TypeMLClsf, err error) {
	var bytes []byte
	bytes, err = ioutil.ReadFile(inFile)
	if err != nil {
		return
	}
	var d TypeMLClsf
	err = json.Unmarshal(bytes, &d)
	if err != nil {
		return
	}

	c = &d
	return
}

// Save : マルチラベル分類器のファイルへの保存
func (c *TypeMLClsf) Save(outFile string) (err error) {

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

// Train : 文書 doc をマルチラベル mLabel として学習する
func (c *TypeMLClsf) Train(doc TypeDoc, mLabel TypeMultiLabel) (err error) {

	if len(mLabel) != c.DegreeMLabel {
		err = fmt.Errorf("degree of multi-label is missmatch")
		return
	}
	for i := 0; i < c.DegreeMLabel; i++ {
		label := TypeCat(mLabel[i])
		c.Clsfs[i].Train(doc, label)
	}

	return
}

// Predict : 与えられた文書のマルチラベルを推定する
func (c *TypeMLClsf) Predict(doc TypeDoc) (mlabel TypeMultiLabel) {
	mlabel = make(TypeMultiLabel, c.DegreeMLabel)
	for i := 0; i < c.DegreeMLabel; i++ {
		mlabel[i] = int(c.Clsfs[i].Predict(doc))
	}
	return
}
