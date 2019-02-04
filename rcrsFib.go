//Fibonacci using recursion
package main

import "fmt"

func fib(n int) int{
  switch n{
    case 0: return 0
    case 1: return 1
    default: return fib(n-1) + fib(n-2)
  }
}

func main(){
  n := 50
  for i:=0; i<n;i++{
    fmt.Print(fib(i)," ")
    if (i+1)%10 == 0 {fmt.Print("\n")}
  }
}