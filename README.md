# opensirene [![Build Status](https://travis-ci.org/jclebreton/opensirene.svg?branch=master)](https://travis-ci.org/jclebreton/opensirene)
French company database based on French government open data

- https://www.data.gouv.fr/fr/datasets/base-sirene-des-entreprises-et-de-leurs-etablissements-siren-siret/
- http://files.data.gouv.fr/sirene

## Update database from scratch
```
go run update/*.go complete --maxworkers=10 --wd=/tmp
```
