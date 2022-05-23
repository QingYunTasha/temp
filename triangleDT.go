func minimumTotal(triangle [][]int) int {
    final := len(triangle)-1
    for i, arrDown := range triangle{
        if i == final{
            break
        }
        arrUp := triangle[i+1]
        
        arrUp[0] += arrDown[0]
        f := len(arrDown) - 1
        arrUp[f+1] += arrDown[f]
        
        for j := 0; j < f; j++{
            if arrDown[j] <= arrDown[j+1]{
                arrUp[j+1] += arrDown[j]
            }else{
                arrUp[j+1] += arrDown[j+1]
            }
        }
    }
    Up := triangle[final]
    re := Up[0]
    for i := 1; i < len(Up); i++{
        if re > Up[i]{
            re = Up[i]
        }
    }
    return re
}
