// A package must be in its own directory, and the directory name should be the same as the package name
package file_operations

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// All functions that should be exported should start with a capital letter
func SaveFloatToFile(val float64, fileName string) {
	data := fmt.Sprint(val)
	os.WriteFile(fileName, []byte(data), 0644)
}

func ReadFloatFromFile(fileName string, defVal float64) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return defVal, errors.New("unable to find data file")
	}

	valText := string(data)
	val, err := strconv.ParseFloat(valText, 64)

	if err != nil {
		return defVal, errors.New("unable to read from file, mismatched versions")
	}

	return val, nil
}
