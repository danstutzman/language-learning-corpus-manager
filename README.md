# How to run automated tests

go test ./...

# How to run locally

`go install -v ./...`

`db/1_populate.sh`

`AWS_PROFILE=personal aws s3 cp db/index.sqlite3 s3://danstutzman-language-learning-corpora/index.sqlite3`

```
S3_TOKEN=MY_TOKEN_HERE \
  S3_SECRET=MY_SECRET_HERE \
  S3_REGION=us-east-1 \
  S3_BUCKET=danstutzman-language-learning-corpora \
  TEMP_DIR=db \
  $GOPATH/bin/backend
```
