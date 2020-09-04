// O(n^2) notation
func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
    i := int32(0)
    ans := int32(0)
    for i < n {
        j := int32(0)
        curr := ar[i]
        ar := ar[i+1:]
        for  j < int32(len(ar)) {
            if pair := curr+ar[j]; pair%k == 0 {
                ans++
            }
            j++   
        }
        i++
    }
    return ans;
}
