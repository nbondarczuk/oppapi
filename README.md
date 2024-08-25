# Online Payment Platform POC

## Purpose

The project shows usage of Gin router and Mongo DB to build a simple payment gateway API.

## API description

The following main API points are available:

### /payment

#### POST

It performs the payment if stated with POST request.

#### GET

It does payment search if started with GET and a payment id in the path parameter.

### /refund

#### POST

It performs the refund if stated with POST request and a payment id in the path paramter.

#### GET

It does payment search if started with GET and a payment id. The payment type must
be REFUND. It is validatedd.

### /transaction

#### POST

It is  mock of the bank of the merchant. IT turns payments or refunds into trnsactions
so to say. It clears them.

For the sake of simplicity the ID of the transaction is the same as the underlaying payment.
The transactions are not stored in the collection. They are returned after POST request
is successful.

### /health

#### GET

It is a mandatory target for docker compose and K8S deployments. It does very little
therefor it is super fast.

## Config

The config/config.yaml file is used a the main configuration source. The config
values are stored in it. They can be changes to adjust the runtime conditions
to the local environment. The config values may be overriden by env
variables.

## Authorization

The X-KEY-API key is used to authenticate the client. It is stored in config file
under path auth/x_api_key. The API entry points like /payment, /refund
are to be used with authentication. Missing or wrong auth key provided in the header
of the request causes abortion of the request processing with HTTP code 401 - Unauthorized.

The X-API-KEY of the bank is sored on the mocked config values of the merchant.

## Usage

The help target is provided so that all targets imported are documented.

### Main targets

The make file provides all building functionalities. With default empty target
an executable is is create in ./bin dir. It can be locally started with make run as well.
It is a prerequisite for running run tests.

Another useful target is make clean. IT is really recommended to run it before
adding files to git repo.

### Docker image & compose

A docker image can be created with target make docker/image. The docker
compose can be started locally with make docker/compose/up.

### Swagger documentation

The swagger file may be built with target make swagger/generate. The swagger
can be started with make swagger/serve so that IPI may be locally tested
assuming it was locally started with make run target.

The swagger doc is the main reference for the API documentation. The swagger
format is used in the dunction headers

The swagger-go project is used. There is no need to install the exec as the docker
is used to perform related swagger targets.

### Minikube deployment

Assuming that the minikube is locally installed a deployment may be created
with target make docker/deploy.

### K8S deployment

A regular K8S cluseter running in docker may be used to do deployment. It can be done
with make target docker/deploy. A set of yamls is provided to do that.

### Unit tests

The unit tests can be started with target make test/unit. They test repository using docker
image of the mongo so that the whole chain of request processing can be simulated.

### Run tests

The run tests are shell scripts accessing main API points. Assuming that
the local executable is started with make run the run tets may be started with
make test/run.

## Required packages/programs

The following additional packages/programs shall be installed in the environment:

- GNU make
- go compiler version 1.22.3
- mongodb
- curl
- docker
- docker compose
- minikube
- kubectl
- govet
- golint
- golangci-lint
- jq
- ab

## Considerations

### Execution instructions

1. Do git clone of the repository.
2. Make sure the mongo db is installed and started. It shall run on: mongodb://localhost:27017.
3. Run make in the oppabi main directory.
4. Run make run in the same location on a separate terminal session.
5. Start the unit tests with make test/unit.
6. Try scripts in test/run to create a payment, read it, create refund of it, read it.
7. Check results in mongo db.

### Dependencies

It shall work in any MAC-OS-X platform or Linux assuming the required packages or at least some
of the m are installed. For the beginning make, mongo, jq, curl and docker are necessary,
not speaking of Golang compiler. It is essential.

### Assumptions

- The bank is just a internal mock. It may be switched off for the test run with config flag.
- The bank interface is not well secured. Probably something like Open Banking API may be used.

## Areas for improvement

* Better validation of the payload in the payment creation.
* Increase unit test quality and coverage
* Use go ver and other schecks in the building
* Add monitoring nd tracing
* Add JWT authentication and authorization layer
* Improve go-swagger comments so that run tests can be done with swagger page
* Add testing docker image to do integration tests
* Use ab to make performance testin in docker producing some artefacts like graphs
* Add of datastore caching with redis
* Consider migrating from hadlers into the model of controller.

## Cloud technologies

It is possible to run it in minikube assuming it is installed. It will work in K8S
as the yamls are provided. Docker Desktop has a simple implementation of K8S. With this
approach it can be easily deployed on any modern cloud platform: Azure, AWS, Google.
The only blocking point is the costs as K8S is not for free.

## Extra

### Authentication and security

It is base on X-API-KEY. A temporary random key is stored in the config.yaml and the same
one is used in the run scripts.

### Audit trail

It is implemented withe on nscreen logging on INFO and/or DEBUG levels.
