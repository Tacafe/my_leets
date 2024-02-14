#
# @lc app=leetcode id=133 lang=python3
#
# [133] Clone Graph
#

# @lc code=start
"""
# Definition for a Node.
class Node:
    def __init__(self, val = 0, neighbors = None):
        self.val = val
        self.neighbors = neighbors if neighbors is not None else []
"""

class Solution:
    def deep_copy(self, node):
        if not node: return None
        cp = Node(node.val, [])
        self.nodes[cp.val] = cp

        for n in node.neighbors:
            if not n.val in self.nodes:
                cp.neighbors.append(self.deep_copy(n))
            else:
                cp.neighbors.append(self.nodes[n.val])

        return cp

    def cloneGraph(self, node: 'Node') -> 'Node':
        self.nodes = {}
        return self.deep_copy(node)


# @lc code=end

