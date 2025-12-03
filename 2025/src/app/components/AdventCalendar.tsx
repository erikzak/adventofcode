import { useState } from 'react'
import CalendarDoor from './CalendarDoor'
import './AdventCalendar.css'
import { dayModules } from '../utils/dayModules'

interface DayResult {
  day: number
  part1: number
  part2: number
  duration: number
}

const AdventCalendar = () => {
  const [openDoors, setOpenDoors] = useState<Set<number>>(new Set())
  const [results, setResults] = useState<Map<number, DayResult>>(new Map())
  const [isRunning, setIsRunning] = useState<number | null>(null)

  const handleDoorClick = async (day: number) => {
    if (isRunning !== null) return
    
    // Check if day is implemented
    if (!(day in dayModules)) {
      alert(`Day ${day} has not been implemented yet!`)
      return
    }
    
    const newOpenDoors = new Set(openDoors)
    newOpenDoors.add(day)
    setOpenDoors(newOpenDoors)
    setIsRunning(day)

    try {
      const startTime = performance.now()
      
      // Dynamically import the day's module and test input
      const dayStr = `d${String(day).padStart(2, '0')}`
      
      // Fetch test input
      const testInputResponse = await fetch(`/src/${dayStr}/test.txt`)
      if (!testInputResponse.ok) {
        throw new Error(`Could not load test input for day ${day}`)
      }
      const testInputText = await testInputResponse.text()
      // Split on both \n and \r\n to handle Windows (CRLF) and Unix (LF) line endings
      const testInput = testInputText.split(/\r?\n/).filter(line => line.trim() !== '')
      console.log(testInput)
      
      // Dynamically import and run the solution using explicit import map
      const importFn = dayModules[day as keyof typeof dayModules]
      const module = await importFn()
      
      if (!module.main || typeof module.main !== 'function') {
        throw new Error(`Day ${day} module does not export a main function`)
      }
      
      // Run the solution
      const [part1, part2] = module.main(testInput)
      const duration = performance.now() - startTime
      
      const newResults = new Map(results)
      newResults.set(day, { day, part1, part2, duration })
      setResults(newResults)
    } catch (error) {
      console.error(`Error running day ${day}:`, error)
      const errorMessage = error instanceof Error ? error.message : String(error)
      alert(`Error running day ${day}: ${errorMessage}`)
      
      // Remove from open doors if there was an error
      const newOpenDoors = new Set(openDoors)
      newOpenDoors.delete(day)
      setOpenDoors(newOpenDoors)
    } finally {
      setIsRunning(null)
    }
  }

  // Generate 12 days
  const days = Array.from({ length: 12 }, (_, i) => i + 1)

  return (
    <div className="advent-calendar">
      <div className="calendar-grid">
        {days.map((day) => (
          <CalendarDoor
            key={day}
            day={day}
            isOpen={openDoors.has(day)}
            isRunning={isRunning === day}
            result={results.get(day)}
            onClick={() => handleDoorClick(day)}
          />
        ))}
      </div>
    </div>
  )
}

export default AdventCalendar
