# opensirene
French company database based on open data from the French government

- https://www.data.gouv.fr/fr/datasets/base-sirene-des-entreprises-et-de-leurs-etablissements-siren-siret/
- http://files.data.gouv.fr/sirene

## Update database from scratch
```
go run update/*.go complete --maxworkers=10 --wd=/tmp
```
