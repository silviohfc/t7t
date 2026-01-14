# t7t

A simple, fast, and no-frills terminal task manager. Built for quick capture and even quicker planning.

> **Vibe-coded with [Claude Code](https://claude.ai/code)**

[Leia em Portugues](README.pt-BR.md)

## Features

- **4 Priority Lists**: Today, This Week, Not Urgent, General
- **Projects**: Group related tasks together
- **Context Tags**: Use `@context` in task names for visual highlighting (e.g., `@work`, `@personal`)
- **Markdown Descriptions**: Full markdown support in task descriptions
- **Local Storage**: All data stored locally in JSON

## Installation

### Requirements

- Go 1.21+

### Build

```bash
git clone https://github.com/yourusername/t7t.git
cd t7t
go build -o t7t ./cmd/t7t
```

### Add to PATH

After building, move the binary to a directory in your PATH or add the current directory to your PATH.

**Option 1: Move to a standard location**

```bash
sudo mv t7t /usr/local/bin/
```

**Option 2: Add to your PATH**

<details>
<summary>Bash (~/.bashrc)</summary>

```bash
echo 'export PATH="$PATH:/path/to/t7t"' >> ~/.bashrc
source ~/.bashrc
```
</details>

<details>
<summary>Zsh (~/.zshrc)</summary>

```bash
echo 'export PATH="$PATH:/path/to/t7t"' >> ~/.zshrc
source ~/.zshrc
```
</details>

<details>
<summary>Fish (~/.config/fish/config.fish)</summary>

```fish
echo 'set -gx PATH $PATH /path/to/t7t' >> ~/.config/fish/config.fish
source ~/.config/fish/config.fish
```
</details>

Replace `/path/to/t7t` with the actual directory where the binary is located.

## Usage

```bash
t7t
```

### Workflow Example

Here's a typical workflow to get you started:

**1. Capture tasks to the General list**

Press `Shift+A` to create tasks directly in the General list. This is your inbox - capture everything here first without worrying about priority.

**2. Create projects**

Press `P` to switch to Projects view, then `a` to create projects (e.g., "Website Redesign", "Q1 Planning").

**3. Associate tasks with projects**

Go back to tasks with `P`, select a task and press `p` to link it to one or more projects.

**4. Prioritize by moving tasks**

Use number keys to move tasks to the right list:
- `1` - Today
- `2` - This Week
- `3` - Not Urgent

**5. Complete today's tasks**

Work through your "Today" list. Press `Space` to mark tasks as done.

**6. Clean up completed tasks**

Press `D` to delete all completed tasks from the current list.

**7. Complete and archive projects**

Switch to Projects with `P`, select a finished project and press `Space` to mark it complete. Press `d` to delete it when no longer needed.

Press `?` at any time to see all available keybindings.

## Context Tags

Add `@context` tags to your task names for visual organization:

```
Review PR @work @urgent
Buy groceries @personal
Fix bug in auth @project-x
```

Context tags are highlighted in a different color, making them easy to spot.

## Built With

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - TUI framework
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) - Style definitions
- [Glamour](https://github.com/charmbracelet/glamour) - Markdown rendering

## License

MIT
