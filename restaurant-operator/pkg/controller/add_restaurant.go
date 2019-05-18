package controller

import (
	"github.com/ruromero/presentations/restaurant-operator/pkg/controller/restaurant"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, restaurant.Add)
}
