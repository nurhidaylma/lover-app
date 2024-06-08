package util

import "errors"

func NewError(fileName, funcName string, err error) error {
	if err != nil {
		return errors.New(fileName + ":" + funcName + " Err:" + err.Error())
	}

	return nil
}
