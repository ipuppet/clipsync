# clipsync

> Sync clip with http api

```
-p [string] The service listening port.
```

## Build

require github.com/josephspurrier/goversioninfo

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
