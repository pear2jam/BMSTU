class InequalitySystem(coefficients: List[List[Double]], constants: List[Double]) {

  def this(coefficients: List[Double], constants: Double) = {
    this(List(coefficients), List(constants))
  }
  
  var inequalities: List[(List[Double], Double)] = coefficients.zip(constants)
  
  def +(other: InequalitySystem): InequalitySystem = {
    val combinedInequalities = this.inequalities ++ other.inequalities
    new InequalitySystem(combinedInequalities.map(_._1), combinedInequalities.map(_._2))
  }

  def /(index: Int): InequalitySystem = {
    println(coefficients.toString)
    val newCoefficients = coefficients.map(row => row.patch(index - 1, Nil, 1))
    new InequalitySystem(newCoefficients, constants)
  }

  def check(point: List[Double]): Boolean = {
    inequalities.forall { case (coeffs, constant) =>
      val sum = (coeffs, point).zipped.map(_ * _).sum
      sum <= constant
    }
  }
}

var sys = new InequalitySystem(List(1.0, 2.0, -3.0), 6.0)
var sys2 = new InequalitySystem(List(4.0, 2.0, -1.0), 2.0)
var sys_sum = sys + sys2
println(sys_sum.inequalities.toString())
var sys_mod = sys_sum / 2
println(sys_mod.inequalities.toString)
println(sys_mod.check(List(1.0, 2.0)).toString)
println(sys_mod.check(List(1.0, 4.0)).toString)

