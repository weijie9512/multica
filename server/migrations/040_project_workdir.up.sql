-- customize: add working_dir to project
-- When set, the daemon spawns the agent with cwd = working_dir instead of
-- the default isolated path under MULTICA_WORKSPACES_ROOT.
ALTER TABLE project ADD COLUMN working_dir TEXT;
