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

