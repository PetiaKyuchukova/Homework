package main

import (
	"errors"
	"fmt"
)

type Action func() error

var actionWithError Action = func() error { return errors.New("Error happened") }
var actionNoError Action = func() error { return nil }
var actionPanic Action = func() error {
	panic("Panic happened")
	return nil
}

func SafeExec(a Action) Action {
	var errWrap error

	defer func() {
		fmt.Printf("3\n")
		if r := recover(); r != nil {
			errWrap = fmt.Errorf("Recovered: %v", r)
		}
	}()

	defer func() {
		err := a
		fmt.Printf("1\n")
		errRes := err()
		fmt.Printf("2\n")
		if errRes != nil {
			errWrap = fmt.Errorf("safe exec: %w", errRes)
		}
	}()

	var errWrapAction Action = func() error { return errWrap }
	fmt.Printf("4\n")
	return errWrapAction
}

func main() {
	errFunc := SafeExec(actionWithError)
	err := errFunc()
	if err == nil {
		fmt.Printf("No error\n")
	} else {
		fmt.Printf("%s \n", err.Error())
	}

	errFunc = SafeExec(actionNoError)
	err = errFunc()
	if err == nil {
		fmt.Printf("No error\n")
	} else {
		fmt.Printf("%s \n", err.Error())
	}

	errFunc = SafeExec(actionPanic)
	err = errFunc()
	fmt.Print(err.Error())

}
