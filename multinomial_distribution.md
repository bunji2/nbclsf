# 多項分布のパラメータ推定

* 最終更新日: 2020/1/19
* AUTHOR：Bunji² (Bunji Square) 

## 1. はじめに

　多項分布のパラメータ推定について勉強したのでまとめる。

　![m](https://latex.codecogs.com/gif.latex?m) 個の目を持つサイコロを例にとってみる。各目の出る確率は与えられていないものとする。

　各目の出る確率が均等であれば、どの目も ![1/m](https://latex.codecogs.com/gif.latex?1/m) の確率で出るはずであるが、均等かどうかすらも不明だとする。

　本稿ではこのサイコロを振った結果から各目の確率を推定する場合の計算方法について説明していく。

## 2. 多項分布

## 2.1 各目の確率

　![j](https://latex.codecogs.com/gif.latex?j) を ![](https://latex.codecogs.com/gif.latex?1\leq&space;j\leq&space;m) とし、
![1](https://latex.codecogs.com/gif.latex?1) から ![m](https://latex.codecogs.com/gif.latex?m) までのそれぞれの目の出る確率を ![](https://latex.codecogs.com/gif.latex?\theta_j) とする。

　![1](https://latex.codecogs.com/gif.latex?1) から ![m](https://latex.codecogs.com/gif.latex?m) までのいずれかの目の出ることからこれらの確率の合計は ![1](https://latex.codecogs.com/gif.latex?1) となる。

> ![](https://latex.codecogs.com/gif.latex?\sum_{j=1}^m\theta_j=1)

　便宜上これらの ![](https://latex.codecogs.com/gif.latex?\theta_j) のセットを次のように ![](https://latex.codecogs.com/gif.latex?\boldsymbol{\theta}) で表すことにする。

> 
> ![](https://latex.codecogs.com/gif.latex?\boldsymbol{\theta}&space;=&space;\theta_1,\cdots,\theta_j,\cdots,\theta_m)
> 


## 2.2 試行

　![j](https://latex.codecogs.com/gif.latex?i) を ![](https://latex.codecogs.com/gif.latex?1\leq&space;i\leq&space;n) とし、
![n](https://latex.codecogs.com/gif.latex?n) 回分の試行全体を次のように表すことにする。

> 
> ![](https://latex.codecogs.com/gif.latex?\boldsymbol{x}&space;=&space;\boldsymbol{x}_1,\cdots,\boldsymbol{x}_i,\cdots,\boldsymbol{x}_n)
> 

　![n](https://latex.codecogs.com/gif.latex?n) 回中の i 番目の試行は 1 から ![m](https://latex.codecogs.com/gif.latex?m) までのいずれかの目の出た結果となり、次のように表すことにする。

> ![](https://latex.codecogs.com/gif.latex?\boldsymbol{x}_i&space;=&space;x_{i1},\cdots,x_{ij},\cdots,x_{im})

　このとき、1 の目から ![m](https://latex.codecogs.com/gif.latex?m) の目までのいずれかが必ず出るので、それぞれの試行は 0 か 1 のいずれかとなり、一つだけ 1 となることから合計は 1 となる。

> ![](https://latex.codecogs.com/gif.latex?x_{ij}\in\left\{0,1\right\})
> 
> 
> ![](https://latex.codecogs.com/gif.latex?\sum_{j=1}^mx_{ij}=1)
> 

　特定の目の出る回数、例えば j の目が出る回数 ![r_j](https://latex.codecogs.com/gif.latex?r_j) は次のように表すことができる。

> 
> ![](https://latex.codecogs.com/gif.latex?r_j=\sum_{i=1}^nx_{ij})
> 
> ![](https://latex.codecogs.com/gif.latex?\sum_{j=1}^mr_j=n)
> 


## 2.3 ![n](https://latex.codecogs.com/gif.latex?n) 回の試行の確率

　![i](https://latex.codecogs.com/gif.latex?i) 回目の試行の確率は次のような多項分布で表される。

> 
> ![](https://latex.codecogs.com/gif.latex?p(\boldsymbol{x}_i|\boldsymbol{\theta})=\prod_{j=1}^m\theta_{j}^{x_{ij}})
> 


　各試行が独立しているとすると、![n](https://latex.codecogs.com/gif.latex?n) 回の試行の確率は次のように表される。

> 
> ![](https://latex.codecogs.com/gif.latex?p(\boldsymbol{x}|\boldsymbol{\theta})=\prod_{i=1}^np(\boldsymbol{x}_i|\boldsymbol{\theta})=\prod_{i=1}^n\prod_{j=1}^m\theta_{j}^{x_{ij}}=\prod_{i=1}^n\theta_{j}^{\sum_{j=1}^mx_{ij}})
> 

全試行中の特定の目の出る回数 ![](https://latex.codecogs.com/gif.latex?r_j) を使えば次のように表される。

> 
> ![](https://latex.codecogs.com/gif.latex?p(\boldsymbol{x}|\boldsymbol{\theta})=\prod_{i=1}^n\theta_{j}^{\sum_{j=1}^mx_{ij}}=\prod_{j=1}^m\theta_{j}^{r_j})
> 
>  ![](https://latex.codecogs.com/gif.latex?r_j=\sum_{i=1}^nx_{ij})
> 

## 3. 各目の出る確率の推定

　各目の出る確率 ![](https://latex.codecogs.com/gif.latex?\boldsymbol{\theta}) を推定するにあたり、最尤推定 (maximum likelihood estimation; MLE) 、最大事後確率推定(maximum a posteriori; MAP 推定)、期待事後確率推定(expected a
posteriori; EAP 推定)でそれぞれ求めてみる。

## 3.1 最尤推定

　先の多項分布を ![](https://latex.codecogs.com/gif.latex?\boldsymbol{\theta}) の尤度関数とみなしそれが最大となるような ![](https://latex.codecogs.com/gif.latex?\boldsymbol{\theta}) を求める。

> 
> ![](https://latex.codecogs.com/gif.latex?L(\boldsymbol{\theta})=p(\boldsymbol{x}|\boldsymbol{\theta})=\prod_{j=1}^m\theta_{j}^{r_j})
> 

　対数化し、

> 
> ![](https://latex.codecogs.com/gif.latex?\log&space;L(\boldsymbol{\theta})=\sum_{j=1}^mr_j\log\theta_j=\sum_{j=1}^{m-1}r_j\log\theta_j&plus;r_m\log(1-\Sigma_{j=1}^{m-1}\theta_j))
> 

> 
> ![](https://latex.codecogs.com/gif.latex?\theta_m=1-\sum_{j=1}^{m-1}\theta_j)
> 

　微分を行う。

> ![](https://latex.codecogs.com/gif.latex?\frac{\partial&space;}{\partial\theta_j}\log&space;L(\boldsymbol{\theta})=\frac{r_j}{\theta_j}-\frac{r_m}{1-\Sigma_{j=1}^{m-1}\theta_j}=\frac{r_j}{\theta_j}-\frac{r_m}{\theta_m}=0)
> 
> ![](https://latex.codecogs.com/gif.latex?\theta_j=\frac{r_j}{r_m}\theta_m)

　ここで、![](https://latex.codecogs.com/gif.latex?\Sigma_{j=1}^m\theta_j=1) であることより、

> 
> ![](https://latex.codecogs.com/gif.latex?\sum_{j=1}^m\theta_j=\sum_{j=1}^m\frac{r_j}{r_m}\theta_m=\frac{n}{r_m}\theta_m=1)
> 
> ![](https://latex.codecogs.com/gif.latex?\theta_j=\frac{r_j}{r_m}\theta_m=\frac{r_j}{r_m}\cdot\frac{r_m}{n}=\frac{r_j}{n})
> 

　以上より、最尤推定値は次のように示される。


> 
> ![](https://latex.codecogs.com/gif.latex?{\hat{\theta_{j}}}_{MLE}=\dfrac{r_j}{n})
> 

　つまり、![](https://latex.codecogs.com/gif.latex?j) の目が出る確率の最尤推定値は、![](https://latex.codecogs.com/gif.latex?j) の目が出た回数を全試行数で割ることで算出される。

## 3.2 最大事後確率推定と期待事後確率推定

　最大事後確率推定(MAP 推定)と期待事後確率推定 (EAP 推定) を行うには事後確率を求める必要があり、そのためには事前確率を用意する必要がある。

## 3.2.1 事前確率

　事後確率を計算しやすい事前確率として「共益事前分布」というものがあり、
多項分布の場合は「ディリクレ分布」がよく用いられる。

> 
> ![](https://latex.codecogs.com/gif.latex?p(\boldsymbol{\theta})=Dir(\boldsymbol{\theta};\boldsymbol{\alpha})=\dfrac{1}{B(\boldsymbol{\alpha})}\prod_{j=1}^{m}\theta_j^{\alpha_j-1})
> 

　ただし、![](https://latex.codecogs.com/gif.latex?\boldsymbol{\alpha}) はディリクレ分布のパラメータであり、![](https://latex.codecogs.com/gif.latex?B(\boldsymbol{\alpha})) は次のようにガンマ関数 ![](https://latex.codecogs.com/gif.latex?\Gamma) で与えられる係数である。

> 
> ![](https://latex.codecogs.com/gif.latex?\boldsymbol{\alpha}=\alpha_1,\cdots,\alpha_j,\cdots,\alpha_m)
> 
> ![](https://latex.codecogs.com/gif.latex?B(\boldsymbol{\alpha})=\dfrac{\Pi_{j=1}^{m}\Gamma(\alpha_j)}{\Gamma(\Sigma_{j=1}^m\alpha_j)})
> 

　![](https://latex.codecogs.com/gif.latex?p(\boldsymbol{\theta})) が一様分布、つまり![](https://latex.codecogs.com/gif.latex?\boldsymbol{\theta}) に依存せず定数となる場合には、パラメータの各値は ![](https://latex.codecogs.com/gif.latex?\alpha_j=1) となる。

> 
> ![](https://latex.codecogs.com/gif.latex?p(\boldsymbol{\theta})=\dfrac{1}{B(\boldsymbol{\alpha})}\prod_{j=1}^{m}\theta_j^{1-1}=\dfrac{\Pi_{j=1}^{m}\Gamma(1)}{\Gamma(\Sigma_{j=1}^m1)}=\dfrac{1}{(m-1)!})
> 


## 3.2.2 事後確率

　多項分布の事前確率にディリクレ分布を用いると、事後確率もディリクレ分布となる。詳細は Appendix を参照のこと。

> 
> ![](https://latex.codecogs.com/gif.latex?p(\boldsymbol{\theta}|\boldsymbol{x})=Dir(\boldsymbol{\theta};\boldsymbol{\alpha'})=\dfrac{1}{B(\boldsymbol{\alpha'})}\prod_{j=1}^{m}\theta_j^{\alpha_j'-1})
> 

　ただし、![](https://latex.codecogs.com/gif.latex?\boldsymbol{\alpha'}) は ![](https://latex.codecogs.com/gif.latex?\alpha_j) と各目の出た回数 ![](https://latex.codecogs.com/gif.latex?r_j) を用いて次のようになる。

> ![](https://latex.codecogs.com/gif.latex?\boldsymbol{\alpha'}=\alpha_1',\cdots,\alpha_j',\cdots,\alpha_m')
> 
> ![](https://latex.codecogs.com/gif.latex?\alpha_{j}'=\alpha_j+r_j)
> 

## 3.2.3 最大事後確率

　最大事後確率推定 (MAP 推定) は事後確率を最大にする ![](https://latex.codecogs.com/gif.latex?\boldsymbol{\theta}) を推定する。
　事後確率を ![](https://latex.codecogs.com/gif.latex?\boldsymbol{\theta}) の尤度関数とみなす。

> 
> ![](https://latex.codecogs.com/gif.latex?L(\boldsymbol{\theta})=p(\boldsymbol{\theta}|\boldsymbol{x})=Dir(\boldsymbol{\theta};\boldsymbol{\alpha'})=\dfrac{1}{B(\boldsymbol{\alpha'})}\prod_{j=1}^{m}\theta_j^{\alpha_j'-1})
> 

　対数化し微分を行う。

> ![](https://latex.codecogs.com/gif.latex?\log&space;L(\boldsymbol{\theta})=-\log&space;B(\boldsymbol{\alpha'})&plus;\sum_{j=1}^{m}(\alpha_j'-1)\log\theta_j)


> 
> ![](https://latex.codecogs.com/gif.latex?\frac{\partial}{\partial\theta_j}\log&space;L(\boldsymbol{\theta})=\dfrac{\alpha_j'-1}{\theta_j}-\dfrac{\alpha_m'-1}{\Sigma_{j=1}^{m-1}\theta_j}=\dfrac{\alpha_j'-1}{\theta_j}-\dfrac{\alpha_m'-1}{\theta_m}=0)
> 
> ![](https://latex.codecogs.com/gif.latex?\theta_j=\dfrac{\alpha_j'-1}{\alpha_m'-1}\theta_m)

> 
> ![](https://latex.codecogs.com/gif.latex?\sum_{j=1}^m\theta_j=\sum_{j=1}^m\dfrac{\alpha_j'-1}{\alpha_m'-1}\theta_m=\frac{\sum_{j=1}^m(\alpha_j'-1)}{\alpha_m'-1}\theta_m=1)
> 
> ![](https://latex.codecogs.com/gif.latex?\theta_m=\frac{\alpha_m'-1}{\sum_{j=1}^m(\alpha_j'-1)})
> 
> ![](https://latex.codecogs.com/gif.latex?\theta_j=\dfrac{\alpha_j'-1}{\alpha_m'-1}\theta_m=\dfrac{\alpha_j'-1}{\alpha_m'-1}\cdot\frac{\alpha_m'-1}{\sum_{j=1}^m(\alpha_j'-1)}=\dfrac{r_j&plus;\alpha_j-1}{n+\Sigma_{j=1}^m(\alpha_j-1)})

> 

　以上のことから、最大事後確率推定値は次のようになる。

> 
> ![MAP](https://latex.codecogs.com/gif.latex?\hat{\theta_j}_{MAP}=\dfrac{r_j&plus;\alpha_j-1}{n+\Sigma_{j=1}^{m}(\alpha_j-1)})
> 

## 3.2.4 期待事後確率推定

　期待事後確率推定 (EAP 推定) は、事後確率の期待値を使う。

> 
> ![](https://latex.codecogs.com/gif.latex?E[p(\boldsymbol{\theta}|\boldsymbol{x})]=E\left[\dfrac{1}{B(\boldsymbol{\alpha'})}\prod_{j=1}^{m}\theta_j^{\alpha_j'-1}\right]=\dfrac{\alpha_j'}{\Sigma_{j=1}^{m}\alpha_j'}=\dfrac{r_j&plus;\alpha_j}{n&plus;\Sigma_{j=1}^{m}\alpha_j})
> 

　以上より、期待事後確率推定は次のようになる。

> 
> ![EAP](https://latex.codecogs.com/gif.latex?{\hat{\theta_j}}_{EAP}=\dfrac{r_j&plus;\alpha_j}{n&plus;\Sigma_{j=1}^{m}\alpha_j})
> 

## 3.3 一様分布の場合の推定値

　事前分布を一様分布とした場合つまり ![](https://latex.codecogs.com/gif.latex?\alpha_j=1) の場合には、最尤推定(MLE)・最大事後確率推定(MAP推定)・期待事後確率推定(EAP推定)は次のようになる。

> ![MLE/MAP](https://latex.codecogs.com/gif.latex?\hat{\theta_j}_{MLE}=\hat{\theta_j}_{MAP}=\dfrac{r_j}{n})
> 
> ![EAP](https://latex.codecogs.com/gif.latex?{\hat{\theta_j}}_{EAP}=\dfrac{r_j&plus;1}{n&plus;m})

## 4. おわりに

　本稿では多項分布のパラメータ推定について ![m](https://latex.codecogs.com/gif.latex?m) 個の目を持つサイコロを例にして、最尤推定・最大事後確率推定・期待事後確率推定の方法についてまとめた。

## Appendix

## A.1 事後確率 ![](https://latex.codecogs.com/gif.latex?p(\boldsymbol{\theta}|\boldsymbol{x})) 

　本文で省略した事後確率の計算式を残しておく。

　ベイズの定理に基づき、事後確率 ![](https://latex.codecogs.com/gif.latex?p(\boldsymbol{\theta}|\boldsymbol{x})) は次がなりたつ。

> ![](https://latex.codecogs.com/gif.latex?p(\boldsymbol{\theta}|\boldsymbol{x})=\frac{p(\boldsymbol{\theta})p(\boldsymbol{x}|\boldsymbol{\theta})}{p(\boldsymbol{x})})

　ここで、![](https://latex.codecogs.com/gif.latex?p(\boldsymbol{\theta})p(\boldsymbol{x}|\boldsymbol{\theta})), ![](https://latex.codecogs.com/gif.latex?p(\boldsymbol{x})) は次のようになる。

> 
> ![](https://latex.codecogs.com/gif.latex?p(\boldsymbol{\theta})p(\boldsymbol{x}|\boldsymbol{\theta})=\dfrac{1}{B(\boldsymbol{\alpha})}\prod_{j=1}^{m}\theta_j^{\alpha_j-1}\cdot\prod_{j=1}^m\theta_{j}^{r_j}=\dfrac{1}{B(\boldsymbol{\alpha})}\prod_{j=1}^{m}\theta_j^{r_j&plus;\alpha_j-1})
> 

　![](https://latex.codecogs.com/gif.latex?p(\boldsymbol{x})) は次のようになる。

> 
> ![](https://latex.codecogs.com/gif.latex?p(\boldsymbol{x})=\int_{0}^1p(\boldsymbol{\theta})p(\boldsymbol{x}|\boldsymbol{\theta})d\boldsymbol{\theta}) 
> 
> ![](https://latex.codecogs.com/gif.latex?=\int_{0}^1\dfrac{1}{B(\boldsymbol{\alpha})}\prod_{j=1}^{m}\theta_j^{r_j&plus;\alpha_j-1}d\boldsymbol{\theta}) 
> 
> ![](https://latex.codecogs.com/gif.latex?=\dfrac{1}{B(\boldsymbol{\alpha})}\int_{0}^1\prod_{j=1}^{m}\theta_j^{r_j&plus;\alpha_j-1}d\theta_1\cdots&space;d\theta_m) 
> 
> ![](https://latex.codecogs.com/gif.latex?=\dfrac{1}{B(\boldsymbol{\alpha})}\dfrac{\Pi_{j=1}^m\Gamma(r_j&plus;\alpha_j)}{\Gamma(\Sigma_{j=1}^m(r_j&plus;\alpha_j))}) 
> 
> ![](https://latex.codecogs.com/gif.latex?=\dfrac{B(\boldsymbol{\alpha}')}{B(\boldsymbol{\alpha})}) 
> 

　ただし、![](https://latex.codecogs.com/gif.latex?\boldsymbol{\alpha'}) は ![](https://latex.codecogs.com/gif.latex?\alpha_j) と各目の出た回数 ![](https://latex.codecogs.com/gif.latex?r_j) を用いて次で示される。

> 
> ![](https://latex.codecogs.com/gif.latex?\boldsymbol{\alpha'}=\alpha_1',\cdots,\alpha_j',\cdots,\alpha_m')
> 
> ![](https://latex.codecogs.com/gif.latex?\alpha_{j}'=\alpha_j+r_j)
> 


　従って事後確率 ![](https://latex.codecogs.com/gif.latex?p(\boldsymbol{\theta}|\boldsymbol{x})) は次のようになる。

> 
> ![](https://latex.codecogs.com/gif.latex?p(\boldsymbol{\theta}|\boldsymbol{x})=\frac{p(\boldsymbol{\theta})p(\boldsymbol{x}|\boldsymbol{\theta})}{p(\boldsymbol{x})})
> 
> ![](https://latex.codecogs.com/gif.latex?=\dfrac{1}{B(\boldsymbol{\alpha})}\prod_{j=1}^{m}\theta_j^{r_j&plus;\alpha_j-1}\cdot\dfrac{B(\boldsymbol{\alpha})}{B(\boldsymbol{\alpha}')})
> 
> ![](https://latex.codecogs.com/gif.latex?=\dfrac{1}{B(\boldsymbol{\alpha'})}\prod_{j=1}^{m}\theta_j^{\alpha_j'-1})
> 
> ![](https://latex.codecogs.com/gif.latex?=Dir(\boldsymbol{\theta};\boldsymbol{\alpha'}))
> 

　従って、事後分布はディリクレ分布となる。


## 参考

* ガンマ関数の積分公式、ディリクレ分布の期待値の算出については、次のページを参考にした。
 
  「ディリクレ分布の期待値・分散・共分散の導出」
  
    https://to-kei.net/distribution/dirichlet-distribution/dirichlet-distribution-derivation/


