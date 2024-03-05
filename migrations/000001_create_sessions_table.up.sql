CREATE TABLE IF NOT EXISTS sessions (
    token TEXT PRIMARY KEY,
    data bytea NOT NULL,
    expiry TIMESTAMPTZ NOT NULL
);

CREATE INDEX sessions_expiry_idx ON sessions (expiry);