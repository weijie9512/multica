-- customize: rollback wiki metadata fields on issue
ALTER TABLE issue DROP COLUMN IF EXISTS wiki_query_hint;
ALTER TABLE issue DROP COLUMN IF EXISTS allow_wiki_writes;
ALTER TABLE issue DROP COLUMN IF EXISTS consult_wiki;
