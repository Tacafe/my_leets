from typing import Optional
# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    def append(self, tree, new_node):
        if tree.val > new_node.val:
            if tree.left:
                tree = self.append(tree.left, new_node)
            else:
                tree.left = new_node
        else:
            if tree.right:
                tree = self.append(tree.right, new_node)
            else:
                tree.right = new_node

        return tree

    def bst(self, tree, nodes):
        if not tree.left and not tree.right:
            return

        if tree.left:
            nodes.append(tree.left)
            if not tree.right:
                nodes.append(None)
        if tree.right:
            if not tree.left:
                nodes.append(None)
            nodes.append(tree.right)

        if tree.left: self.bst(tree.left, nodes)
        if tree.right: self.bst(tree.right, nodes)


    def sortedArrayToBST(self, nums: list[int]) -> Optional[TreeNode]:
        # root_index = len(nums) // 2
        # root = TreeNode(val=nums.pop(root_index), left=None, right=None)

        # for num in nums:
        #     new_node = TreeNode(val=num, left=None, right=None)
        #     self.append(root, new_node)

        # return root

        if not nums:
            return None
        mid = len(nums) // 2
        root = TreeNode(nums[mid])
        root.left = self.sortedArrayToBST(nums[:mid])
        root.right = self.sortedArrayToBST(nums[mid+1:])
        return root


if __name__ == '__main__':
    nums = [-10,-3,0,5,9]
    print(Solution().sortedArrayToBST(nums))
