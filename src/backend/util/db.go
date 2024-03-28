package util

import (
	"database/sql"
	"errors"
)

func ConvertNullArrayToInt(arr []sql.NullInt64) ([]int, error) {
	var newArr []int
	for _, num := range arr {
		if !num.Valid {
			return nil, errors.New("is not possible to parse the array")
		}
		newArr = append(newArr, int(num.Int64))
	}

	return newArr, nil
}
