def subarraySum(nums, k):
        # 求sum(j,i) = k, sum[i]表示从nums[0]...nums[i]的和
        # 如果sum[i] - k == 某一个sum[j] 那么说明从nums[j+1] + ... + num[i]的和就是k
        N = len(nums)
        count = 0
        # 创建一个字典，如果sum[i] - k == 某一个sum[j]，即sum[j]的值在字典中，count+sum[j]在字典中的值
        # 而不是+1，因为可能会有多个j开始一直到i的和为k
        sum_all = {0: 1}
        sum_i = 0
        for i in range(N):
            sum_i += nums[i]
            if sum_i - k in sum_all:
                count += sum_all[sum_i - k]
            if sum_i not in sum_all:
                sum_all[sum_i] = 0
            sum_all[sum_i] += 1

        return count