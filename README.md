# PACER Documentation Index
**PACER — Project Actions, Constraints & Evidence Register**  
This index links every PACER document and the most important sections inside each one.

---

## Core Docs
- **Authoritative Spec:** [pacer-spec-v1.md](pacer-spec-v1.md)  
  - §3 File Format • §4 Data Model • §5 Identifiers • §6 Status Lifecycle • §7 Dependencies • §9 Validation • §10 Operations
- **Field Manual (Ops):** [pacer-field-manual.md](pacer-field-manual.md)  
  - §3 Command Lexicon • §4 Daily Ritual • §5 Reports • §7 Playbooks • §8 Agent Integration
- **Quickstart (1‑page):** [pacer-quickstart.md](pacer-quickstart.md)  
  - Create → Start → Review → Done • Dependency Gate • Validation pointers

## Design & Proof
- **Rationale (Design Notes):** [pacer-rationale.md](pacer-rationale.md)  
  - Flat CSV justification • DAG and DoD reasoning • Anti‑patterns • Comparisons
- **Evidence Pack (Proof Template):** [pacer-evidence.md](pacer-evidence.md)  
  - Operational Log • Metrics & formulas • 3–7 day procedure • Ready‑to‑fill tables

## Reference
- **Patterns & FAQ:** [pacer-faq.md](pacer-faq.md)  
  - Epics • Spikes • Splitting • Overrides • Multi‑agent • Multi‑repo • Troubleshooting
- **JSON Schema:** [pacer.schema.json](pacer.schema.json) — Machine validation (CSV→JSON)
- **CSV Template:** [pacer-template.csv](pacer-template.csv) — Header‑only starter

---

## Start Here
1. **Quickstart** for a 2‑minute setup: [pacer-quickstart.md](pacer-quickstart.md)  
2. **Spec** for the exact rules: [pacer-spec-v1.md](pacer-spec-v1.md)  
3. **Field Manual** for daily operation: [pacer-field-manual.md](pacer-field-manual.md)

**Validate (optional):** Convert your CSV to JSON and check with [pacer.schema.json](pacer.schema.json).

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
- Current version: **PACER v1.0** (see [pacer-spec-v1.md](pacer-spec-v1.md) §15).  
- Backward‑compatible extensions use profiles (Spec §13). Document deviations explicitly.

---

**Maintainers:** Keep this index in sync when adding or updating docs.
