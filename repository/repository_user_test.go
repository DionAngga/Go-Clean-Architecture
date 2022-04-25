package repository

import (
	entity "crud/entity/requests"
	"database/sql"
	"fmt"
	"regexp"
	"strconv"
	"testing"

	"github.com/go-test/deep"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

func (s *Suite) SetupSuite() {
	s.db, s.mock, s.err = sqlmock.New()
	require.NoError(s.T(), s.err)
	dialector := mysql.New(mysql.Config{
		Conn:                      s.db,
		SkipInitializeWithVersion: true,
	})
	s.DB, s.err = gorm.Open(dialector, &gorm.Config{})
	fmt.Println("database ========", s.db)
	fmt.Println("s.Database ========", s.DB)
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

func (s *Suite) Test_repository_Get() {
	var (
		id      = 1
		name    = "test-name"
		age     = 20
		nasabah = "simpedes"
		email   = "testname@gg.com"
	)
	idstr := strconv.Itoa(id)
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "u" WHERE (id = $1)`)).
		WithArgs(idstr).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
			AddRow(idstr, name, age, nasabah, email))

	res, err := s.repository.GetId(idstr)
	fmt.Println("ssssssssssssssssssssssss===========", s.T())
	fmt.Println("ressssssssssssssssssdccs===========", res)
	fmt.Println("----------MODEL----===========", &entity.Userx{Id: id, Name: name, Age: age, Nasabah: nasabah, Email: email})
	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(&entity.Userx{Id: id, Name: name, Age: age, Nasabah: nasabah, Email: email}, res))
}
