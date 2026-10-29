package main

import "fmt"

type ValidationError struct {
	Mesage string
}

func (e *ValidationError) Error() string {
	return e.Mesage
}

type NotFoundError struct {
	Mesage string
}

func (e *NotFoundError) Error() string {
	return e.Mesage
}

func SaveData(id string, data any) error {
	if id == "" {
		return &ValidationError{"validation error"}
	}

	if id != "indra" {
		return &NotFoundError{"data not found"}
	}

	// return ok
	return nil
}

func main() {
	error := SaveData("", nil)
	if error != nil {
		// terjadi error
		//if ValidationError, ok := error.(*ValidationError); ok {
		//	fmt.Println("validation error", ValidationError.Mesage)
		//} else if NotFoundError, ok := error.(*NotFoundError); ok {
		//	fmt.Println("data not found", NotFoundError.Mesage)
		//} else {
		//	fmt.Println("unknown error:", error.Error())
		//}

		switch finalError := error.(type) {
		case *ValidationError:
			fmt.Println("validation error", finalError.Mesage)
		case *NotFoundError:
			fmt.Println("data not found", finalError.Mesage)
		default:
			fmt.Println("unknown error", finalError.Error())
		}
	} else {
		// sukses
		fmt.Println("Success")
	}
}
