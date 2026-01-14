# t7t

Um gerenciador de tarefas para terminal simples, rapido e sem firulas. Feito para capturar rapido e planejar ainda mais rapido.

> **Vibe-codado com [Claude Code](https://claude.ai/code)**

[Read in English](README.md)

## Funcionalidades

- **4 Listas de Prioridade**: Hoje, Essa Semana, Nao Urgente, Lista Geral
- **Projetos**: Agrupe tarefas relacionadas
- **Tags de Contexto**: Use `@contexto` no nome das tarefas para destaque visual (ex: `@trabalho`, `@pessoal`)
- **Descricoes em Markdown**: Suporte completo a markdown nas descricoes
- **Armazenamento Local**: Todos os dados salvos localmente em JSON

## Instalacao

### Requisitos

- Go 1.21+

### Build

```bash
git clone https://github.com/yourusername/t7t.git
cd t7t
go build -o t7t ./cmd/t7t
```

### Adicionar ao PATH

Apos o build, mova o binario para um diretorio no seu PATH ou adicione o diretorio atual ao PATH.

**Opcao 1: Mover para um local padrao**

```bash
sudo mv t7t /usr/local/bin/
```

**Opcao 2: Adicionar ao seu PATH**

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

Substitua `/caminho/para/t7t` pelo diretorio real onde o binario esta localizado.

## Uso

```bash
t7t
```

### Exemplo de Workflow

Aqui esta um fluxo de trabalho tipico para comecar:

**1. Capture tarefas na Lista Geral**

Pressione `Shift+A` para criar tarefas diretamente na Lista Geral. Essa eh sua caixa de entrada - capture tudo aqui primeiro sem se preocupar com prioridade.

**2. Crie projetos**

Pressione `P` para ir para a visualizacao de Projetos, depois `a` para criar projetos (ex: "Redesign do Site", "Planejamento Q1").

**3. Associe tarefas aos projetos**

Volte para tarefas com `P`, selecione uma tarefa e pressione `p` para vincula-la a um ou mais projetos.

**4. Priorize movendo as tarefas**

Use as teclas numericas para mover tarefas para a lista correta:
- `1` - Hoje
- `2` - Essa Semana
- `3` - Nao Urgente

**5. Complete as tarefas de hoje**

Trabalhe na sua lista "Hoje". Pressione `Espaco` para marcar tarefas como concluidas.

**6. Limpe as tarefas concluidas**

Pressione `D` para deletar todas as tarefas concluidas da lista atual.

**7. Complete e arquive projetos**

Mude para Projetos com `P`, selecione um projeto finalizado e pressione `Espaco` para marca-lo como concluido. Pressione `d` para deleta-lo quando nao for mais necessario.

Pressione `?` a qualquer momento para ver todos os atalhos disponiveis.

## Tags de Contexto

Adicione tags `@contexto` nos nomes das suas tarefas para organizacao visual:

```
Revisar PR @trabalho @urgente
Comprar mantimentos @pessoal
Corrigir bug no auth @projeto-x
```

As tags de contexto sao destacadas em uma cor diferente, facilitando a identificacao.

## Construido Com

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - Framework TUI
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) - Definicoes de estilo
- [Glamour](https://github.com/charmbracelet/glamour) - Renderizacao de Markdown

## Licenca

MIT
