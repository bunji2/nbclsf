# ナイーブベイズ分類器 (Naive Bayes Classifier) の GoLang による実装


* 最終更新日：2020/01/18 (数式部分は 2020/05/23 に更新)
* AUTHOR：Bunji² (Bunji Square) 

## 1. はじめに

　ナイーブベイズを用いた分類器の実装方法について勉強したのでまとめておく。
 
 ここで「分類器」とは、ある文書を入力するとその文書が属するカテゴリを推定し出力するシステムを指すものとする。本稿はこの分類器をナイーブベイズを用いたモデルで示し、そして GoLang による実装例を示す。

　本稿では多項分布のパラメータ推定として最尤推定値 (MLE) と 期待事後確率推定値 (EAP 推定値) を使用するが、具体的な計算については[別紙](multinomial_distribution.md)を参照されたい。



## 2. 推定

　本節では、与えられた文書に対するカテゴリの推定について示す。



## 2.1 PredictCat (暫定版)

　ある文書 [D](fig/texclip20200503235724.png)
 が与えられたときに、それがカテゴリ ![C](fig/texclip20200503235813.png)
 に属する確率を [P(C|D)}](fig/texclip20200503235905.png)
 と表す。 

　カテゴリが *n* 個あり、![C_1,\cdots,C_n](fig/texclip20200504000020.png)
 のいずれかのとき、ある文書 *D* の属するカテゴリは ![P(C_i|D)](fig/texclip20200504000133.png)
 が最大となる  ![C_i](fig/texclip20200504000204.png)
  で与える。

例：ニュースのカテゴリ群が 社会,政治,国際,スポーツ,科学 の 5 種類とする。
あるニュースの文書 ![D](teclip20200503235724.png)
 が与えられそれぞれのカテゴリの確率が以下のとき、属するカテゴリは ![P(C_i|D)](fig/texclip20200504000133.png) が最大の値となる「スポーツ」となる。

> 
> 表 2.1.1
> 
> |![C_i](fig/texclip20200504000204.png)|社会|政治|国際|スポーツ|科学|
> |:---:|:---:|:---:|:---:|:---:|:---:|
> |![P(C_i given D)](fig/texclip20200504000254.png)|0.12|0.17|0.05|0.51|0.15|
> 

 ![P(C_i|D)](fig/texclip20200504000133.png) を求める関数を ProbCatGivenDoc とする。

　ある文書が属するカテゴリを推定する関数 PredictCat はすべてのカテゴリについて ProbCatGivenDoc を計算していき、その値が最大となるときのカテゴリを決定すればよいので、実装は次のようになる。


```golang
// ソースコード 1

// catList : カテゴリのリスト
var catList []TypeCat

// PredictCat : 与えられた文書 doc のカテゴリを推定する関数。暫定版。
func PredictCat(doc TypeDoc) (cat TypeCat) {
	cat = catList[0]
	maxValue := ProbCatGivenDoc(doc, cat)
	for i:=1; i<len(catList); i++ {
		result := probCatGivenDoc(doc, catList[i])
		if result > maxValue {
			maxValue = result
			cat = catList[i]
		}
	}
	return
}
```

　上の実装は暫定版である。後ほど最終版を示す。


## 2.2 ProbCatGivenDoc (暫定版)

　ある文書 ![D_k](fig/texclip20200504000421.png) がカテゴリ ![C_i](fig/texclip20200504000204.png) に属する確率 ![P(C_i|D_k)](fig/texclip20200504000337.png) はベイズの定理により次のように表される。

> 
> 式 2.2.1
> 
> ![P(C_i|D_k)=\frac{P(C_i)P(D_k|C_i)}{P(D_k)}](fig/texclip20200504000515.png)
> 

　ここで、![P(C_i),P(D_k|C_i),P(D_k)](fig/texclip20200504000605.png) はそれぞれ次の確率を示す。

> 
> 表 2.2.1
> 
> |式|意味|
> |:---:|:---|
> |![P(C_i)](fig/texclip20200504000705.png)|カテゴリー ![C_i](fig/texclip20200504000204.png) である確率|
> |![P(D_k given C_i)](fig/texclip20200504000754.png)|カテゴリー ![C_i](fig/texclip20200504000204.png) に文書 ![D_k](fig/texclip20200504000421.png) が含まれる確率|
> |![P(D_k)](fig/texclip20200504000901.png)|文書 ![D_k](fig/texclip20200504000421.png) が成立する確率|
> 

　また ![P(D_k)](fig/texclip20200504000901.png) は以下を満たす。

