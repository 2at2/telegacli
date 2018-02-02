[![Build Status](https://travis-ci.org/2at2/telegacli.svg?branch=master)](https://travis-ci.org/2at2/telegacli)

# Telega CLI
Console telegram sender.

## Build
```
make deps
make ok
make build
```

## Send message

```
bin/telegacli send <chatId> <message> --token <botId>:<token>
```
Or

```
TGTOKEN=<botId>:<token> bin/telegacli send <chatId> <message>
```

## Docker

Pull image
```
docker pull strebul/telegacli:latest
```

Run
```
docker run strebul/telegacli:latest send <chatId> <message> --token <botId>:<token>
```