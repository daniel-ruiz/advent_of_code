package adventofcode2024
import org.junit.jupiter.api.Assertions.*
import org.junit.jupiter.api.Test

class ComputeTotalDistanceBetweenLocationListsTest {

    @Test
    fun `Should return the difference of location IDs when right and left lists contain one location`() {
        val rigthList = listOf(5)
        val leftList = listOf(3)

        val action = ComputeTotalDistanceBetweenLocationLists()

        assertEquals(action.execute(rigthList, leftList), 2)
    }
}
