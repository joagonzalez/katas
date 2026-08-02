package hamming

import "errors"

var DistanceError = errors.New("strings must be of equal length")

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, DistanceError
	}

	distance := 0

	for i := range a {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}
