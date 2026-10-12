package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeAWithHighScores(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 95, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 91, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeBands(t *testing.T) {
	testCases := []struct {
		name          string
		grade         int
		expectedGrade string
	}{
		{name: "C", grade: 70, expectedGrade: "C"},
		{name: "D", grade: 60, expectedGrade: "D"},
		{name: "F", grade: 59, expectedGrade: "F"},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			gradeCalculator := NewGradeCalculator()
			gradeCalculator.AddGrade("assignment", testCase.grade, Assignment)
			gradeCalculator.AddGrade("exam", testCase.grade, Exam)
			gradeCalculator.AddGrade("essay", testCase.grade, Essay)

			if actual := gradeCalculator.GetFinalGrade(); actual != testCase.expectedGrade {
				t.Errorf("GetFinalGrade() = %q; want %q", actual, testCase.expectedGrade)
			}
		})
	}
}

func TestEmptyCalculatorReturnsF(t *testing.T) {
	gradeCalculator := NewGradeCalculator()

	if actual := gradeCalculator.GetFinalGrade(); actual != "F" {
		t.Errorf("GetFinalGrade() = %q; want %q", actual, "F")
	}
}

func TestGradeTypeString(t *testing.T) {
	testCases := []struct {
		gradeType GradeType
		want      string
	}{
		{gradeType: Assignment, want: "assignment"},
		{gradeType: Exam, want: "exam"},
		{gradeType: Essay, want: "essay"},
	}

	for _, testCase := range testCases {
		if actual := testCase.gradeType.String(); actual != testCase.want {
			t.Errorf("GradeType(%d).String() = %q; want %q", testCase.gradeType, actual, testCase.want)
		}
	}
}
