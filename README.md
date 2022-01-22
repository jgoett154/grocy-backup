# grocy-backup
A simple backup and restore utility for [grocy](https://github.com/grocy/grocy).

## How it Works
The utility uses the [Grocy API](https://demo.grocy.info/api) to make requests to download the raw data or to re-upload the raw data.
All data is stored in a single flat JSON file. The name of the file will contain the date and time thus allowing you to backup the data over time easily.
The utility should be idempotent which will allow you to re-run the program against the same Grocy server multiple times and will get the same result.
No duplicate data will be created, if the data is there the entry will be skipped.

## How to Use
Currently the utility has 2 commands, backup and restore.
Both commands require 2 flags to be passed.

`--server` defines the address of the Grocy server to connect to. It does not have to be the `/api` endpoint, the utility will handle that if not given.

`--api-key` in order to connect to the Grocy server an API key will need to be created prior to doing any backup or restore functions.

### Backup
Backup does exactly what it sounds like, it sends requests to your Grocy server to download the data and store it in a file for use later.
Backup does not require any additional arguments or flags aside from `--server` and `--api-key`.

### Restore
Restore takes a previous backup file and sends requests to your Grocy server in order to create any missing entries.
Restore requires the name of the file to be passed to it in addition to `--server` and `--api-key`.

## Examples
`grocy-backup --server "http://localhost" --api-key "123456789" backup`

`grocy-backup --server "http://localhost/api" --api-key "123456789" backup`

`grocy-backup --server "http://localhost/" --api-key "123456789" restore "backupfile.json"`

## Additional Notes/TODO
grocy-backup is very basic currently. It relies heavily on how the Grocy API functions in order to not duplicate entries.
The choice to use the API was out of simplicity, it may be better to interact with the `grocy.db` directly in the future.
Eventually I would like to provide a way to backup individual entries to their own files so they can be shared. Such as quantity units/conversions, recipes, products, etc.
