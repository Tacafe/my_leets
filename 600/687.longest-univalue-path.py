#
# @lc app=leetcode id=687 lang=python3
#
# [687] Longest Univalue Path
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def longestUnivaluePath(self, root):
        self.ans = 0

        def arrow_length(node):
            if not node: return 0
            left_length = arrow_length(node.left)
            right_length = arrow_length(node.right)
            left_arrow = right_arrow = 0
            if node.left and node.left.val == node.val:
                left_arrow = left_length + 1
            if node.right and node.right.val == node.val:
                right_arrow = right_length + 1
            self.ans = max(self.ans, left_arrow + right_arrow)
            return max(left_arrow, right_arrow)

        arrow_length(root)
        return self.ans

    # def merge_uni_value_dics(self, uvd_1, uvd_2):
    #     merged = {}
    #     for k in uvd_1.keys():
    #         if k in uvd_2:
    #             merged[k] = max([uvd_1[k], uvd_2[k]])
    #         else:
    #             merged[k] = uvd_1[k]

    #     for k in uvd_2.keys():
    #         if not k in merged:
    #             merged[k] = uvd_2[k]

    #     return merged

    # def traverse(self, node, uni_value_dic):
    #     # left side
    #     new_uni_value_dic = {}

    #     if node.left:
    #         if node.val == node.left.val:
    #             if node.val in uni_value_dic:
    #                 uni_value_dic[node.val] += 1
    #             else:
    #                 uni_value_dic[node.val] = 1

    #             # keep uni_value_dic
    #             uni_value_dic = self.traverse(node.left, uni_value_dic)
    #         else:
    #             # initialize new dict when the path stopped
    #             new_uni_value_dic = self.traverse(node.left, new_uni_value_dic)

    #     uni_value_dic = self.merge_uni_value_dics(uni_value_dic, new_uni_value_dic)

    #     # right side
    #     new_uni_value_dic = {}

    #     if node.right:
    #         if node.val == node.right.val:
    #             if node.val in uni_value_dic:
    #                 uni_value_dic[node.val] += 1
    #             else:
    #                 uni_value_dic[node.val] = 1

    #             # keep uni_value_dic
    #             uni_value_dic = self.traverse(node.right, uni_value_dic)
    #         else:
    #             new_uni_value_dic = self.traverse(node.right, new_uni_value_dic)

    #     uni_value_dic = self.merge_uni_value_dics(uni_value_dic, new_uni_value_dic)

    #     return uni_value_dic

    # def longestUnivaluePath(self, root: Optional[TreeNode]) -> int:
    #     if not root: return 0

    #     uni_value_dic = {}
    #     uni_value_dic = self.traverse(root, uni_value_dic)

    #     if len(uni_value_dic) == 0: return 0

    #     max_len = sorted(uni_value_dic.items(), key=lambda x: x[1], reverse=True)[0][1]
    #     return max_len

# @lc code=end

