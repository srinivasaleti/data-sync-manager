A data synchronization scheduler that connects to Amazon S3, pulls data periodically,
generates local files per sync job, and handles failures by resuming from the last successful checkpoint.

Tech:
- Go lang
- Make

Commands
--------
**Help**: `make help`

**Run tests**: `make test`

**Build app**: `make build`

**Run app**:  `make run CONFIG=<input file location>`

- Note you need provide CONFIG as an yaml file. You can refer config.example.yaml