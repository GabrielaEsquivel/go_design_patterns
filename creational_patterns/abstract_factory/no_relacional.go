package abstract_factory

type AbstractFactoryNoRelacional struct{}

// Concrete factory: CreateStudentsRepository
func (abs *AbstractFactoryNoRelacional) CreateStudentsRepository() (StudentsRepository, error) {
	return &NoRelacionalStudents{}, nil
}

// Concrete factory: CreateCoursesRepository
func (abs *AbstractFactoryNoRelacional) CreateCoursesRepository() (CoursesRepository, error) {
	return &NoRelacionalCourses{}, nil
}

// Concrete Product: NoRelacionalStudents
type NoRelacionalStudents struct{}

func (nor *NoRelacionalStudents) GetStudents() string {
	// do something
	return "Get Students No Relacional"
}

// Concrete Product: NoRelacionalCourses
type NoRelacionalCourses struct{}

func (nor *NoRelacionalCourses) GetCourses() string {
	// do something
	return "Get Courses No Relacional"
}
