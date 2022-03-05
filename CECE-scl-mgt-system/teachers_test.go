package my

import (
	"testing"
)

func TestTeachers_GradeCourse(t *testing.T) {
	testGrade := []struct {
		input  Teachers
		input1 string
		input2 string
		input3 string
		input4 int

		expectedOutput string
	}{
		{Teachers{"Lovey", 23, "female", "French"}, "Lovey", "Isioma", "French", 70, "you have successfully been graded"},
	}

	for _, v := range testGrade {
		got := v.input.GradeCourses(v.input.Name, v.input2, v.input3, v.input4)
		if got != v.expectedOutput {
			t.Errorf("expected %v, got %v", v.expectedOutput, got)
		}
	}
}
