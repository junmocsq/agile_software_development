package ch06

type Game struct {
	firstThrowInFrame bool
	itsCurrentFrame   int

	itsScorer *Scorer
}

func NewGame() *Game {
	return &Game{
		firstThrowInFrame: true,
		itsCurrentFrame:   0,
		itsScorer:         NewScorer(),
	}
}

func (f *Game) Add(pins int) {
	f.itsScorer.Add(pins)
	f.adjustCurrentFrame(pins)
}

func (f *Game) adjustCurrentFrame(pins int) {
	if f.lastBallInFrame(pins) {
		f.advanceFrame()
	} else {
		f.firstThrowInFrame = false
	}
}

func (f *Game) lastBallInFrame(pins int) bool {
	return f.strike(pins) || !f.firstThrowInFrame
}

func (f *Game) strike(pins int) bool {
	return f.firstThrowInFrame && pins == 10
}

func (f *Game) advanceFrame() {
	f.itsCurrentFrame++
	if f.itsCurrentFrame > 10 {
		f.itsCurrentFrame = 10
	}
	f.firstThrowInFrame = true
}

func (f *Game) GetScore() int {
	return f.ScoreForFrame(f.itsCurrentFrame)
}

func (f *Game) GetCurrentFrame() int {
	return f.itsCurrentFrame
}

func (f *Game) ScoreForFrame(frame int) int {
	return f.itsScorer.ScoreForFrame(frame)
}
