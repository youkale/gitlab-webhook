#gitlab web hooks server

### config.json
```
{
  "address": "0.0.0.0:8964",
  "path": "/webhook",
  "secret_token": "2575079e53e0605b24b1bd8df2e2f757",
  "settings": [
    {
      "event": "Job Hook",
      "project_name": "hyd-admin",
      "build_name": "qm-build",
      "build_stage": "build",
      "build_status": "success",
      "ref": "-test",
      "command": [
        "ls -al"
      ]
    }
  ]
}

```

### Compile && execute
```
go build
./gitlab-webhook
```

### LICENSE
 MIT