package offer

func findNumberIn2DArray(matrix [][]int, target int) bool {
    if matrix == nil {
        return false
        }
    li := len(matrix)    
    if li==0 {
        return false
        }
    lj := len(matrix[0])
    if lj==0  {
        return false
        }
    
    for i,j := 0,lj-1;i<li && j>=0 ; {
        if matrix[i][j] == target {
            return true
            }
        if matrix[i][j] > target {
            j--
            } else {
            i++
            }
        }
    return false
}