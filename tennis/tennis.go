package tennis

import "fmt"

// ShowScore for show tennis score
func ShowScore(scoreA, scoreB int) (result string, err error) {
	score := map[int]string{
		0:  "LOVE",
		15: "FIFTEEN",
		30: "THIRTY",
		40: "FORTY",
	}

	if _, ok := score[scoreA]; !ok {
		return "", fmt.Errorf("Invalid score player A expect `|0, 15, 30, 40|` got `%d`", scoreA)
	}
	if _, ok := score[scoreB]; !ok {
		return "", fmt.Errorf("Invalid score player B expect `|0, 15, 30, 40|` got `%d`", scoreB)
	}

	res := fmt.Sprintf("%s - %s", score[scoreA], score[scoreB])

	if scoreA == 40 {
		res += "\t-> A WIN"
		return res, nil
	}

	if scoreB == 40 {
		res += "\t-> B WIN"
		return res, nil
	}

	return res, nil
}
