package domain

func (us *UserSummary) PercCorrect() int {
	return int(float64(us.CorrectPredictions) / float64(us.CorrectPredictions+us.IncorrectPredictions) * 100.0)
}

func (us *UserSummary) PercIncorrect() int {
	return 100 - us.PercCorrect()
}
