# Celo Alert Bot

**Celo Alert Bot** will send an alert about a new proposal in the network.

## Install Prerequisites
- **Go 13.x+**

### We need to install influxDB next

```sh
$ cd $HOME
$ wget -qO- https://repos.influxdata.com/influxdb.key | sudo apt-key add -
$ source /etc/lsb-release
$ echo "deb https://repos.influxdata.com/${DISTRIB_ID,,} ${DISTRIB_CODENAME} stable" | sudo tee /etc/apt/sources.list.d/influxdb.list
```

### Start influxDB

```sh
$ sudo -S apt-get update && sudo apt-get install influxdb
$ sudo -S service influxdb start

```

### Create database:

```sh
$ influx -precision rfc3339
> CREATE DATABASE celo

> SHOW DATABASES
```

## Install and configure the Celo Alert Bot

### Get the code and run it after editing the config file

```sh
go get github.com/chainflow/celo
cd go/src/github.com/chainflow/celo

go get .

```

### Edit the config:

```sh
mv example.config.toml config.toml
nano config.toml

update tg_chat_id, tg_bot_token (Give your telegram chat id and bot token, to which you want to get the alerts)

```
### Start the bot using screen

```sh
Screen -S celobot
go run main.go

To detach the screen (Ctrl+A and then D)
To attach the screen from next time : screen -rd celobot

```