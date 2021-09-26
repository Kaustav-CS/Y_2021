/******************************************************************************

Link:   https://leetcode.com/problems/zigzag-conversion/

6. ZigZag Conversion
*******************************************************************************/

func convert(s string, numRows int) string {
   length := len(s)

    bufferSlice := make([]bytes.Buffer, numRows)

  for i := 0; i < len(bufferSlice); i++ {
	  var buffer bytes.Buffer
	  bufferSlice[i] = buffer
  }
   i := 0
  for i < length {
  	  for index := 0; index < numRows && i < length; index++ {
		  bufferSlice[index].WriteString(string(s[i]))
  	  	   i++
	  }
	  for index := numRows - 2; index >= 1 && i < length; index-- {
		  bufferSlice[index].WriteString(string(s[i]))
	  	i++
	  }
  }
  for index := 1; index < len(bufferSlice); index++{
	  bufferSlice[0].WriteString(bufferSlice[index].String())
  }
  return bufferSlice[0].String()
}
