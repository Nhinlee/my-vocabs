DROP TABLE IF EXISTS vocab;
CREATE TABLE vocab (
    vocab_id TEXT PRIMARY KEY NOT NULL,
    word TEXT NOT NULL,
    image_urls TEXT[],
    next_review TIMESTAMPTZ DEFAULT (DATE_TRUNC('day', NOW()) + INTERVAL '2 days'), -- Truncate to midnight, then add 2 days
    reviewed_time INTEGER DEFAULT 0,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

-- Add constraint that word must be unique
CREATE UNIQUE INDEX vocab_word_idx ON vocab (word);