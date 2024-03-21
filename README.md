# Diseño de Patrones con Golang.
Este repositorio pretende contener distintos ejemplos en los que se pueden implementar los diseños de patrones con el lenguaje Golang.

Para empezar, ¿Qué son los Patrones de Diseño?

Son soluciones aplicables, reutilizables y escalables en la forma en la que se construye una aplicación. Estos patrones se
apoyan en lo que se conoce como los ***Principios Solid***.

- Single Responsibility Principle.
- Open Closed Principle.
- Liskov Principle.
- Interface Segregation Principle.
- Dependency Inversion Principle.


**Single Responsibility Principle**
O también "Principio de Responsabilidad Única" en español.

Va meramente con el grado de cohesión en el que un módulo y su contenido está relacionado entre sí.

Si un módulo tiene una única responsabilidad, decimos que su cohesión es alta.
De lo contrario, tiene cohesión baja.

La alta cohesión de los módulos en una app permite que el código sea mucho más mantenible así como la reutilización
del mismo.

**Open Closed Principle**
Las entidades del software deben ser capaces de crecer (open), pero sin que implique una modificación en el código ya existente (closed), debido a que introducir nuevas integraciones en el código ya existente, puede implicar
nuevos errores no mapeados.

**Liskov Substitution Principle**
Este principio busca resolver un problema frecuente en el polimorfismo de objetos (cuando una clase S es heredada de una clase T), por lo que
impone que todos los atributos de la clase T deben existir e implementarse en la clase S para que sea segura su implementación.

**Interface Segregation Principle**
Dicho principio impone que los clientes de una aplicación no deberían ser forzados a utilizar métodos que no necesitan.

**Dependency Inversion Principle**
Hace referencia a la implementación de dependencias inversas, es decir, de lo más específico, a lo más general.

Los módulos de alto nivel no deben depender de los de bajo nivel. Ambos deben depender de abstracciones.
Las abstracciones no deben depender de los detalles, sino que los detalles deben depender de las abstracciones.




