package my

type Students struct {
	Courses
	Classes
	Name   string
	Age    int
	Gender string
	Fees   int
}

type Classes struct {
	Classrooms map[string]int
}

type Courses struct {
	CoursesAndGrades map[string]int
}

var StudentDB []Students

func init() {
	st := []Students{
		{Name: "Chisom", Age: 14, Gender: "female", Fees: 10750, Classes: Classes{Classrooms: map[string]int{"JSS1": +1}},
			Courses: Courses{CoursesAndGrades: map[string]int{"Physics": 0, "Chemistry": 0, "Economics": 0, "English": 0,
				"Agric": 0, "French": 0}}},
		{Name: "Gbenle", Age: 13, Gender: "male", Fees: 5000, Classes: Classes{Classrooms: map[string]int{"JSS2": +1}},
			Courses: Courses{CoursesAndGrades: map[string]int{"Physics": 0, "Chemistry": 0, "Economics": 0, "Biology": 0, "English": 0,
				"French": 0}}},
		{Name: "Isioma", Age: 15, Gender: "female", Fees: 10750, Classes: Classes{Classrooms: map[string]int{"JSS3": +1}},
			Courses: Courses{CoursesAndGrades: map[string]int{"Physics": 0, "Chemistry": 0, "Economics": 0, "Maths": 0, "English": 0,
				"French": 0}}},
	}
	for _, v := range st {
		StudentDB = append(StudentDB, v)
	}

}

//a method for formating the printout of my STUDENTDB
/*
func (s Students) Print() {
	fmt.Printf("S/N\t%s\tName\t%s\tAge\t%v\tGender\t%s\tFees\t%s\tClasses\t\t%s\tCourses\t\t\n", "|", "|", "|", "|", "|", "|")
	for i, v := range StudentDB {

		fmt.Println(strings.Repeat("-", 114))
		fmt.Printf("%v\t%s\t%s\t%s\t%v\t%s\t%s\t%s\t%v\t%s\t%v\t%s\t%v\t\t\n", i+1, "|", v.Name, "|", v.Age, "|", v.Gender, "|", v.Fees, "|", v.Classes.Classrooms, "|", v.Courses.CoursesAndGrades)
	}
}
*/

var listOfCourses = []string{"Physics", "Chemistry", "Biology", "Economics", "English", "Maths", "Agric", "Yoruba", "Literature", "Commerce", "French"}

func (s Students) TakeACourse(name string, course string) string {
	for i, v := range StudentDB {
		if v.Name == name {
			for _, x := range listOfCourses {
				if x == course {
					v.CoursesAndGrades = map[string]int{course: 0}
					StudentDB[i] = v
					return "you have successfully registered for a course"
				}
			}
		}
	}
	return "sorry, you are not a student of this GREAT institution"
}
