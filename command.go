package agile

import (
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
		sleepTime:     milliseconds *1e6,
		engine:        object,
		wakeupCommand: wakeupCommand,
	}
}

func (a *SleepCommand) execute() error {
	curr := time.Now().UnixNano()
	//fmt.Println(curr,a.startTime,)
	//time.Sleep(time.Millisecond*10)
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
func (w *WakeupCommand) execute()error {
	commandExecuted = true
	return nil
}
