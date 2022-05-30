package api

func AllPage(allr int, limit int) int {
	var (
		allp int
	)
	if limit == 0 {
		limit = 1
	}
	ost := allr % limit
	if ost != 0 {
		allp = allr/limit + 1
	} else {
		allp = allr / limit
	}
	return allp
}
