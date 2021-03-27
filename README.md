
## Baselines

### Package design

We are isolating our business domain to its own package, called domain. The domain
package holds the domain types for our application, 
in case of an online store those would be: users, orders etc.

Our domain package also holds abstractions of the services(logical domain types) that are used by our components to 
communicate between them. This lets us break any direct depenedencies between packages as
every package only has to depend on the domain. If a package has a dependnecy on another package,
it should import an interface from the domain that the dependency package implements.

We are isolating our presentation logic in packages at the same level as the domain package. 
We name these packages after their underlying technology/communication channel e.g. `api`, `http`

We are isolating our infrastructure logic in subpackages of the domain package. The purpose of these
subpackages is to implement functionality defined by the interfaces in the domain.

#### Domain Package 

- Domain package is used to hold high-level information about the problem we
  are trying to solve: Types and interfaces, errors.
  
  e.g. the User type, the UserService interface.

- Domain package is not allowed to have any external dependencies.
  
  We are not importing anything related to implementation, doing our best to use primitive types when possible.

  
#### Domain Subpackages

- Domain Subpackages are an adapter between the domain and the implementation. They offer the actual functionality
defined by the domain interfaces.
  
  e.g. while having an UserService interface in the domain we would like to create an implementation
  backed by PostgreSQL. We create a `psql` package where we define an UserService implementation.
  
  The `psql` package can hold multiple services implementations backed by Postgres, so we could also have
  an OrderService and a PubSub.
  
  This isolates our PostgreSQL dependency which simplifies testing, makes it easy to migrate
  to another db in the future, or use different implementations.
  
- Domain Subpackages have to offer functionality and are not used only to provide a namespace for types, variables or constants. 
  _To be purposeful, packages must provide, not contain_ from Design Philosophy On Packaging by B. Kennedy.

- Service methods should implement transactional boundary so the transaction object will not be shared outside the method.

- Domain Subpackages can not depenend on other subpackages, but they can depend on the domain.
  
  e.q. We develop a budgeting app. We have an User type that has the `ID` and `Budget` properties. We may store our user information
  in Postgres while we may store the budget on a third party service PayPal (for whatever reason). In this case we wrap our Paypal
  dependency with a domain interface, called BudgetService. Now those services are decoupled and communicate solely through domain lanaguage.
```go
// ./internal/domain/user.go
package domain

type User struct {
	ID,
	Budget int64
	
	Name string
}

type UserService interface {
	GetUser(userID int64) User
}

// ./internal/domain/budget.go
package domain

type BudgetService interface {
	GetBudget(userID int64) (budget int64)
}

// ./internal/domain/psql/userservice.go
package psql

type UserService struct {
	db *sql.DB
	budgetService domain.BudgetService
}

func (s *UserService) GetUser(userID int64) domain.User {
	var user domain.User
	
	s.db.QueryRow('SELECT name...').Scan(&user) // get the user from db
	
	// get the budget from the 3rd party depenendcy
	user.Budget = s.budgetService.GetBudget(user.ID)
	
    return user
}
```

What does this achieve?
Packages are uniform, we know what they offer and are isolated, at most depending on the domain.

#### Examples of this design:
- Go Standard Library:
  - [io.Reader](https://github.com/golang/go/blob/master/src/io/io.go) is a domain type for reading types with implementations
    grouped by dependencies: [gzip.Reader](https://github.com/golang/go/blob/0e85fd7561de869add933801c531bf25dee9561c/src/compress/gzip/gunzip.go#L74),
    [tar.Reader](https://github.com/golang/go/blob/f1980efb92c011eab71aa61b68ccf58d845d1de7/src/archive/tar/reader.go#L18).
  - `net/http` builds on top of the abstractions of the `net` package.
  
#### Resources and Inspiration
[Standard Package Layout by Ben Johnson](https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1) an article discussing a DDD approach to project structure.

[WTF Dial by Ben Johnson](https://github.com/benbjohnson/wtf) a GitHub repo containing a working example application
that follows the project structure from the previous article

[Moving Towards Domain Driven Design in Go by Jon Calhoun](https://www.calhoun.io/moving-towards-domain-driven-design-in-go/)

[Design Philosophy On Packaging by Bill Kennedy](https://www.ardanlabs.com/blog/2017/02/design-philosophy-on-packaging.html) 

[Chronohraf](https://github.com/influxdata/chronograf) A real web-app following a similar DDD design