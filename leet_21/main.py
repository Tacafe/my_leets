# Definition for singly-linked list.

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    def mergeTwoLists(self, list1: Optional[ListNode], list2: Optional[ListNode]) -> Optional[ListNode]:
        if list1 == [] and list2 == []:
            return []
        elif list1 == [] and list2 != []:
            return list2
        elif list1 != [] and list2 == []:
            return list1

        merged_list = []
        i = 0
        j = 0
        is_1 = True
        while not list1.next is None and not list2.next is None:
            if not list1[i].next is None:
                if list2[j].next is None:
                    megerged_list += list1[i:]
                    return merged_list
                elif list1[i].val <= list2[j].val:
                    megerged_list.append(list1[i])
                    i += 1
                else:
                    megerged_list.append(list2[j])
                    j += 1
            else:
                megerged_list += list2[j:]
                return merged_list
