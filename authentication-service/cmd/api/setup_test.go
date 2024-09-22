package main

import (
	"authentication-service/data"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	repo := data.NewPostgresTestRepository(nil)
	testApp.Repo = repo
	os.Exit(m.Run())
}
