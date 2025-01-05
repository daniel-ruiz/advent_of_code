package adventofcode2024
import org.junit.jupiter.api.Assertions.*
import org.junit.jupiter.api.Test

class ComputeTotalDistanceBetweenLocationListsTest {

    @Test
    fun `Should return the difference of location IDs when right and left lists contain one location`() {
        val rigthList = listOf(5)
        val leftList = listOf(3)

        val distance = ComputeTotalDistanceBetweenLocationLists().execute(rigthList, leftList)

        assertEquals(distance, 2)
    }

    @Test
    fun `Should return a positive distance when the right location ID is less than the left location ID`() {
        val rightList = listOf(4)
        val leftList = listOf(11)

        val distance = ComputeTotalDistanceBetweenLocationLists().execute(rightList, leftList)

        assertEquals(distance, 7) 
    }

    @Test
    fun `Should return the sum of distances of each pair of location IDs when the lists are ordered and have more than one location`() {
        val rightList = listOf(2, 6)
        val leftList = listOf(1, 11)

        val distance = ComputeTotalDistanceBetweenLocationLists().execute(rightList, leftList)

        assertEquals(distance, 6) 
    }
}
