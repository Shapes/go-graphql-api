package postgres

import (
	"database/sql"
	"fmt"

	// postgres driver
	_ "github.com/lib/pq"
)

// Db is our database struct used for interacting with the database
type Db struct {
	*sql.DB
}

// New makes a new database using the connection string and
// returns it, otherwise returns the error
func New(connString string) (*Db, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	// Check that our connection is good
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Db{db}, nil
}

// ConnString returns a connection string based on the parameters it's given
// This would normally also contain the password, however we're not using one
func ConnString(host string, port int, user string, dbName string, pass string) string {
	return fmt.Sprintf(
		"host=%s password=%s port=%d user=%s dbname=%s sslmode=require",
		host, pass, port, user, dbName,
	)
}

// User shape
type User struct {
	ID        int
	FirstName string
	LastName  string
}

// GetUsersByName is called within our user query for graphql
func (d *Db) GetUsersByName(firstName string) []User {

	// Prepare query, takes a name argument, protects from sql injection
	stmt, err := d.Prepare(`SELECT "id", "firstName", "lastName" FROM "Users" WHERE "firstName"=$1`)

	if err != nil {
		fmt.Println("GetUserByName Preperation Err: ", err)
	}

	// Make query with our stmt, passing in name argument
	rows, err := stmt.Query(firstName)
	if err != nil {
		fmt.Println("GetUserByName Query Err: ", err)
	}

	// Create User struct for holding each row's data
	var r User
	// Create slice of Users for our response
	users := []User{}
	// Copy the columns from row into the values pointed at by r (User)
	for rows.Next() {
		err = rows.Scan(
			&r.ID,
			&r.FirstName,
			&r.LastName,
		)
		if err != nil {
			fmt.Println("Error scanning rows: ", err)
		}
		users = append(users, r)
	}

	return users
}

// Job shape
type Job struct {
	ID    int
	Title string
}

// GetJobsByTitle is called within our user query for graphql
func (d *Db) GetJobsByTitle(title string) []Job {

	// Prepare query, takes a name argument, protects from sql injection
	stmt, err := d.Prepare(`SELECT "id", "title" FROM "Jobs" WHERE "title"=$1`)

	if err != nil {
		fmt.Println("GetJobsByTitle Preperation Err: ", err)
	}

	// Make query with our stmt, passing in name argument
	rows, err := stmt.Query(title)
	if err != nil {
		fmt.Println("GetJobsByTitle Query Err: ", err)
	}

	// Create User struct for holding each row's data
	var r Job
	// Create slice of Users for our response
	jobs := []Job{}
	// Copy the columns from row into the values pointed at by r (User)
	for rows.Next() {
		err = rows.Scan(
			&r.ID,
			&r.Title,
		)
		if err != nil {
			fmt.Println("Error scanning rows: ", err)
		}
		jobs = append(jobs, r)
	}

	return jobs
}
