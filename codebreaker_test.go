package main

import "testing"

func Test_codebreakerAciertaX(t *testing.T){
  //6189
  expected := "X"
  actual := codebreaker("6189")
  if actual != expected{
    t.Errorf("Test failed, Expected: '%s' , Actual: '%s'",expected,actual)
  }
}

func Test_codebreakerAciertaXX(t *testing.T){
  //6189
  expected := "XX"
  actual := codebreaker("6589")
  if actual != expected{
    t.Errorf("Test failed, Expected: '%s' , Actual: '%s'",expected,actual)
  }
}

func Test_codebreakerAcierta_(t *testing.T){
  //6189
  expected := "_"
  actual := codebreaker("1762")
  if actual != expected{
    t.Errorf("Test failed, Expected: '%s' , Actual: '%s'",expected,actual)
  }
}

func Test_codebreakerAciertaX_(t *testing.T){
  //6189
  expected := "X_"
  actual := codebreaker("1264")
  if actual != expected{
    t.Errorf("Test failed, Expected: '%s' , Actual: '%s'",expected,actual)
  }
}

func Test_codebreakerAciertaXXXX(t *testing.T){
  //6189
  expected := "XXXX"
  actual := codebreaker("2584")
  if actual != expected{
    t.Errorf("Test failed, Expected: '%s' , Actual: '%s'",expected,actual)
  }
}

func Test_codebreakerAcierta____(t *testing.T){
  //6189
  expected := "____"
  actual := codebreaker("4852")
  if actual != expected{
    t.Errorf("Test failed, Expected: '%s' , Actual: '%s'",expected,actual)
  }
}
