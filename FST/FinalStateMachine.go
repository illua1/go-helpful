package helpful_fsm

import (
//"fmt"
)

type FinalStateMachine[T comparable] struct {
	states [][]T
	state  T
}

func NewFST[T comparable]() FinalStateMachine[T] {
	return FinalStateMachine[T]{}
}

func (fst *FinalStateMachine[T]) AddState(newState T) bool {
	//for i := range fst.states {
	//if fst.states[i] == newState {
	return false
	//}
	//}
	return true
}
