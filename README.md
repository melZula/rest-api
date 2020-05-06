# rest-api
* Go v1.13.8
* PostgreSQL v12
* Router - [gorilla/mux](https://github.com/gorilla/mux)
* Handlers - [gorilla/handlers](https://github.com/gorilla/handlers)
* Sessions, cookie - [gorilla/sessions](https://github.com/gorilla/sessions)
* Logger - [logrus](https://github.com/sirupsen/logrus)
* Validation - [ozzo-validation](https://github.com/go-ozzo/ozzo-validation)
* *.toml* parser - [BurntSushi/toml](https://github.com/BurntSushi/toml)

## Usage
```make``` build app<br/>
```./apiserver``` run

## Middleware
* authenticate user for private subrouter
* CORS headers
* logger with measuring time of each request
