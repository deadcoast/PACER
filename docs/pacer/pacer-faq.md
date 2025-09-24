# PACER Patterns & FAQ
**Status:** Stable • **Applies to:** PACER v1.1 • **Spec:** [pacer-spec.md](docs/pacer/pacer-spec.md) • **Ops:** [pacer-field-manual.md](docs/pacer/pacer-field-manual.md)

Practical patterns, do/don't guidance, and answers to common questions. For normative rules, see the **[PACER Specification v1.1](docs/pacer/pacer-spec.md)**.

---

## 1) Core Patterns

### 1.1 Epics (Lightweight)
**Goal:** Group related PACs without inventing hierarchy.

**Pattern:** Create an epic PAC (e.g., `PAC-100 “Voting System”`) whose **DoD** is “children complete”. Children reference the epic in `Notes` and optionally **block** the epic.

- Parent `PAC-100` DoD: “Complete `PAC-101..PAC-109`”  
- Children `PAC-101..109`: add `Notes`: “Epic: PAC-100 Voting System”  
- Optionally: set `BlockedBy` on parent to the child IDs (Spec §7.1).

**Link:** See **Spec §4.1 DoD**, **Spec §7 Dependencies**, **Field Manual §7.2 Split a PAC**.

---

### 1.2 Spikes (Timeboxed Research)
**Goal:** De-risk by learning quickly.

**Pattern:** Create a PAC with DoD = “Findings documented; recommendation decided.” Strictly timebox; put links in `Notes`. `Status=REVIEW` → brief read → `DONE`.

**Link:** **Field Manual §7.1 Create a PAC**, **Spec §6 Status**.

---

### 1.3 Split a PAC (Too Big / Stuck)
**Trigger:** `DOING > 48h` or DoD bloats.

**Pattern:** Create `PAC-123A`, `PAC-123B` with narrower DoD; move details into the children. Parent’s DoD becomes “A & B complete”; parent blocks on A,B.

**Link:** **Field Manual §7.2**, **Evidence Pack §4.6 Aging DOING**.

---

### 1.4 Hotfix / Emergency Out-of-Order
**Goal:** Ship a fix while acknowledging the gate.

**Pattern:** Create a new PAC (e.g., `PAC-900`), document the exception in `Notes` (“override dependency order due to incident”). Afterward, reconcile dependencies and avoid repeats.

**Link:** **Spec §7.4 Administrative Overrides**, **Field Manual §7.4**.

---

### 1.5 Labels / Tags (Optional)
**Goal:** Slice the register without new columns.

**Pattern:** Put `#tag` tokens in `Notes` (e.g., `#perf`, `#infra`). Prefer a dedicated `Labels` column if you filter often (Spec §8 Extensibility).

---

### 1.6 Multiple Agents
**Goal:** Several automations operate safely.

**Pattern:** Serialize writes (lock the file) or write atomically via temp files. Agents **must** validate, preserve unknown columns, and refuse `DONE` with open blockers.

**Link:** **Spec §9 Validation**, **Spec §11 Concurrency & Integrity**, **Field Manual §8 Agent Integration**.

---

### 1.7 Multi‑Repo / Multi‑Project
**Options:**
1) One register per repo/project. Cross‑reference by plain text (e.g., “Depends on upstream: `CORE-012`”).  
2) One top‑level register with a `Repo` column (Spec §8).

Choose the smallest thing that works. Keep each register self‑consistent.

---

### 1.8 Archive without Deleting
**Goal:** Keep history intact.

**Pattern:** Avoid row deletion (Spec §10.7). Mark the PAC `DONE` or set `Status=TODO` with `Notes: de‑scoped`. For long‑term archives, move rows to a dated CSV copy.

---

## 2) Do / Don’t

**Do**
- Keep **DoD** objective and short.  
- Keep **≤ 2–3 in DOING** per owner.  
- Append a **one‑line note** after each significant change.  
- Enforce the **dependency gate** strictly.

