package my

import (
	"testing"
)

func TestPrincipal_AdmitApplicant(t *testing.T) {
	//s:= my.Students{
	//	Courses: my.Courses{map[string]int{"Chemistry": 54}},
	//	Classes: my.Classes{map[string]int{"JSS 1B": 32}},
	//	Name:    "Joe B",
	//	Age:     14,
	//	Gender:  "Male",
	//	Fees:    0,
	//}
	//
	//c:= my.Courses{CoursesAndGrades: map[string]int{"Chemistry": 54}}

	var admitApplicant = []struct {
		input  Principal
		input1 Applicants

		expectedOutput string
	}{
		{Principal{Name: "Okon", Age: 43, Gender: "Male"}, Applicants{"Alexis", 18, "Female", "JSS 2A"}, "welcome to cece's high school"},
		{Principal{Name: "Judith", Age: 18, Gender: "Female"}, Applicants{"Joe B", 14, "Male", "JSS 1C"}, "not old enough to be admitted"},
	}
	for _, val := range admitApplicant {
		admitted := val.input.AdmitApplicant(val.input1)
		if admitted != val.expectedOutput {
			t.Errorf("expected output to be %v but got output %v", val.expectedOutput, admitted)
		}
	}
}

func TestPrincipal_ExpelStudent(t *testing.T) {
	var expelStudent = []struct {
		input  Principal
		input1 string

		expectedOutput string
	}{
		{Principal{"Taofeek", 45, "male"}, "Chisom", "Chisom has been expeled from our school"},
	}

	for _, val := range expelStudent {
		expelled := val.input.ExpelStudent(val.input1)
		if expelled != val.expectedOutput {
			t.Errorf("expected output to be %v but got output %v", val.expectedOutput, expelled)
		}
	}
}

func TestPrincipal_SackTeacher(t *testing.T) {
	var sackTeacher = []struct {
		input  Principal
		input1 string

		expectedOutput string
	}{
		{Principal{"Phillip", 56, "Male"}, "Clinton", "Clinton has been sacked from this HONOURABLE institution"},
	}

	for _, v := range sackTeacher {
		sacked := v.input.SackTeacher(v.input1)
		if sacked != v.expectedOutput {
			t.Errorf("expected output to be %v but got output %v", v.expectedOutput, sacked)
		}
	}
}
