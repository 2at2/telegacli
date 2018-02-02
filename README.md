[![Build Status](https://travis-ci.org/2at2/telegacli.svg?branch=master)](https://travis-ci.org/2at2/telegacli)

# Telega CLI
Console telegram sender.

## Send message

```
telegacli send <chatId> <message> --token <botId>:<token>
```
Or

```
TGTOKEN=<botId>:<token> telegacli send <chatId> <message>
```