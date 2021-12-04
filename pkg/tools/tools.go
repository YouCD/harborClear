package tools

func GetPageSize(counter int) int {
	if counter > 100 {
		if counter%100 == 0 {
			return counter / 100
		} else {
			pageSize := counter / 100
			return pageSize + 1
		}
	}
	return 1
}
