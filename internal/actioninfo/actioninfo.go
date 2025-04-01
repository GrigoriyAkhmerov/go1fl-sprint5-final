// Package actioninfo displays main information about all type of workouts.
package actioninfo

import (
	"github.com/GrigoriyAkhmerov/go1fl-sprint5-final/internal/trainings"
	"github.com/GrigoriyAkhmerov/go1fl-sprint5-final/internal/daysteps"
	"github.com/GrigoriyAkhmerov/go1fl-sprint5-final/internal/personaldata"
)

// DataParser interface.
type DataParser interface {
    trainings.Parse()
	trainings.ActionInfo()
    daysteps.Parse()
    daysteps.ActionInfo()
} 

// Info function.
func Info(dataset []string, dp DataParser) {
	for _, v := range dataset {
		DaylyAction := daysteps.Parse(v)
	}
	if err != nil {
		return err
		continue
	}
	fmt.Println(daysteps.ActionInfo())

	for _, v := range dataset {
		DaylyAction := trainings.Parse(v)
	}
	if err != nil {
		return err
		continue
	}
	fmt.Println(trainings.ActionInfo())

	
} 
