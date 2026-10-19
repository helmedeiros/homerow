package scoring

const MaxStars = 5

type Result struct {
	Accuracy float64
	WPM      float64
}

var accuracyBands = []float64{80, 88, 94, 97}

func Stars(r Result, goalWPM float64) int {
	stars := 0
	for _, band := range accuracyBands {
		if r.Accuracy >= band {
			stars++
		}
	}
	if stars == len(accuracyBands) && r.WPM >= goalWPM {
		stars = MaxStars
	}
	return stars
}

func Passed(stars int) bool {
	return stars > 0
}
