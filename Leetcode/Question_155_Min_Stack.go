/******************************************************************************
Link:   https://leetcode.com/problems/min-stack/
155. Min Stack

Your input
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]

Output
[null,null,null,null,-3,null,0,-2]

Expected
[null,null,null,null,-3,null,0,-2]

DATE:    2021, December 31
১৫ পৌষ,১৪২৮
*******************************************************************************/
type MinStack struct {
     Min int
    Val int
    Valid bool
    Next *MinStack
}


func Constructor() MinStack {
    return MinStack{0,0,false,nil}
}


func (this *MinStack) Push(val int)  {
    if this.Valid {
        min := int(math.Min(float64(val), float64(this.Min)))
        m := *this
        this.Min = min
        this.Val = val
        this.Valid = true
        this.Next = &m
    } else {
        this.Min = val
        this.Val = val
        this.Valid = true
    }
}


func (this *MinStack) Pop()  {
    if this.Valid {
        if this.Next != nil {
            *this = *this.Next
        } else {
            this.Val = 0
            this.Min = 0
            this.Valid = false
        }
        
    }
}


func (this *MinStack) Top() int {
    return this.Val
}


func (this *MinStack) GetMin() int {
    return this.Min
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

