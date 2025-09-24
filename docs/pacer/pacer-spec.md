# PACER Specification v1.0 (Authoritative)
**PACER — Project Actions, Constraints & Evidence Register**

This document defines the **authoritative, normative** specification for PACER v1.0.
Implementations **MUST** follow this specification to claim PACER compatibility.

> Normative keywords (MUST, MUST NOT, REQUIRED, SHALL, SHALL NOT, SHOULD, SHOULD NOT, RECOMMENDED, MAY, OPTIONAL) are to be interpreted as described in RFC 2119.

---

## 1. Scope & Non‑Goals

**Scope.** PACER is a minimal, tool‑agnostic format for tracking work as a **single CSV table**. Each row represents a task, called a **PAC**. The format is intentionally flat, deterministic, and easy to automate.

**Non‑Goals.** PACER does **not** define UI, sprint/iteration rituals, estimation models, or vendor‑specific workflows. PACER is not a replacement for Git, issue trackers, or docs; it is a compact **register** that integrates with them.

---

## 2. Concepts & Terminology

- **PAC.** A single unit of work represented by one CSV row.
- **Register.** The PACER CSV file itself (the single source of truth).
- **DoD.** Definition of Done — objective acceptance criteria for a PAC.
- **Blocker.** A PAC listed in another PAC’s `BlockedBy` field.
- **DAG.** Directed Acyclic Graph induced by the `BlockedBy` relation.

---

## 3. File Format

PACER v1.0 is defined primarily for **CSV** (Comma‑Separated Values). A **JSON** representation (array of records) is also permitted for validation and APIs.

### 3.1 CSV Requirements (Normative)

- The register **MUST** be encoded as UTF‑8 without BOM.
- The file **MUST** contain a header row naming columns exactly as specified.
- The file **MUST** contain **one row per PAC**.
- Columns **MUST** appear at least in the **Required Column Set** (Section 4.1). Additional columns are allowed (Section 8).

**Delimiter & quoting.** Comma `,` as delimiter. Fields containing commas, quotes, or newlines **MUST** be quoted per RFC 4180. Double quotes inside a quoted field **MUST** be escaped by doubling them.

### 3.2 JSON Requirements (Informative)

For machine APIs, the CSV MAY be mirrored as an array of JSON objects using the same column names as keys. Validation MAY be performed via the official JSON Schema (`pacer.schema.json`).

---

## 4. Data Model

### 4.1 Required Column Set (Normative)

| Column     | Type        | Constraints                                                                 | Semantics |
|------------|-------------|------------------------------------------------------------------------------|-----------|
| `ID`       | string      | **REQUIRED.** Unique, immutable. Default pattern `^PAC-\d{3,4}$` (configurable prefix). | Stable handle for the PAC. |
| `Title`    | string      | **REQUIRED.** Non‑empty, ≤100 UTF‑8 characters.                             | Human‑readable summary. |
| `Phase`    | enum string | **REQUIRED.** One of: `Foundation`, `Auth & DB`, `Contest Mgmt`, `Voting`, `UX & Polish`, `Admin Tools`, `Integrations`, `Accounts & Profiles`, `Security/Perf/Ops`, `Embeds`, `Release`. | Logical grouping for planning. |
| `Status`   | enum string | **REQUIRED.** One of: `TODO`, `DOING`, `REVIEW`, `DONE`.                    | Current lifecycle state. |
| `DoD`      | string      | **REQUIRED.** Non‑empty. At least **one** objective acceptance criterion.   | Definition of Done. |

### 4.2 Recommended Columns (Normative)

| Column       | Type        | Constraints                                                        | Semantics |
|--------------|-------------|---------------------------------------------------------------------|-----------|
| `BlockedBy`  | string      | Comma‑separated list of valid `ID` values, or empty.               | Hard dependencies that must be DONE first. |
| `Assignee`   | string      | Free text, MAY be an `@handle`.                                    | Owner. |
| `StartedAt`  | timestamp   | ISO 8601 (e.g., `2025-09-23T21:17:00Z`). Set on first entry to DOING. | Work start time. |
| `DoneAt`     | timestamp   | ISO 8601. Set when entering DONE.                                  | Completion time. |
| `Notes`      | string      | Free text. KEEP concise; prefer one line per update.               | Context/comments. |

