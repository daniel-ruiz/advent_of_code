package adventofcode2024
class ComputeTotalDistanceBetweenLocationLists {
  fun execute(rightList: List<Int>, leftList: List<Int>): Int {
    val rightLocationID = rightList[0]
    val leftLocationID = leftList[0]

    return rightLocationID - leftLocationID
  }
}
