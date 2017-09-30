# opensirene [![Build Status](https://travis-ci.org/jclebreton/opensirene.svg?branch=master)](https://travis-ci.org/jclebreton/opensirene)
French company database based on French government open data

- https://www.data.gouv.fr/fr/datasets/base-sirene-des-entreprises-et-de-leurs-etablissements-siren-siret/
- http://files.data.gouv.fr/sirene


## Getting Started

This project is not yet operational. It is still under development.

### Update database from scratch
```
$ go run update/*.go complete
```

### Update database with daily upgrade file
```
$ go run update/*.go daily
```

### Options
```
$ go run update/*.go --help

Opensirene

French company database based on French government open data.
Github: https://github.com/jclebreton/opensirene

Usage:
  update daily [--wd=<path>] [--debug]
  update complete [--wd=<path>] [--maxworkers=<int>] [--month=<string>] [--debug]
  update -h | --help

Options:
  --wd=<path>        Working directory path (by default: /tmp/tmp[0-9]{8,})
  --maxworkers=<int> Maximum number of workers to use for processing files (min: 1, max: 31)
  --month=<string>   Month to download (ex: Sep)
  --debug            Enable debugging
  -h --help          Show this screen.
```
