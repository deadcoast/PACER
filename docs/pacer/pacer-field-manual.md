# PACER Field Manual (Ops Guide)

**Status:** Stable • **Applies to:** PACER v1.1 • **Spec:** [pacer-spec.md](pacer-spec.md)

This manual describes **how to operate PACER day-to-day** for both humans and AI agents. It implements the rules defined in the **authoritative spec** and adds practical workflows optimized for AI/LLM consumption.

> If a rule here appears to conflict with the spec, the spec wins. See: **[PACER Specification v1.1](pacer-spec.md)**.

---

## 0. Quick Orientation

- **Register:** the single CSV file that lists all PACs (one row = one task).
- **Lifecycle:** `TODO → DOING → REVIEW → DONE` (see **Spec §6: Status Lifecycle**).
- **Completion gate:** A PAC may be `DONE` **iff** all IDs in `BlockedBy` are `DONE` (see **Spec §7.2**).
- **Evidence:** DoD lives in the row; it must be objective (see **Spec §4.1**).
- **AI-First:** Enhanced fields for AI agents: Context, Instructions, Dependency Intelligence (see **Spec §4.3**).

Shortcuts:

- **Quickstart (1 page):** [pacer-quickstart.md](pacer-quickstart.md)
- **Design notes:** [pacer-rationale.md](pacer-rationale.md)
- **Evidence pack:** [pacer-evidence.md](pacer-evidence.md)
- **FAQ & patterns:** [pacer-faq.md](pacer-faq.md)

---

## 1. Operating Principles

1. **Single source of truth.** Maintain exactly one register; version it (Git recommended).
2. **IDs are immutable.** Never edit an `ID` after creation. (See **Spec §5.1**)
3. **Deterministic transitions only.** Use the allowed lifecycle moves; timestamp correctly. (See **Spec §6.1–6.2**)
4. **No DONE with open blockers.** Enforce the dependency gate. (See **Spec §7.2**)
5. **Objective DoD.** A PAC is complete when its DoD is satisfied—no gut feel.
6. **Small WIP.** Keep ≤ 2–3 PACs in `DOING` per person/agent.

---

## 2. Roles (Humans & Agents)

- **Owner (Assignee).** Responsible for moving the PAC through its lifecycle and updating Notes.
- **Reviewer.** For `REVIEW` transitions; may be the same person in solo workflows.
- **Agent.** An AI or automation that reads/writes the register with deterministic rules.

**Authority:** Only trusted humans/agents MAY write to the register. Use code reviews or protected branches where possible.

---

## 3. Command Lexicon (Natural Language → Deterministic Edit)

These patterns are intentionally terse so humans and AIs can operate consistently.

> **Note:** All transitions and timestamp semantics must follow **Spec §6**. Blocker checks must follow **Spec §7**.

### 3.1 Lifecycle

- **Start work**  
  Phrase: `Start PAC-021`  
  Edit: `Status=DOING`, set `StartedAt=<UTC now>` *iff* `StartedAt` empty.

- **Send to review**  
  Phrase: `Review PAC-021` or `PAC-021 to review`  
  Edit: `Status=REVIEW`.

- **Complete**  
  Phrase: `PAC-021 done`  
  Edit: verify blockers all `DONE` → set `Status=DONE`, `DoneAt=<UTC now>`.  
  If a blocker is not `DONE`, **refuse** and append to `Notes`: `blocked by: PAC-0xx`.

- **Rollback**  
  Phrase: `Rollback PAC-021 to doing`  
  Edit: `REVIEW → DOING` (or `DONE → REVIEW`) and append `Notes` why.

### 3.2 Ownership & Dependencies

- **Assign**  
  Phrase: `Assign PAC-040 to @alex`  
  Edit: `Assignee=@alex`.

- **Block on**  
  Phrase: `Block PAC-055 on 060,065`  
  Edit: `BlockedBy=PAC-060,PAC-065` (validate existence).

- **Unblock**  
  Phrase: `Unblock PAC-055 remove 065`  
  Edit: Remove `PAC-065` from `BlockedBy` with rationale in `Notes`.

### 3.3 DoD & Notes

- **Set/Update DoD**  
  Phrase: `DoD PAC-032: server uniqueness; already-voted error; confirmation UI`  
  Edit: Replace `DoD` with the provided text (retain objectivity).

- **Append status note**  
  Phrase: `Note PAC-032: wired uniqueness check; 2 tests failing`  
  Edit: Append a new line to `Notes`, prefixed with timestamp and author/agent if available.

### 3.4 Batch Ops

- **Start multiple**  
  Phrase: `Start PAC-010, PAC-011` → apply **Start work** to each (respect WIP limit).
- **Mark many done**  
  Phrase: `Done PAC-030..PAC-033` → attempt `DONE` in order; refuse where blockers aren’t resolved.

---

## 4. Daily Ritual (5–10 minutes)

1. **Triage (Incoming TODO)**  
   - Validate `Title`, `Phase`, `DoD`.  
   - Add `BlockedBy` if needed.

2. **Plan (Pull into DOING)**  
   - Pull ≤ 2–3 PACs per owner.  
   - Set `StartedAt` on first entry to DOING.

3. **Review**  
   - Move finished work to `REVIEW`.  
   - If solo, quickly verify DoD → `DONE` (respect blockers).

4. **Health Checks**  
   - **Blocked report:** Any PAC with non-DONE blockers.  
   - **Aging DOING:** `DOING` older than 48h → split or unblock.  
   - **Idle TODO:** High-priority items untouched for a week → reconsider or de-scope.

---

## 5. Reports (Simple Filters)

These can be executed by a human in a spreadsheet or by an agent script.

- **Blocked PACs**  
  Show rows where `BlockedBy` contains any ID whose `Status != DONE`.

