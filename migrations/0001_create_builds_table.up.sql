-- Create the builds table for storing CI pipeline results.
CREATE TABLE IF NOT EXISTS builds (
    id          BIGSERIAL PRIMARY KEY,
    repo        TEXT NOT NULL,                    -- e.g. "user/repo"
    branch      TEXT NOT NULL,                    -- e.g. "main"
    commit_hash TEXT NOT NULL,                    -- SHA of the commit
    status      TEXT NOT NULL CHECK (status IN ('pending', 'success', 'failure')),
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
    );

-- Indexes for fast filtering by status and time range.
CREATE INDEX idx_builds_status ON builds(status);
CREATE INDEX idx_builds_created_at ON builds(created_at);