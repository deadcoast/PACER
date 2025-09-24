# PACER Evidence Pack (Proof Template)
**Status:** Stable • **Applies to:** PACER v1.0 • **Spec:** [pacer-spec.md](docs/pacer/pacer-spec.md) • **Ops:** [pacer-field-manual.md](docs/pacer/pacer-field-manual.md)

This pack shows how to collect **small, credible, repeatable evidence** that PACER improves execution—without a large study. It provides templates, formulas, and a 3–7 day procedure you can run solo.

> If any procedure appears to conflict with the spec, the spec wins. See: **[PACER Specification v1.0](docs/pacer/pacer-spec.md)**.

---

## 1. Objectives
Demonstrate—using your own PACER register—that:
1. Work moves predictably through **TODO → DOING → REVIEW → DONE** (Spec §6).
2. The **dependency gate** (Spec §7.2) prevents out‑of‑order completion.
3. **Embedded DoD** (Spec §4.1) reduces rework / back‑and‑forth.
4. Small **WIP** correlates with shorter cycle times (Little’s Law).

---

## 2. Minimal Data to Capture
From your register (CSV):
- `ID, Title, Status, StartedAt, DoneAt, BlockedBy, Notes`
- Optional: `Assignee` (if >1 person/agent)

You already have this if you follow the Spec and Field Manual. Keep timestamps in **UTC ISO 8601**.

---

## 3. Daily Operational Log (Template)
Record notable transitions or decisions (10–20 entries is enough).

| Timestamp (UTC) | Action | ID | Detail |
|---|---|---|---|
| 2025‑09‑23T10:12Z | start | PAC‑021 | contest CRUD |
| 2025‑09‑23T12:35Z | review | PAC‑021 | DoD met; awaiting final check |
| 2025‑09‑23T13:02Z | done | PAC‑021 | blockers satisfied |
| 2025‑09‑23T16:41Z | refusal | PAC‑032 | could not close; blocked by PAC‑011 |
| 2025‑09‑24T09:08Z | unblock | PAC‑032 | PAC‑011 done; resuming |

**Notes**
- “refusal” entries are valuable; they prove the dependency gate is effective (Spec §7.2).
- Keep each line short and objective; the full context can live in the PAC `Notes`.

---

## 4. Metrics (Formulas & Targets)

> Run these over a **3–7 day** window. A handful of completed PACs is enough to see patterns.

### 4.1 Throughput / Day
- **Definition:** number of PACs with `Status = DONE` per day.
- **Formula:** `throughput = count(DONE) / days_observed`
- **Use:** Baseline cadence; compare weeks.

### 4.2 Cycle Time (per PAC)
- **Definition:** elapsed time a PAC spends from `DOING` to `DONE`.
- **Formula:** `cycle_time = DoneAt − StartedAt` (hours).
- **Aggregate:** median, p75, p90.
- **Target:** Consistent median; shrinking tails over time.

### 4.3 WIP (Work‑In‑Progress)
- **Definition:** number of PACs in `DOING` at any point.
- **Guideline:** keep ≤ **3** per person/agent (Field Manual §6.4).
- **Observation:** As WIP ↓, cycle time tends to ↓ (Little’s Law).

### 4.4 Review Latency
- **Definition:** time from entering `REVIEW` to `DONE`.
- **Use:** If long, add a checklist or automate checks in DoD.

### 4.5 Blocked Ratio
- **Definition:** proportion of elapsed time a PAC was blocked by unmet dependencies.
- **Approximation:** count days where a PAC in DOING/REVIEW had at least one blocker not DONE; divide by total days in that state.
- **Use:** High ratio → simplify the graph or reorder work.

### 4.6 Aging DOING
- **Definition:** PACs in `DOING` for more than N hours (suggest **48h**).
- **Use:** Split or escalate; indicates hidden blockers or oversized scope.

---

## 5. Simple Procedure (3–7 Days)

1. **Day 0 (Setup)**
   - Ensure the register follows **Spec §4.1** header and enum values.
   - Decide your observation window (3–7 days).
   - Start the **Operational Log** table (Section 3).

