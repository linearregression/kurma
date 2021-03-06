[![License][License-Image]][License-URL] [![Build][Build-Status-Image]][Build-Status-URL] [![Release][Release-Image]][Release-URL]

# Kurma

Kurma is a container runtime built with extensibility and flexibility in
mind. It focuses on "everything is a container" and uses this to enable plugins
that run within Kurma, leaving Kurma easy and simple to deploy and
manage. Configuring networking plugins or customizing how containers are
instrumented is easily extensible.

Kurma implements the [App Container (appc)](https://github.com/appc/spec)
specification, and leverages
[libcontainer](https://github.com/opencontainers/runc/tree/master/libcontainer)
from the [Open Container Initiative (OCI)](https://www.opencontainers.org/).

For more information about Kurma, please visit our [website](https://kurma.io)
and see our [documentation](https://kurma.io/documentation).

### Building Kurma

Building Kurma is made possible by leveraging Docker. You may say "your
container engine is built with another container?" Yep. Docker has great tooling
for augmenting a developer workflow, such as
[Docker Toolbox](https://www.docker.com/products/docker-toolbox). We leverage
that rather than reinventing.

To develop and build against Kurma, all you will need is the following:

* Docker (if not on Linux, we recommend using [Docker Toolbox](https://www.docker.com/products/docker-toolbox))
* Go 1.6

Most compilation is done through Docker to ensure all compiling is done through
our pre-built build environment images. This allows you to easily build Kurma on
Linux even if you're running OS X. The CLI is still natively compiled, so if
you're on OS X, you get a CLI built for OS X.

Our `Makefile` will automatically ensure you're running within Docker and map in
the code at the necessary location, set the GOPATH, and ensure users match.

For a list of some of the most common tasks, you can run `make help` to see a
list and description of some of the tasks.

For local development, you should check out the `kurma` repository within your
`$GOPATH` so that it is at `$GOPATH/src/github.com/kurma`. You'll typically
need to run `make download` to fetch some of the pre-compiled assets (busybox
image, CNI networking image), then run `make local` to compile local binaries
and `make run` to run the daemon mode. If on OS X, this will start the daemon in
Docker and print out how to connect. To shut down the server, simply press
Control-C and it will tear down all pods and exit.

```shell
$ git clone git@github.com:apcera/kurma.git $GOPATH/src/github.com/apcera/kurma
$ cd $GOPATH/src/github.com/apcera/kurma
$ make download
...
$ make local
...
$ make run
Running kurmad
==============================================================
Kurma remote API will be available at 172.17.0.2:12312

To connect with kurma-cli, please run:
  export KURMA_HOST=172.17.0.2
==============================================================
...
```

### Downloading Kurma

The latest release images can be found [on our website](https://kurma.io/download).

### Running Tests

Kurma has both unit tests and integration tests. The unit tests are implemented
in Go and can be ran using normal Go tooling.

The easiest way to run full unit tests would be to run `make test`. This will
run the tests within a Docker image so that it can test on Linux and in the
proper build environment.

The integration tests focus on testing Kurma from the outside in. They are
implemented in Ruby using the well known [RSpec](http://rspec.info/)
framework. The integration tests are focused on testing the daemon functionality
via the API and CLI. It focuses on verifying the end-to-end functionality works
as expected. Integration tests are ran on every pull request and master change,
and test compatibility across each distribution that is officially supported.

The integration tests live under `testing/integration` directory. They should
generally work on any recent version of Ruby, Rubygems, and Bundler. Our own
testing process involves using Ruby 2.2.4, Rubygems 2.6.1, and Bundler 1.11.2.

To run the integration tests, can run the following:

```shell
$ cd testing/integration
$ bundle install
$ bundle exec rake spec
```

The integration tests handle launching the currently compiled Kurma binaries
from the `bin` directory. It is necessary to run the tests from a Linux host, as
it does actively launch the daemon rather than connect to an outside one.



[License-URL]: https://opensource.org/licenses/Apache-2.0
[License-Image]: https://img.shields.io/:license-apache-blue.svg
[Build-Status-URL]: http://ci.kurma.io
[Build-Status-Image]: https://img.shields.io/teamcity/http/ci.kurma.io/s/Kurma_UnitTests_2.svg
[Release-URL]: https://kurma.io/download
[Release-Image]: https://img.shields.io/badge/release-v0.4.0-1eb0fc.svg
