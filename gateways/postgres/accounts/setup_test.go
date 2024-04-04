package accounts

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/golang-migrate/migrate"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"
)

var testCtx context.Context
var dbPGXPool *pgxpool.Pool

func TestMain(m *testing.M) {
	os.Exit(testMain(m))
}

func testMain(m *testing.M) int {
	// TODO create docker pool in extensions
	dockerPool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	if err = dockerPool.Client.Ping(); err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	resource, err := dockerPool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "latest",
		Env:        []string{"POSTGRES_USER=postgres", "POSTGRES_PASSWORD=postgres"},
	}, func(hc *docker.HostConfig) {
		hc.AutoRemove = true
		hc.RestartPolicy = docker.NeverRestart()
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	if err = dockerPool.Retry(func() error {
		dbPGXPool, err = pgxpool.New(context.Background(), fmt.Sprintf("postgres://postgres:postgres@localhost:%s/%s?sslmode=disable", resource.GetPort("5432/tcp"), "postgres"))
		if err != nil {
			return fmt.Errorf("creating pgx pool: %w", err)
		}

		if err = dbPGXPool.Ping(context.Background()); err != nil {
			return fmt.Errorf("pinging to postgres: %w", err)
		}

		return nil
	}); err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}

	// The migrations will be performed on the template1 database, so when a new database is created it will
	// already have the migrations applied.
	if err = runMigrations(); err != nil {
		log.Fatalf("running migrations: %s", err)
	}

	// You can't defer this because os.Exit doesn't care for defer
	teardownFn := func() {
		if err := dockerPool.Purge(resource); err != nil {
			log.Fatalf("Could not purge resource: %s", err)
		}
	}

	defer teardownFn()

	testCtx = context.Background()

	return m.Run()
}

func getGoModuleRoot() (string, error) {
	cmd := exec.Command("go", "env", "GOMOD")
	cmd.Env = os.Environ()

	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("getting go env GOMOD output: %w", err)
	}

	return filepath.Dir(string(output)), nil
}

func runMigrations() error {
	migrationPath := "gateways/postgres/migrations"
	connString := dbPGXPool.Config().ConnString()

	rootPath, err := getGoModuleRoot()
	if err != nil {
		return fmt.Errorf("getting go module root: %w", err)
	}

	path := filepath.Join(rootPath, migrationPath)
	m, err := migrate.New("file://"+path, connString)
	if err != nil {
		return fmt.Errorf("creating migrate instance: %w", err)
	}

	if err = m.Up(); err != nil {
		return fmt.Errorf("running up migrations: %w", err)
	}

	serr, err := m.Close()
	if serr != nil {
		return fmt.Errorf("closing the source: %w", serr)
	}

	if err != nil {
		return fmt.Errorf("closing pg connection: %w", err)
	}

	return nil
}
