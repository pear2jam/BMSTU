% Лабораторная работа № 3 «Обобщённые классы в Scala»
% 30 апреля 2024 г.
% Сергей Наумов, ИУ9-62Б

# Цель работы
Целью данной работы является приобретение навыков разработки\
обобщённых классов на языке Scala с использованием неявных \
преобразований типов.

# Индивидуальный вариант
Класс Vector[T <: Product], представляющий вектор в двух или трёхмерном пространстве.\
Если тип T представляет пару или тройку, элементы которой имеют одинаковый числовой \
тип, то для Vector[T <: Product] должны быть доступны операции сложения и скалярного\
умножения. Дополнительно, для троек должна быть реализована операция векторного \
умножения.

# Реализация

Вектор может быть представлен различными типами данных, такими как \
целые числа, дробные числа или кортежи координат.

**Класс Vector[T <: Product]**:
- Класс Vector параметризован типом T, который ограничен сверху типом 
- Product. Тип Product является супертипом для всех классов, имеющих метод
productArity, который возвращает размерность объекта.
- У класса есть поле components, которое представляет собой компоненты вектора.
- Метод dimension возвращает размерность вектора.
- Метод + выполняет операцию сложения двух векторов.
- Метод dotProduct вычисляет скалярное произведение двух векторов.
- Метод crossProduct вычисляет векторное произведение двух трехмерных векторов.
- Метод toString возвращает строковое представление вектора.


**Трейт VectorOps[T <: Product]**:
- Трейт определяет операции над векторами, такие как сложение и скалярное произведение.


**Объект VectorOps**:
- Объект содержит неявные реализации VectorOps для различных типов векторов, \
таких как Int, Double и кортежи координат.
- Каждая реализация предоставляет методы add и dotProduct для выполнения \
соответствующих операций над векторами указанных типов.\


Для каждого типа вектора (Int, Double, кортежи координат) определены соответствующие \
операции сложения и скалярного произведения. Когда эти операции вызываются для\
объекта класса Vector, компилятор Scala автоматически выбирает соответствующую \
реализацию из объекта VectorOps, что позволяет работать с векторами различных \
типов данных без явного указания типа.

Код реализации:
```scala
class Vector[T <: Product](val components: T) {
  // Метод для получения размерности вектора
  def dimension: Int = components.productArity

  // Метод для сложения двух векторов
  def +(other: Vector[T])(implicit ops: VectorOps[T]): Vector[T] =
    new Vector(ops.add(components, other.components))

  // Метод для скалярного умножения двух векторов
  def dotProduct(other: Vector[T])(implicit ops: VectorOps[T]): Double =
    ops.dotProduct(components, other.components)

  // Метод для векторного умножения двух трехмерных векторов
  def crossProduct(other: Vector[(Double, Double, Double)]): (Double, Double, Double) = {
    val (x1, y1, z1) = components.asInstanceOf[(Double, Double, Double)]
    val (x2, y2, z2) = other.components
    val crossX = y1 * z2 - z1 * y2
    val crossY = z1 * x2 - x1 * z2
    val crossZ = x1 * y2 - y1 * x2
    (crossX, crossY, crossZ)
  }

  override def toString: String = {
    val values = components.productIterator.map(_.toString).mkString(", ")
    s"($values)"
  }
}

// Трейт определяющий операции над векторами
trait VectorOps[T <: Product] {
  def add(a: T, b: T): T
  def dotProduct(a: T, b: T): Double
}

object VectorOps {
  // Неявные объекты для различных типов векторов
  implicit object intVectorOps extends VectorOps[(Int, Int)] {
    def add(a: (Int, Int), b: (Int, Int)): (Int, Int) = (a._1 + b._1, a._2 + b._2)
    def dotProduct(a: (Int, Int), b: (Int, Int)): Double = a._1 * b._1 + a._2 * b._2
  }

  implicit object doubleVectorOps extends VectorOps[(Double, Double)] {
    def add(a: (Double, Double), b: (Double, Double)): (Double, Double) = (a._1 + b._1, a._2 + b._2)
    def dotProduct(a: (Double, Double), b: (Double, Double)): Double = a._1 * b._1 + a._2 * b._2
  }

  implicit object tripleDoubleVectorOps extends VectorOps[(Double, Double, Double)] {
    def add(a: (Double, Double, Double), b: (Double, Double, Double)): (Double, Double, Double) =
      (a._1 + b._1, a._2 + b._2, a._3 + b._3)
    def dotProduct(a: (Double, Double, Double), b: (Double, Double, Double)): Double =
      a._1 * b._1 + a._2 * b._2 + a._3 * b._3
  }
}
```

# Тестирование
Протестируем программу на работе векторами.
Входные данные:
```scala
val vec1 = new Vector((1, 2))
val vec2 = new Vector((3, 4))
val vec3 = new Vector((1.0, 2.0))
val vec4 = new Vector((3.0, 4.0))
val vec5 = new Vector((1.0, 2.0, 3.0))
val vec6 = new Vector((4.0, 5.0, 6.0))

println(vec1 + vec2)  // (4, 6)
println(vec3 + vec4)  // (4.0, 6.0)
println(vec1.dotProduct(vec2))  // 11.0
println(vec3.dotProduct(vec4))  // 11.0
println(vec5.crossProduct(vec6))  // (-3.0, 6.0, -3.0)
println(vec5 + vec6) // (5.0, 7.0, 9.0)
println(vec5.dotProduct(vec6)) // 32
```
Выходные данные:
```
(4, 6)
(4.0, 6.0)
11.0
11.0
(-3.0,6.0,-3.0)
(5.0, 7.0, 9.0)
32.0
```
Программа отработала корректно

# Вывод
В процессе выполнения данной лабораторной работы я углубил свои знания о\
работе с обобщёнными классами в Scala, а также о применении неявных \
преобразований типов. Интересным моментом было изучение механизма неявных\
объектов и их применение для реализации операций над векторами различных \
типов данных. Это позволило создать универсальный класс Vector, способный\
работать как с целыми числами и дробными числами, так и с кортежами \
координат. Теперь я более уверенно могу применять обобщённые классы и неявные\
преобразования в своих Scala-проектах, что расширяет мой инструментарий для \
разработки.
