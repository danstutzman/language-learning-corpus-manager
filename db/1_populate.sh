#!/bin/bash -ex

cd `dirname $0`

sqlite3 db.sqlite3 <<EOF
DROP TABLE IF EXISTS corpora;
CREATE TABLE corpora (
  id   INTEGER PRIMARY KEY NOT NULL,
  name TEXT
);
CREATE UNIQUE INDEX idx_corpora_name ON corpora(name);
INSERT INTO corpora (name) VALUES ('spintx');
EOF
