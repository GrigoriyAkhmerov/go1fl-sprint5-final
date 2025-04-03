package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/GrigoriyAkhmerov/go1fl-sprint5-final/internal/personaldata"
	"github.com/GrigoriyAkhmerov/go1fl-sprint5-final/internal/spentenergy"
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

	StepsDuration := strings.Split(datastring, ",") //  Point 1. get slice of data after splitting string.

	if len(StepsDuration) != 2 { //  Point 2. need only two variables: number of steps and duration.
		return err
	}

	ds.Steps, err = strconv.Atoi(StepsDuration[0]) //  Point 3. convert the first element of the slice to int type and set Steps field in DaySteps structure.
	if err != nil {
		return err
	}

	ds.Duration, err = time.ParseDuration(StepsDuration[1]) // Point 4. convert the second element of the slice to time.Duration type and set Duration field in DaySteps structure.
	if err != nil {
		return err
	}
	return
}

// ActionInfo() method displays workout info.
func (ds DaySteps) ActionInfo() (string, error) {

	if ds.Duration <= 0 {
		return "", spentenergy.ErrNotPositiveNumber
	}

	DistanceKm := spentenergy.Distance(ds.Steps) // Point 2. calculate distance using Distance function.

	if (ds.Weight <= 0) && (ds.Height <= 0) { // Point 3.
		return "", spentenergy.ErrNotPositiveNumber
	}
	CaloriesBurnedWalk, nil := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration) // Point 3.

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, DistanceKm, CaloriesBurnedWalk), nil // Point 4.

}
