// +build integration

package databaseintegrationtest

import (
	"context"
	"testing"

	"github.com/soverdrive/anygo/src/testcollections/databaseintegrationtest/database"
	"github.com/soverdrive/anygo/src/testcollections/databaseintegrationtest/testingutil"
)

func TestInsertBook(t *testing.T) {
	t.Parallel()
	testDb, schemaName, cleanup := database.CreateTestDatabase(t)
	defer cleanup()

	loadTestData(t, testDb, schemaName, "book")
	insertBookQuery = addSchemaPrefix(schemaName, insertBookQuery) // override the query and add schema to the query

	title := "New title"
	author := "New author"
	err := InsertBook(context.Background(), testDb, title, author) // will execute insertBookQuery
	testingutil.Ok(t, err)
}

func TestGetBooks(t *testing.T) {
	t.Parallel()
	testDb, schemaName, cleanup := database.CreateTestDatabase(t)
	defer cleanup()

	loadTestData(t, testDb, schemaName, "book")
	getBooksQuery = addSchemaPrefix(schemaName, getBooksQuery) // override the query and add schema to the query

	books, err := GetBooks(context.Background(), testDb) // will execute getBooksQuery
	testingutil.Ok(t, err)
	testingutil.Equals(t, 2, len(books))
}
