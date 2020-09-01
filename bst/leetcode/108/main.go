package main

func main() {

}

//https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree/
func sortedArrayToBST(nums []int) *TreeNode {
	return createBST(nums, 0, len(nums)-1)
}

func createBST(nums []int, l, r int) *TreeNode {
	if l > r {
		return nil
	}
	mid := (l + r + 1) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = createBST(nums, l, mid-1)
	root.Right = createBST(nums, mid+1, r)
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
