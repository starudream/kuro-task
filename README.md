<h1>TO: https://github.com/starudream/sign-task</h1>

---

# Kuro-Task

![golang](https://img.shields.io/github/actions/workflow/status/starudream/kuro-task/golang.yml?style=for-the-badge&logo=github&label=golang)
![release](https://img.shields.io/github/v/release/starudream/kuro-task?style=for-the-badge)
![license](https://img.shields.io/github/license/starudream/kuro-task?style=for-the-badge)

## Config

- `global` [doc](https://github.com/starudream/go-lib/blob/v2/README.md) - [example](https://github.com/starudream/go-lib/blob/v2/app.example.yaml)

以下参数暂时需要电脑登录 [库街区](https://www.kurobbs.com/mc/home/9) 手动获取，再通过下方 [Account](#account) 新增

```yaml
accounts:
  - phone: "手机号码，作为唯一标识使用"
    dev_code: "设备识别码，LocalStorage 中的 dc"
    token: "授权令牌，LocalStorage 中的 auth_token"
    source: "来源，目前有 android 和 h5，可空默认 h5"
    version: "版本，LocalStorage 中的 av，可空默认 2.2.0"
cron:
  spec: "签到奖励执行时间，默认 5 4 8 * * * 即每天 08:04:05"
  startup: "是否启动时执行一次，默认 false"
```

## Usage

```
> kuro-task -h
Usage:
  kuro-task [command]

Available Commands:
  account     Manage accounts
  config      Manage config
  cron        Run as cron job
  notify      Manage notify
  service     Manage service
  sign        Run sign task

Flags:
  -c, --config string   path to config file
  -h, --help            help for kuro-task
  -v, --version         version for kuro-task

Use "kuro-task [command] --help" for more information about a command.
```

### Account

下载项目中 [geetest.html](./geetest.html) 文件，本地打开进行验证后，复制极验结果，并使用 `login` 命令。

- [GeeTest](https://github.com/geetestweb/gt4-public-client-demo/blob/master/slide-popup.html)

```shell
# list accounts
kuro-task account list
# add account from h5 website token
kuro-task account add <account phone>
# add account by phone verification code
kuro-task account login <account phone>
```

### SignGame `库街区游戏签到`

```shell
kuro-task sign game <account phone>
```

### Cron

```shell
kuro-task cron
```

### Service

```shell
# register as system service
kuro-task service --user --config kuro-task.yaml install
kuro-task service start
kuro-task service status
```

## Docker

```shell
mkdir kuro && touch kuro/app.yaml
docker run -it --rm -v $(pwd)/kuro:/kuro -e DEBUG=true starudream/kuro-task /kuro-task -c /kuro/app.yaml account add <account phone>
docker run -it --rm -v $(pwd)/kuro:/kuro -e DEBUG=true starudream/kuro-task /kuro-task -c /kuro/app.yaml sign game <account phone>
```

## Docker Compose

```yaml
version: "3"
services:
  kuro:
    image: starudream/kuro-task
    container_name: kuro
    restart: always
    command: /kuro-task -c /kuro/app.yaml cron
    volumes:
      - "./kuro/:/kuro"
    environment:
      DEBUG: "true"
      app.log.console.level: "info"
      app.log.file.enabled: "true"
      app.log.file.level: "debug"
      app.log.file.filename: "/kuro/app.log"
      app.cron.spec: "26 7 8 * * *"
```

## [License](./LICENSE)
