/******************************************************************************
Link:   https://leetcode.com/problems/gas-station/
134. Gas Station

Your input
[1,2,3,4,5]
[3,4,5,1,2]

Output
3

Expected
3

DATE:    2021, November 17
০১ অগ্রহায়ণ,১৪২৮
*******************************************************************************/
func canCompleteCircuit(gas []int, cost []int) int {
    start := 0
    total := 0
    tank := 0
    for i := 0; i < len(gas); i++ {
        tank += gas[i] - cost[i]
        if tank < 0 {
            start = i + 1
            total += tank
            tank = 0
        }
    }
    if total + tank < 0 {
        return -1
    }
    return start
}

