import random
from typing import List


def shuffle(nums: List[int]) -> List[int]:
    length = len(nums)
    for i in range(length):
        j = random.randrange(i, length)
        nums[i], nums[j] = nums[j], nums[i]
    return nums


res = shuffle([1, 2, 3, 4, 5, 6, 7])
print(res)
