package graph

import (
	"database/sql"

	"github.com/joseporres/sonarcloud_example/graph/connection"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	DB *sql.DB
}

func (r *Resolver) InitializePool(){
	r.DB = connection.FetchConnection()
}
