# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def inner_traversal(self, node: Optional[TreeNode], values: list[int]) -> list[int]:
        if node is None: return
        self.inner_traversal(node.left, values)
        values.append(node.val)
        self.inner_traversal(node.right, values)

    def inorderTraversal(self, root: Optional[TreeNode]) -> List[int]:
        values = []
        self.inner_traversal(root, values)
        return values


