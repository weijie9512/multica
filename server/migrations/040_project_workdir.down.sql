-- customize: rollback working_dir on project
ALTER TABLE project DROP COLUMN IF EXISTS working_dir;
