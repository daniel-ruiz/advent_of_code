package adventofcode2024
import kotlin.math.abs

class ComputeTotalDistanceBetweenLocationLists {
  fun execute(rightList: List<Int>, leftList: List<Int>): Int {
    var totalDistance = 0
    val rightList: List<Int> = rightList.sorted()
    val leftList: List<Int> = leftList.sorted()

    for (i in 0..rightList.count()-1) {
      val rightLocationID = rightList[i]
      val leftLocationID = leftList[i]

      totalDistance += calculateDistanceBetweenLocationIDs(rightLocationID, leftLocationID)
    }

    return totalDistance
  }

  private fun calculateDistanceBetweenLocationIDs(right: Int, left: Int): Int {
    return abs(right - left)
  }
}
