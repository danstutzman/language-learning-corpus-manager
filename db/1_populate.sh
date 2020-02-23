#!/bin/bash -ex

cd `dirname $0`

sqlite3 db.sqlite3 <<EOF
  DROP TABLE IF EXISTS corpora;
  CREATE TABLE corpora (
    id   INTEGER PRIMARY KEY NOT NULL,
    name TEXT NOT NULL
  );
  CREATE UNIQUE INDEX idx_corpora_name ON corpora(name);
  INSERT INTO corpora (name) VALUES ('spintx');

  DROP TABLE IF EXISTS files;
  CREATE TABLE files (
    id       INTEGER PRIMARY KEY NOT NULL,
    filename TEXT NOT NULL,
    size     INTEGER NOT NULL
  );
  CREATE UNIQUE INDEX idx_files_filenames ON files(filename);
EOF
