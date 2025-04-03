// Package trainings splits string into slice of data, calculate and display workout information.
package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/GrigoriyAkhmerov/go1fl-sprint5-final/internal/personaldata"
	"github.com/GrigoriyAkhmerov/go1fl-sprint5-final/internal/spentenergy"
)

// Training structure.
type Training struct {
	Steps        int
	TrainingType string // "running or walking"
	Duration     time.Duration
	personaldata.Personal
}

// Parse() method parse string in format "3456,Ходьба,3h00m" and set Training structure fields.
func (t *Training) Parse(datastring string) (err error) {
	StepsActivityDuration := strings.Split(datastring, ",") //  Point 1. get slice of data after splitting string.

	if len(StepsActivityDuration) != 3 { //  Point 2. need only three variables: number of steps, activity and duration.
		return err
	}

	t.Steps, err = strconv.Atoi(StepsActivityDuration[0]) //  Point 3. convert the first element of the slice to int type and set Steps field in Training structure.
	if err != nil {
		return err
	}
	if (StepsActivityDuration[1] != "Бег") && (StepsActivityDuration[1] != "Ходьба") { // Point 4. Must be "Бег" or "Ходьба" only.
		return err
	}

	t.TrainingType = StepsActivityDuration[1] //  Point 5. set TrainingType field in Training structure.

	t.Duration, err = time.ParseDuration(StepsActivityDuration[2]) // Point 6. convert the third element of the slice to time.Duration type set Duration field in Training structure.
	if err != nil {
		return err
	}
	return
}

var ErrUnknownTrainingType = errors.New("unknown training type") // "for points 1 and 2 in WalkingSpentCalories function"

// ActionInfo() method set string with workout data.
func (t Training) ActionInfo() (string, error) {

	FinalDist := spentenergy.Distance(t.Steps) //  Point 1. Get distance using Distance() function from spentenergy package.

	if t.Duration <= 0 { //  Point 2. must be greater then zero.
		return "", spentenergy.ErrNotPositiveNumber
	}

	FinalSpeed := spentenergy.MeanSpeed(t.Steps, t.Duration) // Point 3. Get average speed using MeanSpeed() function from spentenergy package.

	switch t.TrainingType { // Point 4. calculate calories burned for every type of workout.

	case "Ходьба":
		FinalCaloriesWalk, nil := spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)

		DurationConverted1 := float64(time.Duration(t.Duration * time.Hour))

		return fmt.Sprintf(`Тип тренировки: %s
							Длительность: %.2f ч.
							Дистанция: %.2f км.
							Скорость: %.2f км/ч
							Сожгли калорий: %.2f`, t.TrainingType, DurationConverted1, FinalDist, FinalSpeed, FinalCaloriesWalk), nil // Point 5.

	case "Бег":
		FinalCaloriesRun, nil := spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Duration)

		DurationConverted2 := float64(time.Duration(t.Duration * time.Hour))

		return fmt.Sprintf(`Тип тренировки: %s
							Длительность: %.2f ч.
							Дистанция: %.2f км.
							Скорость: %.2f км/ч
							Сожгли калорий: %.2f`, t.TrainingType, DurationConverted2, FinalDist, FinalSpeed, FinalCaloriesRun), nil // Point 5.

	default:
		return fmt.Sprintln("неизвестный тип тренировки"), ErrUnknownTrainingType // Point 6.
	}
}
