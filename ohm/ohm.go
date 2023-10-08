package ohm

func Ohm(v, r, i float32) float32 {
	if v > 0 && r > 0 && i > 0{
		return 0
	} else if r != 0 && i != 0{
		result := r * i
        return result
	} else if v != 0 && r != 0{
		result := v / r
        return result
	} else if v != 0 && i != 0{
		result := v / i
        return result
	} else {
		return 0
	}
}