# cronmon

## Overview

`cronmon` is a simple utility designed to run scheduled jobs and provide desktop notifications upon completion.

Background: i love the cronjob syntax and hate systemd timers but i want to know when cronjobs fail.

## Usage

```
NAME:
   cronmon - monitors a job and notifies according to your wishes 🔔

USAGE:
   cronmon [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --command value, -c value
   --cron-name value, --cn value
   --notify-success, --success    (default: notifies also when the command was successfully executed)
   --help, -h                     show help
```

## Crontab example

```cron
0 * * * * cronmon -c "calendarsync -config ~/.config/sync.yaml" -success -cn "CalendarSync"
```

## License

MIT
