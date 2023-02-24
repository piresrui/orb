# Virtual Orb

Virtual orb is a library to simulate an Orb, periodically pinging a manager API.

## Configuration

You can set the required environment variables to run the orb. There are default values and example values for a mocked API server in the `conf/config.yaml` file.

```yaml
API_HOST="mock-api"
API_PORT=1080
ASSET_DIR="./assets"
SIGNUP_PATH="/signup"
REPORT_PATH="/status"
REPORT_PERIOD=10
SIGNUP_PERIOD=10
```

## Building and Running with Mocked API

To run locally in a docker environment you can run the following commands. This will deploy a mocked API server and the orb. Follow the orb logs to see the requests being made.

```bash
make build-orb

make run-with-mock

docker compose logs -f orb
```

## Building and Running against any API

Set up the desired configuration and pass it to the docker image as environment file.

```bash
make build-orb

docker run -it --env-file=./conf/config.yaml --rm orb
```

## Notes

For the sake of velocity, some design decisions were done with a PoC mindset. i.e Proper logging isn't implemented, error handling is basic, etc...

In a complete, "production ready" solution, I'd improve the following:

1. Setup a proper logger, with defined levels (Warn, Debug, Error)
2. Use a dependency injection lib to facilitate mock generation and reduce boilerplate code (e.g wire)
3. Define error types and their propagation levels (e.g some errors should be user visible, other should show stacktraces)
4. Add a smarter periodic system, with randomized periods.