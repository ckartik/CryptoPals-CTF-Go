package hammingdist

import (
	"container/heap"
	"math"
)

var englishModel = []float64{
	0.0651738, 0.0124248, 0.0217339, 0.0349835, //'A', 'B', 'C', 'D',...
	0.1041442, 0.0197881, 0.0158610, 0.0492888,
	0.0558094, 0.0009033, 0.0050529, 0.0331490,
	0.0202124, 0.0564513, 0.0596302, 0.0137645,
	0.0008606, 0.0497563, 0.0515760, 0.0729357,
	0.0225134, 0.0082903, 0.0171272, 0.0013692,
	0.0145984, 0.0007836, 0.1918182}

const (
	captialStart = 0x61
	captialEnd   = 0x7a
	lowerStart   = 0x41
	lowerEnd     = 0x5a
)

// plaintextscore finds the manhatan distance between text and the english model.
// Our formula finds the p = 0.5 norm of the difference between the english model and the normailzed count of each letter.
// It also adds a lambda weight factor to penalize the occurence of non-english output.
func PlaintextScore(text string) float64 {
	stringVector := make([]float64, len(englishModel))
	lambda := 0.0
	plaintextSize := len(text)
	for i := range text {
		hex := byte(text[i])
		if hex >= captialStart && hex <= captialEnd {
			stringVector[int(hex-captialStart)] += 0.99
		} else if hex >= lowerStart && hex <= lowerEnd {
			stringVector[int(hex-lowerStart)]++
		} else if hex == byte(' ') || hex == byte('.') {
		} else {
			lambda += float64(plaintextSize)
		}
	}

	for i := range stringVector {
		stringVector[i] /= float64(plaintextSize)
	}

	return pnorm(stringVector, englishModel, 0.5) + (lambda / float64(plaintextSize))
}

// pnorm find the norm with pvalue p of the difference between vector1 and vector2.
// p=0.5 is the most effective because similarty on a few features is a strong
// inidicator of similarty between the enlgish model and the small piece of text.
func pnorm(vector1, vector2 []float64, p float64) float64 {
	dist := 0.0
	for i := range vector1 {
		dist += math.Pow(math.Abs(vector1[i]-vector2[i]), p)
	}

	return math.Pow(dist, 1/p)
}

// CalculateDistance will find the hamming distance between two strings.
// The length of the paramters str1 and str2 should be the same.
// It will return the hamming distance in 64 byte floating point.
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

/********* START OF MIN-HEAP DEFN *************/
// TODO(@ckartik): Test the min-heap.
type Guess struct {
	distanceMeasure float64
	keysize         int
}

type GuessMinHeap []Guess

func (h GuessMinHeap) Len() int      { return len(h) }
func (h GuessMinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// We want this to be a min-heap.
func (h GuessMinHeap) Less(i, j int) bool { return h[i].distanceMeasure < h[j].distanceMeasure }

func (h *GuessMinHeap) Push(x interface{}) {
	*h = append(*h, x.(Guess))
}

func (h *GuessMinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	*h = old[0 : n-1]

	return old[n-1]
}

/********* END OF MIN-HEAP DEFN *************/

// GuessKeySize returns a best guess of the keysiz returns a best guess of the keysize
// Try using a min-heap of size 5.
// TODO: Fix this, it's not providing the correct values.
func GuessKeySize(cipher []byte) GuessMinHeap {
	// Init paramters for discovery.
	h := &GuessMinHeap{{math.Inf(1), 0},
		{math.Inf(1), 0}, {math.Inf(1), 0},
		{math.Inf(1), 0}, {math.Inf(1), 0}}

	heap.Init(h)
	// Discover value for the keysize.
	for keySize := 1; keySize < 41; keySize++ {
		distanceMeasure := 0.0
		base := keySize
		chunk1 := string(cipher[base : base+keySize])
		base += keySize
		chunk2 := string(cipher[base : base+keySize])
		base += keySize
		chunk3 := string(cipher[base : base+keySize])
		base += keySize
		chunk4 := string(cipher[base : base+keySize])

		distanceMeasure += CalculateDistance(chunk1, chunk2)
		distanceMeasure += CalculateDistance(chunk1, chunk4)
		distanceMeasure += CalculateDistance(chunk2, chunk3)
		distanceMeasure += CalculateDistance(chunk2, chunk4)
		distanceMeasure += CalculateDistance(chunk3, chunk4)
		distanceMeasure += CalculateDistance(chunk1, chunk3)

		result := float64(distanceMeasure / float64(keySize))
		guess := Guess{result, keySize}
		h.Push(guess)
		h.Pop()
	}

	return *h
}
