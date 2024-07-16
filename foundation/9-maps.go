package foundation

import "fmt"

func Maps() {
	fmt.Println("\n--> Maps")
	StudentAge := make(map[string] int)

	StudentAge["Wagner"] = 31
	StudentAge["Anaysa"] = 28

	fmt.Println(StudentAge)
	fmt.Println(StudentAge["Wagner"])

	superhero := map[string]map[string]string{
		"Superman" : {
			"RealName" : "Clark Kent",
			"City" : "Metropolis",
		},

		"Batman" : {
			"RealName" : "Bruce Wayne",
			"City" : "Gotham City",
		},
	}

	if temp, hero := superhero["Superman"]; hero {
		fmt.Println(temp["RealName"], temp["City"])
	}
}