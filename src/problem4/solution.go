package main
import "fmt"

func sum_to_n_a(n int) int {
	// your code here
	result := 0
	for i := 1; i <= n; i++ {
		result += i
	}
	return result
}

func sum_to_n_b(n int) int {
	// your code here
	return n * (n + 1) / 2
}

func sum_to_n_c(n int) int {
	//your code

    cache := make(map[int]int)

    var dp func(int) int

    dp = func(n int) int {
        if n == 0 {
            return 0
        }

        if result, ok := cache[n]; ok {
            return result
        }

        result := dp(n-1) + n
        cache[n] = result

        return result
    }

    return dp(n)
}

func main() {
    fmt.Println(sum_to_n_a(10)) // 55
    fmt.Println(sum_to_n_b(10)) // 55
    fmt.Println(sum_to_n_c(10)) // 55
}

//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