// Package spentenergy contain functions for calculating calories burned while walking,
// running, functions for calculating average speed and distance.
package spentenergy

import (
	"errors"
	"time"
)

// Main constants for calculating.
const (
	lenStep   = 0.65  // "average step lenght"
	mInKm     = 1000  // "meters in km"
	minInH    = 60    // "minutes in 1 hour"
	kmhInMsec = 0.278 // "coefficient for converting km/h to meters/sec"
	cmInM     = 100   // "cm in meter"
	speed     = 1.39  // "average speed in meters/sec units"
)

// Distance function takes the number of steps and returns the distance (in km units) that the user covered during workout.
func Distance(steps int) float64 {
	return (float64(steps)) * (float64(lenStep)) / float64(mInKm) // steps int - steps number.
}

// MeanSpeed function takes the number of steps, the duration of activity and returns the average speed during workout.
func MeanSpeed(steps int, duration time.Duration) float64 {
	if duration <= 0 { // Point 1.
		return 0
	}

	DistConv := Distance(steps) // Point 2. calculate distance using Distance() function.

	DurationConv := float64(time.Duration(duration * time.Hour)) // Convert time.Duration to float64.

	Speed := DistConv / DurationConv // Point 3. calculate and return the average speed.

	return Speed
}

// Constants for calculating calories burned while walking.
const (
	walkingCaloriesWeightMultiplier = 0.035 // "body mass multiplier"
	walkingSpeedHeightMultiplier    = 0.029 // "height multiplier"
)

var ErrNotPositiveNumber = errors.New("must be greater then zero") // "for points 1 and 2 in WalkingSpentCalories function"

// WalkingSpentCalories takes the number of steps, weight and height of user,
// duration of workout and returns calories burned while walking.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if (weight <= 0) && (height <= 0) { // Point 1.
		return 0, ErrNotPositiveNumber
	}

	if duration <= 0 { // Point 2.
		return 0, ErrNotPositiveNumber
	}

	meanSpeedWalk := MeanSpeed(steps, duration) // Point 3. Calculate average speed  using MeanSpeed()

	WalkingBurnedCalories := ((walkingCaloriesWeightMultiplier * weight) + (meanSpeedWalk*meanSpeedWalk/height)*walkingSpeedHeightMultiplier) * float64(time.Duration(duration*time.Hour)) * minInH // Point 4.

	return WalkingBurnedCalories, nil
}

// Constants for calculating calories burned while running.
const (
	runningCaloriesMeanSpeedMultiplier = 18.0 // "average speed multiplier"
	runningCaloriesMeanSpeedShift      = 20.0 // "average number of calories burned while running"
)

// RunningSpentCalories takes the number of steps, weight of user,
// duration of workout and returns calories burned while running.
func RunningSpentCalories(steps int, weight float64, duration time.Duration) (float64, error) {
	if weight <= 0 {
		return 0, ErrNotPositiveNumber // Point 1.
	}
	if duration <= 0 {
		return 0, ErrNotPositiveNumber // Point 2.
	}
	meanSpeedRun := MeanSpeed(steps, duration) // Point 3. Calculate average speed  using MeanSpeed()

	RunningBurnedCalories := ((runningCaloriesMeanSpeedMultiplier * meanSpeedRun) - runningCaloriesMeanSpeedShift) * weight // Point 4.

	return RunningBurnedCalories, nil
}
