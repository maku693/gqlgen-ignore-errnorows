package db

import "context"

type Todo struct {
	ID   int
	Text string
	Done bool
}

func ListTodos(ctx context.Context, db DB) ([]*Todo, error) {
	rows, err := db.QueryContext(ctx, "SELECT id, text, done FROM todos ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := make([]*Todo, 0)
	for rows.Next() {
		todo := new(Todo)
		if err := rows.Scan(&todo.ID, &todo.Text, &todo.Done); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func GetTodoByID(ctx context.Context, db DB, id int) (*Todo, error) {
	row := db.QueryRowContext(ctx, "SELECT id, text, done FROM todos WHERE id = ?", id)

	todo := new(Todo)
	if err := row.Scan(&todo.ID, &todo.Text, &todo.Done); err != nil {
		return nil, err
	}

	return todo, nil
}
