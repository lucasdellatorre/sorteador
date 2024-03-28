package util

import (
	"database/sql"
	"testing"
)

func TestConvertNullArrayToIntOK(t *testing.T) {
	want := []int{1, 2, 3, 4, 5}

	sut := []sql.NullInt64{
		{Int64: 1, Valid: true},
		{Int64: 2, Valid: true},
		{Int64: 3, Valid: true},
		{Int64: 4, Valid: true},
		{Int64: 5, Valid: true},
	}

	got, _ := ConvertNullArrayToInt(sut)

	for i := 0; i < len(got); i++ {
		if want[i] != got[i] {
			t.Errorf("Error while parsing number %d", got[i])
		}
	}
}

func TestConvertNullArrayToIntNotOK(t *testing.T) {
	sut := []sql.NullInt64{
		{Int64: 1, Valid: true},
		{Int64: 2, Valid: false},
		{Int64: 3, Valid: true},
		{Int64: 4, Valid: true},
		{Int64: 5, Valid: true},
	}

	_, err := ConvertNullArrayToInt(sut)

	if err.Error() != "is not possible to parse the array" {
		t.Errorf("Wrong error message in %s", err)
	}
}
