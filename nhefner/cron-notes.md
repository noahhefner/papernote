# `Cron`

`cron` is a daemon that runs on all main distributions on Unix and Linux. The purpose of the daemon is to execute commands at a given time. These jobs are commonly referred to as cronjobs. Cronjobs are often used to automate execution of tasks or scripts.

## Creating A New Crontab

The crontab defines the commands to be run and the time at which the command should be run.

To create or edit the crontab:

`crontab -e`

## Other Crontab Commands

- `crontab -e`: Edit or create a crontab file if doesn’t already exist.
- `crontab -l`: To Display the crontab file.
- `crontab -r`: To Remove the crontab file.
- `crontab -v`: To Display the last time you edited your crontab file. (This option is only available on a few systems.)

## Crontab Formatting

```
* * * * * command to be executed
– – – – –
| | | | |
| | | | +—– day of week (0 – 6) (Sunday=0)
| | | +——- month (1 – 12)
| | +——— day of month (1 – 31)
| +———– hour (0 – 23)
+————- min (0 – 59)
```

### Special Characters

| Character | Meaning |
|-----------|---------|
| `*` | For all occurrences of the parameter |
| `,` | Used for creating a list when edclaring two or mor execution times of a command |
| `-` | Specifies a rand of time in which the script can run |
| `/` | Denotes specified intervals of time within a range |

### Special Strings

| String | Meaning |
|--------|---------|
| `@reboot` | Run once, at system startup |
| `@yearly` | Run once every year, “0 0 1 1 *” |
| `@annually` | (same as @yearly) |
| `@monthly` | Run once every month, “0 0 1 * *” |
| `@weekly` | Run once every week, “0 0 * * 0” |
| `@daily` | Run once each day, “0 0 * * *” |
| `@midnight` | (same as @daily) |
| `@hourly` | Run once an hour, “0 * * * *” |

## Examples

Every minute of every day:

` * * * * * /home/user/script.sh`

Every 10 minutes every day:

`*/10 * * * * /home/user/script.sh`

Every 5 minutes of the 6am hour starting at 6:30:

`30-59/5 06 * * * /home/user/script.sh`

Every day at midnight:

`0 0 * * * /home/user/script.sh`

Every day at 3:30am:

`30 03 * * * /home/user/script.sh`

Weekends at 6am:

`0 06 * * 6,7 /home/user/script.sh`

Every hour:

`@hourly /home/usr/script.sh`

## Store command output in a log file

Use the `>` operator to send command output to a log file. Or use `>>` to append output instead of replacing the entire file each run:

Example:

`10 * * * * /user/bin/python /home/usr/script.py > /var/log/cron.log`