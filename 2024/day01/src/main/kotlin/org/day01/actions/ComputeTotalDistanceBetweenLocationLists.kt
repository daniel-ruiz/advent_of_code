package org.day01.actions
import org.day01.domain.Locations
import kotlin.math.abs

class ComputeTotalDistanceBetweenLocationLists {
  fun execute(locations: Locations): Int {
    var totalDistance = 0
    val rightList: List<Int> = locations.rightList.sorted()
    val leftList: List<Int> = locations.leftList.sorted()

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
