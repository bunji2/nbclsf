package nbclsf_test

import (
  "fmt"
)

func DummyTrainData()(docs[]TypeDoc,cats[]TypeCat) {
  docs = []TypeDoc{
    TypeDoc{0:1,1,23,2:44},
    // ...
    TypeDoc{0:4,2:4,3:12},
  }
  cats = []TypeCat{
    TypeCat(0),
    // ...
    TypeCat(1),
  }
  return
}

func Example() {
  docs,cats := DummyTrainData()
  clsf := nbclsf.NewClsf()
  for i:=0; i<len(docs); i++ {
    clsf.Train(docs[i], cats[i])
  }
  data := TypeDoc{0:1,1,23,2:43}
  fmt.Println("predicted Category =", clsf.Predict(data))
}
