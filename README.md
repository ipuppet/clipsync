# clipsync

default account:
```json
{
    "account": "admin@erots.com",
    "password": "admin"
}
```

```
--config [string] Set config file path.
--log    [string] Set log file path.
```


## API

### file

- POST `/api/file/image/:module`

    使用 FORM 上传图片，图片文件的键为 `image`。

    成功时返回 `json` 对象包含 `path` 字段，使用域名 + `path` 即可访问该图片。

- DELETE `/api/file/image/:module/:date/:name`

    删除图片，url 参数中的 `:date` 和 `:name` 可从上传时返回的链接中取。

- POST `/api/file/app/:name`

    上传 App 文件，成功时和上传图片一样，返回一个访问路径。
