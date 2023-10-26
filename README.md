# env-from-mr

`env-from-mr`是一个命令行工具，用于从GitLab Merge Request的描述中提取`env`代码块并其打印为导出环境变量的格式。可以在Merge Request中为Pipeline脚本动态地提供环境变量。

## 安装

你可以通过以下方式获取`env-from-mr`：

```shell
go get github.com/raojinlin/env-from-mr
```

## 使用
env-from-mr 接受以下参数：

- `mr-iid`：Merge Request的内部ID（如果没有指定，取CI_MERGE_REQUEST_IID环境变量）。
- `pid`：GitLab项目的ID（必需）。
- `token`：GitLab访问令牌（必需）。
- `url`：GitLab的地址（必需）。

示例:

打印描述中的环境变量
```bash
$ env-from-mr -mr-iid 123 -pid 456 -token yourtoken -url yourgitlaburl
export ENV_A='a'
export SVC_A_HOST='localhost'
```

将描述中的环境变量导出
```bash
$ eval $(env-from-mr -mr-iid 123 -pid 456 -token yourtoken -url yourgitlaburl) 
$ echo $ENV_A
a
$ echo $SVC_A_HOST
localhost
```


在Pipeline中使用
```yaml
pages-job:
  stage: deploy
  script:
    - eval $(env-from-mr -mr-iid $CI_MERGE_REQUEST_IID -pid $CI_PROJECT_ID -url $CI_SERVER_URL -token yourtoken)
    - echo $ENV_A
    - echo $SVC_A_HOST
```


## 工作流程

1. 获取指定Merge Request的描述。
2. 在描述中查找env代码块并提取其中的环境变量内容。
3. 使用eval命令导出提取的环境变量，可以在后续Pipeline步骤中使用。


## 注意事项

- 请确保你的GitLab访问令牌具有足够的权限来访问项目和Merge Request信息。
- 该工具需要连接到GitLab的API，因此需要网络连接。