package hammingdist

func CalculateDistance(str1, str2 string) float64 {
	if len(str1) != len(str2) {
		panic("Invalid String Length")
	}
	b1 := []byte(str1)
	b2 := []byte(str2)
	checks := []byte{
		byte(1), byte(2),
		byte(4), byte(8),
		byte(16), byte(32),
		byte(64), byte(128),
	}
	distance := 0.0

	for i := range b1 {
		tmp := b1[i] ^ b2[i]
		for j := range checks {
			if uint32(tmp&checks[j]) > 0 {
				distance += 1
			}
		}
	}

	return distance
}
