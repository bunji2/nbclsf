![C1...Cn](https://latex.codecogs.com/gif.latex?C_1,&space;...,&space;C_i,&space;...,&space;C_n)

![](https://latex.codecogs.com/gif.latex?P(C_i|D)=\frac{P(C_i)P(D|C_i)}{P(D)})

#ナイーブベイズ分類器 (Naive Bayes Classifier) の GoLang による実装


最終更新日：2020/01/07
AUTHOR：Bunji Square




##1. はじめに

　「分類器」とは、ある文書を入力するとその文書が属するカテゴリを推定し出力するシステムを指すものとする。本稿はこの分類器をナイーブベイズを用いたモデルで示し、そして GoLang による実装例を示す。

　本稿では多項分布のパラメータ推定として最尤推定値 (MLE) と 期待事後確率推定値 (EAP 推定値) を使用するが、具体的な計算については別紙を参照されたい。



##2. 推定

　本節では、与えられた文書に対するカテゴリの推定について示す。



##2.1 PredictCat (暫定版)

　ある文書  が与えられたときに、それがカテゴリ  に属する確率を  と表す。 

　カテゴリが  個あり、 のいずれかのとき、ある文書  の属するカテゴリは  が最大となる  で与える。

例：ニュースのカテゴリ群が 社会,政治,国際,スポーツ,科学 の 5 種類とする。
あるニュースの文書  が与えられそれぞれのカテゴリの確率が以下のとき、  の属するカテゴリは   が最大の値となる「スポーツ」となる。

表 2.1.1

社会
政治
国際
スポーツ
科学

0.12
0.17
0.05
0.51
0.15


　 を求める関数を ProbCatGivenDoc とする。

　ある文書がどのカテゴリに属するかを推定する関数 PredictCat はすべてのカテゴリについて ProbCatGivenDoc を計算していき、その値が最大となるときのカテゴリを決定すればよいので、実装は次のようになる。


ソースコード 1
```
// catList : カテゴリのリスト
var catList []TypeCat

// PredictCat : 与えられた文書 doc のカテゴリを推定する関数。暫定版。
func PredictCat(doc TypeDoc) (cat TypeCat) {
	cat = catList[0]
	maxValue := ProbCatGivenDoc(doc, cat)
	for i:=1; i&lt;len(catList); i++ {
		result := probCatGivenDoc(doc, catList[i])
		if result &gt; maxValue {
			maxValue = result
			cat = catList[i]
		}
	}
	return
}
```

　上の実装は暫定版である。後ほど最終版を示す。


##2.2 ProbCatGivenDoc (暫定版) --- 

　ある文書  がカテゴリ  に属する確率   はベイズの定理により次のように表される。



　ここで、 はそれぞれ次の確率を示す。

表 2.2.1

カテゴリー  である確率

カテゴリー  に文書  が含まれる確率

文書  が成立する確率

　また  は以下を満たす。



　 は  に関係なく固定の値であること、また、カテゴリの推定は  の大小関係のみに基づいていることから、次の比例関係に注目すればよい。



　 を計算する関数 ProbCat と  を計算する関数 ProbDocGivenCat があるとすれば、 の「比」を計算する関数 ProbCatGivenDoc の実装は次のような形になる。


ソースコード 2
```
// ProbCatGivenDoc : 文書 doc がカテゴリ cat に含まれる確率の比を求める関数。
// 暫定版。
func ProbDocGivenCat(doc TypeDoc, cat TypeCat) float64 {
	return ProbCat(cat) * ProbDocGivenCat(doc, cat)
}
```

　上の実装は暫定版である。後ほど最終版を示す。



##2.3 ProbCat --- 

　 はカテゴリー  である確率である。ここでは全文書における、カテゴリー  に属す文書の割合とみなし、単純に文書の個数の割合で考えることにする。

 を求める関数 ProbCat は次の実装で与えることができる。

ソースコード 3
```
var numAllDocs int // すべての文書の数
var numDocsCat map[TypeCat]int // 各カテゴリごとの文書の数

// ProbCat : カテゴリ cat の確率＝カテゴリ cat の文書の全文書に対する割合
func ProbCat(cat TypeCat) float64 {
    return float64(numDocsCat[cat]) / float64(numAllDocs) 
}
```


##2.4 ProbDocGivenCat --- 


はカテゴリー  （に属する文書群）に文書  が含まれる確率である。

　しかし文書  がカテゴリーに属する文書と一致するケースがほぼないと考えられるため、文書数だけでは計算することができない。


　ここで文書  に出現する単語群  に注目し、次の仮定をおくことにする。

文書は単語の並びである
文書中にある単語が現れる確率は他の単語が現れる確率に依存せず独立である
文書中にある単語が現れる確率は文書中の位置に依存しない

　確率  を、文書  に含まれる単語  がカテゴリ  に出現する確率とし、単語  が文書  に出現する個数を  とすれば、  は次のような多項分布関数で表される。

 


　確率  を求める関数を ProbWordGivenCat とするとき、 を求める関数 ProbDocGivenCat は次のようになる。

ソースコード 4
```
// TypeDoc : 文書の型。各単語の出現個数のmap。
type TypeDoc map[TypeWord]int
// 例：doc[word] = 単語 word が文書 doc に含まれる個数

// ProbDocGivenCat : 文書 doc がカテゴリ cat に含まれる確率
func ProbDocGivenCat(doc TypeDoc, cat TypeCat) (r float64) {
	r = 1.0
	for word, num := range doc {
		r *= math.Pow(ProbWordGivenCat(word, cat), float64(num))
	}
	return
}
```




##2.5 ProbWordGivenCat --- 

　単語の出現確率  は のパラメータである。実測した単語  の出現数をもとにこのパラメータを推定するにあたり、まず最尤推定値 (Maximum Likelihood Estimator) を使う。

 を下に示すような尤度関数  とみなし、これが最大になる を推定することになる。



　確率  を推定する関数 ProbWordGivenCat の実装を考えるが、ここで２つの案が考えられる。

案1：カテゴリー  に属する文書に出現する単語群に、文書  に出現する単語群がどれだけ含まれるかで考える。

j= Ciに属する文書における単語jの出現回数の合計Ciに属する文書における全単語の出現回数の合計
    
案2：文書  に含まれる各単語を含む文書がどれだけカテゴリー  に属しているかで考える。
j= 単語jを含みかつCiに属する文書の個数の合計Ciに属する文書の個数の合計

　一つの文書には異なる単語が複数含まれることが簡単に予想されることから、案 2 では  の条件を満たすことができない。ここでは案 1 を実装する。

　numWordInCat[][] を文書  に含まれる単語  がカテゴリ  に含まれる個数とすれば、 を計算する関数 ProbWordGivenCat は次のように与えることができる。

ソースコード 5
```
// numWordInCat: ある単語があるカテゴリに含まれる個数
var numWordInCat map[TypeCat]map[TypeWord]int
// 例：numWordInCat[cat][word] = 単語 word がカテゴリ cat に含まれる個数

// numAllWordsInCat : カテゴリに含まれる個数
var numAllWordsInCat map[TypeCat]int
// 例：numAllWordsInCat[cat] = カテゴリ cat に含まれる単語の個数

// ProbWordGivenCat : 単語 word がカテゴリ cat に含まれる確率
func ProbWordGivenCat (word TypeWord, cat TypeCat) float64 {
    return float64(numWordInCat[cat][word])/float64(numAllWordsInCat[cat])
}
```

##2.6 ProbWordGivenCat (スムージング拡張版) --- 

　上の実装では、文書の中に一つでもカテゴリ  に含まれない単語が存在すると、他の単語の確率が高いものだったとしても、全体として  が 0 となってしまうという問題がある。
　これを回避するため「加算スムージング」（あるいは「ラプラススムージング」）を使う。重複のない全単語の個数を  とする。 


j= Ciに属する文書における単語jの出現回数の合計  + 1 Ciに属する文書における全単語の出現回数の合計  +  m

　これは期待事後確率推定値 (Expected a Posterior Estimator; EAP 推定値) に相当する。
この推定値は特に標本数が少ない場合に効果があり、標本数が増えるにつれて先の最尤推定値に近づいていく。

ソースコード 6

```
var numAllWords int // 全単語数

// ProbWordGivenCat : 単語 word がカテゴリ cat に含まれる確率(スムージング拡張版)
func ProbWordGivenCat (word TypeWord, cat TypeCat) float64 {
    num := float64(numWordInCat[cat][word] + 1)
    sum := float64(numAllWordsInCat + numAllWords)
    return num/sum
}
```

　上記スムージングを施しても  の条件を満たすことに注意。

##2.7 LogProbDocGivenCat --- 

　上の関数 ProbDocGivenCat の実装では、単語数が多いと分母の値が非常に大きくなりアンダーフローが起きる恐れがあるので、これを回避すべく対数をとる。



　このようにして ProbDocGivenCat の対数を計算する LogProbDocGivenCat の実装は継のようになる。





ソースコード 6

```
// TypeDoc : 文書の型。各単語の出現個数のmap。
type TypeDoc map[TypeWord]int
// 例：doc[word] = 単語 word が文書 doc に含まれる個数

// LogProbDocGivenCat : 文書 doc がカテゴリ cat に含まれる確率の対数
func LogProbDocGivenCat(doc TypeDoc, cat TypeCat) (r float64) {
	r = 0.0
	for word, num := range doc {
		r += float64(num) * math.Log(ProbWordGivenCat(word, cat))
	}
	return
}
```


##2.7 LogProbCatGivenDoc --- 

　以上を踏まえると、冒頭に示した関数 ProbCatGivenDoc のアンダーフローを考慮した対数版 LogProbCatGivenDoc は次のようになる。

ソースコード 7

```
// LogProbCatGivenDoc : 文書 doc がカテゴリ cat に含まれる確率の比の対数
func LogProbDocGivenCat(doc TypeDoc, cat TypeCat) float64 {
	return math.Log(ProbCat(cat)) + LogProbDocGivenCat(doc, cat)
}
```



##2.8 Predict (最終版)

　冒頭に示した関数 PredictCat は LogProbCatGivenDoc を使って次のように実装される。

ソースコード 8
```
// PredictC : 与えられた文書 doc のカテゴリを推定する
func Predict(doc TypeDoc) (cat TypeCat) {
	cat = catList[0]
	maxValue := LogProbCatGivenDoc(doc, cat)
	for i:=1; i&lt;len(catList); i++ {
		result := LogProbCatGivenDoc(doc, catList[i])
		if result &gt; maxValue {
			maxValue = result
			cat = catList[i]
		}
	}
	return
}
```

##3. 学習

　本節では 2 節に示したカテゴリの推定に必要なデータを作成するための「学習」について示す。


##3.1 Train

　教師データとして、文書 doc とそのカテゴリ cat が与えられたとする。使用する変数に対して、以下の表に示す処理を実施する必要がある。

表 3.1.1
変数
概要
処理
直接的に依存する関数
var catList []TypeCat
カテゴリのリスト
cat が初出のときのみリストに追加
PredictCat
var numAllDocs int
すべての文書の数
numAllDocs の値をインクリメント
ProbCat
var numDocsCat map[TypeCat]int
各カテゴリごとの文書の数
numDocsCat[cat] の値をインクリメント
ProbCat
var numWordInCat map[TypeCat]map[TypeWord]int
ある単語があるカテゴリに含まれる個数
文書 doc に含まれるすべての単語 word について、numWordInCat[cat][word] の値をインクリメント
ProbWordGivenCat
var numAllWordsInCat map[TypeCat]int
カテゴリに含まれる単語の個数
文書 doc に含まれる単語の個数だけ、numAllWordsInCat[cat] の値をインクリメント
ProbWordGivenCat
var numAllWords int
全単語数 (重複なし)
文書 doc について wordList を更新後、len(wordList) の値を代入。つまり、これまでに出現した重複のないすべての単語の個数を代入
ProbWordGivenCat
var wordList map[TypeWord]int
単語のリスト、各単語の出現数
文書 doc に含まれるすべての単語 word について、wordList[word] の値をインクリメント
Train


　与えられた文書とカテゴリで学習する関数 Train の実装例は次のようになる。


ソースコード 9
```
// wordList : 単語のリスト
var wordList map[TypeWord]int

// Train : 文書 doc をカテゴリ cat として学習する
func Train(doc TypeDoc, cat TypeCat) {

	// カテゴリ cat が初出かどうか検査する
	_, ok := numDocsCat[cat]
	if !ok { // カテゴリ cat が初出の場合
		// カテゴリリストに追加
		catList = append(catList, cat)
		// カテゴリの文書を初期化
		numWordInCat[cat] = map[TypeWord]int{}
	}

	// すべての文書の数をインクリメント
	numAllDocs++

	// カテゴリ cat の文書の数をインクリメント
	numDocsCat[cat] = numDocsCat[cat] + 1

	// 文書 doc に出現する単語 word についてそれぞれ処理
	for word, num := range doc {
		// カテゴリ cat に含まれる単語の個数をインクリメント
		numAllWordsInCat[cat] = numAllWordsInCat[cat] + num  
		// 単語 word がカテゴリ cat に含まれる個数をインクリメント
		numWordInCat[cat][word] = numWordInCat[cat][word] + num
		// 単語 word を単語リストに追加。
		wordList[word] = wordList[word] + num // 単語の出現回数
	}

	// すべての単語の重複のない個数を計算
	numAllWords = len(wordList)
}
```

##4. 評価

　本稿に示したカテゴリの推定方式について、Livedoor ニュースコーパスを用いて精度を評価した。Livedoor ニュースコーパスには９つのカテゴリがあり、7367 の文書からなる。今回はこのコーパスの 80% を教師データとし、残り 20% を試験データとした場合の、指標を算出した。

　9 つのカテゴリの Precision, Recall, F-Measure, Accuracy はそれぞれ次のようになった。

表 4.1
カテゴリ









Precision
0.958084
0.898734
0.860465
0.930233
0.889535
0.932886
0.857143
0.937500
0.952941
Recall
1.000000
0.940397
0.907975
0.909091
0.987097
0.822485
0.814815
0.967742
0.885246
F-Measure
0.978593
0.919094
0.883582
0.919540
0.935780
0.874214
0.835443
0.952381
0.917847
Accuracy
0.995251
0.983039
0.973541
0.981004
0.985753
0.972863
0.964722
0.989824
0.980326

　また全体の指標は次のようになった。

表 4.2
Micro Precision
0.913161
Micro Recall
0.913161
Micro F-Measure
0.913161
Macro Precision
0.913058
Macro Recall
0.914983
Macro F-Measure
0.914019
Overall Accuracy
0.980703

　単純な実装内容にも関わらず、90% 超の精度でカテゴリ推定できることがわかった。

##5. おわりに

　本稿では、文書のカテゴリを推定する分類器をナイーブベイズを用いたモデルで示し、そして GoLang による実装例を示した。また、Livedoor ニュースコーパスを用いて評価した結果を示した。今後はマルチラベル分類器に拡張していく予定である。



##参考

Livedoor ニュースコーパス
https://www.rondhuit.com/download.html#ldcc


![C1...Cn](https://latex.codecogs.com/gif.latex?C_1,&space;...,&space;C_i,&space;...,&space;C_n)

![](https://latex.codecogs.com/gif.latex?P(C_i|D)=\frac{P(C_i)P(D|C_i)}{P(D)})