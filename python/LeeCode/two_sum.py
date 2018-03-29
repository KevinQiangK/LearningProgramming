class Solution:
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """

        for i in range(len(nums)):
            remain = target - nums[i]
            for j in range(i+1, len(nums)):
                if nums[j] == remain:
                    return [i, j]
        return []

    def twoSum2(self, nums, target):
        cache = {}
        for i in range(len(nums)):
            complement = target - nums[i]
            if complement in cache:
                return [cache[complement], i]
            cache[nums[i]] = i


if __name__ == "__main__":
    solution = Solution()
    print(solution.twoSum([2,7,11,15], 9))
    print(solution.twoSum([0,4,3,0], 0))
    print(solution.twoSum([-1,-2,-3,-4,-5], -8))

    print(solution.twoSum2([2, 7, 11, 15], 9))
    print(solution.twoSum2([0, 4, 3, 0], 0))
    print(solution.twoSum2([-1, -2, -3, -4, -5], -8))