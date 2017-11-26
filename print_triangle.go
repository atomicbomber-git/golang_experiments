package main

import "os"
import "fmt"
import "strconv"

func main() {
    /* 
        This program is supposed to print asterisks in such manner:
        1 asterisk in line 1, 2 asterisk in line 2, so on and so on.
    */

    if len(os.Args) < 2 {
        /* Error */
        fmt.Println("Error: There must be at least an argument.")
        return
    }

    amount, error := strconv.Atoi(os.Args[1])

    if (error != nil) || (amount < 1) {
        fmt.Println("Argument must be an integer larger than 0.")
        return
    }

    for i := 0; i < amount; i++ {
        for j := 0; j < i; j++ {
            fmt.Print("*")
        }
        fmt.Print("\n")
    }
}