> 
> 式 2.2.2
> 
> ![P(D_k)=\sum_{i=1}^{n}P(C_i)P(D_k|C_i)](fig/texclip20200504010147.png)
> 

　![P(D_k)](fig/texclip20200504000901.png) は ![C_i](fig/texclip20200504000204.png) に関係なく固定の値であること、また、カテゴリの推定は ![P(C_i)P(D|C_i)](fig/texclip20200504002425.png) の大小関係のみに基づいていることから、次の比例関係に注目すればよい。

> 
> 式 2.2.3
> 
> ![P(C_i|D)\propto{P(C_i)P(D|C_i)}](fig/texclip20200504002452.png)
> 

　![P(C_i)](fig/texclip20200504000705.png) を計算する関数 ProbCat と ![P(D|C_i)](fig/texclip20200504002342.png) を計算する関数 ProbDocGivenCat があるとすれば、 の「比」を計算する関数 ProbCatGivenDoc の実装は次のような形になる。


```go
// ソースコード 2

// ProbCatGivenDoc : 文書 doc がカテゴリ cat に含まれる確率の比を求める関数。
// 暫定版。
func ProbDocGivenCat(doc TypeDoc, cat TypeCat) float64 {
	return ProbCat(cat) * ProbDocGivenCat(doc, cat)
}
```

　上の実装は暫定版である。後ほど最終版を示す。



## 2.3 ProbCat

　
![P(C_i)](fig/texclip20200504000705.png)
 はカテゴリー 
![C_i](fig/texclip20200504000204.png)
 である確率である。ここでは全文書における、カテゴリー 
![C_i](fig/texclip20200504000204.png)
 に属す文書の割合とみなし、単純に文書の個数の割合で考えることにする。

 
![P(C_i)](fig/texclip20200504000705.png)
 を求める関数 ProbCat は次の実装で与えることができる。

```go
// ソースコード 3

var numAllDocs int // すべての文書の数
var numDocsCat map[TypeCat]int // 各カテゴリごとの文書の数

// ProbCat : カテゴリ cat の確率＝カテゴリ cat の文書の全文書に対する割合
func ProbCat(cat TypeCat) float64 {
    return float64(numDocsCat[cat]) / float64(numAllDocs) 
}
```


## 2.4 ProbDocGivenCat

 
![P(D_k|C_i)](fig/texclip20200504000754.png)
 はカテゴリー 
![C_i](fig/texclip20200504000204.png)
 （に属する文書群）に文書 
![D_k](fig/texclip20200504000421.png)
 が含まれる確率である。

　しかし文書 
![D_k](fig/texclip20200504000421.png)
 がカテゴリーに属する文書と一致するケースがほぼないと考えられるため、文書数だけでは計算することができない。


　ここで文書に出現する単語群に注目し、次の仮定をおくことにする。

* 文書は単語の集合である
* 文書中にある単語が現れる確率は他の単語が現れる確率に依存せず独立である
* 文書中にある単語が現れる確率は文書中の位置に依存しない

　確率 
![\theta_{ij}=P(w_j|C_i)](fig/texclip20200504002610.png)
 を、文書 
![D_k](fig/texclip20200504000421.png)
 に含まれる単語 
![w_j](fig/texclip20200504002811.png)
 がカテゴリ 
![C_i](fig/texclip20200504000204.png)
 に出現する確率とし、単語 
![w_j](fig/texclip20200504002811.png)
 が文書 
![D_k](fig/texclip20200504000421.png)
 に出現する個数を 
![n_{kj}](fig/texclip20200504002902.png)
 とすれば、  は次のような多項分布関数で表される。

 > 
 > 式 2.4.1
 > 
 > ![P(D_k|C_i)=\prod_{j=1}^{m}\theta_{ij}^{n_{kj}}](fig/texclip20200504002545.png)
 > 
 > ただし以下を満たすものとする
 > 
 > ![0<\theta_{ij}<1](fig/texclip20200504002716.png)
 > 
 > ![\sum_{j=1}^m\theta_{ij}=1](fig/texclip20200504002745.png)
 > 

　確率 
![\theta_{ij}=P(w_j|C_i)](fig/texclip20200504002610.png)
 を求める関数を ProbWordGivenCat とするとき、確率 
![P(D_k|C_i)=\prod_{j=1}^{m}\theta_{ij}^{n_{kj}}](fig/texclip20200504002545.png)
 を求める関数 ProbDocGivenCat は次のようになる。

