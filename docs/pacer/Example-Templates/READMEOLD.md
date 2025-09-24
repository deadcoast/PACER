# PACER Documentation Index
**PACER — Project Actions, Constraints & Evidence Register**  
This index links every PACER document and the most important sections inside each one.

> **AI/LLM Priority**: PACER is optimized for AI consumption and automation. Human readability is a bonus, not a primary feature.

---

## Core Docs
- **Authoritative Spec:** [pacer-spec.md](docs/pacer/pacer-spec.md)  
  - §3 File Format • §4 Data Model • §5 Identifiers • §6 Status Lifecycle • §7 Dependencies • §9 Validation • §10 Operations
- **Field Manual (Ops):** [pacer-field-manual.md](docs/pacer/pacer-field-manual.md)  
  - §3 Command Lexicon • §4 Daily Ritual • §5 Reports • §7 Playbooks • §8 Agent Integration
- **Quickstart (1‑page):** [pacer-quickstart.md](docs/pacer/pacer-quickstart.md)  
  - Create → Start → Review → Done • Dependency Gate • Validation pointers

## Design & Proof
- **Rationale (Design Notes):** [pacer-rationale.md](docs/pacer/pacer-rationale.md)  
  - Flat CSV justification • DAG and DoD reasoning • Anti‑patterns • Comparisons
- **Evidence Pack (Proof Template):** [pacer-evidence.md](docs/pacer/pacer-evidence.md)  
  - Operational Log • Metrics & formulas • 3–7 day procedure • Ready‑to‑fill tables

## Reference
- **Patterns & FAQ:** [pacer-faq.md](docs/pacer/pacer-faq.md)  
  - Epics • Spikes • Splitting • Overrides • Multi‑agent • Multi‑repo • Troubleshooting
- **JSON Schema:** [pacer.schema.json](docs/pacer/machine/pacer.schema.json) — Machine validation (CSV→JSON)
- **CSV Template:** [pacer-template.csv](docs/pacer/pacer-template.csv) — Header‑only starter

## Machine-Readable Artifacts (AI/LLM)
- **Full Spec (JSON):** [pacer-machine.json](docs/pacer/pacer-machine.json) — Complete specification for AI agents
- **Full Spec (YAML):** [pacer-machine.yaml](docs/pacer/pacer-machine.yaml) — YAML equivalent for AI preference
- **Command Patterns:** [pacer-commands.jsonl](docs/pacer/machine/pacer-commands.jsonl) — Natural language → deterministic edits
- **API Contract:** [pacer.agent.api.json](docs/pacer/machine/pacer.agent.api.json) — Method signatures and types
- **Grammar:** [pacer.agent.grammar.ebnf](docs/pacer/machine/pacer.agent.grammar.ebnf) — EBNF command grammar
- **Contract:** [pacer.agent.contract.json](docs/pacer/machine/pacer.agent.contract.json) — Agent behavior rules

---

## Start Here
1. **Quickstart** for a 2‑minute setup: [pacer-quickstart.md](docs/pacer/pacer-quickstart.md)  
2. **Spec** for the exact rules: [pacer-spec.md](docs/pacer/pacer-spec.md)  
3. **Field Manual** for daily operation: [pacer-field-manual.md](docs/pacer/pacer-field-manual.md)

**Validate (optional):** Convert your CSV to JSON and check with [pacer.schema.json](docs/pacer/machine/pacer.schema.json).

---

## Cross‑Doc Map (Fast Lookups)
- **IDs / immutability:** Spec §5 • FAQ §4 Q1  
- **Allowed transitions:** Spec §6.1 • Quickstart §3–6 • Field Manual §3.1  
- **Timestamps:** Spec §6.2 • Field Manual §3.1  
- **Dependency gate:** Spec §7.2 • Quickstart §4/6 • Field Manual §3.1  
- **Acyclic guidance:** Spec §7.3 • Rationale §3  
- **DoD requirements:** Spec §4.1 • Rationale §4 • FAQ §2  
- **WIP guidance:** Field Manual §6.4 • Rationale §5 • Evidence §4.3/4.6  
- **Validation:** Spec §9 • Schema • Quickstart §9  
- **Concurrency / atomic writes:** Spec §11 • Field Manual §8

