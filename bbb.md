# 多項分布のパラメータ推定

## 多項分布

![](https://latex.codecogs.com/gif.latex?p(\boldsymbol{x}|\boldsymbol{\theta})=\prod_{j=1}^m\theta_{j}^{r_j})

## 最尤推定

![](https://latex.codecogs.com/gif.latex?\hat{\theta_{j}}&space;=&space;\dfrac{r_j}{n})

## 事前確率

> 
> ![](https://latex.codecogs.com/gif.latex?p(\boldsymbol{\theta})&space;=&space;\dfrac{1}{B(\boldsymbol{\alpha})}\prod_{j=1}^{m}\theta_j^{\alpha_j&space;-1})
> 
> where
> 
> ![](https://latex.codecogs.com/gif.latex?B(\boldsymbol{\alpha})&space;=&space;\dfrac{\Pi_{j=1}^{m}\Gamma(\alpha_j)}{\Gamma(\Sigma_{j=1}^m\alpha_j)})
> 
> ![](https://latex.codecogs.com/gif.latex?\boldsymbol{\alpha}&space;=&space;\alpha_1,\cdots,\alpha_j,\cdots,\alpha_m)

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
