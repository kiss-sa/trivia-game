import { useEffect, useState } from 'react'
import Button from './components/Button'
import styles from './App.module.css'

function App() {
  const [boardCount, setBoardCount] = useState<number | null>(null)

  useEffect(() => {
    fetch('/api/boards')
      .then((res) => res.json())
      .then((data) => setBoardCount(data.length))
  }, [])

  return (
    <div className={styles.page}>
      <p>Boards available: {boardCount ?? '...'}</p>
      <Button>Start Game</Button>
    </div>
  )
}

export default App
