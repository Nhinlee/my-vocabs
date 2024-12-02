DROP TABLE IF EXISTS vocab;
CREATE TABLE vocab (
    vocab_id TEXT PRIMARY KEY NOT NULL,
    word TEXT NOT NULL,
    image_urls TEXT[],
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);
