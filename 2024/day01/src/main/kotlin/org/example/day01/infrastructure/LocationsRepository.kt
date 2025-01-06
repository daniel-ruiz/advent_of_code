package org.example.day01.infrastructure
import org.example.day01.domain.Locations
import java.io.Reader

class LocationsRepository {
  fun getFromReader(reader: Reader): Locations {
    var right: MutableList<Int> = mutableListOf()
    var left: MutableList<Int> = mutableListOf()
    val separator = "   "

    for (line in reader.readLines()) {
      val values = line.split(separator)
      val rightLocation = values.first().toInt()
      right.add(rightLocation)
      val leftLocation = values.last().toInt()
      left.add(leftLocation)
    }
    
    return Locations(right, left)
  }
}
