import { useState, useEffect, useRef } from 'react'
import './App.css'
import AdventCalendar from './components/AdventCalendar'

function App() {
  const [isPlaying, setIsPlaying] = useState(false)
  const audioRef = useRef<HTMLAudioElement>(null)

  const toggleMusic = () => {
    if (audioRef.current) {
      if (isPlaying) {
        audioRef.current.pause()
      } else {
        audioRef.current.play().catch(err => {
          console.error('Failed to play audio:', err)
          setIsPlaying(false)
        })
      }
      setIsPlaying(!isPlaying)
    }
  }

  return (
    <div className="app">
      <audio 
        ref={audioRef} 
        loop
        preload="auto"
        src="/christmas-jazz-christmas-holiday-347485.mp3"
      />
      <button className="music-toggle" onClick={toggleMusic} title={isPlaying ? 'Pause Music' : 'Play Music'}>
        {isPlaying ? '🔊 🎵' : '🔇 🎵'}
      </button>
      <header>
        <h1>🎄 Advent of Code 2025 🎄</h1>
      </header>
      <main>
        <AdventCalendar />
      </main>
    </div>
  )
}

export default App
