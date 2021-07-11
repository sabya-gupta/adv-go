package slices

import (
	"fmt"
	"structs"
)

//SLICE PASSES BY REFERENCE
func WorkWithSlices() {
	var team []structs.TeamMember

	tm1 := structs.TeamMember{"Jack", 23, "200501", 5.6}

	tm2 := structs.TeamMember{
		Name:   "Jones",
		Age:    24,
		EmpId:  "200502",
		Rating: 6.5,
	}

	var tm3 structs.TeamMember
	tm3.Name = "Mary"
	tm3.Age = 25
	tm3.EmpId = "20050"
	tm3.Rating = 6.7

	team = append(team, tm1, tm2, tm3)

	team2 := team //MAKES A COPY OF THE REFERENCE AND ALL VALUES BY REFERENCE

	//team2 = append(team2, tm3)

	//MODIFY EXISTING VALUE AND SEE
	team2[0].Age = 100

	fmt.Println("========SLICES============")
	fmt.Println("team = ", team)
	fmt.Println("team2 = ", team2)

}
