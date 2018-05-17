#gitlab web hooks server

### config.json
```
{
  "address":"0.0.0.0:8964",  //server and port
  "path":"/webhook",   // http listen path
  "repositories":[
    {
      "name":"mofa-core", //repository name
      "event":"Pipeline Hook", // gitlab webhook event
      "command":[
        "ls -al"   // exec cmd
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