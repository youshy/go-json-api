package main

import "fmt"

var currentId int

var effects Effects

// Give us some seed data
func init() {
	RepoCreateEffect(Effect{Brand: "MXR", Name: "Blue Box", Type: "Fuzz"})
	RepoCreateEffect(Effect{Brand: "Way Huge", Name: "Swollen Pickle", Type: "Fuzz"})
}

func RepoFindEffect(id int) Effect {
	for _, t := range effects {
		if t.Id == id {
			return t
		}
	}
	// return empty Effect if not found
	return Effect{}
}

// TODO: Google if that's production-ready
func RepoCreateEffect(t Effect) Effect {
	currentId += 1
	t.Id = currentId
	effects = append(effects, t)
	return t
}

func RepoDeleteEffect(id int) error {
	for i, t := range effects {
		if t.Id == id {
			effects = append(effects[:i], effects[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Effect with id of %d to delete", id)
}
