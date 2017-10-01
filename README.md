# opensirene [![Build Status](https://travis-ci.org/jclebreton/opensirene.svg?branch=master)](https://travis-ci.org/jclebreton/opensirene)
French company database based on French government open data

- https://www.data.gouv.fr/fr/datasets/base-sirene-des-entreprises-et-de-leurs-etablissements-siren-siret/
- http://files.data.gouv.fr/sirene


## Getting Started

This project is not yet operational. It is still under development.

### Build
```
$ glide install
$ go build -o opensirene-cli github.com/jclebreton/opensirene/cli
```


### Update database from scratch
```
$ ./opensirene-cli complete
```

### Update database with daily upgrade file
```
$ ./opensirene-cli daily
```

### Options
```
$ ./opensirene-cli --help

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

### Show logs

```
$ ./opensirene-cli complete --month=Sep --debug
$ less opensirene.log

{"Number of workers":31,"Working directory":"/tmp/tmp304978929","level":"debug","msg":"","time":"2017-10-01T11:26:20+02:00"}
{"Filenames":"sirene_2017244_E_Q.zip, sirene_2017247_E_Q.zip, sirene_2017248_E_Q.zip, sirene_2017249_E_Q.zip, sirene_2017250_E_Q.zip, sirene_2017251_E_Q.zip, sirene_2017254_E_Q.zip, sirene_2017255_E_Q.zip, sirene_2017256_E_Q.zip, sirene_2017257_E_Q.zip, sirene_2017258_E_Q.zip, sirene_2017261_E_Q.zip, sirene_2017262_E_Q.zip, sirene_2017263_E_Q.zip, sirene_2017264_E_Q.zip, sirene_2017265_E_Q.zip, sirene_2017268_E_Q.zip, sirene_2017269_E_Q.zip, sirene_2017270_E_Q.zip, sirene_2017271_E_Q.zip, sirene_2017272_E_Q.zip","Number of files":21,"level":"info","msg":"Zip files to dowload","time":"2017-10-01T11:26:20+02:00"}
{"Errors":"Remote file not found: sirene_2017272_E_Q.zip","Number":1,"level":"error","msg":"Errors during processing files","time":"2017-10-01T11:26:22+02:00"}
{"Filenames":"sirc-17804_9075_14211_2017265_E_Q_20170923_022601923.csv, sirc-17804_9075_14211_2017250_E_Q_20170908_022239452.csv, sirc-17804_9075_14211_2017256_E_Q_20170914_022558547.csv, sirc-17804_9075_14211_2017255_E_Q_20170913_022333676.csv, sirc-17804_9075_14211_2017244_E_Q_20170902_022234303.csv, sirc-17804_9075_14211_2017251_E_Q_20170909_022245403.csv, sirc-17804_9075_14211_2017257_E_Q_20170915_022352087.csv, sirc-17804_9075_14211_2017249_E_Q_20170907_022217821.csv, sirc-17804_9075_14211_2017248_E_Q_20170906_022221313.csv, sirc-17804_9075_14211_2017264_E_Q_20170922_022642892.csv, sirc-17804_9075_14211_2017270_E_Q_20170928_022421138.csv, sirc-17804_9075_14211_2017268_E_Q_20170926_022431764.csv, sirc-17804_9075_14211_2017263_E_Q_20170921_022308465.csv, sirc-17804_9075_14211_2017261_E_Q_20170919_022218720.csv, sirc-17804_9075_14211_2017258_E_Q_20170916_022508788.csv, sirc-17804_9075_14211_2017254_E_Q_20170912_022213794.csv, sirc-17804_9075_14211_2017269_E_Q_20170927_022608879.csv, sirc-17804_9075_14211_2017271_E_Q_20170929_022338993.csv, sirc-17804_9075_14211_2017262_E_Q_20170920_022413518.csv, sirc-17804_9075_14211_2017247_E_Q_20170905_022240981.csv","Number of files":20,"level":"info","msg":"CSV files extracted","time":"2017-10-01T11:26:22+02:00"}
```
