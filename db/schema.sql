CREATE TABLE blogs (
  id          INTEGER PRIMARY KEY,
  title       VARCHAR(100)  NOT NULL,
  body        TEXT  NOT NULL,
  created_at  DATETIME  NOT NULL
);
