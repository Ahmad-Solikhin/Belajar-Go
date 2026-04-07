package simple

type Database struct {
	Name string
}

type DatabasePostgresSQL Database
type DatabaseMongoDB Database

func NewDatabaseMongoDB() *DatabaseMongoDB {
	return &DatabaseMongoDB{Name: "MongoDB"}
}

func NewDatabasePostgresSQL() *DatabasePostgresSQL {
	return &DatabasePostgresSQL{Name: "PostgresSQL"}
}

type DatabaseRepository struct {
	DatabasePostgresSQL *DatabasePostgresSQL
	DatabaseMongoDB     *DatabaseMongoDB
}

func NewDatabaseRepository(postgresSQL *DatabasePostgresSQL, mongoDB *DatabaseMongoDB) *DatabaseRepository {
	return &DatabaseRepository{DatabasePostgresSQL: postgresSQL, DatabaseMongoDB: mongoDB}
}
