package utils

import (
	"errors"
	"reflect"
	"time"
)

func Datetime(v interface{}, param string) error {
	st := reflect.ValueOf(v)
	_, err := time.Parse(DOB_DATE_FORMAT, st.String())
	if err != nil {
		return errors.New("date of birth format does not match yyyy-MM-dd")
	}
	return nil
}