```go
// ソースコード 4

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




## 2.5 ProbWordGivenCat

　単語の出現確率 
![\theta_{ij}=P(w_j|C_i)](fig/texclip20200504002610.png)
 は確率 
![P(D|C_i)](fig/texclip20200504002342.png)
 のパラメータである。単語 
![w_j](fig/texclip20200504002811.png)
 の出現する個数を 
![n_j](fig/texclip20200504003028.png)
 を実測してパラメータを推定するにあたり、まず最尤推定値 (Maximum Likelihood Estimator) を使う。

 
![P(D_k|C_i)](fig/texclip20200504000754.png)
 を下に示すような尤度関数とみなし、これが最大になるパラメータを推定することになる。

> 
> 式 2.5.1
> 
> ![P(D_k|C_i)=L(\theta_{i1}\cdots\theta_{im};n_{k1}\cdots n_{km})=\prod_{j=1}^{m}\theta_{ij}^{n_kj}](fig/texclip20200504003057.png)
> 

カテゴリ 
![C_i](fig/texclip20200504000204.png)
 に属する各文書 
![D_k](fig/texclip20200504000421.png)
 の個数を 
![n_{kj}](fig/texclip20200504002902.png)
 とし、単語 
![w_j](fig/texclip20200504002811.png)
 がカテゴリ 
![C_i](fig/texclip20200504000204.png)
 に出現する個数を 
![n_{ij}](fig/texclip20200504002927.png)
 とすれば 
![\theta_{ij}](fig/texclip20200504002640.png)
 の最尤推定値は次のようになる。

> 
> 式 2.5.2
> 
> ![MLE\hat{\theta_{ij}}=\frac{\sum_{k=1}^ln_{kj}}{\sum_{j=1}^{m}\sum_{k=1}^ln_{kj}}=\frac{n_{ij}}{\sum_{j=1}^{m}n_{ij}}](fig/texclip20200504003133.png)
> 
> ![n_{ij}=\sum_{k=1}^{l}n_{kj}](fig/texclip20200504002959.png)

　numWordInCat[cat][word] を単語 word がカテゴリ cat に含まれる個数とすれば、確率 
![\theta_{ij}=P(w_j|C_i)](fig/texclip20200504002610.png)
 を計算する関数 ProbWordGivenCat は次のように与えることができる。

```go
// ソースコード 5

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

## 2.6 ProbWordGivenCat (スムージング拡張版)

　上の実装では、文書の中に一つでもカテゴリ 
![C_i](fig/texclip20200504000204.png)
 に含まれない単語が存在すると、他の単語の確率が高いものだったとしても、全体として 
![P(D|C_i)](fig/texclip20200504002342.png)
 が 0 となってしまうという問題がある。
 
　これを回避するため「加算スムージング」（あるいは「ラプラススムージング」）を使う。重複のない全単語の個数を 
![m](fig/texclip20200504002835.png)
 とすると次のように示される。

> 
> 式 2.6.1
> 
> ![EAP\hat{\theta_{ij}}=\frac{n_{ij}+1}{\sum_{j=1}^{m}n_{ij}+m}](fig/texclip20200504003210.png)
> 

　これは事前分布を一様分布と想定した場合の期待事後確率推定値 (Expected a Posterior Estimator; EAP 推定値) に相当する。
この推定値は特に標本数が少ない場合に効果があり、標本数が増えるにつれて先の最尤推定値に近づいていく。

```go
// ソースコード 6

var numAllWords int // 全単語数

// ProbWordGivenCat : 単語 word がカテゴリ cat に含まれる確率(スムージング拡張版)
func ProbWordGivenCat (word TypeWord, cat TypeCat) float64 {
    num := float64(numWordInCat[cat][word] + 1)
    sum := float64(numAllWordsInCat + numAllWords)
    return num/sum
}
```

　上記スムージングを施しても 
![\sum_{j=1}^m\theta_{ij}=1](fig/texclip20200504002745.png)
 の条件を満たすことに注意。

## 2.7 LogProbDocGivenCat

　上の関数 ProbDocGivenCat の実装では、単語数が多いと分母の値が非常に大きくなりアンダーフローが起きる恐れがあるので、これを回避すべく対数をとる。

> 
> 式 2.7.1
> 
> ![\log{P(D_k|C_i)}=\log{L(\theta_{i1}\cdots\theta_{im};n_{k1}\cdots n_{km})}=\sum_{j=1}^{m}n_{kj}\log{\theta_{ij}}](fig/texclip20200504003310.png)
> 


　このようにして ProbDocGivenCat の対数を計算する LogProbDocGivenCat の実装は次のようになる。

```go
// ソースコード 6

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


## 2.8 LogProbCatGivenDoc

　以上を踏まえると、冒頭に示した関数 ProbCatGivenDoc のアンダーフローを考慮した対数版 LogProbCatGivenDoc は次のようになる。

```go
// ソースコード 7