2. **Each Day (5–10 min)**
   - Move ≤ 2–3 PACs into `DOING` (Field Manual §4).
   - Append one line to `Notes` after each status change.
   - Record notable **refusals** to mark `DONE` due to blockers (Spec §7.2).

3. **Day N (Wrap‑up)**
   - Compute metrics (Section 4).
   - Capture 3 short **anecdotes** where PACER governed behavior (e.g., agent refused to close, DoD prevented rework).

4. **Publish**
   - Store results alongside the register (e.g., `docs/pacer/evidence-YYYY‑MM‑DD.md`).
   - Include the metrics table, the Operational Log, and the three anecdotes.

---

## 6. Ready‑to‑Fill Tables

### 6.1 Metrics Summary
| Metric | Value | Notes |
|---|---:|---|
| Throughput/day |  |  window: ___ days |
| Cycle time median (h) |  |  p75:  , p90:  |
| Review latency median (h) |  |  |
| Avg WIP (count in DOING) |  |  |
| Blocked ratio (%) |  |  method: approx |
| Aging DOING (count > 48h) |  |  |

### 6.2 PAC Sample (DONE rows only)
| ID | Title | StartedAt | DoneAt | Cycle (h) | BlockedBy | Notes (summary) |
|---|---|---|---|---:|---|---|

### 6.3 Review Latency Sample
| ID | Entered REVIEW | DoneAt | Review (h) | Notes |
|---|---|---|---:|---|

---

## 7. Optional Visuals (How‑To)

### 7.1 Histogram of Cycle Times
- Bin edges (hours): `<12`, `12–24`, `24–48`, `>48`
- Count DONE PACs per bin and draw a simple bar chart (any tool).

### 7.2 Cumulative Flow (Lite)
- For each day, count PACs in `TODO`, `DOING`, `REVIEW`, `DONE`.
- Plot stacked areas; steady bands indicate healthy flow.

---

## 8. Minimal Scripts (Optional)

> You can compute these with a spreadsheet. If you prefer scripts, below are simple patterns (pseudocode).

**Cycle time per DONE row (pseudo):**
```
for row in rows where Status == DONE:
  cycle_hours = hours_between(row.DoneAt, row.StartedAt)
  collect(cycle_hours)
median = median(cycle_hours)
p75 = percentile(cycle_hours, 75)
p90 = percentile(cycle_hours, 90)
```

**Throughput/day (pseudo):**
```
by_day = group(rows where Status == DONE, day(row.DoneAt))
throughput = sum(count(by_day)) / days_observed
```

**Aging DOING (pseudo):**
```
aging = count(rows where Status == DOING and StartedAt < now - 48h)
```

---

## 9. Linking Evidence to Claims
- **Predictable flow** → show throughput/day and CFD (Spec §6).
- **Dependency enforcement** → list refusal log entries (Spec §7.2).
- **Reduced rework via DoD** → include a before/after anecdote where DoD clarified acceptance (Spec §4.1).
- **Small WIP → shorter cycles** → compare median cycle time on a day with WIP ≤3 vs. a day with WIP >3.

---

## 10. FAQ (Evidence)
**Q: Is 3 days enough?**  
A: For a working note, yes. Add more windows over time to strengthen the trend.

**Q: What if nothing finishes?**  
A: Report WIP and Aging DOING; the signal is that work is too large or blocked. Split or re‑order.

**Q: How do I handle timezone drift?**  
A: Use UTC ISO timestamps everywhere (Spec §6.2).

**Q: Can I claim PACER “improves” delivery?**  
A: Make modest, falsifiable claims: “Reduced cycle time variability after enforcing WIP≤3,” “Zero out‑of‑order completions due to dependency gate.”

---

## 11. References
- **Spec:** [pacer-spec.md](docs/pacer/pacer-spec.md) — §4 Data Model, §6 Status Lifecycle, §7 Dependencies, §9 Validation
- **Field Manual:** [pacer-field-manual.md](docs/pacer/pacer-field-manual.md) — reports & daily ops
- **Quickstart:** [pacer-quickstart.md](docs/pacer/pacer-quickstart.md)
- **Schema:** [pacer.schema.json](docs/pacer/pacer.schema.json)

---

**End of PACER Evidence Pack**