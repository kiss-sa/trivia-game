package api

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// LoadBoardsFromCSV reads a CSV file with columns board_id,category,points,question,answer
// and returns a slice of Board structs.
func LoadBoardsFromCSV(path string) ([]*Board, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open csv: %w", err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.TrimLeadingSpace = true

	records, err := r.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("read csv: %w", err)
	}

	if len(records) < 2 {
		return nil, fmt.Errorf("csv has no data rows")
	}

	// skip header row
	rows := records[1:]

	// boardID -> categoryName -> questions
	type boardData struct {
		categories map[string][]*Question
		order      []string // insertion order for categories
	}
	boards := map[string]*boardData{}
	boardOrder := []string{}

	for i, row := range rows {
		if len(row) < 5 {
			return nil, fmt.Errorf("row %d: expected 5 columns, got %d", i+2, len(row))
		}
		boardID := row[0]
		category := row[1]
		points, err := strconv.Atoi(row[2])
		if err != nil {
			return nil, fmt.Errorf("row %d: invalid points %q: %w", i+2, row[2], err)
		}
		question := row[3]
		answer := row[4]

		bd, exists := boards[boardID]
		if !exists {
			bd = &boardData{categories: map[string][]*Question{}}
			boards[boardID] = bd
			boardOrder = append(boardOrder, boardID)
		}

		if _, ok := bd.categories[category]; !ok {
			bd.order = append(bd.order, category)
		}
		bd.categories[category] = append(bd.categories[category], &Question{
			QuestionText: question,
			Answer:       answer,
			Points:       points,
		})
	}

	result := make([]*Board, 0, len(boardOrder))
	for _, boardID := range boardOrder {
		bd := boards[boardID]
		cats := make([]Category, 0, len(bd.order))
		for i, catName := range bd.order {
			cats = append(cats, Category{
				ID:        int64(i + 1),
				Name:      catName,
				Questions: bd.categories[catName],
			})
		}
		result = append(result, &Board{
			ID:         boardID,
			Categories: cats,
		})
	}

	return result, nil
}
