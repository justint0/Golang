/*
Enumerated types (enums) are a special case of sum types. An enum is a type that has a fixed number of possible values, each with a distinct name. Go doesn't have an enum type as a distinct language feature, but enums are simple to implement using existing language idioms.
*/

package main

import "fmt"

// This line creates a new type called ServerState, based on the underlying type int. This means each state will be represented by an integer.
type ServerState int

// This block defines the possible values for the ServerState type as constants.
const (
		// This assigns the first state (StateIdle) to the value 0. iota is a special keyword in Go that starts at 0 and automatically increments for each constant within a const block.
    StateIdle ServerState = iota
    StateConnected
    StateError
    StateRetrying
)

// This creates a map called stateName that associates each ServerState value with a human-readable string.
var stateName = map[ServerState]string{
    StateIdle:      "idle",
    StateConnected: "connected",
    StateError:     "error",
    StateRetrying:  "retrying",
}

// This function attaches a method called String() to the ServerState type. This method allows you to easily print the state as a string instead of its underlying integer value. It does this by looking up the state in the stateName map.
func (ss ServerState) String() string {
    return stateName[ss]
}

func main() {
    ns := transition(StateIdle)
    fmt.Println(ns)

    ns2 := transition(ns)
    fmt.Println(ns2)
}

// This function simulates a state transition for the server. It takes the current state as input and returns the next state based on a set of rules defined in the switch statement.
func transition(s ServerState) ServerState {
    switch s {
    case StateIdle:
        return StateConnected
    case StateConnected, StateRetrying:
        return StateIdle
    case StateError:
        return StateError
    default:
        panic(fmt.Errorf("unknown state: %s", s))
    }
}