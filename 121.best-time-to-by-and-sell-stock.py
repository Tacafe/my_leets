class Solution:
    def maxProfit(self, prices: list[int]) -> int:
        profit = 0
        i = 0
        min_price = prices[0]

        for i in range(1, len(prices)):
            if prices[i] < min_price:
                min_price = prices[i]
            else:
                profit = max(profit, prices[i] - min_price)
        return profit

if __name__ == '__main__':
    test = [2,1,2,0,1]
    print(Solution().maxProfit(test))
