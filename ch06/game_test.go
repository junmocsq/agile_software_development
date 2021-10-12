package ch06

import "testing"

func TestGame(t *testing.T) {
	g := NewGame()
	if 0 != g.GetScore() {
		t.Errorf("没有投掷时 wanted:0 actual:%d \n", g.GetScore())
	}

	g = NewGame()
	g.Add(5)
	g.Add(4)
	if 9 != g.GetScore() {
		t.Errorf("两次投掷时 wanted:9 actual:%d \n", g.GetScore())
	}

	g = NewGame()
	g.Add(5)
	g.Add(4)
	g.Add(7)
	g.Add(2)
	if 18 != g.GetScore() {
		t.Errorf("四次投掷时 wanted:18 actual:%d \n", g.GetScore())
	}
	if 9 != g.ScoreForFrame(1) {
		t.Errorf("第一轮投掷得分 wanted:9 actual:%d \n", g.ScoreForFrame(1))
	}
	if 18 != g.ScoreForFrame(2) {
		t.Errorf("第二轮投掷得分 wanted:18 actual:%d \n", g.ScoreForFrame(2))
	}

	g = NewGame()
	g.Add(3)
	g.Add(7)
	g.Add(3)
	g.Add(2)
	if 13 != g.ScoreForFrame(1) {
		t.Errorf("补中投掷第一轮得分 wanted:13 actual:%d \n", g.ScoreForFrame(1))
	}
	if 18 != g.ScoreForFrame(2) {
		t.Errorf("补中投掷第二轮得分 wanted:18 actual:%d \n", g.ScoreForFrame(2))
	}
	if 18 != g.GetScore() {
		t.Errorf("得分 wanted:18 actual:%d \n", g.GetScore())
	}

	g = NewGame()
	g.Add(10)
	g.Add(3)
	g.Add(6)
	if 19 != g.ScoreForFrame(1) {
		t.Errorf("全中投掷第一轮得分 wanted:19 actual:%d \n", g.ScoreForFrame(1))
	}
	if 28 != g.GetScore() {
		t.Errorf("得分 wanted:28 actual:%d \n", g.GetScore())
	}
	g = NewGame()
	for i := 0; i < 12; i++ {
		g.Add(10)
	}
	if 300 != g.GetScore() {
		t.Errorf("得分 wanted:300 actual:%d \n", g.GetScore())
	}

	g = NewGame()
	for i := 0; i < 9; i++ {
		g.Add(0)
		g.Add(0)
	}
	g.Add(2)
	g.Add(8)
	g.Add(10)
	if 20 != g.GetScore() {
		t.Errorf("得分 wanted:20 actual:%d \n", g.GetScore())
	}
}
