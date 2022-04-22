CREATE TABLE IF NOT EXISTS states (
  client VARCHAR(16) PRIMARY KEY NOT NULL,
  state VARCHAR(32) NOT NULL
);

CREATE TABLE IF NOT EXISTS log (
  uuid UUID PRIMARY KEY NOT NULL,
  client VARCHAR(16),
  timestamp TIMESTAMP NOT NULL,
  event JSON NOT NULL
);