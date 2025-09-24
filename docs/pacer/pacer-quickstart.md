# PACER Quickstart (1‑Page)
**Status:** Stable • **Applies to:** PACER v1.0 • **Spec:** [pacer-spec.md](docs/pacer/pacer-spec.md) • **Ops:** [pacer-field-manual.md](docs/pacer/pacer-field-manual.md)

PACER = one CSV, one row per task (“PAC”). This page shows the **minimum** you need to start, work, and finish—fast.

---

## 1) Create the Register
- Copy `pacer-template.csv` into your repo and name it (e.g., `pacer.csv`).
- Keep it in version control. One file = the **single source of truth**.

**Headers (required):**
```
ID,Title,Phase,Status,BlockedBy,Assignee,StartedAt,DoneAt,DoD,Notes
```

---

## 2) Add Your First PAC
Create a new row with:
- **ID**: `PAC-001` (unique, immutable) — see **Spec §5**  
- **Title**: short action summary  
- **Phase**: one of the enums — see **Spec §4.1**  
- **Status**: `TODO`  
- **DoD**: objective acceptance criteria (at least one)

Example:
```
PAC-001,Initialize repo,Foundation,TODO,,@you,,,Create repo; CI runs green;
```

---

## 3) Start Work
Command (natural language): **“Start PAC-001”**  
Action: set `Status=DOING` and stamp `StartedAt` (UTC).  
Rule: Follow allowed transitions — **Spec §6.1–6.2**.

---

## 4) Track Dependencies (Optional)
If PAC-010 depends on PAC-001 and PAC-005:
```
BlockedBy=PAC-001,PAC-005
```
Gate: A PAC can be **DONE** **iff** all `BlockedBy` IDs are **DONE** — **Spec §7.2**.

---

## 5) Send to Review
Command: **“Review PAC-001”**  
Action: set `Status=REVIEW`.

Solo dev? Treat REVIEW as a short verification step against the **DoD**.

---

## 6) Complete
Command: **“PAC-001 done”**  
Action: verify blockers → set `Status=DONE` and `DoneAt` (UTC).  
If a blocker is not DONE, **refuse** and add a short `Notes` line explaining which blocker is open. (See **Spec §7.2**.)

---

## 7) Daily Ritual (5 minutes)
- **Pull:** Move ≤2–3 PACs to `DOING`.  
- **Review:** Move finished work to `REVIEW`, then `DONE` if DoD is satisfied.  
- **Check:** Run the **Blocked** & **Aging DOING** filters (see **Field Manual §5**).

---

## 8) Common Commands
- “Assign PAC-040 to @alex” → set `Assignee`
- “Block PAC-055 on 060,065” → set `BlockedBy=PAC-060,PAC-065`
- “Note PAC-032: tests green, need copy” → append to `Notes`
- “DoD PAC-032: server uniqueness; already‑voted error; confirmation UI” → replace `DoD`

More patterns: **[pacer-field-manual.md](pacer-field-manual.md)** and **[pacer-faq.md](pacer-faq.md)**.

---

## 9) Validate (Optional but Recommended)
- Convert CSV to JSON and validate against **[pacer.schema.json](pacer.schema.json)**.  
- Minimum checks: header present, ID unique/immutable, enums valid, dependency gate enforced — **Spec §9**.

---

## 10) That’s It
- One file. One row per task. Deterministic rules.
- When in doubt, the **Spec** rules. For daily operation, use the **Field Manual**.

**Next:** Read **[pacer-field-manual.md](pacer-field-manual.md)** for workflows and reports.
