package mysqldb

import "cleancodego/internal/domain/repository"

type mysqlRepository struct {

	// Our db connections will go

}

func NewMysqlRepository() repository.UserRepository {

	return &mysqlRepository{}

}

func (s *mysqlRepository) Save() string {
	return "Repository Mysql Got Called"

}
