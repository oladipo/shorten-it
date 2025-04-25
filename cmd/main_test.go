package main

import (
	"os"
	"testing"
)

func TestSetupApp_BoltDB(t *testing.T) {
	app, err := SetupApp("boltdb")
	if err != nil {
		t.Errorf("SetupApp(boltdb) returned error: %v", err)
	}
	if app == nil {
		t.Error("SetupApp(boltdb) returned nil app")
	}
}

func TestSetupApp_Redis(t *testing.T) {
	app, err := SetupApp("redis")
	if err != nil {
		t.Errorf("SetupApp(redis) returned error: %v", err)
	}
	if app == nil {
		t.Error("SetupApp(redis) returned nil app")
	}
}

func TestSetupApp_Postgres(t *testing.T) {
	os.Setenv("DATABASE_URL", "postgres://user:pass@localhost:5432/db")
	app, err := SetupApp("postgres")
	if err != nil {
		t.Errorf("SetupApp(postgres) returned error: %v", err)
	}
	if app == nil {
		t.Error("SetupApp(postgres) returned nil app")
	}
}

func TestSetupApp_Unknown(t *testing.T) {
	app, err := SetupApp("unknown")
	if err == nil {
		t.Error("SetupApp(unknown) should return error for unknown store type")
	}
	if app != nil {
		t.Error("SetupApp(unknown) should return nil app for unknown store type")
	}
}
