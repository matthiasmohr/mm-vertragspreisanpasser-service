# Repository (Data Source layer)
Repository is for communication with the database.  
Databases accept and return only domain models.
The _Repository Implementation_ implements the _Domain Repository Interface_. It is in charge of combining one or multiple Data Sources.
 
The _Domain Repository Interface_ is described in `repository/store.go`. 
It describes the functions every repository has to implement. These fall into two categories. The general functions include standard database operations like
```go
BeginTransaction() (Store, error)
Commit() error
Rollback() error
```

In order to work with the data in the database tables, some more functions are necessary. These form the second category and are usually project specific. If for example there's a domain entity `Customer` in the project, the domain repository interface would include 
```go
Customer() Customer 
```
