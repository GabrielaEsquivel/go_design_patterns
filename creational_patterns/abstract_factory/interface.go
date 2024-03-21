package abstract_factory

// Abstract Product: StudentsRepository
type StudentsRepository interface {
	GetStudents() string
}

// Abstract Product: CoursesRepository
type CoursesRepository interface {
	GetCourses() string
}

// Abtract Factory: AbstractFactory
type AbstractFactory interface {
	/** Crearemos una interface para crear una familia de objetos sin necesidad de indicar su tipo concreto, en este ejemplo, si
	 si es que se trata de una BD SQL o NoSQL
	**/
	CreateStudentsRepository() (StudentsRepository, error)
	CreateCoursesRepository() (CoursesRepository, error)
}
