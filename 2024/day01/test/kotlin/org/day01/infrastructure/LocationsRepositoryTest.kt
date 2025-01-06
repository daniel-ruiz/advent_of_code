package org.day01.infrastructure
import org.day01.domain.Locations
import org.junit.jupiter.api.Assertions.*
import org.junit.jupiter.api.Test

class LocationsRepositoryTest {
  @Test
  fun `Should return locations from an input stream reader`() {
    val reader = "3   4\n4   3".reader()
    val expectedLocations = Locations(rightList=listOf(3, 4), leftList=listOf(4, 3))

    val locations = LocationsRepository().getFromReader(reader)

    assertEquals(locations, expectedLocations)
  }
}
