package repository

import (
	entity "crud/entity/requests"
	"database/sql"
	"regexp"
	"testing"
	"time"

	"github.com/go-test/deep"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock
	db   *sql.DB
	err  error

	repository Repository
	Person     *entity.Userx
}

var u = &entity.User{
	Model: gorm.Model{
		ID:        1,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
		DeletedAt: nil,
	},
	Name:     "Dion",
	Age:      30,
	Nasabah:  "Simpedes",
	Email:    "Dion@gg.com",
	Password: "sssss",
}

var a = &entity.Userx{
	Id:       1,
	Name:     "Dion",
	Age:      30,
	Nasabah:  "Simpedes",
	Email:    "Dion@gg.com",
	Password: "sssss",
}

func (s *Suite) SetupSuite() {
	s.db, s.mock, s.err = sqlmock.New()
	require.NoError(s.T(), s.err)
	s.DB, s.err = gorm.Open("postgres", s.db)
	//require.NoError(s.T(), s.err)

	//s.DB.LogMode(true)

	s.repository = NewRepository(s.DB)
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) Test_repository_Get() {
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."deleted_at" IS NULL AND ((id = $1))`)).
		WithArgs(u.Model.ID).
		WillReturnRows(sqlmock.NewRows([]string{"name", "age", "nasabah", "email"}).
			AddRow(u.Name, u.Age, u.Nasabah, u.Email))
	result, err := s.repository.GetId(int(u.Model.ID))
	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(&entity.User{Name: u.Name, Age: u.Age, Nasabah: u.Nasabah, Email: u.Email}, result))
	require.NotNil(s.T(), result)
}

func (s *Suite) Test_repository_Getx() {
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users" WHERE (id = $1)`)).
		WithArgs(a.Id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "age", "nasabah", "email"}).
			AddRow(a.Id, a.Name, a.Age, a.Nasabah, a.Email))

	res, _ := s.repository.GetIdx(a.Id)
	//require.NoError(s.T(), err)
	require.NotNil(s.T(), *res)
}
