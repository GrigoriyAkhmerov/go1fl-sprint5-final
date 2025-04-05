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

var (
	ErrUnknownTrainingType = errors.New("unknown training type") // "for points 1 and 2 in WalkingSpentCalories function and Parse() method"
	ErrConverting          = errors.New("converting error")
)

// Parse() method parse string in format "3456,Ходьба,3h00m" and set Training structure fields.
func (t *Training) Parse(datastring string) (err error) {
	stepsActivityDuration := strings.Split(datastring, ",") //  Point 1. get slice of data after splitting string.
	if len(stepsActivityDuration) != 3 {                    //  Point 2. need only three variables: number of steps, activity and duration.
		return errors.New("slice isn't equal 3")
	}
	t.Steps, err = strconv.Atoi(stepsActivityDuration[0]) //  Point 3. convert the first element of the slice to int type and set Steps field in Training structure.
	if err != nil {
		return ErrConverting
	}
	if t.Steps <= 0 {
		return spentenergy.ErrNotPositiveNumber
	}
	if (stepsActivityDuration[1] != "Бег") && (stepsActivityDuration[1] != "Ходьба") { // Point 4. Must be "Бег" or "Ходьба" only.
		return ErrUnknownTrainingType
	}
	t.TrainingType = stepsActivityDuration[1]                      //  Point 5. set TrainingType field in Training structure.
	t.Duration, err = time.ParseDuration(stepsActivityDuration[2]) // Point 6. convert the third element of the slice to time.Duration type set Duration field in Training structure.
	// fmt.Println(t.Duration)
	if err != nil {
		return ErrConverting
	}
	if t.Duration <= 0 {
		return spentenergy.ErrNotPositiveNumber
	}
	return nil
}

// InfoMessage function for AtionInfo() method.
func InfoMessage(training string, duration, distance, speed float64, calories float64) string {
	//	fmt.Println(duration)
	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f", training, duration, distance, speed, calories) // Point 5.
}

// ActionInfo() method set string with workout data.
func (t Training) ActionInfo() (string, error) {

	finalDist := spentenergy.Distance(t.Steps) //  Point 1. Get distance using Distance() function from spentenergy package.

	if finalDist <= 0 {
		return "", spentenergy.ErrNotPositiveNumber
	}

	if t.Duration <= 0 { //  Point 2. must be greater then zero.
		return "", spentenergy.ErrNotPositiveNumber
	}

	finalSpeed := spentenergy.MeanSpeed(t.Steps, t.Duration) // Point 3. Get average speed using MeanSpeed() function from spentenergy package.

	switch t.TrainingType { // Point 4. calculate calories burned for every type of workout.

	case "Ходьба":
		finalCaloriesWalk, nil := spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		return InfoMessage(t.TrainingType, t.Duration.Hours(), finalDist, finalSpeed, finalCaloriesWalk), nil // Point 5.
	case "Бег":
		finalCaloriesRun, nil := spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Duration)
		return InfoMessage(t.TrainingType, t.Duration.Hours(), finalDist, finalSpeed, finalCaloriesRun), nil
	default:
		return fmt.Sprintln("неизвестный тип тренировки"), ErrUnknownTrainingType // Point 6.
	}
}
