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

Support code to the article published at https://oak.dev and video at: https://youtu.be/OcMOXMXYRXo
