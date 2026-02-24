import { useEffect, useState } from 'react'
import { useParams } from 'react-router-dom'
import Button from '../components/Button'
import styles from './Board.module.css'

export interface Board {
    id: string;
    categories: Category[];
}

export interface Category {
    id: number;
    name: string;
    questions: Question[];
}

export interface Question {
    id: number;
    points: number;
    questionText: string;
    answer: string;
}

export default function BoardPage() {
    const { id } = useParams<{ id: string }>()
    const [board, setBoard] = useState<Board | null>(null)

    useEffect(() => {
        fetch(`/api/boards/${id}`)
            .then((res) => res.json())
            .then((board: Board) => setBoard(board))
    }, [id])

    if (!board) return null

    const numCategories = board.categories.length
    const numQuestions = Math.max(...board.categories.map(c => c.questions.length))

    return (
        <div
            className={styles.grid}
            style={{ gridTemplateColumns: `repeat(${numCategories}, 1fr)` }}
        >
            {board.categories.map(category => (
                <div key={category.id} className={styles.categoryHeader}>
                    {category.name}
                </div>
            ))}

            {Array.from({ length: numQuestions }, (_, rowIndex) =>
                board.categories.map(category => {
                    const question = category.questions[rowIndex]
                    return (
                        <Button
                            key={`${category.id}-${rowIndex}`}
                        >
                            {question ? question.points : ''}
                        </Button>
                    )
                })
            )}
        </div>
    )
}
