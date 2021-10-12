package ch13

import (
	"fmt"
	"time"
)

type Command interface {
	execute() error
}

type ActiveObject struct {
	arr []Command
}

func NewActiveObject() *ActiveObject {
	return &ActiveObject{}
}

func (a *ActiveObject) addCommand(c Command) {
	a.arr = append(a.arr, c)
}

func (a *ActiveObject) Run() {
	for len(a.arr) != 0 {
		c := a.arr[0]
		a.arr = a.arr[1:]
		c.execute()
		//fmt.Println("command",commandExecuted)
	}
}

type SleepCommand struct {
	wakeupCommand        Command
	engine               *ActiveObject
	sleepTime, startTime int64
	started              bool
}

func NewSleepCommand(milliseconds int64, object *ActiveObject, wakeupCommand Command) *SleepCommand {
	return &SleepCommand{
		sleepTime:     milliseconds * 1e6,
		engine:        object,
		wakeupCommand: wakeupCommand,
	}
}

func (a *SleepCommand) execute() error {
	curr := time.Now().UnixNano()
	//fmt.Println(curr,a.startTime,)
	time.Sleep(time.Millisecond)
	if !a.started {
		a.started = true
		a.startTime = curr
		a.engine.addCommand(a)
	} else if curr-a.startTime < a.sleepTime {
		a.engine.addCommand(a)
	} else {
		a.engine.addCommand(a.wakeupCommand)
	}
	return nil
}

var commandExecuted bool = false

type WakeupCommand struct {
}

func NewWakeupCommand() *WakeupCommand {
	return &WakeupCommand{}
}
func (w *WakeupCommand) execute() error {
	commandExecuted = true
	return nil
}

var stop bool = false

type StopCommand struct {
}

func NewStopCommand() *StopCommand {
	return &StopCommand{}
}
func (w *StopCommand) execute() error {
	stop = true
	return nil
}

type DelayedTyper struct {
	itsDelay int64
	itsChar  uint8
	engine   *ActiveObject
}

func NewDelayedTyper(engine *ActiveObject, delay int64, c uint8) *DelayedTyper {
	return &DelayedTyper{
		itsDelay: delay,
		itsChar:  c,
		engine:   engine,
	}
}

func (d *DelayedTyper) delayAndRepeat() {
	d.engine.addCommand(NewSleepCommand(d.itsDelay, d.engine, d))
}
func (d *DelayedTyper) execute() error {
	fmt.Printf("%c", d.itsChar)
	if !stop {
		d.delayAndRepeat()
	}
	return nil
}