### 4.3 Optional Columns (Informative)

Implementations MAY add columns (e.g., `Priority`, `Labels`) provided they do **not** change the meaning of required fields. Names SHOULD avoid collisions with required columns.

---

## 5. Identifiers

### 5.1 ID Grammar (Normative)
- Default pattern: ``^PAC-\d{3,4}$`` (e.g., `PAC-021`, `PAC-1001`).  
- The textual **prefix** (`PAC`) MAY be changed project‑wide, but all IDs **MUST** remain unique and stable.
- **IDs MUST NOT change** after creation. If scope changes, create a new PAC and cross‑reference in `Notes`.

### 5.2 Uniqueness (Normative)
- All `ID` values in a single register **MUST** be unique (case‑sensitive).

---

## 6. Status Lifecycle

### 6.1 Allowed Transitions (Normative)
- `TODO → DOING`
- `DOING → REVIEW` or `DOING → TODO` (rollback)
- `REVIEW → DONE` or `REVIEW → DOING` (changes requested)
- `DONE` is terminal. To undo, you **MUST** add a `Notes` entry explaining the rollback and move to `REVIEW` or `DOING`.

### 6.2 Timestamp Semantics (Normative)
- On first transition into `DOING`, set `StartedAt` (if empty).
- On transition into `DONE`, set `DoneAt`.
- Timestamps **SHOULD** be in UTC ISO 8601 (`YYYY-MM-DDThh:mm:ssZ`).

### 6.3 WIP Guidance (Informative)
Keeping ≤2–3 PACs in `DOING` tends to reduce cycle time (Little’s Law).

---

## 7. Dependencies

### 7.1 Dependency Encoding (Normative)
- `BlockedBy` is a comma‑separated list of `ID`s (no spaces recommended), or empty.
- Each referenced `ID` **MUST** exist in the register.

### 7.2 Completion Rule (Normative)
A PAC **MAY** transition to `DONE` **iff** **every** referenced `ID` in `BlockedBy` currently has `Status = DONE`.

### 7.3 Acyclicity (RECOMMENDED)
The dependency graph **SHOULD** be acyclic (a DAG). Implementations SHOULD warn on cycles. Cycles make `DONE` unattainable without administrative override.

### 7.4 Administrative Overrides (Informative)
Projects MAY allow explicit, logged overrides for emergencies. Overrides SHOULD append a `Notes` entry explaining rationale and impact.

---

## 8. Extensibility & Compatibility

- New columns MAY be added if they do not alter the meaning of existing fields.
- Enumerations (e.g., `Phase`) MAY be extended by profile, but MUST NOT overload existing values.
- Tools SHOULD ignore unknown columns and preserve them on write (round‑tripping).

---

## 9. Validation

### 9.1 JSON Schema (Normative)
PACER provides an official JSON Schema: `docs/pacer/pacer.schema.json`. CSV registers MAY be converted to JSON and validated against the schema.

### 9.2 Register‑Level Checks (Normative)
Implementations **MUST** perform at least:
1. **Header check:** Required columns present.
2. **Row check:** `ID`, `Title`, `Phase`, `Status`, `DoD` non‑empty and valid.
3. **Uniqueness:** `ID` uniqueness across file.
4. **Dependency existence:** Every `BlockedBy` ID exists.
5. **Completion rule:** Refuse or flag `Status = DONE` when any blocker is not `DONE`.

### 9.3 Severity & Handling (RECOMMENDED)
- **ERROR** → refuse write/transition.
- **WARN** → permit write but flag (e.g., long `Notes`, missing `Assignee`).
- **INFO** → stylistic suggestions (e.g., Title > 80 chars).

---

## 10. Operations

This section defines normative behavior for common mutations.

### 10.1 Create PAC (Normative)
- **Inputs:** `ID`, `Title`, `Phase`, `Status=TODO`, `DoD`, optional `BlockedBy`, `Assignee`.
- **Preconditions:** `ID` unused; required fields valid.
- **Postconditions:** Row appended. `StartedAt`/`DoneAt` empty.

