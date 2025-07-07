package main

import "time"

// Mantains test container alive so it can attach
func main() {
	for {
		time.Sleep(time.Hour * 1)
	}
}
