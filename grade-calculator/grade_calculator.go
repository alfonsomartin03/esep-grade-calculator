package esepunittests

type GradeCalculator struct {
	grades []Grade
}

type GradeType int

const (
	Assignment GradeType = iota
	Exam
	Essay
)

var gradeTypeName = map[GradeType]string{
	Assignment: "assignment",
	Exam:       "exam",
	Essay:      "essay",
}

func (gt GradeType) String() string {
	return gradeTypeName[gt]
}

type Grade struct {
	Name  string
	Grade int
	Type  GradeType
}

func NewGradeCalculator() *GradeCalculator {
	return &GradeCalculator{
		grades: make([]Grade, 0),
	}
}

func (gc *GradeCalculator) GetFinalGrade() string {
	numericalGrade := gc.calculateNumericalGrade()

	if numericalGrade >= 90 {
		return "A"
	} else if numericalGrade >= 80 {
		return "B"
	} else if numericalGrade >= 70 {
		return "C"
	} else if numericalGrade >= 60 {
		return "D"
	}

	return "F"
}

func (gc *GradeCalculator) AddGrade(name string, grade int, gradeType GradeType) {
	switch gradeType {
	case Assignment, Exam, Essay:
		gc.grades = append(gc.grades, Grade{
			Name:  name,
			Grade: grade,
			Type:  gradeType,
		})
	}
}

func (gc *GradeCalculator) calculateNumericalGrade() float64 {
	var assignmentTotal, examTotal, essayTotal int
	var assignmentCount, examCount, essayCount int

	for _, grade := range gc.grades {
		switch grade.Type {
		case Assignment:
			assignmentTotal += grade.Grade
			assignmentCount++
		case Exam:
			examTotal += grade.Grade
			examCount++
		case Essay:
			essayTotal += grade.Grade
			essayCount++
		}
	}

	assignmentAverage := computeAverage(assignmentTotal, assignmentCount)
	examAverage := computeAverage(examTotal, examCount)
	essayAverage := computeAverage(essayTotal, essayCount)

	return assignmentAverage*0.50 + examAverage*0.35 + essayAverage*0.15
}

func computeAverage(total, count int) float64 {
	if count == 0 {
		return 0
	}

	return float64(total) / float64(count)
}
