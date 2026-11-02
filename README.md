# Go Learning Projects

A step-by-step collection of Go examples, starting with basic syntax and building toward concurrency, HTTP services, layered APIs, and authentication. The numbered folders are intended to be explored in order, but each example can be opened and run independently.

## Repository Layout

Most numbered folders are separate Go modules and have their own `go.mod` file. Folder `33filestructure` contains a larger project in `course-api`, and its Go module is nested there. Run Go commands from the module directory shown in each example below; commands run from the repository root will not automatically select a lesson's module.

The repository includes short command-line examples, local HTTP servers, a PostgreSQL-backed API, and an in-memory authentication API. Some examples are demonstrations of a concept rather than production-ready implementations. Notes below call out examples that need user input, network access, another local service, or extra caution.

## Requirements

- Go version `1.27.1`, as declared by the modules. A newer compatible Go release should also work.
- Internet access the first time a module needs to download dependencies. The Gorilla Mux package is vendored in `28mymodules`; other modules declare their dependencies in `go.mod` and `go.sum`.
- PostgreSQL only for the database-backed project in `33filestructure/course-api`.
- A compatible local HTTP server on port `3000` for `26webrequests2`.

To download dependencies for a module, change to its directory and run:

```sh
go mod download
```

## Learning Roadmap

### Go Basics

#### `01hello` - Hello World

Prints a greeting to the terminal. This is the smallest example and introduces the `main` package, the `main` function, importing a package, and writing output with `fmt.Println`.

Run with `cd 01hello && go run .`.

#### `02variables` - Variables and Values

Declares values using explicit types and type inference, prints values and their types, and declares a constant. Use it to compare Go's short declaration syntax with `var` declarations and to see that constants are fixed at compile time.

Run with `cd 02variables && go run .`.

#### `03userinput` - Reading Input

Reads a name from standard input with a buffered reader and prints it. This introduces `bufio`, `os.Stdin`, string handling, and a simple interactive command-line flow.

Run with `cd 03userinput && go run .`. The program waits for input in the terminal.

#### `04conversion` - Converting Input

Reads a rating from the user, converts the entered text to an integer, and prints an incremented value. The example connects string input to numeric parsing and shows basic conversion error handling.

Run with `cd 04conversion && go run .`. Enter a value when prompted.

#### `05mytime` - Dates and Times

Prints the current time and a fixed date in UTC using Go time layouts. It demonstrates `time.Now`, `time.Date`, time zones, and Go's reference-time formatting convention.

Run with `cd 05mytime && go run .`. The folder also contains compiled binaries; running the source with `go run .` is the portable option.

#### `06mypointers` - Pointers

Prints a value's address and dereferenced value, then changes the value through a pointer. This demonstrates address-of (`&`), pointer types, and dereferencing (`*`).

Run with `cd 06mypointers && go run .`.

#### `07arrays` - Arrays

Creates zero-valued and initialized arrays and prints their contents and lengths. It shows fixed-size collections and that an array's length is part of its type.

Run with `cd 07arrays && go run .`.

#### `08slices` - Slices

Creates and changes slices, appends values, and removes an item by joining the parts before and after its index. This is a practical introduction to Go's resizable slice view over an underlying array.

Run with `cd 08slices && go run .`.

#### `09maps` - Maps

Creates a map of languages, looks up and adds entries, and iterates over its key/value pairs. It demonstrates map literals, map indexing, and `range` iteration.

Run with `cd 09maps && go run .`. Map iteration order is not guaranteed.

#### `10Structs` - Structs

Defines and constructs a `User` value, then prints the struct and its fields. This introduces custom data types, fields, and keyed struct literals.

Run with `cd 10Structs && go run .`.

#### `11ifelse` - Conditional Logic

Uses conditional branches to report whether a number is even or odd and compare it with a threshold. It demonstrates boolean expressions, `if`/`else if`/`else`, and declarations scoped to a conditional.

Run with `cd 11ifelse && go run .`.

#### `12switchcase` - Switch Statements

Rolls a pseudo-random die and uses a `switch` to choose output for the result. It combines random number generation with multi-branch control flow.

Run with `cd 12switchcase && go run .`.

#### `13loops` - Loops

Shows Go's `for` loop in its common forms, including iterating over values and controlling loop execution. Go uses `for` for traditional counter loops as well as condition-based repetition.

