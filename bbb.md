# 多項分布のパラメータ推定

## ![m](https://latex.codecogs.com/gif.latex?m) 個の目を持つサイコロ

![m](https://latex.codecogs.com/gif.latex?m) 個の目を持つサイコロを想定する。

![j](https://latex.codecogs.com/gif.latex?j) を ![](https://latex.codecogs.com/gif.latex?1\leq&space;j\leq&space;m) とし、
1 から ![m](https://latex.codecogs.com/gif.latex?m) までのそれぞれの目の出る確率を ![](https://latex.codecogs.com/gif.latex?\theta_j) とする。

1 から ![m](https://latex.codecogs.com/gif.latex?m) までのいずれかの目の出ることからこれらの確率の合計は 1 となる。

> ![](https://latex.codecogs.com/gif.latex?\sum_{j=1}^m\theta_j=1)

便利上これらの ![](https://latex.codecogs.com/gif.latex?\theta_j) のセットを次のように![](https://latex.codecogs.com/gif.latex?\boldsymbol{\theta}) で表すことにする。

> 
> &space;=&space;\theta_1,\cdots,\theta_j,\cdots,\theta_m)
> 


## 試行

![n](https://latex.codecogs.com/gif.latex?n) 回分の試行を次のように表すことにする。

> 
> ![](https://latex.codecogs.com/gif.latex?\boldsymbol{x}&space;=&space;\boldsymbol{x}_1,\cdots,\boldsymbol{x}_i,\cdots,\boldsymbol{x}_n)
> 

![n](https://latex.codecogs.com/gif.latex?n) 回中の i 番目の試行を次のように表すことにする。

> ![](https://latex.codecogs.com/gif.latex?\boldsymbol{x}_i&space;=&space;x_{1j},\cdots,x_{ij},\cdots,x_{nj})

このとき、1 の目から ![m](https://latex.codecogs.com/gif.latex?m) の目までのいずれかが出るので、それぞれの試行は 0 か 1 のいずれかとなり、一つだけ 1 となることから合計は 1 となる。

> ![](https://latex.codecogs.com/gif.latex?x_{ij}&space;\in&space;\{0,1\})
> 
> ![](https://latex.codecogs.com/gif.latex?\sum_{j=1}^mx_{ij}=1)
> 

j の目が出る回数 ![r_j](https://latex.codecogs.com/gif.latex?r_j) は次のように表すことができる。

> 
> ![](https://latex.codecogs.com/gif.latex?r_j=\sum_{i=1}^nx_{ij})
> 
> ![](https://latex.codecogs.com/gif.latex?\sum_{j=1}^mr_j=n)
> 


## 多項分布

> 
> ![](https://latex.codecogs.com/gif.latex?p(\boldsymbol{x}|\boldsymbol{\theta})=\prod_{j=1}^m\theta_{j}^{r_j})
> 

## 最尤推定

> 
> ![](https://latex.codecogs.com/gif.latex?\hat{\theta_{j}}&space;=&space;\dfrac{r_j}{n})
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

