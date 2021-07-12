package sets

import "fmt"

type MySet map[string]struct{}

func UseMySet() {

	s := make(MySet)

	s["item1"] = struct{}{}
	s["item2"] = struct{}{}

	fmt.Println(getValues(s))
}

func getValues(s MySet) []string {
	var retVal []string

	for k, _ := range s {
		retVal = append(retVal, k)
	}
	return retVal
}
