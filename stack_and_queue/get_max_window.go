package stack_and_queue

func getMaxWindowNum(arr []int, w int) []int {
	var qMax []int
	var res []int
	for i := range arr {
		if len(qMax) > 0 {
			for j := len(qMax) - 1; j >= 0; j-- {
				idx := qMax[j]
				if arr[idx] > arr[i] {
					break
				}
				qMax = qMax[:j]
			}
		}
		qMax = append(qMax, i)
		if qMax[0] < i-w+1 {
			qMax = qMax[1:]
		}
		if i >= w-1 {
			res = append(res, arr[qMax[0]])
		}
	}
	return res
}
