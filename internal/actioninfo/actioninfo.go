// Package actioninfo.
package actioninfo

import (
	"fmt"
)

// DataParser interface.
type DataParser interface {
	Parse(s string) error
	ActionInfo() (string, error)
}

// Info function.
func Info(dataset []string, dp DataParser) {
	for _, v := range dataset { // Point 1. Loop through all values ​​of the dataset slice.
		err := dp.Parse(v) // Point 2. Parsing all values with Parse () method.
		if err != nil {
			fmt.Println(err) // Point 3. Displaying error.
			continue
		}

		infoString, err := dp.ActionInfo()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(infoString) // Point 4. Displaying information about activity.
	}
}
