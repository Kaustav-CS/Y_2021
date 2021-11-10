/******************************************************************************
Link:   https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/
108. Convert Sorted Array to Binary Search Tree

Your input
[-10,-3,0,5,9]

Output
[0,-10,5,null,-3,null,9]

Expected
[0,-3,9,-10,null,5]

DATE:    2021, November 10
২৪ কার্তিক,১৪২৮

*******************************************************************************/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func sortedArrayToBST(nums []int) *TreeNode {
    if len(nums) == 0 { return nil }
    
    m := (len(nums)-1)/2
    return &TreeNode{
                     nums[m],
                     sortedArrayToBST(nums[0:m]),
                     sortedArrayToBST(nums[m+1:]),
                    }
}
