package ch24

import "testing"

func TestDriver(t *testing.T) {
	sour := NewMockTimeSource()
	sink := NewMockTimeSink()
	//dr:= NewClockDriver(sour)
	//dr.addTimeSink(sink)
	//dr.addTimeSink(sink)
	//dr.addTimeSink(sink)
	sour.RegisterObserver(sink)
	sour.RegisterObserver(sink)
	sour.RegisterObserver(sink)
	sour.SetTime(12, 33, 44)
	sour.SetTime(15, 05, 06)
	//sink.PrintTime()
}
