package sequence

// Fibonacci returns the nth Fibonacci number.
func Fibonacci(n uint) uint {
    if n == 0 {
        return 0
    } else if n == 1 {
        return 1
    }
    var a, b uint = 0, 1
    for i := 2; i <= int(n); i++ {
        a, b = b, a+b
    }
    return b
}
