package ch06

type Scorer struct {
	itsThrows []int
	ball      int
}

func NewScorer() *Scorer {
	return &Scorer{}
}

func (f *Scorer) Add(pins int) {
	f.itsThrows = append(f.itsThrows, pins)
}

func (f *Scorer) ScoreForFrame(frame int) int {
	score := 0
	f.ball = 0
	for currentFrame := 0; currentFrame < frame; currentFrame++ {
		if f.strike() {
			score += 10 + f.nextTwoBallsForStrike()
			f.ball++
		} else if f.spare() {
			score += 10 + f.nextBallForSpare()
			f.ball += 2
		} else {
			score += f.twoBallInFrame()
			f.ball += 2
		}
	}
	return score
}

func (f *Scorer) strike() bool {
	return f.itsThrows[f.ball] == 10
}

func (f *Scorer) nextTwoBallsForStrike() int {
	return f.itsThrows[f.ball+1] + f.itsThrows[f.ball+2]
}

func (f *Scorer) spare() bool {
	return f.itsThrows[f.ball]+f.itsThrows[f.ball+1] == 10
}

func (f *Scorer) nextBallForSpare() int {
	return f.itsThrows[f.ball+2]
}

func (f *Scorer) twoBallInFrame() int {
	return f.itsThrows[f.ball] + f.itsThrows[f.ball+1]
}
