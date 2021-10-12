package agile

import (
	"fmt"
	"testing"
	"time"
)

func TestSleepCommand(t *testing.T)  {
	wakeup :=  NewWakeupCommand()
	activeObject := NewActiveObject()
	sleepCommand := NewSleepCommand(2000,activeObject,wakeup)
	activeObject.addCommand(sleepCommand)

	start := time.Now().UnixNano()
	activeObject.Run()
	stop := time.Now().UnixNano()
	sleepTime := stop-start
	fmt.Println(sleepTime,commandExecuted)
}