-- customize: rollback — re-add wiki_query_hint column
--
-- Rolling back to the pre-042 schema. The column is nullable with no default,
-- so existing rows come back with NULL — the original contents are gone
-- unless you restored from a backup.
ALTER TABLE issue ADD COLUMN IF NOT EXISTS wiki_query_hint TEXT;