### 10.2 Start Work (Normative)
- **Action:** Set `Status=DOING`. If `StartedAt` empty, set to current UTC ISO timestamp.
- **Preconditions:** `Status` is `TODO` (or `REVIEW` if rework).

### 10.3 Submit for Review (Normative)
- **Action:** Set `Status=REVIEW`.
- **Preconditions:** `Status` is `DOING`.

### 10.4 Complete (Normative)
- **Action:** Set `Status=DONE`, set `DoneAt` (UTC ISO).
- **Preconditions:** All `BlockedBy` rows have `Status=DONE`. `DoD` satisfied.
- **Failure:** Refuse transition; append `Notes` explaining unmet blockers.

### 10.5 Rollback (Normative)
- **Action:** Move `DONE → REVIEW` or `REVIEW → DOING`.
- **Requirement:** Append `Notes` describing the reason.

### 10.6 Edit Non‑Key Fields (Normative)
- **Action:** Update `Title`, `Phase`, `DoD`, `Notes`, `Assignee`, `BlockedBy`.
- **Constraint:** `ID` **MUST NOT** change. Updates **SHOULD** be atomic and logged (e.g., via Git diff).

### 10.7 Delete (RECOMMENDED: Avoid)
- Deletion of rows **SHOULD** be avoided. Prefer `Status=TODO` with `Notes` “de‑scoped” or an archival mechanism. If deletion occurs, implementations SHOULD log the removal and ensure no other `BlockedBy` references remain dangling.

---

## 11. Concurrency & Integrity (Informative)

- Prefer **atomic writes** (write temp file, then replace).
- Use **file locks** or serialize updates to avoid lost updates.
- On conflict, **merge** by newest timestamped change, preserving both `Notes` entries.

---

## 12. Security & Privacy (Informative)

- Do not store secrets in `Notes` or `Title`.
- Timestamps SHOULD be UTC and may exclude PII.
- Keep the register in version control; limit write access to trusted agents.

---

## 13. Profiles (Informative)

Implementations MAY define **profiles** to tailor enums or columns for a domain (e.g., “Solo Dev”, “Open Source”). Profiles MUST cite PACER v1.0 and list deviations explicitly.

---

## 14. Examples

### 14.1 Minimal CSV (With Header)
```
ID,Title,Phase,Status,BlockedBy,Assignee,StartedAt,DoneAt,DoD,Notes
PAC-021,Create/Edit/Delete Contest,Contest Mgmt,TODO,,,,"","Form with title/desc/dates/rules; status cycle draft→active→voting→completed",""
```

### 14.2 Typical Lifecycle
1. Create → `Status=TODO`
2. Start → `Status=DOING`, set `StartedAt`
3. Review → `Status=REVIEW`
4. Done → verify blockers; `Status=DONE`, set `DoneAt`

---

## 15. Versioning & Compatibility

- This document defines **PACER v1.0**. Minor errata MAY be published without changing the version if they do not alter normative behavior.
- Backward‑compatible extensions (new optional columns, additional `Phase` values under a profile) are allowed.
- Breaking changes require a new major version (e.g., PACER v2.0).

---

## 16. Compliance Checklist (For Implementers)

- [ ] CSV header includes required columns (Section 4.1).
- [ ] `ID` uniqueness enforced; IDs immutable.
- [ ] Status transitions enforced (Section 6.1).
- [ ] Timestamps set per transition (Section 6.2).
- [ ] Dependency completion rule enforced (Section 7.2).
- [ ] JSON Schema validation available (Section 9.1).
- [ ] Atomic writes or equivalent conflict handling (Section 11).
- [ ] Deletion discouraged; archival preferred (Section 10.7).

---

## 17. Glossary

- **PAC** — a single task row in the register.  
- **Register** — the CSV file acting as the system of record.  
- **DoD** — Definition of Done; objective acceptance criteria.  
- **Blocker** — a dependency PAC that must be done first.  
- **DAG** — Directed Acyclic Graph induced by `BlockedBy` references.

---

*End of PACER Specification v1.0*
