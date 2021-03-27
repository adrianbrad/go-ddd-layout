// Package psql contains implementations for PubSub and UserService interfaces
// backed by PostgreSQL.
//
// The methods of the services implemented are expected to implement a transactional boundary
// in order to not leak the transaction object outside this package.
package psql
