# 多項分布のパラメータ推定

## ![m](https://latex.codecogs.com/gif.latex?m) 個の目を持つサイコロ

![m](https://latex.codecogs.com/gif.latex?m) 個の目を持つサイコロを想定する。

![j](https://latex.codecogs.com/gif.latex?j) を ![](https://latex.codecogs.com/gif.latex?1\leq&space;j\leq&space;m) とし、
1 から ![m](https://latex.codecogs.com/gif.latex?m) までのそれぞれの目の出る確率を ![](https://latex.codecogs.com/gif.latex?\theta_j) とする。

1 から ![m](https://latex.codecogs.com/gif.latex?m) までのいずれかの目の出ることからこれらの確率の合計は 1 となる。

> ![](https://latex.codecogs.com/gif.latex?\sum_{j=1}^m\theta_j=1)

便宜上これらの ![](https://latex.codecogs.com/gif.latex?\theta_j) のセットを次のように ![](https://latex.codecogs.com/gif.latex?\boldsymbol{\theta}) で表すことにする。

> 
> ![](https://latex.codecogs.com/gif.latex?\boldsymbol{\theta}&space;=&space;\theta_1,\cdots,\theta_j,\cdots,\theta_m)
> 


## 試行

![j](https://latex.codecogs.com/gif.latex?i) を ![](https://latex.codecogs.com/gif.latex?1\leq&space;i\leq&space;n) とし、
![n](https://latex.codecogs.com/gif.latex?n) 回分の試行全体を次のように表すことにする。

> 
> ![](https://latex.codecogs.com/gif.latex?\boldsymbol{x}&space;=&space;\boldsymbol{x}_1,\cdots,\boldsymbol{x}_i,\cdots,\boldsymbol{x}_n)
> 

![n](https://latex.codecogs.com/gif.latex?n) 回中の i 番目の試行は 1 から ![m](https://latex.codecogs.com/gif.latex?m) までのいずれかの目の出た結果となり、次のように表すことにする。

> ![](https://latex.codecogs.com/gif.latex?\boldsymbol{x}_i&space;=&space;x_{i1},\cdots,x_{ij},\cdots,x_{im})

このとき、1 の目から ![m](https://latex.codecogs.com/gif.latex?m) の目までのいずれかが必ず出るので、それぞれの試行は 0 か 1 のいずれかとなり、一つだけ 1 となることから合計は 1 となる。

> ![](https://latex.codecogs.com/gif.latex?x_{ij}&space;\in&space;\{0,1\})
> 
> ![](https://latex.codecogs.com/gif.latex?\sum_{j=1}^mx_{ij}=1)
> 

特定の目の出る回数、例えば j の目が出る回数 ![r_j](https://latex.codecogs.com/gif.latex?r_j) は次のように表すことができる。

> 
> ![](https://latex.codecogs.com/gif.latex?r_j=\sum_{i=1}^nx_{ij})
> 
> ![](https://latex.codecogs.com/gif.latex?\sum_{j=1}^mr_j=n)
> 


## 多項分布

![n](https://latex.codecogs.com/gif.latex?n) 回の試行の確率は次のように多項分布で表される。

> 
> ![](https://latex.codecogs.com/gif.latex?p(\boldsymbol{x}|\boldsymbol{\theta})=\prod_{j=1}^m\theta_{j}^{r_j})
> 

## 各目の出る確率の推定

ここで、先の各目の出る確率 ![](https://latex.codecogs.com/gif.latex?\boldsymbol{\theta}) が不明であり、これらを推定する必要がある。

以下、これらの確率を、最尤推定、最大事後確率推定、平均事後確率推定でそれぞれ求めてみる。

## 最尤推定

先の多項分布を ![](https://latex.codecogs.com/gif.latex?\boldsymbol{\theta}) の尤度関数とみなしそれが最大となるような ![](https://latex.codecogs.com/gif.latex?\boldsymbol{\theta}) を求める。

> 
> ![](https://latex.codecogs.com/gif.latex?L(\boldsymbol{\theta})=p(\boldsymbol{x}|\boldsymbol{\theta})=\prod_{j=1}^m\theta_{j}^{r_j})
> 

対数化し、微分を行う。

> 
> ![](https://latex.codecogs.com/gif.latex?\log&space;L(\boldsymbol{\theta})=\sum_{j=1}^mr_j\log\theta_j=\sum_{j=1}^{m-1}r_j\log\theta_j&plus;r_(1-\Sigma_{j=1}^m\theta_j)\log(1-\Sigma_{j=1}^m\theta_j))
> 
> ![](https://latex.codecogs.com/gif.latex?\frac{\partial&space;}{\partial\theta_j}\log&space;L(\boldsymbol{\theta})=\frac{r_j}{\theta_j}-\frac{n-\Sigma_{j=1}^{m-1}r_j}{1-\Sigma_{j=1}^{m-1}\theta_j}=\frac{r_j}{\theta_j}-\frac{r_m}{\theta_m}=0)
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

## 事前確率

> 
> ![](https://latex.codecogs.com/gif.latex?p(\boldsymbol{\theta})&space;=&space;\dfrac{1}{B(\boldsymbol{\alpha})}\prod_{j=1}^{m}\theta_j^{\alpha_j&space;-1})
> 
> where
> 
> ![](https://latex.codecogs.com/gif.latex?B(\boldsymbol{\alpha})&space;=&space;\dfrac{\Pi_{j=1}^{m}\Gamma(\alpha_j)}{\Gamma(\Sigma_{j=1}^m\alpha_j)})
> 
> ![](https://latex.codecogs.com/gif.latex?\boldsymbol{\alpha}&space;=&space;\alpha_1,\cdots,\alpha_j,\cdots,\alpha_m)

　一様分布の場合には ![](https://latex.codecogs.com/gif.latex?\alpha_j&space;=&space;1) となることから、

> ![](https://latex.codecogs.com/gif.latex?B(\boldsymbol{\alpha})&space;=&space;\dfrac{\Pi_{j=1}^{m}\Gamma(1)}{\Gamma(\Sigma_{j=1}^m1)}&space;=&space;\dfrac{1}{(m-1)!})
> 

## 事後確率

> 
> ![](https://latex.codecogs.com/gif.latex?p(\boldsymbol{\theta}|\boldsymbol{x})=\dfrac{1}{B(\boldsymbol{\alpha'})}\prod_{j=1}^{m}\theta_j^{\alpha_j'-1})
> 
> where
> 
> ![](https://latex.codecogs.com/gif.latex?\boldsymbol{\alpha'}&space;=&space;\alpha_1',\cdots,\alpha_j',\cdots,\alpha_m')
> 
> ![](https://latex.codecogs.com/gif.latex?\alpha_{j}'&space;=&space;\alpha_j+r_j)
> 

## MAP 推定値

![MAP](https://latex.codecogs.com/gif.latex?\hat{\theta_j}_{MAP}=\dfrac{r_j&space;&plus;&space;\alpha_j&space;-1}{n+\Sigma_{j=1}^{m}r_j-m})

## EAP 推定値

![](https://latex.codecogs.com/gif.latex?{\hat{\theta_j}}_{EAP}=\dfrac{r_j&plus;\alpha_j}{n&plus;\Sigma_{j=1}^{m}r_j})

