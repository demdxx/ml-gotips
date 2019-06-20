package pairwise

func minus(params ...float64) (res float64) {
	res = params[0]
	for i := 1; i < len(params); i++ {
		res -= params[i]
	}
	return res
}
