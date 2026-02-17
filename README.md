# Trivia Game

A Jeopardy-style trivia game for two teams. A moderator controls the game flow from a single PC, displayed on a TV screen. Teams take turns picking questions from a board with multiple categories and point values.

## Prerequisites

- [Go](https://go.dev/) 1.23+
- [Node.js](https://nodejs.org/) 18+

## Setup

**Backend:**

```sh
cd backend
go mod download
```

**Frontend:**

```sh
cd ui
npm install
```

## Running

Start both the backend and frontend:

```sh
# Terminal 1 – API server (runs on :8080)
cd backend
go run ./cmd/main.go

# Terminal 2 – Dev server (runs on :5173, proxies /api to backend)
cd ui
npm run dev
```

Then open http://localhost:5173 in your browser.

## Adding Boards

Boards are loaded from `backend/data/boards.csv`. Each row defines one question:

```
board_id,category,points,question,answer
```

Add new rows to create additional boards or categories.
