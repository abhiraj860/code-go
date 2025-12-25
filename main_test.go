package main

import "testing"

func TestRaceCondition(t *testing.T) {
	var state int32

	for i := 0; i < 10; i++ {
		go func() { 
			state += int32(i) 
		}()
	}
}

// func TestPackItems(t *testing.T) {
// totalItems := PackItems(0)
// expectedTotal := 2000
// if totalItems != expectedTotal {
// t.Errorf("Expected total: %d, Actual value %d", expectedTotal, totalItems)
// }
// }
