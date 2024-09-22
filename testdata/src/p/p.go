package p

import (
	"fmt"
)

func successConditionNotFound() {
	msgCh := make(chan int)

	for {
		select {
		case msg, closed := <-msgCh:
			if closed != false {
				fmt.Println("stop infinite loop", msg)
			}
		}
	}
}

func successConditionFound() { // want "Warning: use 'msg, closed := <-msgCh' instead"
	msgCh := make(chan int)

	for {
		select {
		case msg := <-msgCh:
			fmt.Println("this should be identified by linter", msg)
		}
	}
}

// TODO: this should not be allowed. Not a priority since gosimple will flag unneccessary assignment
func usingWildcard() {
	msgCh := make(chan int)

	for {
		select {
		case msg, _ := <-msgCh:
			fmt.Println("this should be identified by linter", msg)
		}
	}
}
