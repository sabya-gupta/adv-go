package maps

import (
	"fmt"
	"structs"
)

func MakeSomeMaps() {
	tm1 := structs.TeamMember{"Jack", 23, "200501", 5.6}

	// tm2 := structs.TeamMember{
	// 	Name:   "Jones",
	// 	Age:    24,
	// 	EmpId:  "200502",
	// 	Rating: 6.5,
	// }

	var tm3 structs.TeamMember
	tm3.Name = "Mary"
	tm3.Age = 25
	tm3.EmpId = "20050"
	tm3.Rating = 6.7

	//NOW THE maps
	var map1 map[string]structs.TeamMember
	map1 = make(map[string]structs.TeamMember)
	map1["200501"] = tm1
	map1["20050"] = tm3
	fmt.Println("===========MAPS==========")
	fmt.Println(map1)

	map2 := map1

	v, _ := map2["200501"]
	v.Age = 200
	fmt.Println(map1)

}
