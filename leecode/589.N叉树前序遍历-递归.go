package main

/*
给定一个 N 叉树，返回其节点值的前序遍历。

例如，给定一个 3叉树 :

 



 

返回其前序遍历: [1,3,5,6,2,4]。

 

说明: 递归法很简单，你可以使用迭代法完成此题吗?

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

 func preorder(root *Node) []int {
    var list []int
    list = dfs(root, list)
    return list
}

func dfs(root *Node, list []int) []int {
    if root != nil {
        list = append(list, root.Val)
        for _,node := range root.Children {
            list = dfs(node,list)
        }
    }
    return list
}