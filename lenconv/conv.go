package lenconv

func FtToMet(ft Foot) Meter {
	return Meter(ft / 3.281)
}

func MetToFt(m Meter) Foot {
	return Foot(m * 3.281)
}
