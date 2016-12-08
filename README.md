# Liferaybeat

Welcome to Liferaybeat.

Ensure that this folder is at the following location:
`${GOPATH}/github.com/packt`

## Getting Started with Liferaybeat

### Requirements

* [Golang](https://golang.org/dl/) 1.7

### Init Project
To get running with Liferaybeat and also install the
dependencies, run the following command:

```
make setup
```

It will create a clean git history for each major step. Note that you can always rewrite the history if you wish before pushing your changes.

To push Liferaybeat in the git repository, run the following commands:

```
git remote set-url origin https://github.com/packt/liferaybeat
git push origin master
```

For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).

### Build

To build the binary for Liferaybeat run the command below. This will generate a binary
in the same directory with the name liferaybeat.

```
make
```


### Run

To run Liferaybeat with debugging output enabled, run:

```
./liferaybeat -c liferaybeat.yml -e -d "*"
```


### Test

To test Liferaybeat, run the following command:

```
make testsuite
```

alternatively:
```
make unit-tests
make system-tests
make integration-tests
make coverage-report
```

The test coverage is reported in the folder `./build/coverage/`

### Update

Each beat has a template for the mapping in elasticsearch and a documentation for the fields
which is automatically generated based on `etc/fields.yml`.
To generate etc/liferaybeat.template.json and etc/liferaybeat.asciidoc

```
make update
```


### Cleanup

To clean  Liferaybeat source code, run the following commands:

```
make fmt
make simplify
```

To clean up the build directory and generated artifacts, run:

```
make clean
```


### Clone

To clone Liferaybeat from the git repository, run the following commands:

```
mkdir -p ${GOPATH}/github.com/packt
cd ${GOPATH}/github.com/packt
git clone https://github.com/packt/liferaybeat
```


For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).


## Packaging

The beat frameworks provides tools to crosscompile and package your beat for different platforms. This requires [docker](https://www.docker.com/) and vendoring as described above. To build packages of your beat, run the following command:

```
make package
```

This will fetch and create all images required for the build process. The hole process to finish can take several minutes.
