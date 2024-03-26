CREATE TABLE IF NOT EXISTS poll (
    pollSession TEXT PRIMARY KEY,
    title TEXT NOT NULL,
    answers TEXT[] NOT NULL,
    results INTEGER[] NOT NULL,
    expires_at timestamp(0) NOT NULL
);