import { useEffect, useState } from 'react'
import { useNavigate } from 'react-router-dom'
import Button from '../components/Button'
import appStyles from '../App.module.css'
import type { Board } from './Board'

export default function Home() {
  const [boards, setBoards] = useState<Board[]>([])
  const navigate = useNavigate()

  useEffect(() => {
    fetch('/api/boards')
      .then((res) => res.json())
      .then(setBoards)
  }, [])

  return (
    <div className={appStyles.page}>
      <h1 className={appStyles.title}>Trivia Night</h1>
      <p>Boards available: {boards.length || '...'}</p>
      <Button
        disabled={boards.length === 0}
        onClick={() => navigate(`/board/${boards[0].id}`)}
      >
        Start Game
      </Button>
    </div>
  )
}
