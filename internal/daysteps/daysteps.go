// Package daysteps.
package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/GrigoriyAkhmerov/go1fl-sprint5-final/internal/personaldata"
	"github.com/GrigoriyAkhmerov/go1fl-sprint5-final/internal/spentenergy"
	"github.com/GrigoriyAkhmerov/go1fl-sprint5-final/internal/trainings"
)

const (
	StepLength = 0.65
)

// DaySteps structure.
type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// Parse() method parse string in format "678,0h50m" and set Training structure fields.
func (ds *DaySteps) Parse(datastring string) (err error) {

	stepsDuration := strings.Split(datastring, ",") //  Point 1. get slice of data after splitting string.

	if len(stepsDuration) != 2 { //  Point 2. need only two variables: number of steps and duration.
		return errors.New("slice isn't equal 2")
	}

	ds.Steps, err = strconv.Atoi(stepsDuration[0]) //  Point 3. convert the first element of the slice to int type and set Steps field in DaySteps structure.
	if err != nil {
		return trainings.ErrConverting
	}

	ds.Duration, err = time.ParseDuration(stepsDuration[1]) // Point 4. convert the second element of the slice to time.Duration type and set Duration field in DaySteps structure.
	if err != nil {
		return trainings.ErrConverting
	}
	if ds.Duration <= 0 {
		return spentenergy.ErrNotPositiveNumber
	}
	return nil
}

// ActionInfo() method displays workout info.
func (ds DaySteps) ActionInfo() (string, error) {

	if ds.Duration <= 0 {
		return "", spentenergy.ErrNotPositiveNumber
	}

	distanceKm := spentenergy.Distance(ds.Steps) // Point 2. calculate distance using Distance function.
	if ds.Duration <= 0 {
		return "", spentenergy.ErrNotPositiveNumber
	}
	if (ds.Weight <= 0) && (ds.Height <= 0) { // Point 3.
		return "", spentenergy.ErrNotPositiveNumber
	}
	caloriesBurnedWalk, nil := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration) // Point 3.

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distanceKm, caloriesBurnedWalk), nil // Point 4.

}