// LogProbCatGivenDoc : 文書 doc がカテゴリ cat に含まれる確率の比の対数
func LogProbDocGivenCat(doc TypeDoc, cat TypeCat) float64 {
	return math.Log(ProbCat(cat)) + LogProbDocGivenCat(doc, cat)
}
```



## 2.9 Predict (最終版)

　冒頭に示した関数 PredictCat は LogProbCatGivenDoc を使って次のように実装される。

```go
// ソースコード 8

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

## 3. 学習

　本節では 2 節に示したカテゴリの推定に必要なデータを作成するための「学習」について示す。


## 3.1 Train

　教師データとして、文書 doc とそのカテゴリ cat が与えられたとする。使用する変数に対して、以下の表に示す処理を実施する必要がある。

> 表 3.1.1
> 
> |変数|概要|処理|直接的に依存する関数|
> |:---|:---|:---|:---|
> |var catList []TypeCat|カテゴリのリスト|cat が初出のときのみリストに追加|PredictCat|
> |var numAllDocs int|すべての文書の数|numAllDocs の値をインクリメント|ProbCat|
> |var numDocsCat map[TypeCat]int|各カテゴリごとの文書の数|numDocsCat[cat] の値をインクリメント|ProbCat|
> |var numWordInCat map[TypeCat]map[TypeWord]int|ある単語があるカテゴリに含まれる個数|文書 doc に含まれるすべての単語 word について、numWordInCat[cat][word] の値をインクリメント|ProbWordGivenCat|
> |var numAllWordsInCat map[TypeCat]int|カテゴリに含まれる単語の個数|文書 doc に含まれる単語の個数だけ、numAllWordsInCat[cat] の値をインクリメント|ProbWordGivenCat|
> |var numAllWords int|全単語数 (重複なし)|文書 doc について wordList を更新後、len(wordList) の値を代入。つまり、これまでに出現した重複のないすべての単語の個数を代入|ProbWordGivenCat|
> |var wordList map[TypeWord]int|単語のリスト、各単語の出現数|文書 doc に含まれるすべての単語 word について、wordList[word] の値をインクリメント|-|
> 


　与えられた文書とカテゴリで学習する関数 Train の実装例は次のようになる。


```go
// ソースコード 9

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

## 4. 評価

　本稿に示したカテゴリの推定方式について、Livedoor ニュースコーパスを用いて精度を評価した。Livedoor ニュースコーパスには９つのカテゴリがあり、7367 の文書からなる。今回はこのコーパスの 80% を教師データとし、残り 20% を試験データとした場合の、適合率(Precision)・再現率(Recall)・F値(F-Measure)・正解率(Accuracyの４つの指標値を算出した。

　9 つのカテゴリの指標値はそれぞれ次のようになった。

> 
> 表 4.1
> 
> |カテゴリ|C1|C2|C3|C4|C5|C6|C7|C8|C9|
> |:--|:--|:--|:--|:--|:--|:--|:--|:--|:--|
> |Precision|0.958084|0.898734|0.860465|0.930233|0.889535|0.932886|0.857143|0.937500|0.952941|
> |Recall|1.000000|0.940397|0.907975|0.909091|0.987097|0.822485|0.814815|0.967742|0.885246|
> |F-Measure|0.978593|0.919094|0.883582|0.919540|0.935780|0.874214|0.835443|0.952381|0.917847|
> |Accuracy|0.995251|0.983039|0.973541|0.981004|0.985753|0.972863|0.964722|0.989824|0.980326|
> 

また指標のマクロ平均・マイクロ平均は次のようになった。

> 
> 表 4.2
> 
> |指標|算出値|
> |:--|:--|
> |Micro Precision|0.913161|
> |Micro Recall|0.913161|
> |Micro F-Measure|0.913161|
> |Macro Precision|0.913058|
> |Macro Recall|0.914983|
> |Macro F-Measure|0.914019|
> |Overall Accuracy|0.980703|
> 

　単純な実装内容にも関わらず、適合率・再現率とも 90% 超の精度でカテゴリ推定できることがわかった。

## 5. おわりに

　本稿では、文書のカテゴリを推定する分類器をナイーブベイズを用いたモデルで示し、そして GoLang による実装例を示した。また、Livedoor ニュースコーパスを用いて評価した結果を示した。今後はマルチラベル分類器に拡張していく予定である。


## 参考

Livedoor ニュースコーパス
https://www.rondhuit.com/download.html#ldcc

nbclsf (Naive Bayes Classifier)
https://godoc.org/github.com/bunji2/nbclsf