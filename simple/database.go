package simple

type Database struct {
	Name string
}

type DatabasePostgresql Database
type DatabaseMongoDB Database

func NewDatabasePostgresql() *DatabasePostgresql {
	return (*DatabasePostgresql)(&Database{Name: "Postgresql"})
}

func NewDatabaseMongoDB() *DatabaseMongoDB {
	return (*DatabaseMongoDB)(&Database{Name: "MongoDB"})
}

type DatabaseRepository struct {
	DatabasePostgresql *DatabasePostgresql
	DatabaseMongoDB    *DatabaseMongoDB
}

func NewDatabaseRepository(postgreSQL *DatabasePostgresql, mongoDB *DatabaseMongoDB) *DatabaseRepository {
	return &DatabaseRepository{DatabasePostgresql: postgreSQL, DatabaseMongoDB: mongoDB}
}
