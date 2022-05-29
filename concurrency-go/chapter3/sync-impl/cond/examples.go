package cond

import "sync"

func CondExample() {
	c := sync.NewCond(&sync.Mutex{})
	//we lock the Locker for this condition.
	// This is necessary because the call to Wait automatically calls Unlock on the Locker when entered.

	c.L.Lock()
	for conditionTrue() == false {
		//we wait to be notified that the condition has occurred. This is a blocking call and the goroutine will be suspended.
		c.Wait()
	}
	// we unlock the Locker for this condition. This is necessary because when the call to Wait exits,
	// it calls Lock on the Locker for the condition.
	c.L.Unlock()

}

func conditionTrue() bool {
	return true
}
