# Contacts Go
A Contacts RestFul API application using Golang Gin Framework.

## Execute on Development
1. Install the dependencies
```bash
$ go mod download
```

2. Rename ```.env.template``` file to ```.env``` and fill the fields

3. Use ```docker-compose``` to create postgres database instance
```bash
$ docker-compose up -d
```

4. Enter to the database instance and install the ```uuid-ossp``` extension
```postgresql
CREATE EXTENSION "uuid-ossp";
```

5. Start the application
```bash
$ go run .
```