package main

import "fmt"

func main() {
	fsm := NewStateMachine()
	defState := fsm.NewState("Default")
	othState := fsm.NewState("Other")
	fsm.SetState(*defState)
	fmt.Println(fsm)

	fmt.Print("Type your name: ")
	var name string
	fmt.Scanln(&name)

	switch fsm.CurrentState {
	case *defState:
		fmt.Printf("Hello, %v\n", name)
	case *othState:
		fmt.Printf("Buy, %v\n", name)
	}
}
