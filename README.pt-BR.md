# t7t

Um gerenciador de tarefas para terminal simples, rápido e sem firulas. Feito para capturar rápido e planejar ainda mais rápido.

> **Vibe-codado com [Claude Code](https://claude.ai/code)**

[Read in English](README.md)

## Funcionalidades

- **4 Listas de Prioridade**: Hoje, Essa Semana, Não Urgente, Lista Geral
- **Projetos**: Agrupe tarefas relacionadas
- **Tags de Contexto**: Use `@contexto` no nome das tarefas para destaque visual (ex: `@trabalho`, `@pessoal`)
- **Descrições em Markdown**: Suporte completo a markdown nas descrições
- **Armazenamento Local**: Todos os dados salvos localmente em JSON
- **Suporte Multi-idioma**: Disponível em Português (Brasil) e Inglês

## Instalação

### Requisitos

- Go 1.21+

### Build

```bash
git clone https://github.com/silviohfc/t7t.git
cd t7t
go build -o t7t .
```

### Adicionar ao PATH

Após o build, mova o binário para um diretório no seu PATH ou adicione o diretório atual ao PATH.

**Opção 1: Mover para um local padrão**

```bash
sudo mv t7t /usr/local/bin/
```

**Opção 2: Adicionar ao seu PATH**

<details>
<summary>Bash (~/.bashrc)</summary>

```bash
echo 'export PATH="$PATH:/caminho/para/t7t"' >> ~/.bashrc
source ~/.bashrc
```
</details>

<details>
<summary>Zsh (~/.zshrc)</summary>

```bash
echo 'export PATH="$PATH:/caminho/para/t7t"' >> ~/.zshrc
source ~/.zshrc
```
</details>

<details>
<summary>Fish (~/.config/fish/config.fish)</summary>

```fish
echo 'set -gx PATH $PATH /caminho/para/t7t' >> ~/.config/fish/config.fish
source ~/.config/fish/config.fish
```
</details>

Substitua `/caminho/para/t7t` pelo diretório real onde o binário está localizado.

## Uso

```bash
t7t
```

### Exemplo de Workflow

Aqui está um fluxo de trabalho típico para começar:

**1. Capture tarefas na Lista Geral**

Pressione `Shift+A` para criar tarefas diretamente na Lista Geral. Essa é sua caixa de entrada - capture tudo aqui primeiro sem se preocupar com prioridade.

**2. Crie projetos**

Pressione `P` para ir para a visualização de Projetos, depois `a` para criar projetos (ex: "Redesign do Site", "Planejamento Q1").

**3. Associe tarefas aos projetos**

Volte para tarefas com `P`, selecione uma tarefa e pressione `p` para vinculá-la a um ou mais projetos.

**4. Priorize movendo as tarefas**

Use as teclas numéricas para mover tarefas para a lista correta:
- `1` - Hoje
- `2` - Essa Semana
- `3` - Não Urgente

**5. Complete as tarefas de hoje**

Trabalhe na sua lista "Hoje". Pressione `Espaço` para marcar tarefas como concluídas.

**6. Limpe as tarefas concluídas**

Pressione `D` para deletar todas as tarefas concluídas da lista atual.

**7. Complete e arquive projetos**

Mude para Projetos com `P`, selecione um projeto finalizado e pressione `Espaço` para marcá-lo como concluído. Pressione `d` para deletá-lo quando não for mais necessário.

Pressione `?` a qualquer momento para ver todos os atalhos disponíveis.

## Tags de Contexto

Adicione tags `@contexto` nos nomes das suas tarefas para organização visual:

```
Revisar PR @trabalho @urgente
Comprar mantimentos @pessoal
Corrigir bug no auth @projeto-x
```

As tags de contexto são destacadas em uma cor diferente, facilitando a identificação.

## Idioma

t7t suporta múltiplos idiomas. Pressione `L` (Shift+L) para abrir o modal de seleção de idioma e escolha entre:

- **Português (Brasil)**
- **English**

Sua preferência de idioma é salva automaticamente em `~/.t7t/language.json`.

## Construído Com

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - Framework TUI
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) - Definições de estilo
- [Glamour](https://github.com/charmbracelet/glamour) - Renderização de Markdown

## Licença

MIT
