import { useState } from 'react'
import './App.css'
import AdventCalendar from './components/AdventCalendar'

function App() {
  return (
    <div className="app">
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
