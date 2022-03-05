package my

import (
	"fmt"
)

var ApplicantDB []Applicants

type Principal struct {
	Name   string
	Age    int
	Gender string
}

type Applicants struct {
	Name   string
	Age    int
	Gender string
	Class  string
}

//func init() {
//	app := []Applicants{
//		{Name: "Frank", Age: 26, Gender: "female", Class: "JSS3"},
//		{Name: "David", Age: 10, Gender: "male", Class: "Jss1"},
//		{Name: "Loveth", Age: 20, Gender: "female", Class: "ss1"},
//		{Name: "Ogo", Age: 15, Gender: "female", Class: "Jss2"},
//		{Name: "Ebube", Age: 26, Gender: "female", Class: "SS2"},
//	}
//	for _, v := range app {
//		ApplicantDB = append(ApplicantDB, v)
//	}
//	log.Println(ApplicantDB)
//}

func (p *Principal) AdmitApplicant(a Applicants) string {

	if a.Age < 16 {
		return "not old enough to be admitted"
	} else {
		student := Students{
			Name:    a.Name,
			Age:     a.Age,
			Gender:  a.Gender,
			Fees:    2000,
			Classes: Classes{Classrooms: map[string]int{a.Class: +1}},
		}
		StudentDB = append(StudentDB, student)
		return "welcome to cece's high school"

	}

}

func (p Principal) ExpelStudent(studentName string) string {
	for i, v := range StudentDB {
		if studentName == v.Name {
			StudentDB = append(StudentDB[:i], StudentDB[i+1:]...)
			s := fmt.Sprintf("%s has been expeled from our school", studentName)
			return s
		}
	}
	return ""
}

func (p Principal) SackTeacher(teacherName string) string {
	for i, v := range TeacherDB {
		if teacherName == v.Name {
			TeacherDB = append(TeacherDB[:i], TeacherDB[i+1:]...)
			t := fmt.Sprintf("%s has been sacked from this HONOURABLE institution", teacherName)
			return t
		}
	}
	return ""
}
