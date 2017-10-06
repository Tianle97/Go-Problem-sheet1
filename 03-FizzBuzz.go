// Write by Tianle Shu
package main

import "fmt"

func main(){
    var a = "Fizz"          //   a  negative Fizz
    var b = "BUZZ"           //  b negative Buzz
    var c = "FizzBuzz"          // c negative FizzBuzz
 
    for i := 1; i <=100; i++{        //for-loop 1<=i<=100
        if i%3 == 0 {                //if i divisible by 3 then print a(Fizz)
                
            fmt.Println(a)            
        }
        if i%5 == 0 {                //if i divisible by 5 then print b(Buzz)
                
            fmt.Println(b)
        }
        if i%3 == 0 && i%5 == 0{       //if i divisible by 3 and 5 then print c(FizzBuzz)             
            fmt.Println(c)
    }
}
}