package structs

import "fmt"

type TeamMember struct {
	Name   string
	Age    int
	EmpId  string
	Rating float32
}

func ProduceSomeTeamMembers() {
	tm1 := TeamMember{"Jack", 23, "200501", 5.6}

	tm2 := TeamMember{
		Name:   "Jones",
		Age:    24,
		EmpId:  "200502",
		Rating: 6.5,
	}

	var tm3 TeamMember
	tm3.Name = "Mary"
	tm3.Age = 25
	tm3.EmpId = "20050"
	tm3.Rating = 6.7

	fmt.Println("tm1 = ", tm1)
	fmt.Println("tm2= ", tm2)
	fmt.Println("tm3 = ", tm3)

	tm3Copy := tm3
	tm3Copy.Name = "tm3Copy"

	fmt.Println("tm1 = ", tm1)
	fmt.Println("tm2= ", tm2)
	fmt.Println("tm3 = ", tm3)
	fmt.Println("tm3Copy = ", tm3Copy)

	tm2ref := &tm2
	tm2ref.Name = "ChangedInRef"
	fmt.Println("tm1 = ", tm1)
	fmt.Println("tm2= ", tm2)
	fmt.Println("tm3 = ", tm3)
	fmt.Println("tm3Copy = ", tm3Copy)
}
