package repository

import (
	"fmt"
	"time"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/eduardogr/webser-go/internal/api"
	"github.com/eduardogr/webser-go/internal/config"
)

type Storage interface {
	Initialize() error
	CloseConnections()

	GetAll() ([]api.Number, error)
	Get(id int) (api.Number, error)
	Create(n api.NewNumberRequest) error
	Update(n api.Number, id int) error
	Remove(id int) error
}

type storage struct {
	db *sql.DB
}

func ProvideNumberRepository(s Storage) api.NumberRepository {
	n := new(api.NumberRepository)
	*n = s
	return *n
}

func NewStorage(db *sql.DB) Storage {
	s := &storage{db: db}
	err := s.Initialize()

	if err != nil {
		return nil
	}

	return s
}

func (s *storage) Initialize() error {
	// Create schema
	create, err := s.db.Query("CREATE TABLE IF NOT EXISTS numbers ( id INTEGER, timestamp VARCHAR(30) )")

	if err != nil {
		return err
	}

	defer create.Close()

	return nil
}

func (s *storage) CloseConnections() {
	s.db.Close()
}

func (s *storage) GetAll() ([]api.Number, error) {
	results, err := s.db.Query("SELECT id, timestamp FROM numbers")

	if err != nil {
		return nil, err
	}

	defer results.Close()

	var Numbers []api.Number
	for results.Next() {
		var n api.Number
		err := results.Scan(&n.ID, &n.Timestamp)

		if err != nil {
			return nil, err
		}

		Numbers = append(Numbers, n)
	}

	return Numbers, nil
}

func (s *storage) Get(id int) (api.Number, error) {
	results, err := s.db.Query("SELECT id, timestamp FROM numbers WHERE id = ?", id)

	if err != nil {
		return api.Number{}, err
	}

	if results.Next() {
		var n api.Number
		err := results.Scan(&n.ID, &n.Timestamp)

		if err != nil {
			return api.Number{}, err
		}

		if n.ID == id {
			return n, nil
		}
	}

	defer results.Close()

	return api.Number{}, nil
}

func (s *storage) Create(n api.NewNumberRequest) error {
	timestamp := time.Now()
	create, err := s.db.Query("INSERT INTO numbers VALUES ( ?, ? )", n.ID, timestamp)

	if err != nil {
		panic(err.Error())
	}

	// be careful deferring Queries if you are using transactions
	defer create.Close()

	return nil
}

func (s *storage) Update(n api.Number, id int) error {
	update, err := s.db.Query("UPDATE numbers SET id=? WHERE id=?", n.ID, id)

	if err != nil {
		return err
	}

	// be careful deferring Queries if you are using transactions
	defer update.Close()

	return nil
}

func (s *storage) Remove(id int) error {
	remove, err := s.db.Query("DELETE FROM numbers WHERE id = ?", id)

	if err != nil {
		return err
	}

	defer remove.Close()

	return nil
}

func SetupDatabase() (*sql.DB, error) {
	db, err := getConnection()

	if err != nil {
		return nil, err
	}

	// ping the DB to ensure that it is connected
	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}

func getConnection() (*sql.DB, error) {
	c := config.GetConfiguration(config.STRATEGY)
	dbConnection := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		c.DbUser, c.DbPassword,
		c.DbHost, c.DbPort, c.DbDatabase)

	db, err := sql.Open("mysql", dbConnection)

	if err != nil {
		return nil, err
	}

	return db, nil
}
