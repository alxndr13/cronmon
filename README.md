# cronmon

## Overview

`cronmon` is a simple utility designed to run scheduled jobs and provide desktop notifications upon completion.

Background: i love the cronjob syntax and hate systemd timers but i want to know when cronjobs fail.

## Crontab example

```cron
0 * * * * cronmon -c "calendarsync -config ~/.config/sync.yaml" -success -cn "CalendarSync"
```

## License

MIT
