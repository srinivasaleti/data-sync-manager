A data synchronization scheduler that connects to Amazon S3, pulls data periodically,
generates local files per sync job, and handles failures by resuming from the last successful checkpoint.

Tech:
- Go lang
- Make

Commands To Run In Local
-------------------------
**Help**: `make help`

**Run tests**: `make test`

**Build app**: `make build`

**Run app**:  `make run CONFIG=<input file location>`

- Note you need provide CONFIG as an yaml file. You can refer config.example.yaml

To Execute app in docker
-------------------------
- Copy `config.example.yaml` into `config.yaml` and provide aws configuration
- Note: At the moment supported types : `s3` as `source.type`, `filesystem` as `target.type`


    **Run:** `make docker-run` to run the app inside docker container. It make sure running tests, also it downloads files from given source and sync files to target.