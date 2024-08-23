# Online Payment Platform POC

## Purpose

The project shows usage of Gin router to build a simple payment gateway API.

## Usage

The help target is prodided so that all targets imported are documented.

### Main targets

The make file provides all building functionalities. With default empty target
an executable is is create in ./bin dir. It can be locally started with make run.

### Docker image & compose

A docker image can be created with target make docker/image. The docker
compose can be started locally with make docker/compose/up.

### Swagger testing

The swagger file may be built with target make swagger/generate. The swagger
can be started with make swagger/serve so that IPI may be locally tested
assuming it was locally started with make run target.

### Minikube deployment

Assuming that the minikube is locally installed a deployment may be created
with target make docker/deploy.

### Unit tests

The unit tests can be started with target make test/unit.

### Run tests

The run tests are shell scripts accessing main API points. Assuming that
the local executable is started with make run the run tets may be started with
make test/run.

