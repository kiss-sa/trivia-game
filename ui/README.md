# Trivia Game — UI

React + TypeScript frontend for the trivia game, built with Vite. Communicates with the Go backend via a proxied `/api` route.

## Tech choices

**React Router** — used for client-side navigation between pages. Keeps URLs meaningful during development (you can navigate directly to `/board/board_01`) and scales cleanly as new pages are added.

**CSS Modules** — each component has its own `.module.css` file, scoped automatically by Vite. No extra dependencies, no runtime cost, and full CSS freedom. Global design tokens (colours, fonts) are defined as custom properties in `index.css` and referenced from any module.

## Structure

```
src/
  components/     # Reusable UI components (Button, ...)
  pages/          # Route-level page components (Home, Board)
  App.tsx         # Route definitions
  main.tsx        # App entry point
  index.css       # Global resets and CSS custom properties
```

## Start

```sh
npm install
npm run dev
```

Runs on http://localhost:5173. Requires the backend to be running on `:8080`.