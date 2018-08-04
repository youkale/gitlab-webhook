gitlab webhook
-----

说明： 能跟gitlab-ci很好地集成， 根据事件触发hook

配置: config.json
```
{
  "address": "0.0.0.0:8964",  //监听地址
  "path": "/webhook",   //监听路径
  "secret_token": "2575079e53e0605b24b1bd8df2e2f757",  //gitlab 请求头token
  "settings": [
    {
      "event": "Job Hook",   //事件名称
      "project_name": "hyd-admin",   //工程名
      "build_name": "qm-build",  //构建名
      "build_stage": "build",  //构建策略
      "build_status": "success",  //构建状态
      "ref": "-test",   //tag后缀
      "command": [
        "ls -al"   //执行命令，可以多条
      ]
    }
  ]
}

```

编译执行

```
go build
./gitlab-webhook
```

LICENSE

MIT
