package main

import (
	"fmt"
)

func main() {
        numGenerator := generator();
        for i := 0; i < 5; i++ {
	    fmt.Print(numGenerator(), "\t");
	}
}

func generator() func() int{
   var i = 0;
   return func() int{
       i  += 1;  
       return i;
    } 

}

