# **Machine-readable PACER artifacts** 

Point your IDE’s AI directly at:

* **Canonical JSON (full spec for agents):**
  [docs/pacer/pacer-machine.json](sandbox:/mnt/data/docs/pacer/pacer-machine.json)

* **YAML equivalent (some AIs prefer YAML):**
  [docs/pacer/pacer-machine.yaml](sandbox:/mnt/data/docs/pacer/pacer-machine.yaml)

* **Commands pack (JSONL) for parser testing / few-shot priming:**
  [docs/pacer/pacer-commands.jsonl](sandbox:/mnt/data/docs/pacer/pacer-commands.jsonl)

### What’s inside the machine file (designed for deterministic AI behavior)

* **Register contract:** column set, enum values, `ID` regex, invariants, CSV rules (RFC 4180), extensibility.
* **Lifecycle model:** allowed transitions, timestamp semantics.
* **Dependency gate:** “DONE iff all `BlockedBy` are DONE” with failure behavior (`refuse_and_note`).
* **Operations:** `create`, `start_work`, `to_review`, `complete`, `rollback`, `edit_nonkey` with pre/post-conditions.
* **Validation suite:** required header/fields, ID uniqueness/pattern, enum checks, dependency existence, completion-gate enforcement.
* **Agent guidelines:** atomic writes / serialize, UTC timestamps, preserve unknown columns, WIP hint, notes format.
* **Command lexicon:** regex patterns → deterministic edits (start/review/done/assign/block/unblock/note/set DoD).
* **Examples:** minimal row + notes line format.

### How to wire this in Cursor / IDE AIs

1. **Pin the JSON:** Put `docs/pacer/pacer-machine.json` in your repo.
2. **System prompt (one-time):**

   * “Load `docs/pacer/pacer-machine.json` and obey its `register`, `operations`, `validation`, and `commands` rules. Refuse to mark DONE if blockers aren’t DONE. On every update, write a one-line Note.”
3. **At runtime:** Give natural commands that match the patterns:

   * `Start PAC-021`
   * `Review PAC-021`
   * `Done PAC-021`
   * `Block PAC-055 on 060,065`
   * `Assign PAC-040 to @alex`
   * `Note PAC-040: image upload routed; awaiting copy`

If you want a **project profile** (e.g., different ID prefix, extra phases/columns), I’ll emit a `pacer-machine.profile.json` overlay that overrides only those fields so your agents keep strict compatibility with the base spec.
