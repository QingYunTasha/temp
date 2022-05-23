func minSimple(x, y int) int {
    if x < y {
        return x
    }
    return y
}

func minimumTotal(triangle [][]int) int {
    dp := make([]int, len(triangle) + 1)
    for row := len(triangle) - 1; row >= 0; row-- {
        for idx, currNum := range triangle[row] {
            dp[idx] = currNum + minSimple(dp[idx], dp[idx + 1])

        }
    }
    return dp[0]
}
