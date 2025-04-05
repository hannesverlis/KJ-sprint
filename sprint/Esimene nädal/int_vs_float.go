package sprint

func IntVsFloat(i int, f float32) string {

	if float32(i) == float32(f) { // peale if kõigepealt statement, siis loogiline sulg ja väljund
		return "Same"
	}

	if float32(i) < float32(f) {
		return "Float"
	}

	if float32(i) > float32(f) {
		return "Integer"
	}

	return "error"
}
