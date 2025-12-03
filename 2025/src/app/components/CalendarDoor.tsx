import './CalendarDoor.css'

interface DayResult {
  day: number
  part1: number
  part2: number
  duration: number
}

interface CalendarDoorProps {
  day: number
  isOpen: boolean
  isRunning: boolean
  result?: DayResult | undefined
  onClick: () => void
}

const CalendarDoor = ({ day, isOpen, isRunning, result, onClick }: CalendarDoorProps) => {
  const dayStr = String(day).padStart(2, '0')
  
  return (
    <div 
      className={`calendar-door ${isOpen ? 'open' : ''} ${isRunning ? 'running' : ''}`}
      onClick={onClick}
    >
      <div className="door-front">
        <div className="door-number">{day}</div>
        <div className="door-decoration">❄️</div>
      </div>
      <div className="door-back">
        {isRunning ? (
          <div className="loading">
            <div className="spinner">⏳</div>
            <p>Running...</p>
          </div>
        ) : result ? (
          <div className="result">
            <h3>Day {day}</h3>
            <div className="result-item">
              <span className="label">Part 1:</span>
              <span className="value">{result.part1}</span>
            </div>
            <div className="result-item">
              <span className="label">Part 2:</span>
              <span className="value">{result.part2}</span>
            </div>
            <div className="result-time">
              {result.duration.toFixed(2)}ms
            </div>
          </div>
        ) : (
          <div className="empty">
            <p>Click to run!</p>
          </div>
        )}
      </div>
    </div>
  )
}

export default CalendarDoor
