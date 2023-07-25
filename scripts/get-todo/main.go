package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/maku693/gqlgen-ignore-errnorows/db"
)

func main() {
	if err := ignoreErrNoRows(cmd); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func cmd() error {
	database, err := db.InitializeDB()
	if err != nil {
		return err
	}

	if len(os.Args) < 2 {
		return errors.New("no id provided")
	}

	id, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return err
	}

	todo, err := db.GetTodoByID(context.Background(), database, id)
	if err != nil {
		return err
	}

	fmt.Printf("%#v\n", todo)

	return nil
}

func ignoreErrNoRows(f func() error) error {
	err := f()
	if errors.Is(err, sql.ErrNoRows) {
		return nil
	}
	return err
}
