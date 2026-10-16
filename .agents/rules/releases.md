# Releases and docs

Refines `~/.claude/rule-details/versioning-a-project.md` and `~/.claude/rule-details/writing-docs.md`.

- **REL1** The git tag is the version. Before 1.0 a breaking change to the exported API bumps the minor version; everything else is a patch. A regeneration that renames or removes generated methods or types is a breaking change.
- **REL2** Every release prepends a `## vX.Y.Z, YYYY-MM-DD` section to `CHANGELOG.md`. A release that regenerates the client names the spec version and source it was generated from.
- **REL3** `.agents/.codex-plugin/plugin.json` carries `version` and moves to the release version in the same commit. `.agents/.claude-plugin/plugin.json` has no `version`.
- **REL4** The agent skill in `.agents/skills/go-julep/` changes together with any generated method or type it names, and every example in it must compile against the code.
- **REL5** Commits, tags and pushes go through `git-update.sh` only. Tag pushes publish the skill to ClawHub through the `clawhub` job in `pipeline.yml`.
