# Go SQL Injection Demo

Demo code for SQL injection in Go using MariaDB / MySQL with multiStatements parameter set to true. 
Includes solution which is simply using SQL statement parameters.

Important highlights:

SQL injection code: 

```
'); truncate messages; -- 
```

MySQL driver configuration parameter: 

```
multiStatements=true
```

Support code to the article and video published at:

- https://oak.dev/2024/07/17/golang-sql-injection-in-mariadb-mysql/
- https://youtu.be/OcMOXMXYRXo
