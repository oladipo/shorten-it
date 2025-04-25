package storage

// Storage is an interface for URL storage backends
// Implementations: BoltDB, Redis, Postgres

type Storage interface {
	// Placeholder: Define required methods
}

// NewBoltDB returns a new BoltDB storage (placeholder)
func NewBoltDB(path string) Storage {
	return nil
}

// NewRedis returns a new Redis storage (placeholder)
func NewRedis(addr, password string) Storage {
	return nil
}

// NewPostgres returns a new Postgres storage (placeholder)
func NewPostgres(url string) Storage {
	return nil
}
