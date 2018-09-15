# Caldera

A command line utility which provide a code generated boilerplate for the services inside of your container. It should helps to save 2 and more days of developers routine work whom decided to create them first (micro) service.

## Features of Caldera

- Using of configuration file that contains your saved preferences for a new boilerplate service
- Interactive mode to select preferred features for a new service
- Using of CLI flags to create new service quickly

### Features of boilerplate service

- Implementation of the health checks
- Configuring the service using config file, environment variables or flags
- Processing of graceful shutdown for every registered component
- Database interfaces with migration features
- CI/CD pipelines integrated into Makefile
- Helm charts for deploying the service in Kubernetes environment
- Container SSL certificates integration for using a secure client
- Integration of the package manager
- Versioning automation

## Requirements

- [Go compiler](https://golang.org/dl/) v1.9 or newer
- [GNU make utility](https://en.wikipedia.org/wiki/Make_(software)) that probably already installed on your system

### Requirements for boilerplate service

- Docker service, version 18.03 or newer

## Setup

```sh
go get -u github.com/takama/caldera
cd $GOPATH/src/github.com/takama/caldera
make
```

## Usage of Caldera

### Interactive mode

In this mode, you'll be asked about the general properties associated with the new service. The configuration file will be used for all other data, such as the host, port, etc., if you have saved it before. Otherwise, the default settings will be used.

```txt
./caldera
Caldera boilerplate version: v0.0.1 build date: 2018-09-15T12:02:17+07

Provide name for your Github account: my-account
Do you want to deploy your service to the Google Kubernetes Engine? (y/n): y
Provide ID of your project on the GCP (my-project-id):
Provide compute zone of your project on the GCP (europe-west1-b):
Provide cluster name in the GKE (my-cluster-name):
Provide name for your service (service): my-service
Provide description for your service (my-service description):
Do you need API for the service? (y/n): y
What kind of API do you need? (rest,grpc): grpc
Do you need one more API for the service? (y/n): y
What kind of API do you need? (rest): rest
Do you need gRPC client for the service? (y/n): y
Do you need storage driver? (y/n): y
What kind of storage driver do you need? (postgres,mysql): postgres
Templates directory (~/go/src/github.com/takama/caldera/.templates):
New service directory (~/go/src/github.com/my-account/my-service):
```

### CLI mode

In this mode, you'll be not asked about everything. The configuration file will be used for all other data, such as the host, port, etc., if you have saved it before. Otherwise, the default settings will be used.

```sh
./caldera new [ --service <name> --description <description> --github <account> --grpc-client ]
```

### Save configuration for future use

For example of save a `storage` parameters in Caldera configuration file:

```sh
./caldera storage [flags]

Flags:
  -h, --help       help for storage
      --enabled    A Storage modules using
      --mysql      A mysql module using
      --postgres   A postgres module using
```

Save a `storage` parameters for database driver in Caldera configuration file:

```sh
./caldera storage driver [flags]

Flags:
  -h, --help              help for driver
      --host string       A host name (default "postgres")
      --port int          A port number (default 5432)
      --name string       A database name (default "postgres")
  -u, --username string   A name of database user (default "postgres")
  -p, --password string   An user password (default "postgres")
      --max-conn int      Maximum available connections (default 10)
      --idle-conn int     Count of idle connections (default 1)
```

Save an API parameters for `REST/gRPC` (REST always used gRCP gateway):

```sh
./caldera api [flags]

Flags:
  -h, --help      help for api
      --enabled   An API modules using
      --grpc      A gRPC module using
      --rest      A REST module using
```

Save a common API parameters:

```sh
./caldera api protocol [flags]

Flags:
  -h, --help       help for protocol
      --port int   A service port number (default 8000)
```

## Health checks

Service should have [health checks](https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/) for successful execution in containers environment. It should helps with correct orchestration of the service.

## Configuring

The [twelve factors](https://12factor.net/config) service must be configured using environment variables. The service has a built-in library for automatically recognizing and allocating environment variables that are stored inside `struct` of different types. As additional methods, a configuration file and flags are used. All these methods of setting are directly linked to each other in using of configuration variables.

## System signals

The service has an ability to intercept system signals and transfer actions to special methods for graceful shutdown, maintenance mode, reload of configuration, etc.

```go
type Signals struct {
    shutdown    []os.Signal
    reload      []os.Signal
    maintenance []os.Signal
}
```

## Build automation

In the CI/CD pipeline, there is a series of commands for the static cross-compilation of the service for the specified OS. Build a docker image and push it into the container registry. Optimal and compact `docker` image `FROM SCRATCH`.

```Dockerfile
FROM scratch

ENV MY_SERVICE_SERVER_PORT 8000
ENV MY_SERVICE_INFO_PORT 8080
ENV MY_SERVICE_LOGGER_LEVEL 0

EXPOSE $MY_SERVICE_SERVER_PORT
EXPOSE $MY_SERVICE_INFO_PORT

COPY certs /etc/ssl/certs/
COPY migrations /migrations/
COPY bin/linux-amd64/service /

CMD ["/service", "serve"]
```

## SSL support

Certificates support for creating a secure SSL connection in the `Go` client. Attaching the certificate to the docker image.

## Testing

The command `make test` is running set of checks and tests:

- tool `go fmt` used on package sources
- set of linters used on package sources (20+ types of linters)
- tests used on package sources excluding vendor
- a testing coverage of new boilerplate service
- compile and check of Helm charts

## Helm charts and Continuous Delivery

A set of basic templates for the deployment of the service in Kubernetes has been prepared. Only one `make deploy` command loads the service into Kubernetes. Wait for the successful result, and the service will be ready to go.

## Package manager

To properly work with dependencies, we need to select a package manager. [dep](https://github.com/golang/dep) is one of the popular dependency management tools for Go.

## Versioning automation

Using a special script to increment the release version

```sh
./bumper.sh
Current version 0.0.1.
Please enter new version [0.0.2]:
```

## Contributing to the project

See the [contribution guidelines](docs/CONTRIBUTING.md) for information on how to participate in the Caldera project to submitting a pull request or creating a new issue.

## Versioned changes

All changes in the project described in [changelog](docs/CHANGELOG.md)

## License

[MIT Public License](https://github.com/takama/caldera/blob/master/LICENSE)
