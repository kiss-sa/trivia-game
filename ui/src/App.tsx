import { useEffect, useState } from 'react'

function App() {
  const [boardCount, setBoardCount] = useState<number | null>(null)

  useEffect(() => {
    fetch('/api/boards')
      .then((res) => res.json())
      .then((data) => setBoardCount(data.length))
  }, [])

  return (
    <div>
      <p>Boards available: {boardCount ?? '...'}</p>
      <button>Start Game</button>
    </div>
  )
}

export default App
