package workflow

import (
	"fmt"
	"os"
	"time"

	"github.com/dymm/orchestrators/messageQ/pkg/messaging"
)

//StartSessionTimeoutChecking start the timeout control loop
func StartSessionTimeoutChecking(queue messaging.Queue, timeoutDestination string) {

	go timeoutLoop(queue, timeoutDestination)
}

func timeoutLoop(queue messaging.Queue, timeoutDestination string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in timeoutLoop", r)
			os.Exit(-2)
		}
	}()

	for {
		workitemToSend := listAllTimeoutedSession(queue, timeoutDestination)
		for _, workitem := range workitemToSend {
			err := queue.Send(timeoutDestination, workitem)
			if err != nil {
				fmt.Println("Can't send a timeout message", err)
			}
		}
		time.Sleep(1 * time.Second)
	}

}

func listAllTimeoutedSession(queue messaging.Queue, timeoutDestination string) []messaging.WorkItem {

	sessionListMutex.Lock()
	defer sessionListMutex.Unlock()

	sessionCount := len(sessionList)
	fmt.Printf("%d session in memory\n", len(sessionList))

	workItemToSend := make([]messaging.WorkItem, 0, sessionCount)
	now := time.Now()
	for _, session := range sessionList {
		if session.CurrentStep.TimeoutTime.IsZero() == false && now.After(session.CurrentStep.TimeoutTime) {
			session.CurrentStep.Workitem.GetValues()["error"] = `{"message":"Timeout"}`
			workItemToSend = append(workItemToSend, session.CurrentStep.Workitem)
		}
	}
	return workItemToSend
}
