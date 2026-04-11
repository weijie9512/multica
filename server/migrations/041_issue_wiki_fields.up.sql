-- customize: add wiki metadata fields to issue
--
-- Three fields that the Phase 5 sidecar reads to decide how to wire an
-- issue into the llm-wiki service:
--
--   consult_wiki       when true, sidecar fetches wiki context and posts
--                      a comment before the agent starts work
--   allow_wiki_writes  when true, sidecar writes the agent's final summary
--                      back to the wiki as a new entry on task completion
--   wiki_query_hint    optional free-form hint to bias the wiki search;
--                      falls back to the issue title when null/empty
--
-- All nullable with safe defaults so existing rows stay valid.
ALTER TABLE issue ADD COLUMN consult_wiki BOOLEAN NOT NULL DEFAULT FALSE;
ALTER TABLE issue ADD COLUMN allow_wiki_writes BOOLEAN NOT NULL DEFAULT FALSE;
ALTER TABLE issue ADD COLUMN wiki_query_hint TEXT;
