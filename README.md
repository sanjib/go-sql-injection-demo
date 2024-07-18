# Go SQL Injection Demo

Demo code for SQL injection in Go using MariaDB / MySQL with multiStatements parameter set to true. 
Includes solution which is simply using SQL statement parameters.

## Installation

1. Clone repository: git clone https://github.com/sanjib/go-sql-injection-demo
2. cd "go-sql-injection-demo"
3. Run `go mod tidy` to install dependencies: one direct dependency:
   "go-sql-driver"
4. Create MariaDB / MySQL table using db.sql file.
5. Edit main.go: func openDB(): replace the database credentials with your 
   own: 

```go
db, err = sql.Open("mysql", "root@/va_test1?parseTime=true&multiStatements=true")
```

Replace `root@/va_test1` with your own username, password and database name:
`your_username:your_password@/your_database_name`

## How to Run

1. Toggle the following 2 lines of code in main.go: func
   insertMessage(msg string):

```go
// Allows SQL injection.
// Uncomment following 2 lines to demo SQL injection.
q := fmt.Sprintf(`INSERT INTO messages (message) VALUES ('%s')`, msg)
_, err := db.Exec(q)

// Prevents SQL injection.
// Uncomment following 2 lines to prevent SQL injection.
//q := `INSERT INTO messages (message) VALUES (?)`
//_, err := db.Exec(q, msg)
```

2. Inside directory "go-sql-injection-demo", run: `go run .` You should see the
   message `starting server on port 3000...` 

3. Point your browser to: `http://localhost:3000/`

4. Insert some demo data. Then enter `'); truncate messages; -- ` to see SQL
   injection truncate / wipe out all records in `messages` table.

5. Toggle the following 2 lines of code again in main.go: func
   insertMessage(msg string):

```go
// Allows SQL injection.
// Uncomment following 2 lines to demo SQL injection.
//q := fmt.Sprintf(`INSERT INTO messages (message) VALUES ('%s')`, msg)
//_, err := db.Exec(q)

// Prevents SQL injection.
// Uncomment following 2 lines to prevent SQL injection.
q := `INSERT INTO messages (message) VALUES (?)`
_, err := db.Exec(q, msg)
```

6. Repeat step 4 above. SQL injection should no longer be possible.

## Important highlights
There are 3 criteria to be able to demo SQL injection in this program. 

1. The SQL injection code itself (note space after -- below): 

```
'); truncate messages; -- 
```

2. MySQL driver configuration parameter: 

```
multiStatements=true
```

3. Non-Parameterized code

```go
q := fmt.Sprintf(`INSERT INTO messages (message) VALUES ('%s')`, msg)
_, err := db.Exec(q)
```

Support code to the article and video published at:

- https://oak.dev/2024/07/17/golang-sql-injection-in-mariadb-mysql/
- https://youtu.be/OcMOXMXYRXo
