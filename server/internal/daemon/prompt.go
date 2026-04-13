package daemon

import (
	"fmt"
	"strings"
)

// BuildPrompt constructs the task prompt for an agent CLI.
// Keep this minimal — detailed instructions live in CLAUDE.md / AGENTS.md
// injected by execenv.InjectRuntimeConfig.
func BuildPrompt(task Task) string {
	if task.ChatSessionID != "" {
		return buildChatPrompt(task)
	}
	var b strings.Builder
	b.WriteString("You are running as a local coding agent for a Multica workspace.\n\n")
	fmt.Fprintf(&b, "Your assigned issue ID is: %s\n\n", task.IssueID)
	fmt.Fprintf(&b, "Start by running `multica issue get %s --output json` to understand your task, then complete it.\n", task.IssueID)

	// customize: per-task worktree branch hint. Tells the agent which branch
	// it's on so it can push and open a draft PR.
	if task.WorktreeBranch != "" {
		b.WriteString("\n## Git worktree\n\n")
		fmt.Fprintf(&b, "You are working in an isolated git worktree on branch `%s`.\n", task.WorktreeBranch)
		b.WriteString("When your work is complete:\n")
		b.WriteString("1. Commit your changes\n")
		fmt.Fprintf(&b, "2. Push the branch: `git push origin %s`\n", task.WorktreeBranch)
		b.WriteString("3. Open a **draft** pull request against the default branch\n")
		b.WriteString("4. Post the PR link as a comment on the issue\n")
		b.WriteString("Do NOT push directly to the default branch.\n")
	}

	// customize: per-task memex integration hints. These live in the prompt
	// (not in CLAUDE.md) because the flags vary per issue — writing them to
	// CLAUDE.md would leave stale state when the workdir is reused across
	// tasks with different flag combinations.
	if task.ConsultWiki || task.AllowWikiWrites {
		b.WriteString("\n## Memex integration\n\n")
		b.WriteString("This task has memex integration enabled. Read the `memex` skill for the HTTP API details, path conventions, and caveats; the bullets below tell you when to invoke it for *this* task.\n\n")
		if task.ConsultWiki {
			b.WriteString("- **consult_wiki is ENABLED** — before starting substantive work on this task, use the `memex` skill to query memex for prior context relevant to this issue. Treat results as background, not ground truth.\n")
		}
		if task.AllowWikiWrites {
			b.WriteString("- **allow_wiki_writes is ENABLED** — when your work is complete and *before* you mark the task done, use the `memex` skill to save a summary of what you did so future tasks on this project can find it. Write once, skip on blocked/cancelled outcomes.\n")
		}
	}

	return b.String()
}

// buildChatPrompt constructs a prompt for interactive chat tasks.
func buildChatPrompt(task Task) string {
	var b strings.Builder
	b.WriteString("You are running as a chat assistant for a Multica workspace.\n")
	b.WriteString("A user is chatting with you directly. Respond to their message.\n\n")
	fmt.Fprintf(&b, "User message:\n%s\n", task.ChatMessage)
	return b.String()
}