Run with `cd 13loops && go run .`.

#### `14functions` - Functions and Return Values

Calls ordinary and variadic add functions and demonstrates a function that returns more than one value. It introduces parameters, return values, and accepting a variable number of arguments.

Run with `cd 14functions && go run .`.

#### `15methods` - Methods

Defines methods on a `User` struct and calls them. This demonstrates receiver declarations and how methods associate behavior with a type; the example uses value receivers, which operate on a copy of the receiver.

Run with `cd 15methods && go run .`.

#### `16defer` - Deferred Calls

Schedules function calls with `defer` and prints when they execute. Deferred calls run when the surrounding function returns, in last-in-first-out order.

Run with `cd 16defer && go run .`.

#### `17interfaces` - Interfaces

Passes different animal values, including dogs and cats, to a shared function that calls their sound behavior. It demonstrates interface types and Go's implicit interface satisfaction: a type implements an interface by providing its methods.

Run with `cd 17interfaces && go run .`.

### Concurrency and I/O

#### `18goroutines` - Goroutines and Shared State

Starts concurrent HTTP status-code requests for several hardcoded public URLs and waits for the workers to finish. It demonstrates goroutines, a `sync.WaitGroup`, and mutex-protected shared data.

Run with `cd 18goroutines && go run .`. Requires outbound network access.

#### `19mutexandawaitgroups` - Wait Groups and Mutexes

Starts two goroutines that increment a shared counter and waits for both goroutines to finish. It demonstrates `sync.WaitGroup` and is intended to accompany the mutex example, but the active counter increments are not protected by a mutex. This creates a data race; `go run -race .` can help expose it.

Run with `cd 19mutexandawaitgroups && go run .`.

#### `20channels` - Channels

A producer sends the integers 1 through 5 on a channel, closes it, and a consumer ranges over the received values. It demonstrates unbuffered channels, sending and receiving, channel direction in function parameters, and closing a channel after production is complete.

Run with `cd 20channels && go run .`.

#### `21channelsSelect` - Select

Starts two delayed goroutines that send values on separate channels, then uses `select` to receive whichever channel is ready. The surrounding loop continues after both one-time messages have been consumed and blocks waiting for more values, so the program does not exit by itself.

Run with `cd 21channelsSelect && go run .`. Stop it with `Ctrl+C` after observing the receives.

#### `22workerpool` - Worker Pool

Starts three workers that read jobs, square each number, and send results. It demonstrates distributing work through channels, coordinating workers with a wait group, and collecting concurrent results. Result order may vary between runs.

Run with `cd 22workerpool && go run .`.

#### `23files` - File Operations

Creates or truncates `mygofile.txt`, writes data, and includes a helper for reading a file. It demonstrates filesystem APIs, byte counts, and error handling.

Run with `cd 23files && go run .`. The program writes in the current directory and overwrites the example file.

### HTTP, URLs, JSON, and Modules

#### `24webrequests` - Making an HTTP Request

Makes a GET request to `https://example.com` and prints the response. It introduces `net/http`, response handling, and reading a response body.

Run with `cd 24webrequests && go run .`. Requires outbound network access.

#### `25urls` - URL Parsing and Construction

Parses a hardcoded URL, prints URL components and query parameters, and constructs a URL. It demonstrates `net/url`, query parsing, and query encoding. Values in the example are sample data, not live credentials.

Run with `cd 25urls && go run .`.

#### `26webrequests2` - HTTP Methods and Request Bodies

Sends GET, JSON POST, and form POST requests to a local service on port `3000`. It demonstrates building requests, encoding JSON and form data, and reading responses.

Run a compatible local mock server first, then run `cd 26webrequests2 && go run .`. The server must provide `/get`, `/post`, and `/postform` routes.

#### `27bitmorejson` - JSON

Decodes a hardcoded JSON example into Go values. The code also contains an encoder example that is not called by the active `main` function. It demonstrates struct tags, `json.Valid`, unmarshalling, `omitempty`, and ignoring a field with `json:"-"`.

Run with `cd 27bitmorejson && go run .`.

#### `28mymodules` - Modules and Vendoring

Starts a small Gorilla Mux web server with a home route. The folder demonstrates using an external package in a Go module and includes a `vendor` directory containing the dependency source.

Run with `cd 28mymodules && go run .`. The Gorilla Mux dependency is vendored. The folder also contains a compiled binary; `go run .` builds from source.

