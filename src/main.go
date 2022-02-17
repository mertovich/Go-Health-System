package main

import (
	"fmt"
	"main/doctor"
	"main/patient"
)

func main() {
	// dcotor data
	d1 := doctor.Doctor{`Jhon`, 1}
	d2 := doctor.Doctor{`Anna`, 2}
	docList := []doctor.Doctor{}
	docList = append(docList, d1)
	docList = append(docList, d2)

	// patient variable
	var HealthNumber int
	var name string
	var surname string
	var expertise int

	// panel
	fmt.Print(`HealthNumber: `)
	fmt.Scan(&HealthNumber)
	fmt.Print(`Name: `)
	fmt.Scan(&name)
	fmt.Print(`Surname: `)
	fmt.Scan(&surname)
	fmt.Println(`[1] General Surgery`)
	fmt.Println(`[2] Ophthalmologist`)
	fmt.Print(`Expertise: `)
	fmt.Scanf(`%d`, &expertise)

	// create patient
	p1 := patient.Patient{}
	p1.HealthNumber = HealthNumber
	p1.Name = name
	p1.Surname = surname
	p1.Expertise = expertise

	// to compare
	for _, d := range docList {
		if d.Expertise == p1.Expertise {
			fmt.Println(`Your doctor's name: ` + d.Name)
		}
	}
}
