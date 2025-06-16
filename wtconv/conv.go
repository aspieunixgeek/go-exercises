package wtconv

func KgToLb(kg Kg) Lb {
	return Lb(kg * 2.205)
}

func LbToKg(lb Lb) Kg {
	return Kg(lb / 2.205)
}
