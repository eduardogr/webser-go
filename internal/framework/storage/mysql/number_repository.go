package mysql

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/eduardogr/webser-go/internal/adapters/interfaces/repositories"
	"github.com/eduardogr/webser-go/internal/application/config"
	"github.com/eduardogr/webser-go/internal/domain"
)

type Mysql struct {
	db *sql.DB
}

func NewMysqlNumberRepository(db *sql.DB) repositories.NumberRepository {
	s := &Mysql{db: db}
	err := s.Initialize()

	if err != nil {
		return nil
	}

	return s
}

func (s *Mysql) Initialize() error {
	// Create schema
	// TODO: Actually this is just used in development. A production service will not create tables.
	create, err := s.db.Query("CREATE TABLE IF NOT EXISTS numbers ( id INTEGER, timestamp VARCHAR(30) )")

	if err != nil {
		return err
	}

	defer create.Close()

	return nil
}

func (s *Mysql) CloseConnections() {
	s.db.Close()
}

func (s *Mysql) GetAll() ([]domain.Number, error) {
	results, err := s.db.Query("SELECT id, timestamp FROM numbers")

	if err != nil {
		return nil, err
	}

	defer results.Close()

	var Numbers []domain.Number
	for results.Next() {
		var n domain.Number
		err := results.Scan(&n.ID, &n.Timestamp)

		if err != nil {
			return nil, err
		}

		Numbers = append(Numbers, n)
	}

	return Numbers, nil
}

func (s *Mysql) Get(id int) (domain.Number, error) {
	results, err := s.db.Query("SELECT id, timestamp FROM numbers WHERE id = ?", id)

	if err != nil {
		return domain.Number{}, err
	}

	if results.Next() {
		var n domain.Number
		err := results.Scan(&n.ID, &n.Timestamp)

		if err != nil {
			return domain.Number{}, err
		}

		if n.ID == id {
			return n, nil
		}
	}

	defer results.Close()

	return domain.Number{}, nil
}

func (s *Mysql) Create(n domain.NewNumberRequest) error {
	timestamp := time.Now()
	create, err := s.db.Query("INSERT INTO numbers VALUES ( ?, ? )", n.ID, timestamp)

	if err != nil {
		panic(err.Error())
	}

	// be careful deferring Queries if you are using transactions
	defer create.Close()

	return nil
}

func (s *Mysql) Update(n domain.Number, id int) error {
	update, err := s.db.Query("UPDATE numbers SET id=? WHERE id=?", n.ID, id)

	if err != nil {
		return err
	}

	// be careful deferring Queries if you are using transactions
	defer update.Close()

	return nil
}

func (s *Mysql) Remove(id int) error {
	remove, err := s.db.Query("DELETE FROM numbers WHERE id = ?", id)

	if err != nil {
		return err
	}

	defer remove.Close()

	return nil
}

func SetupMysqlDatabase() (*sql.DB, error) {
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
