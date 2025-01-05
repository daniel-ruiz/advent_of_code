package adventofcode2024
import kotlin.math.abs

class ComputeTotalDistanceBetweenLocationLists {
  fun execute(rightList: List<Int>, leftList: List<Int>): Int {
    val rightLocationID = rightList[0]
    val leftLocationID = leftList[0]

    return abs(rightLocationID - leftLocationID)
  }
}
