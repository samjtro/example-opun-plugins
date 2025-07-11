# Opun Examples

This directory contains simple, clear examples of each Opun component type.

## Structure

```
examples/
├── plugin/          # Go plugin example
├── action/          # Simple command wrapper
├── workflow/        # Sequential workflow
└── prompt/          # Prompt template
```

## 1. Plugin Example (`plugin/`)

**File**: `code-analyzer` - A Go plugin that analyzes code complexity

This example demonstrates:
- Command registration (`/analyze`)
- Tool registration (`count-lines`)
- Configuration schema
- Input/output validation

To install:
```bash
opun add --plugin --path examples/plugin
```

Usage:
```bash
# In chat:
/analyze main.go
```

## 2. Action Example (`action/`)

**File**: `find-todos.yaml` - Searches for TODO comments in code

This example shows:
- Simple command wrapping (uses `ripgrep`)
- Provider compatibility
- Basic configuration

To install:
```bash
opun add --action --path examples/action/find-todos.yaml
```

Usage:
```bash
# In chat:
/find-todos
```

## 3. Workflow Example (`workflow/`)

**File**: `test-and-fix.yaml` - Runs tests, analyzes failures, suggests fixes

This example demonstrates:
- Sequential agent execution
- Output chaining between agents
- Conditional execution
- Mixed provider usage (Claude + Gemini)
- Variable substitution

To install:
```bash
opun add --workflow --path examples/workflow/test-and-fix.yaml
```

Usage:
```bash
opun run test-and-fix
# Or in chat: /testfix
```

## 4. Prompt Example (`prompt/`)

**File**: `code-review.md` - Template for comprehensive code reviews

This example shows:
- Prompt metadata (variables, tags)
- Conditional templating
- Variable substitution
- Structured output format

To install:
```bash
opun add --prompt --path examples/prompt/code-review.md
```

Usage:
```bash
opun prompt code-review --file-path main.go
```

## Quick Start

1. **Install all examples**:
   ```bash
   # From the opun root directory
   opun add --plugin --path examples/plugin
   opun add --action --path examples/action/find-todos.yaml
   opun add --workflow --path examples/workflow/test-and-fix.yaml
   opun add --prompt --path examples/prompt/code-review.md
   ```

2. **List installed components**:
   ```bash
   opun list --all
   ```

3. **Try them out**:
   ```bash
   # Start a chat session
   opun chat
   
   # Use the action
   /find-todos
   
   # Use the plugin
   /analyze main.go
   
   # Exit chat and run workflow
   opun run test-and-fix
   ```

## Key Concepts

### Actions vs Plugins

- **Actions**: Simple, stateless command wrappers (YAML only)
- **Plugins**: Complex extensions with state, lifecycle, and multiple tools

### Workflows

- Orchestrate multiple AI agents
- Pass outputs between agents
- Support conditional execution
- Can use different providers for different steps

### Prompts

- Reusable templates with variables
- Support conditionals and loops
- Stored in PromptGarden
- Can be referenced by workflows

## Testing

Each example is designed to work out of the box:

1. The plugin can analyze its own source file
2. The action will find TODOs in the codebase
3. The workflow can test any Go project
4. The prompt can review any code file

## Next Steps

- Modify these examples for your use cases
- Create your own components based on these patterns
- Combine components (e.g., use the prompt in a workflow)
- Share your creations with the community