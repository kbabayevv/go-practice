package pointer

import "fmt"

func Pointer() {
	age := 32
	agePointer := &age
	fmt.Println("Age is", *agePointer)

	getAdultYears(agePointer)
	fmt.Println("Adult years", age)
}

func getAdultYears(age *int) {
	*age = *age - 18
}