### Web APIs and Application Structure

#### `29buildapi` - In-Memory Course API

Runs a Gorilla Mux HTTP API on port `4000`, seeds example courses, and provides course CRUD handlers backed by in-memory data. It demonstrates JSON request/response handling, routing, HTTP handlers, and generated IDs. Data is not persisted when the process stops.

Run with `cd 29buildapi && go run .`. The module needs its declared dependencies. A compiled binary is also present in the folder; running the source is more portable.

#### `30Context` - Request Context and Cancellation

The active example starts a `GET /slow` server on port `4000`. Its simulated database operation waits five seconds and can stop early when the request context is cancelled. A separate worker/cancellation example is commented out.

The current module is named `Context`, which conflicts case-insensitively with Go's standard `context` package in this example. As checked in, the project does not build with `go run .`; rename the module to a distinct lowercase name before using it as a runnable example.

#### `31middleware` - HTTP Middleware

Runs an HTTP server on port `4000` with logging, request-ID, panic-recovery, and permissive CORS middleware. It has public course routes, a protected profile route, and admin routes. It demonstrates middleware chaining, subrouters, and passing request-scoped values through context.

Run with `cd 31middleware && go run .`. The example's authentication middleware only checks that an `Authorization` header is present, then assigns a fixed user ID and admin role. It does not validate credentials and must not be treated as real access control.

#### `32httpstatuscode` - Status Codes and API Errors

Runs an in-memory course CRUD API on port `4000`, with sample courses and explicit HTTP status/error responses. It demonstrates routing, JSON handlers, status codes, and CRUD behavior. Like `29buildapi`, its data is not persisted.

Run with `cd 32httpstatuscode && go run .`. The module uses Gorilla Mux and UUID dependencies.

#### `33filestructure` - Layered Course API

This project is under `33filestructure/course-api`. It organizes a course API into handlers, services, repository interfaces, models, and PostgreSQL and in-memory repository implementations. The server exposes course CRUD routes on port `4000`. A SQL migration defines the database schema, while the separate `loader` entry point is used for schema-loading work.

Run the API with `cd 33filestructure/course-api && go run ./cmd/server`. This project requires PostgreSQL on port `5432` and a database configured for the connection details in the source. The current connection code uses a local hardcoded DSN rather than reading the parent `.env` file's `DATABASE_URL`. Its dependencies include GORM, the PostgreSQL driver, Gorilla Mux, UUID, and Atlas tooling. The Atlas development configuration also uses a Docker PostgreSQL database.

#### `34auth-jwt` - JWT Authentication API

Runs an authentication API on port `4000`. It provides `POST /signup`, `POST /signin`, `POST /signout`, and `POST /refresh`; `GET /api/profile` is protected by authentication middleware, and `GET /api/admin/dashboard` additionally requires the `admin` role. It demonstrates password hashing, JWT claims, refresh tokens, cookies, middleware, and role-based authorization.

Run with `cd 34auth-jwt && go run ./cmd/server`. Users and refresh tokens are stored in memory, so they are lost when the process stops. The server seeds `admin@example.com` with password `admin123`. The JWT secret is hardcoded and cookies are configured with `Secure: false`; these are local-learning defaults only and must be replaced before any real deployment.

## Running Examples

For most lessons, enter the folder and run its module:

```sh
cd 01hello
go run .
```

For a project with a nested server entry point, run the package path from the module directory:

```sh
cd 34auth-jwt
go run ./cmd/server
```

Each HTTP server occupies its configured port while running. Several examples use port `4000`, so stop one server before starting another on that port. Interactive examples wait for terminal input; network examples need connectivity or a compatible local service as noted above.

## Notes and Limitations

- These folders are educational examples and are not all production-ready applications.
- `19mutexandawaitgroups` contains an unprotected shared counter and therefore has a data race.
- `21channelsSelect` continues blocking after its one-time channel messages are consumed.
- `23files` overwrites `mygofile.txt` in its working directory.
- `30Context` currently has a module-name conflict that prevents it from building as written.
- `31middleware` uses placeholder authorization logic, not verified authentication.
- `33filestructure` expects a local PostgreSQL setup; inspect the source/configuration before changing database settings.
- `34auth-jwt` includes seeded demo credentials and a hardcoded signing secret. Do not expose it as-is.