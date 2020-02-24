#!/bin/bash -ex

cd `dirname $0`

sqlite3 index.sqlite3 <<EOF
  DROP TABLE IF EXISTS corpora;
  CREATE TABLE corpora (
    id   INTEGER PRIMARY KEY NOT NULL,
    name TEXT NOT NULL
  );
  CREATE UNIQUE INDEX idx_corpora_name ON corpora(name);
  INSERT INTO corpora (name) VALUES ('spintx');

  DROP TABLE IF EXISTS files;
  CREATE TABLE files (
    id     INTEGER PRIMARY KEY NOT NULL,
    s3_key TEXT NOT NULL,
    size   INTEGER NOT NULL
  );
  CREATE UNIQUE INDEX idx_files_s3_key ON files(s3_key);
EOF
