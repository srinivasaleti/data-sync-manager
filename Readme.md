A data synchronization scheduler that connects to Amazon S3, pulls data periodically,
generates local files per sync job, and handles failures by resuming from the last successful checkpoint.

TechStack:
-----
- Go lang
- Make
- Docker

# Running the Application

## Note:
- To run application in local you should install `go`
- To work with docker you should install `docker`
- The configuration for application is driven through `yaml` files
- Copy `config.example.yaml` into `config.yaml` and provide aws configuration
- At the moment code supports only `s3` as source connector, `filesystem` as target connector. So prefer not changing the connector types.
- If you are running the app using docker please don't change `outdir`. The files synced to `/app/out` in docker a volume attached in `docker-compose`. If you want change `outdir` you need update in `docker-compose.yaml` too.

    ```
    volumes:
        - ../../out:/app/out
    ```

## Commands To Run In Local
**Help**: `make help`

**Run tests**: `make test`

**Run test coverage**: `make test-coverage`

**Build app**: `make build`

**Run app**:  `make run CONFIG=<input file location>`

- Note you need provide CONFIG as an yaml file. You can refer config.example.yaml

## To Execute app in docker
**Run:** `make docker-run` to run the app inside docker container. It make sure running tests, also it downloads files from given source and sync files to target.