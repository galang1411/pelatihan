import (
    "fmt"
    "math/big"
)

func main() {
    for i := 2; i < 1000; i++ {
        if big.NewInt(int64(i)).ProbablyPrime(20) {
            fmt.Printf("%d is probably prime\n", i)
        } else {
            fmt.Printf("%d is definitely not prime\n", i)
        }
    }
    
}