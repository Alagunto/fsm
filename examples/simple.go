// +build ignore

package main

import (
	"fmt"
	"github.com/Alagunto/fsm"
)

func main() {
	fsm := fsm.NewFSM(
		"A",
		fsm.Events{
			{Name: "go", Src: []string{"A"}, Dst: "B"},
			{Name: "loop", Src: []string{"C"}, Dst: "C"},
		},
		fsm.Callbacks{
			"before_go": func(e *fsm.Event) {
				e.FSM.SetDst("C")
			},
			"enter_C": func(e *fsm.Event) {
				fmt.Println("Actually, i'm triggered")
			},
		},
	)

	fmt.Println(fsm.Current())

	err := fsm.Event("go")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(fsm.Current())

	err = fsm.Event("loop")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fsm.Current())


	//err = fsm.Event("lol")
	//if err != nil {
	//	fmt.Println(err)
	//}

	//fmt.Println(fsm.Current())
}
