func getSneakyNumbers(nums []int) []int {
     var array []int
    for i:=0; i<len(nums);i++{
        for j:=i+1; j<len(nums); j++{
            if nums[i]==nums[j]{
                array = append(array, nums[i])
            }
        }
    }

    return array
    
}