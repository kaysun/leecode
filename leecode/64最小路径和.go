package main

func main() {

}

/*
    给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

    说明：每次只能向下或者向右移动一步。

    示例:

    输入:
    [
      [1,3,1],
    [1,5,1],
    [4,2,1]
    ]
    输出: 7
    解释: 因为路径 1→3→1→1→1 的总和最小。

    来源：力扣（LeetCode）
    链接：https://leetcode-cn.com/problems/minimum-path-sum
    著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func minPathSum(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    dp := make([][]int, m)
    for i := range dp {
        dp[i] = make([]int, n)
        for j := range dp[i] {
            if i == 0 && j == 0 {
                dp[i][j] = grid[i][j]
            } else if i == 0 && j != 0 {
                dp[i][j] = dp[i][j-1] + grid[i][j] 
            } else if i != 0 && j == 0 {
                dp[i][j] = dp[i-1][j] + grid[i][j] 
            } else {
                if dp[i-1][j] < dp[i][j-1] {
                    dp[i][j] = dp[i-1][j] + grid[i][j] 
                } else {
                    dp[i][j] = dp[i][j-1] + grid[i][j] 
                }
            }
        }
    }
    return dp[m-1][n-1]
}