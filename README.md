<div align="center">

# PACER
## Project Actions, Constraints & Evidence Register

**AI-First Project Tracking • Deterministic • Machine-Readable**

[![Version](https://img.shields.io/badge/version-1.1-blue.svg)](docs/pacer/pacer-spec.md)
[![Format](https://img.shields.io/badge/format-CSV-green.svg)](docs/pacer/pacer-template.csv)
[![AI Optimized](https://img.shields.io/badge/AI-Optimized-orange.svg)](#machine-readable-artifacts)

*One CSV file. One row per task. Built for AI consumption with human readability as a bonus.*

</div>

---

## What is PACER?

PACER is a **minimal, deterministic project tracking format** optimized for AI/LLM consumption. It uses a single CSV file where each row represents a task (called a **PAC**). The format enforces strict rules for dependencies, status transitions, and completion gates.

### Key Features
- **AI-First Design** - Optimized for machine consumption and automation
- **Deterministic Rules** - Clear state transitions and dependency gates  
- **Single Source of Truth** - One CSV file, no desync
- **Tool Agnostic** - Works with any editor, Git, or automation
- **Version Controlled** - Text-based, diffable, auditable

---

## Quick Start

### 1. Create Your Register
```bash
# Copy the template
cp docs/pacer/pacer-template.csv my-project.csv
```

### 2. Add Your First Task
```csv
ID,Title,Phase,Status,BlockedBy,Assignee,StartedAt,DoneAt,DoD,Notes
PAC-001,Initialize project,Foundation,TODO,,@you,,,Create repo; CI runs green,
```

### 3. Start Working
```bash
# Natural language commands (AI-friendly)
"Start PAC-001"     # → Status=DOING, StartedAt=now
"PAC-001 done"      # → Status=DONE, DoneAt=now (if no blockers)
```

### 4. Track Dependencies
```csv
PAC-002,Add authentication,Auth & DB,TODO,PAC-001,,,,"OAuth flow; user sessions; logout",
```

---

## Core Concepts

### Task Lifecycle
```
TODO → DOING → REVIEW → DONE
  ↑       ↑        ↑
  └───────┴────────┘
    (rollback allowed)
```

### Required Fields
| Field | Type | Description |
|-------|------|-------------|
| `ID` | string | Unique identifier (e.g., `PAC-001`) |
| `Title` | string | Human-readable summary |
| `Phase` | enum | Logical grouping (Foundation, Auth & DB, etc.) |
| `Status` | enum | Current state (TODO, DOING, REVIEW, DONE) |
| `DoD` | string | Definition of Done - objective criteria |

### Dependency Rule
**A task can only be DONE if all its blockers are DONE.**

```csv
# PAC-010 cannot be DONE until PAC-001 and PAC-005 are DONE
PAC-010,Deploy app,Release,TODO,PAC-001,PAC-005,,,,"Live on production; health checks pass",
```

---

## AI/LLM Integration

### Machine-Readable Artifacts

| File | Purpose | Format |
|------|---------|--------|
| [pacer-machine.json](docs/pacer/pacer-machine.json) | Complete specification | JSON |
| [pacer-machine.yaml](docs/pacer/pacer-machine.yaml) | Complete specification | YAML |
| [pacer-commands.jsonl](docs/pacer/machine/pacer-commands.jsonl) | Command patterns | JSONL |
| [pacer.agent.api.json](docs/pacer/machine/pacer.agent.api.json) | API contract | JSON |
| [pacer.agent.grammar.ebnf](docs/pacer/machine/pacer.agent.grammar.ebnf) | Command grammar | EBNF |
| [pacer.agent.contract.json](docs/pacer/machine/pacer.agent.contract.json) | Behavior rules | JSON |

### Natural Language Commands
```bash
# Status transitions
"Start PAC-021"                    # → DOING
"Review PAC-021"                   # → REVIEW  
"PAC-021 done"                     # → DONE (if no blockers)

# Dependencies
"Block PAC-055 on 060,065"         # → BlockedBy=PAC-060,PAC-065
"Unblock PAC-055 remove 065"       # → Remove PAC-065 from BlockedBy

# Assignments & Notes
"Assign PAC-040 to @alex"          # → Assignee=@alex
"Note PAC-021: tests passing"      # → Append to Notes
"DoD PAC-021: server returns 200"  # → Update DoD
```

---

## Documentation

### Human-Readable Docs
| Document | Purpose | Best For |
|----------|---------|----------|
| [**Quickstart**](docs/pacer/pacer-quickstart.md) | 1-page setup guide | Getting started |
| [**Specification**](docs/pacer/pacer-spec.md) | Authoritative rules | Implementation |
| [**Field Manual**](docs/pacer/pacer-field-manual.md) | Daily operations | Day-to-day use |
| [**FAQ**](docs/pacer/pacer-faq.md) | Patterns & troubleshooting | Common questions |
| [**Rationale**](docs/pacer/pacer-rationale.md) | Design decisions | Understanding why |

### Machine-Readable Docs
| Document | Purpose | Best For |
|----------|---------|----------|
| [**Machine Spec (JSON)**](docs/pacer/pacer-machine.json) | Complete AI specification | AI agents |
| [**Machine Spec (YAML)**](docs/pacer/pacer-machine.yaml) | Complete AI specification | AI agents |
| [**Command Patterns**](docs/pacer/machine/pacer-commands.jsonl) | Natural language → actions | Command parsing |
| [**API Contract**](docs/pacer/machine/pacer.agent.api.json) | Method signatures | API implementation |
| [**Grammar**](docs/pacer/machine/pacer.agent.grammar.ebnf) | Command grammar | Parser generation |

---

## Usage Examples

### Basic Workflow
```csv
# 1. Create task
PAC-001,Setup database,Foundation,TODO,,@dev,,,PostgreSQL running; migrations applied,

# 2. Start work  
# Command: "Start PAC-001"
PAC-001,Setup database,Foundation,DOING,,@dev,2025-01-27T10:00:00Z,,PostgreSQL running; migrations applied,

# 3. Complete work
# Command: "PAC-001 done" 
PAC-001,Setup database,Foundation,DONE,,@dev,2025-01-27T10:00:00Z,2025-01-27T11:30:00Z,PostgreSQL running; migrations applied,
```

### Dependency Management
```csv
# PAC-002 depends on PAC-001
PAC-001,Setup database,Foundation,DONE,,@dev,2025-01-27T10:00:00Z,2025-01-27T11:30:00Z,PostgreSQL running; migrations applied,
PAC-002,Add user auth,Auth & DB,TODO,PAC-001,,,,"OAuth flow; user sessions; logout",
```

### AI Agent Integration
```python
# Example AI agent workflow
def process_command(command: str, register: str) -> str:
    if command.startswith("Start "):
        pac_id = extract_pac_id(command)
        return transition_to_doing(pac_id, register)
    elif command.startswith("done "):
        pac_id = extract_pac_id(command)
        return complete_pac(pac_id, register)
    # ... more command patterns
```

---

## Advanced Features

### Custom Phases
```csv
# Add your own phases
PAC-100,Research phase,Research,TODO,,,,"Literature review; prototype built",
```

### Batch Operations
```bash
# Start multiple tasks
"Start PAC-010, PAC-011, PAC-012"

# Complete multiple tasks  
"Done PAC-030..PAC-033"
```

### Emergency Overrides
```csv
# Document exceptions in Notes
PAC-999,Hotfix security issue,Security,TODO,,,,"Override dependency order due to incident",
```

---

## Validation & Quality

### Schema Validation
```bash
# Convert CSV to JSON and validate
csv-to-json pacer.csv | jq . | validate-json docs/pacer/machine/pacer.schema.json
```

### Common Checks
- All required fields present
- IDs are unique and immutable  
- Status transitions are valid
- Dependencies exist and are acyclic
- DoD is objective and testable

---

## Getting Started

1. **Read**: [Quickstart Guide](docs/pacer/pacer-quickstart.md)
2. **Learn**: [Field Manual](docs/pacer/pacer-field-manual.md)  
3. **Understand**: [Specification](docs/pacer/pacer-spec.md)
4. **Validate**: [JSON Schema](docs/pacer/machine/pacer.schema.json)
5. **Test**: [Command Patterns](docs/pacer/machine/pacer-commands.jsonl)
6. **Integrate**: [Agent Contract](docs/pacer/machine/pacer.agent.contract.json)

### For AI Agents
1. **Load**: [Machine Specification](docs/pacer/pacer-machine.json)
2. **Parse**: [Command Grammar](docs/pacer/machine/pacer.agent.grammar.ebnf)
3. **Implement**: [API Contract](docs/pacer/machine/pacer.agent.api.json)

---

## Benefits

### For AI/LLM Systems
- **Deterministic Processing** - Clear rules, no ambiguity
- **Natural Language Commands** - Easy to parse and execute
- **Structured Data** - Machine-readable format with validation
- **Atomic Operations** - Safe concurrent access
- **Extensible Schema** - Preserve unknown fields

### For Devs
- **Single Source of Truth** - No scattered information
- **Clear Dependencies** - Visual blocking relationships
- **Objective Completion** - DoD eliminates "done" ambiguity
- **Version Control Friendly** - Text-based, diffable
- **Tool Agnostic** - Works with any editor or system

---

## Contributing

### Documentation
- Keep human docs in `docs/pacer/`
- Keep machine specs in `docs/pacer/machine/`
- Update this README when adding features

### Format Changes
- Follow [Specification](docs/pacer/pacer-spec.md) for normative rules
- Use [Profiles](docs/pacer/pacer-spec.md#13-profiles) for customizations
- Document deviations explicitly

---

## License

PACER format and documentation are open source. Use freely in your projects.

---

<div align="center">

**Built for AI • Simple for Humans • Deterministic by Design**

[Get Started](docs/pacer/pacer-quickstart.md) • [Read the Spec](docs/pacer/pacer-spec.md) • [Machine Docs](docs/pacer/machine/)

</div>
