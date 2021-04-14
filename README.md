# API LOGIN GO

simple api to login using jwt and implemented using golang & manggoDB

## Clone

clone this project to your local.

```bash
clone https://github.com/martinyonathann/api-login-go.git
```

## Setup ManggoDB

Open config/db/db.go then change : 

```bash
const (
	password = "your password"
	dbname   = "db name"
)
```


## Run

running the project

```bash
go run main.go
```

## Usage

#### Registrasi

```
POST /registrasi HTTP/1.1
Host: localhost:8080
Content-Type: application/json
Content-Length: 107

{
    "username":"username",
    "firstname":"firstname",
    "lastname":"lastname",
    "password":"pwd"
}
```

#### Login

```
POST /login HTTP/1.1
Host: localhost:8080
Content-Type: application/json
Content-Length: 58

{
    "username":"username",
    "password":"pwd"
}
```

#### get Profil

```
GET /profile HTTP/1.1
Host: localhost:8080
Authorization: <<response Token from Registration>>
```