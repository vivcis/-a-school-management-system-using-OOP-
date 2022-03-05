package my

type NonAcademicStaffs struct {
	Name   string
	Age    int
	Gender string
}

func (n NonAcademicStaffs) LateComersToSchool(s Students) {
	//this is for the security
	//scl assembly starts by 8am
	//any students later than that is assumed a latecomer and punished by cutting grass or washing toilet

	//earlyStudents := "8:00:00am"

}