---

## Governance & Versioning
- Current version: **PACER v1.1** (see [pacer-spec.md](docs/pacer/pacer-spec.md) §15).  
- Backward‑compatible extensions use profiles (Spec §13). Document deviations explicitly.

---

**Maintainers:** Keep this index in sync when adding or updating docs.

# PACER — Project Actions, Constraints & Evidence Register

PACER is a tiny, tool‑agnostic tracker: **one CSV, one row per task**. It's built to be deterministic for AIs and rock‑simple for humans.

> This README links the spec and shows exactly how to use PACER day‑to‑day.

---

## Files

- **docs/pacer/pacer-spec.md** — authoritative contract (IDs, columns, rules)
- **docs/pacer/pacer-field-manual.md** — daily ops (commands & etiquette)
- **docs/pacer/pacer-quickstart.md** — 1‑page starter
- **docs/pacer/pacer-rationale.md** — why it works
- **docs/pacer/pacer-evidence.md** — how to capture small proof
- **docs/pacer/pacer-faq.md** — patterns & answers
- **docs/pacer/machine/pacer.schema.json** — JSON Schema (optional validation)
- **docs/pacer/pacer-template.csv** — header‑only starter CSV

If you already have a live tracker CSV, keep it as your **single source of truth** (e.g., `docs/pacer/pac_backlog_tracker.csv`).

---

## ID & Columns (at a glance)

- **ID**: `PAC-###` (unique, immutable)  
- **Status**: `TODO → DOING → REVIEW → DONE`  
- **BlockedBy**: comma‑separated IDs that must be **DONE** first  
- **DoD**: **D**efinition **o**f **D**one — objective acceptance criteria

Full details: see **pacer-spec.md**.

---

## Quickstart

1. **Create a task (row)**  
   Add: `ID, Title, Phase, Status=TODO, DoD, (optional) BlockedBy`  
   _Template:_ `docs/pacer/pacer-template.csv`

2. **Start work**  
   Say or log: “Start PAC-021” → set `Status=DOING`, stamp `StartedAt`.

3. **Finish work**  
   “PAC-021 done” → set `Status=DONE`, stamp `DoneAt`.  
   **Rule:** You may only set DONE if all `BlockedBy` IDs are DONE.

4. **See what’s blocked**  
   Filter rows where any `BlockedBy` task is not `DONE`.

5. **Keep it clean**  
   - One row = one task
   - Don’t change IDs
   - Keep DoD objective and testable
   - Use Notes for brief status after each change

---

## Daily Ops (AI + humans)

- "Start PAC-021" → `DOING` + `StartedAt=now`
- "Mark PAC-032 done" → `DONE` + `DoneAt=now`
- "Block PAC-055 on 060,065" → `BlockedBy=PAC-060,PAC-065`
- "Assign PAC-040 to @me; note 'needs hero copy'" → set `Assignee`, append to `Notes`

> Tip: Keep ≤ 2–3 items in **DOING**. It shortens cycle time and keeps focus.

---

## Dependency Rule (enforced)

A task can be **DONE** **iff** every ID in **BlockedBy** is already **DONE**.  
This creates a natural order (a DAG) and prevents out‑of‑sequence work.

---

## Where to go next

- Read **pacer-spec.md** if you want exact rules.
- Skim **pacer-quickstart.md** to onboard contributors fast.
- Use **pacer-field-manual.md** for daily queries and updates.
- Collect light proof in **pacer-evidence.md** (logs & micro‑metrics).

---

## FAQ (short)

- **Can I skip REVIEW?** For solo, REVIEW can be brief; it still separates “coded” from “merged/shipped.”  
- **Can I rename `PAC-`?** Yes, change the prefix project‑wide and keep IDs unique.  
- **Can I add columns?** Yes, if they don’t change the meaning of existing fields.

For more, see **pacer-faq.md**.

---

### License / Attribution
PACER format and docs © you. This repo includes a JSON Schema to help validate the CSV; use or ignore as you like.