- **Old DOING**  
  `Status=DOING` and `StartedAt < now - 48h`.

- **Review Queue**  
  `Status=REVIEW` sorted by `StartedAt` desc.

- **Throughput**  
  Count rows with `Status=DONE` grouped by day/week (use `DoneAt`).

- **Cycle Time**  
  For DONE rows, compute `DoneAt - StartedAt`. Show median/percentiles.

---

## 6. Conventions

### 6.1 Titles

Keep ≤ 80 chars, lead with action verb. Good: “Create contest CRUD form”. Bad: “contest stuff”.

### 6.2 Commit Messages

Include the ID in commits/PRs: `feat: PAC-021 contest CRUD form`.  
If a PR closes a PAC, move to `REVIEW`; merge should set `DONE`.

### 6.3 Notes Format

Each entry is one line, newest on top, optionally prefixed:

```pace
[2025-09-23T20:10Z] agent@gpt: validated DoD; moving to REVIEW
```

### 6.4 WIP Policy

Unless overridden, each owner keeps ≤ 3 active PACs in DOING.

---

## 7. Common Playbooks

### 7.1 Create a PAC

1. Add a row: `ID, Title, Phase, Status=TODO, DoD, (optional) BlockedBy, Assignee`  
2. Validate fields per **Spec §4.1**.  
3. If immediately actionable, `Start PAC-xxx`.

### 7.2 Split a PAC

- Create `PAC-123A`, `PAC-123B`; copy relevant DoD fragments.  
- Update parent `Notes` to reference children.  
- Optionally block parent on both children; parent to DONE when children are DONE.

### 7.3 Rescope Without ID Change

- Update `Title`/`DoD`; add a `Notes` entry explaining scope change.  
- Keep `ID` stable (see **Spec §5.1**).

### 7.4 Emergency Out-of-Order

- Create a new PAC (e.g., hotfix), document exception in `Notes`.  
- Afterward, reconcile dependencies; avoid repeated overrides (see **Spec §7.4**).

### 7.5 Import/Export

- CSV is the source; JSON export MAY be used for validation against `machine/pacer.schema.json`.  
- Keep headers exact. Preserve unknown columns on round-trip (see **Spec §8**).

---

## 8. Agent Integration

### 8.1 Read → Act → Write Loop

1. **Read:** Parse register; validate per **Spec §9**.  
2. **Select:** Choose the target PAC(s) by ID or filter (e.g., oldest REVIEW).  
3. **Act:** Perform the work (outside the register).  
4. **Write:** Update `Status`, timestamps, and append one-line `Notes`.  
5. **Enforce:** Refuse `DONE` if blockers not DONE. Respect allowed transitions.

### 8.2 Determinism Requirements (for Agents)

- No partial updates; apply atomic edits per PAC.  
- Use UTC timestamps (ISO 8601).  
- Never edit `ID`.  
- Preserve unknown columns when writing the file.  
- If a rule prevents completion, **refuse** and write an explanatory `Notes` entry.

### 8.3 Example Agent Note

```pace
[2025-09-23T18:02Z] agent@ide: PAC-032 tests passing; dependency PAC-011 not DONE → refusal to close
```

---

## 9. AI-First Operations (v1.1)

### 9.1 Context & Memory Management

- **Context**: Provide background information the AI needs to understand the task
- **PreviousAttempts**: Document what has been tried before to prevent repetition
- **RelatedWork**: Link to similar PACs for pattern recognition
- **LearningNotes**: Record insights gained from working on this PAC

### 9.2 Dependency Intelligence

- **DependencyType**: Use `hard`, `soft`, or `optional` to help AI understand constraint severity
- **DependencyReason**: Explain why dependencies exist for AI reasoning
- **UnblockingStrategy**: Provide alternative approaches when blocked

### 9.3 Instruction Clarity

- **Instructions**: Write clear, step-by-step guidance for AI execution
- **ExpectedOutput**: Define what success looks like
- **ValidationCriteria**: Specify how to verify completion
- **ErrorHandling**: Provide recovery strategies for common failures

### 9.4 AI Agent Workflow

1. **Read Context**: Review `Context`, `PreviousAttempts`, `RelatedWork`
2. **Understand Dependencies**: Check `DependencyType`, `DependencyReason`, `UnblockingStrategy`
3. **Execute Instructions**: Follow `Instructions`, validate against `ExpectedOutput`
4. **Handle Errors**: Use `ErrorHandling` strategies if issues arise
5. **Update Learning**: Record insights in `LearningNotes`

---

## 10. Troubleshooting

- **Symptom:** Can’t mark DONE.  
  **Check:** A `BlockedBy` ID isn’t DONE yet. See **Spec §7.2**.

- **Symptom:** Lost updates / garbled file.  
  **Check:** Concurrent writes. Serialize updates or use file locks. See **Spec §11**.

- **Symptom:** Validation fails on schema.  
  **Check:** Header names, `ID` pattern, enum values. See **Spec §9.1**.

- **Symptom:** Long DOING times.  
  **Check:** WIP > 3; split PAC or unblock dependencies. See **§6.3**.

---

## 10. References

- **Authoritative Spec:** [pacer-spec.md](pacer-spec.md) — see **§4 Data Model**, **§6 Status**, **§7 Dependencies**, **§9 Validation**.
- **Quickstart:** [pacer-quickstart.md](pacer-quickstart.md)
- **Rationale:** [pacer-rationale.md](pacer-rationale.md)
- **Evidence Pack:** [pacer-evidence.md](pacer-evidence.md)
- **FAQ & Patterns:** [pacer-faq.md](pacer-faq.md)
- **JSON Schema:** [pacer.schema.json](machine/pacer.schema.json)

---
