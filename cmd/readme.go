// An application may produce multiple binaries
// so we’ll use the Go convention of placing our main package as a subdirectory of the cmd package.
// For example, our project may have a myapp server binary
// but also a myappctl client binary for managing the server from the terminal.
// We’ll layout our main packages like this:
// myapp/
//    cmd/
//        myapp/
//            main.go
//        myappctl/
//            main.go
package cmd
