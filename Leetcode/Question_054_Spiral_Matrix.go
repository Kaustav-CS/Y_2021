/******************************************************************************
Link:   https://leetcode.com/problems/spiral-matrix/
54. Spiral Matrix

Your input
[[1,2,3],[4,5,6],[7,8,9]]

Output
[1,2,3,6,9,8,7,4,5]

Expected
[1,2,3,6,9,8,7,4,5]


DATE:    2021, October 20
০৩ কার্তিক,১৪২৮
*******************************************************************************/
func spiralOrder(matrix [][]int) []int {
    rows := len(matrix)
    var columns int
	if rows == 0 {
		columns = 0
	} else {
		columns = len(matrix[0])
	}
	numberOfIterations := rows * columns
	ret := []int{}
	xAxis, yAxis := 0, 0
	nodesTravelled := make(map[string]bool)
	direction := "east"
	for i := 0; i < numberOfIterations; i++ {
		ret = append(ret, matrix[yAxis][xAxis])
		nodesTravelled[strconv.Itoa(xAxis)+"_"+strconv.Itoa(yAxis)] = true
		nextXaxis, nextYaxis := marchForward(direction, xAxis, yAxis)
		if (nextXaxis < 0 || nextXaxis >= columns) || (nextYaxis < 0 || nextYaxis >= rows) || nodesTravelled[strconv.Itoa(nextXaxis)+"_"+strconv.Itoa(nextYaxis)] {
			changeDir(&direction)
			nextXaxis, nextYaxis = marchForward(direction, xAxis, yAxis)
		}
		xAxis, yAxis = nextXaxis, nextYaxis
	}

	return ret
}

func marchForward(direction string, xAxis int, yAxis int) (int, int) {
	switch direction {
	case "east":
		xAxis++
		break
	case "south":
		yAxis++
		break
	case "west":
		xAxis--
		break
	case "north":
		yAxis--
		break
	}
	return xAxis, yAxis
}

func changeDir(direction *string) {
	switch *direction {
	case "east":
		*direction = "south"
		break
	case "south":
		*direction = "west"
		break
	case "west":
		*direction = "north"
		break
	case "north":
		*direction = "east"
		break
	}
}
