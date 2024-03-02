package main

import (
 "fmt"
 "reflect"
)

func choto(arr interface{}) interface{} {
 arrValue := reflect.ValueOf(arr)
 if arrValue.Kind() != reflect.Slice {
  return nil
 }

 arrLen := arrValue.Len()
 reversedArr := reflect.MakeSlice(arrValue.Type(), arrLen, arrLen)

 for i := 0; i < arrLen; i++ {
  reversedArr.Index(i).Set(arrValue.Index(arrLen - 1 - i))
 }

 return reversedArr.Interface()
}

func main() {
 chis := []int{1, 2, 3, 4, 5}
 fmt.Println(choto(chis)) // [5 4 3 2 1]

 slov := []string{"Hello", "how", "are", "you", "?"}
 fmt.Println(choto(slov)) // [? you are how Hello]

	//var slov []string  = []string {"Hello", "how", "are", "you", "?"} // 
	//fmt.Println(slov)
	//var slove []string =  []string { "?", "you","are", "how", "Hello"}
	//fmt.Println(slove)
 
}
