package abstract_factory

type AbstractFactoryRelacional struct{}

// Concrete factory: CreateStudentsRepository
func (abs *AbstractFactoryRelacional) CreateStudentsRepository() (StudentsRepository, error) {

	return &RelacionalStudents{}, nil
}

// Concrete factory: CreateCoursesRepository
func (abs *AbstractFactoryRelacional) CreateCoursesRepository() (CoursesRepository, error) {

	return &RelacionalCourses{}, nil
}

// Concrete Product: RelacionalStudents
type RelacionalStudents struct{}

func (nor *RelacionalStudents) GetStudents() string {
	// do something
	return "Get Students Relacional"
}

// Concrete Product: RelacionalCourses
type RelacionalCourses struct{}

func (nor *RelacionalCourses) GetCourses() string {
	// do something
	return "Get Courses Relacional"
}
