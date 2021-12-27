/******************************************************************************
Link:   https://leetcode.com/problems/max-points-on-a-line/
149. Max Points on a Line

Your input
[[1,1],[2,2],[3,3]]

Output
3

Expected
3

DATE:    2021, December 27
১১ পৌষ,১৪২৮
*******************************************************************************/
func maxPoints(points [][]int) int {
    n := len(points)
    if n < 2 {
        return n
    }
    res := 0 
    for i:=0; i < n; i ++ {
        for j := i+1 ; j < n; j++ {
            if r := find(points[i], points[j], points); r > res {
                res = r
            }
        }
    }
    return res
}



func find(p1 []int, p2 []int, points [][]int) int {
    res := 0 
    
    // same points
    if p1[0] == p2[0] && p1[1] == p2[1] {
        for _, v := range points {
            if v[0] == p1[0] && v[1] == p1[1] {
                res++
            }
        }
        return res
    }
    
    // diff points
    for _, v := range points {
        if sameLine(p1, p2, v) {
            res++
        }
    }
    
    return res
}


// three points in one line
func sameLine(p1 []int, p2 []int, target []int) bool {
	if target[0] == p1[0] && target[1] == p1[1] {
		return true
	}
	if target[0] == p2[0] && target[1] == p2[1] {
		return true
	}
    
	if (target[0]-p1[0])*(target[1]-p2[1]) != (target[0]-p2[0])*(target[1]-p1[1]) {
		return false
	}

	return true
}
