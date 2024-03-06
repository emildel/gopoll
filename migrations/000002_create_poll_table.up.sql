CREATE TABLE IF NOT EXISTS poll (
    pollSession TEXT PRIMARY KEY,
    title TEXT NOT NULL,
    answers text[] NOT NULL,
    expires_at timestamp(0) NOT NULL
);