**Don’t**
- Change **IDs** after creation.  
- Close a PAC with **open blockers**.  
- Let DOING **age** silently past 48h—split or unblock.  
- Scatter copies of the register (avoid multi‑source truth).

---

## 3) Troubleshooting

**Can’t mark DONE**
- A `BlockedBy` ID isn’t `DONE`. Check those rows. See **Spec §7.2**.

**Validation fails**
- Confirm header names (Spec §4.1), enum values for `Phase`/`Status`, `ID` pattern, and non‑empty `DoD`. See **Spec §9.1**.

**Lost updates / corruption**
- Concurrent writes or lack of atomic replace. See **Spec §11**.

**Too many “epics”**
- If a parent DoD becomes an essay, convert children into stand‑alone PACs and keep the parent minimal.

**Scope shifted mid‑stream**
- Update `Title` and `DoD`; add a `Notes` line “rescope”. Don’t change `ID` (Spec §5.1).

**Too many tags**
- If `Notes` becomes a tag soup, add a dedicated `Labels` column (Spec §8).

---

## 4) FAQ

**Q1: Can IDs be renumbered or re‑prefixed?**  
- **Re‑prefix:** Yes, project‑wide (e.g., `PAC-021` → `APP-021`), but keep the numeric part and mapping; document in a `Notes` entry or a “Profile” doc.  
- **Renumber:** Avoid. It breaks history. Create new PACs instead.

**Q2: Do I need REVIEW if I’m solo?**  
- Keep it brief. REVIEW separates “coded” from “merged/shipped”. If you skip it, note why in `Notes`.

**Q3: How do I handle subtasks?**  
- Either split into multiple PACs or list sub‑items inside **DoD** as an acceptance checklist.

**Q4: Can I model priorities and deadlines?**  
- Yes—add `Priority` and `Due` columns (Spec §8). Don’t overload `Notes`.

**Q5: What about sensitive information?**  
- Don’t put secrets or PII in `Notes` or `Title` (Spec §12 Security & Privacy).

**Q6: Multiple CSVs for the same project?**  
- Avoid. If you must, designate one as the **register of record** and link the rest read‑only.

**Q7: Can agents auto‑close based on tests?**  
- Yes—if they verify the DoD and the dependency gate. They **must** refuse when blockers aren’t `DONE` (Spec §7.2).

**Q8: What if a PAC is blocked for a week?**  
- Escalate: split, reorder, or re‑scope. Log the decision in `Notes`.

**Q9: How big should a PAC be?**  
- Aim for ≤ 2 days of work. If it threatens to grow, split (Pattern 1.3).

**Q10: Can I attach artifacts (links, screenshots)?**  
- Put URLs in `Notes` (or add an `Artifacts` column if frequent; Spec §8).

**Q11: How do AI agents use the new AI-First fields?**  
- **Context**: Provides background for AI understanding
- **Instructions**: Step-by-step guidance for execution
- **DependencyType**: Helps AI reason about constraints
- **LearningNotes**: Accumulates knowledge for future tasks

**Q12: What's the difference between hard, soft, and optional dependencies?**  
- **Hard**: Must be DONE before this PAC can be DONE (blocks completion)
- **Soft**: Preferred but not blocking (AI can work around)
- **Optional**: Nice to have but not required (AI can skip)

---

## 5) Cross‑Doc Map
- **Spec:** [pacer-spec.md](pacer-spec.md) — Normative behavior (IDs, lifecycle, dependencies, validation).  
- **Field Manual:** [pacer-field-manual.md](pacer-field-manual.md) — Commands, daily ritual, playbooks, agent rules.  
- **Quickstart:** [pacer-quickstart.md](pacer-quickstart.md) — 1‑page starter.  
- **Rationale:** [pacer-rationale.md](pacer-rationale.md) — Design choices, DAG/DoD arguments.  
- **Evidence Pack:** [pacer-evidence.md](pacer-evidence.md) — Templates, metrics, procedures.  
- **Schema:** [pacer.schema.json](docs/pacer/machine/pacer.schema.json) — Validation.

---

**End of PACER Patterns & FAQ**