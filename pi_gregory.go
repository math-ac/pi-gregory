/*
 * Author: Matheus Amorim Constancio
 * 
 * Estimate Pi using Gregory's Series
 * 
 * pi / 4 = 1 - (1 / 3) + (1 / 5) - (1 / 7) + ...
 */

package main

import "fmt"
import "os"
import "os/exec"

func gregory(n int) float64 {
    var pi float64
    var signal = 1
     
    for i := 0; i < n; i++ {
        pi += float64(signal) / float64(2 * i + 1)
        signal *= -1
    }
    
    return pi * 4.
}

func main() {
    var pi float64
    var n int
    
    c := exec.Command("clear")
    c.Stdout = os.Stdout
    c.Run()
    
    fmt.Printf("Number of iterations: ")
    fmt.Scan(&n)
    
    pi = gregory(n)
    err := (pi - 3.141592653589793) / 3.141592653589793
    fmt.Printf("Pi = %0.15f\nError = %0.15f\n", pi, err)
}
