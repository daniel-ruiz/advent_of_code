package org.example.day01.actions
import org.example.day01.domain.Locations
import org.junit.jupiter.api.Assertions.*
import org.junit.jupiter.api.Test

class ComputeTotalDistanceBetweenLocationListsTest {

    @Test
    fun `Should return the difference of location IDs when right and left lists contain one location`() {
        val locations = Locations(rightList=listOf(5), leftList=listOf(3))

        val distance = ComputeTotalDistanceBetweenLocationLists().execute(locations)

        assertEquals(distance, 2)
    }

    @Test
    fun `Should return a positive distance when the right location ID is less than the left location ID`() {
        val locations = Locations(rightList=listOf(4), leftList=listOf(11))

        val distance = ComputeTotalDistanceBetweenLocationLists().execute(locations)

        assertEquals(distance, 7) 
    }

    @Test
    fun `Should return the sum of distances of each pair of location IDs when the lists are ordered and have more than one location`() {
        val locations = Locations(rightList=listOf(2, 6), leftList=listOf(1, 11))

        val distance = ComputeTotalDistanceBetweenLocationLists().execute(locations)

        assertEquals(distance, 6) 
    }

    @Test
    fun `Should return the sum of distances of each pair of location IDs when the lists are not ordered`() {
        val locations = Locations(rightList=listOf(2, 6), leftList=listOf(11, 1))

        val distance = ComputeTotalDistanceBetweenLocationLists().execute(locations)

        assertEquals(distance, 6) 
    }
}
