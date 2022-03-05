package my

type Teachers struct {
	Name    string
	Age     int
	Gender  string
	Subject string
}

var TeacherDB []Teachers

func init() {
	t := []Teachers{
		{Name: "Franklyn", Gender: "male", Age: 76, Subject: "Chemistry"},
		{Name: "Cecilia", Gender: "female", Age: 88, Subject: "English"},
		{Name: "Clinton", Gender: "male", Age: 32, Subject: "Physics"},
		{Name: "Lovey", Gender: "female", Age: 23, Subject: "French"},
	}
	for _, v := range t {
		TeacherDB = append(TeacherDB, v)
	}

}

//a method to format the printing of my TEACHERDB
/* func (t Teachers) Printer() {
	fmt.Printf("S/N\t%s\tName\t%v\tGender\t%s\tAge\t\t%s\tSubject\t\t\n", "|", "|", "|", "|")
	for i, v := range TeacherDB {

		fmt.Println(strings.Repeat("-", 90))
		fmt.Printf("%v\t%s\t%s\t%s\t%v\t%s\t%v\t%s\t%v\t\t\n", i+1, "|", v.Name, "|", v.Gender, "|", v.Age, "|", v.Subject)
	}
}
*/

func (t Teachers) GradeCourses(teachersName string, StudentsName string, course string, grade int) string {

	for _, x := range TeacherDB {
		if x.Name == teachersName {
			if x.Subject == course {

				for i, v := range StudentDB {
					if v.Name == StudentsName {
						if _, found := v.CoursesAndGrades[course]; found {
							v.CoursesAndGrades[course] = grade
							StudentDB[i] = v
							return "you have successfully been graded"
						}
					}
				}

			}
		}
	}
	return "not to be graded by you"
}

//func (t Teachers) GetAverageGrades(name string, course string, grade int) float32 {
//	sum := 0
//	for _, v := range StudentDB {
//		if v.Name == name {
//			if _, found := v.CoursesAndGrades[course]; found {
//				v.CoursesAndGrades[course] = grade
//				sum += v
//			}
//		}
//	}
//	return float32(sum) / float32(len(v.CoursesAndGrades[course]))
//}
//
//func (s Students) GetMaxGrades() int {
//	currentMax := 0
//
//	for _, v := range StudentDB {
//		if currentMax < v {
//			currentMax = v
//		}
//	}
//	return currentMax
//}
