package abstract_factory

import "fmt"

func MainAbstract() {

	//RELACIONAL
	abstractFactoryRelacional := AbstractFactoryNoRelacional{}

	//....students
	relacionalStudentsRepository, err := abstractFactoryRelacional.CreateStudentsRepository()

	if err != nil {
		panic(err)
	}

	fmt.Println(relacionalStudentsRepository.GetStudents())

	//....courses
	relacionalCoursesRepository, err := abstractFactoryRelacional.CreateCoursesRepository()

	if err != nil {
		panic(err)
	}

	fmt.Println(relacionalCoursesRepository.GetCourses())

	//NO RELACIONAL
	abstractFactoryNoRelacional := AbstractFactoryNoRelacional{}

	//....students
	noRelacionalStudentsRepository, err := abstractFactoryNoRelacional.CreateStudentsRepository()

	if err != nil {
		panic(err)
	}

	fmt.Println(noRelacionalStudentsRepository.GetStudents())

	//....courses
	noRelacionalCoursesRepository, err := abstractFactoryNoRelacional.CreateCoursesRepository()

	if err != nil {
		panic(err)
	}

	fmt.Println(noRelacionalCoursesRepository.GetCourses())
}
