# clipsync

> Sync clip with http api

```bash
-port [string] The service listening port.
-addr [string] The service listening address.
-conf [string] The config file path.
```

## Build

require github.com/josephspurrier/goversioninfo

`go install github.com/josephspurrier/goversioninfo/cmd/goversioninfo@latest`

`make build_win`

## API

### file

- POST `/api/clip`

    Set clipboard.

    Body:

    ```json
    {
        "data": "Hello world!"
    }
    ```

    Response:

    ```json
    {
        "status": true
    }
    ```

- GET `/api/clip`

    Get clipboard.

    Response:

    ```json
    {
        "status": true,
        "data": "Hello world!"
    }
    ```
