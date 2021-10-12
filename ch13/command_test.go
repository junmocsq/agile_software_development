package ch13

import (
	"fmt"
	"testing"
	"time"
)

func TestSleepCommand(t *testing.T) {
	wakeup := NewWakeupCommand()
	activeObject := NewActiveObject()
	sleepCommand := NewSleepCommand(1000, activeObject, wakeup)
	activeObject.addCommand(sleepCommand)

	start := time.Now().UnixNano()
	activeObject.Run()
	stop := time.Now().UnixNano()
	sleepTime := stop - start
	fmt.Println(sleepTime, commandExecuted)

	activeObject.addCommand(NewDelayedTyper(activeObject, 100, '1'))
	activeObject.addCommand(NewDelayedTyper(activeObject, 300, '3'))
	activeObject.addCommand(NewDelayedTyper(activeObject, 500, '5'))
	activeObject.addCommand(NewDelayedTyper(activeObject, 700, '7'))
	activeObject.addCommand(NewSleepCommand(20000, activeObject, NewStopCommand()))
	activeObject.Run()
}
