package org.adventofcode2024
import org.junit.jupiter.api.Assertions.*
import org.junit.jupiter.api.Test

class PetTest {
    @Test
    fun `Should create a pet with name and type`() {
        val pet = Pet("Buddy", "Dog")
        assertEquals("Buddy", pet.name)
        assertEquals("Doge", pet.type)
    }
}
