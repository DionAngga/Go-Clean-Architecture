package repository

import (
	entity "crud/entity/requests"
	"database/sql"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock
	db   *sql.DB
	err  error

	repository Repository
}

func (s *Suite) SetupSuite() {
	s.db, s.mock, s.err = sqlmock.New()
	require.NoError(s.T(), s.err)
	dialector := mysql.New(mysql.Config{
		Conn:                      s.db,
		SkipInitializeWithVersion: true,
	})
	s.DB, s.err = gorm.Open(dialector, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	require.NoError(s.T(), s.err)

	//s.DB.LogMode(true)

	s.repository = NewRepository(s.DB)
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

var u = &entity.User{
	Model: gorm.Model{
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
		DeletedAt: gorm.DeletedAt{},
	},
	Name:     "Dion",
	Age:      30,
	Nasabah:  "Simpedes",
	Email:    "Dion@gg.com",
	Password: "sssss",
}

var a = &entity.User{
	Model: gorm.Model{
		ID:        1,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
		DeletedAt: gorm.DeletedAt{},
	},
	Name:     "Dion",
	Age:      30,
	Nasabah:  "Simpedes",
	Email:    "Dion@gg.com",
	Password: "sssss",
}

func (s *Suite) Test_repository_Get() {
	s.mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `users` WHERE `users`.`id` = ?")).
		WithArgs(a.Model.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "age", "nasabah", "email", "password"}).
			AddRow(a.Model.ID, a.Name, a.Age, a.Nasabah, a.Email, a.Password))

	res, err := s.repository.GetId(int(a.Model.ID))
	require.NoError(s.T(), err)
	require.NotNil(s.T(), res)
}

func (s *Suite) Test_repository_Create() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`name`,`age`,`nasabah`,`email`,`password`) VALUES (?,?,?,?,?,?,?,?)")).
		WithArgs(u.CreatedAt, u.UpdatedAt, u.Model.DeletedAt, u.Name, u.Age, u.Nasabah, u.Email, u.Password).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	user, err := s.repository.Create(u)

	require.NoError(s.T(), err)
	require.NotNil(s.T(), user)
}
