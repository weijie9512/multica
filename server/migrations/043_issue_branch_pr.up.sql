-- customize: add branch_name and pr_url fields to issues for worktree
-- isolation and PR visibility on the board. Populated by the daemon on
-- task completion when the agent works in a git worktree.
ALTER TABLE issue ADD COLUMN branch_name TEXT;
ALTER TABLE issue ADD COLUMN pr_url TEXT;
