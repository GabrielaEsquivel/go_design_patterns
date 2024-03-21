El **_Abstract Factory Pattern_** se utiliza cuando se desea crear una colección de clases, más no se quiere especificar el tipo concreto de clases a utilizar ya que puede variar de una situación a otra.

Este pattern propone proporcionar una interfaz para crear una familia de objetos relaciones o dependientes *sin* especificar su clase concreta.

**Problema**

Supongamos que tenemos que crear una aplicación que consuma una fuente de datos para obtener la información de los alumnos inscritos en un colegio y que a su vez, podamos consultar los cursos que se ofrecen en dicho colegio.
Para esta app, los clientes quieren relacionar una base de datos que es relacional y otra que no lo es, a primera instancia podríamos pensar que necesitamos dos tipos de implementaciones, en donde en cierto punto se decida a qué
Base de Datos (relacional o no) se hará la consulta, pero esto implicaría duplicar código de forma innecesaria además de que, si en un futuro se tuviera que consultar una fuente de datos distinta, tendríamos que modificar la
implementación previamente creada y adaptarla para esa nueva fuente de datos.

Es aquí donde el Patrón de Diseño Abstract Factory entra. 