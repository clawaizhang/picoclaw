<div align="center">
  <img src="assets/logo.webp" alt="PicoClaw" width="512">

  <h1>PicoClaw: Assistente de IA Ultra-Eficiente em Go</h1>

  <h3>Hardware de $10 · <10MB de RAM · Boot em <1s · 皮皮虾，我们走！</h3>
  <p>
    <img src="https://img.shields.io/badge/Go-1.25+-00ADD8?style=flat&logo=go&logoColor=white" alt="Go">
    <img src="https://img.shields.io/badge/Arch-x86__64%2C%20ARM64%2C%20MIPS%2C%20RISC--V%2C%20LoongArch-blue" alt="Hardware">
    <img src="https://img.shields.io/badge/license-MIT-green" alt="License">
    <br>
    <a href="https://picoclaw.io"><img src="https://img.shields.io/badge/Website-picoclaw.io-blue?style=flat&logo=google-chrome&logoColor=white" alt="Website"></a>
    <a href="https://docs.picoclaw.io/"><img src="https://img.shields.io/badge/Docs-Official-007acc?style=flat&logo=read-the-docs&logoColor=white" alt="Docs"></a>
    <a href="https://deepwiki.com/sipeed/picoclaw"><img src="https://img.shields.io/badge/Wiki-DeepWiki-FFA500?style=flat&logo=wikipedia&logoColor=white" alt="Wiki"></a>
    <br>
    <a href="https://x.com/SipeedIO"><img src="https://img.shields.io/badge/X_(Twitter)-SipeedIO-black?style=flat&logo=x&logoColor=white" alt="Twitter"></a>
    <a href="./assets/wechat.png"><img src="https://img.shields.io/badge/WeChat-Group-41d56b?style=flat&logo=wechat&logoColor=white"></a>
    <a href="https://discord.gg/V4sAZ9XWpN"><img src="https://img.shields.io/badge/Discord-Community-4c60eb?style=flat&logo=discord&logoColor=white" alt="Discord"></a>
  </p>

[中文](README.zh.md) | [日本語](README.ja.md) | **Português** | [Tiếng Việt](README.vi.md) | [Français](README.fr.md) | [Italiano](README.it.md) | [Bahasa Indonesia](README.id.md) | [English](README.md)

</div>

---

> **PicoClaw** é um projeto open-source independente iniciado pela [Sipeed](https://sipeed.com). É escrito inteiramente em **Go** — não é um fork do OpenClaw, NanoBot ou qualquer outro projeto.

🦐 PicoClaw é um assistente pessoal de IA ultra-leve inspirado no [NanoBot](https://github.com/HKUDS/nanobot), reescrito do zero em Go por meio de um processo de auto-inicialização (self-bootstrapping), onde o próprio agente de IA conduziu toda a migração de arquitetura e otimização de código.

⚡️ Roda em hardware de $10 com <10MB de RAM: Isso é 99% menos memória que o OpenClaw e 98% mais barato que um Mac mini!

<table align="center">
  <tr align="center">
    <td align="center" valign="top">
      <p align="center">
        <img src="assets/picoclaw_mem.gif" width="360" height="240">
      </p>
    </td>
    <td align="center" valign="top">
      <p align="center">
        <img src="assets/licheervnano.png" width="400" height="240">
      </p>
    </td>
  </tr>
</table>

> [!CAUTION]
> **🚨 DECLARAÇÃO DE SEGURANÇA & CANAIS OFICIAIS**
>
> * **SEM CRIPTOMOEDAS:** O PicoClaw **NÃO** possui nenhum token/moeda oficial. Todas as alegações no `pump.fun` ou outras plataformas de negociação são **GOLPES**.
>
> * **DOMÍNIO OFICIAL:** O **ÚNICO** site oficial é o **[picoclaw.io](https://picoclaw.io)**, e o site da empresa é o **[sipeed.com](https://sipeed.com)**
> * **Aviso:** Muitos domínios `.ai/.org/.com/.net/...` foram registrados por terceiros.
> * **Aviso:** O PicoClaw está em fase inicial de desenvolvimento e pode ter problemas de segurança de rede não resolvidos. Não implante em ambientes de produção antes da versão v1.0.
> * **Nota:** O PicoClaw recentemente fez merge de muitos PRs, o que pode resultar em maior consumo de memória (10–20MB) nas versões mais recentes. Planejamos priorizar a otimização de recursos assim que o conjunto de funcionalidades estiver estável.

## 📢 Novidades

2026-03-17 🚀 **v0.2.3 Lançado!** Interface de bandeja do sistema (Windows & Linux), rastreamento de status de sub-agentes (`spawn_status`), hot-reload experimental do gateway, portões de segurança para cron e 2 correções de segurança. PicoClaw agora com **25K ⭐**!

2026-03-09 🎉 **v0.2.1 — Maior atualização até agora!** Suporte ao protocolo MCP, 4 novos canais (Matrix/IRC/WeCom/Discord Proxy), 3 novos provedores (Kimi/Minimax/Avian), pipeline de visão, armazenamento de memória JSONL e roteamento de modelos.

2026-02-28 📦 **v0.2.0** lançado com suporte a Docker Compose e launcher Web UI.

2026-02-26 🎉 PicoClaw atingiu **20K stars** em apenas 17 dias! Orquestração automática de canais e interfaces de capacidade implementadas.

<details>
<summary>Novidades anteriores...</summary>

2026-02-16 🎉 PicoClaw atingiu 12K stars em uma semana! Papéis de maintainers da comunidade e [roadmap](ROADMAP.md) publicados oficialmente.

2026-02-13 🎉 PicoClaw atingiu 5000 stars em 4 dias! Roadmap do Projeto e Grupo de Desenvolvedores em preparação.

2026-02-09 🎉 **PicoClaw Lançado!** Construído em 1 dia para trazer Agentes de IA para hardware de $10 com <10MB de RAM. 🦐 PicoClaw, Partiu!

</details>

## ✨ Funcionalidades

🪶 **Ultra-Leve**: Consumo de memória <10MB — 99% menor que o OpenClaw para funcionalidades essenciais.*

💰 **Custo Mínimo**: Eficiente o suficiente para rodar em hardware de $10 — 98% mais barato que um Mac mini.

⚡️ **Inicialização Relâmpago**: Tempo de inicialização 400X mais rápido, boot em <1 segundo mesmo em CPU single-core de 0.6GHz.

🌍 **Portabilidade Real**: Um único binário auto-contido para RISC-V, ARM, MIPS e x86. Um clique e já era!

🤖 **Auto-Construído por IA**: Implementação nativa em Go de forma autônoma — 95% do núcleo gerado pelo Agente com refinamento humano no loop.

🔌 **Suporte MCP**: Integração nativa com o [Model Context Protocol](https://modelcontextprotocol.io/) — conecte qualquer servidor MCP para estender as capacidades do agente.

👁️ **Pipeline de Visão**: Envie imagens e arquivos diretamente ao agente — codificação base64 automática para LLMs multimodais.

🧠 **Roteamento Inteligente**: Roteamento de modelos baseado em regras — consultas simples vão para modelos leves, economizando custos de API.

_*Versões recentes podem usar 10–20MB devido a merges rápidos de funcionalidades. Otimização de recursos está planejada. Comparação de inicialização baseada em benchmarks de single-core a 0.8GHz (veja tabela abaixo)._

|                               | OpenClaw      | NanoBot                  | **PicoClaw**                              |
| ----------------------------- | ------------- | ------------------------ | ----------------------------------------- |
| **Linguagem**                 | TypeScript    | Python                   | **Go**                                    |
| **RAM**                       | >1GB          | >100MB                   | **< 10MB***                               |
| **Inicialização**</br>(CPU 0.8GHz) | >500s         | >30s                     | **<1s**                                   |
| **Custo**                     | Mac Mini $599 | Maioria dos SBC Linux </br>~$50 | **Qualquer placa Linux**</br>**A partir de $10** |

<img src="assets/compare.jpg" alt="PicoClaw" width="512">

> 📋 **[Lista de Compatibilidade de Hardware](docs/hardware-compatibility.md)** — Veja todas as placas testadas, de RISC-V de $5 a Raspberry Pi e telefones Android. Sua placa não está listada? Envie um PR!

## 🦾 Demonstração

### 🛠️ Fluxos de Trabalho Padrão do Assistente

<table align="center">
  <tr align="center">
    <th><p align="center">🧩 Engenharia Full-Stack</p></th>
    <th><p align="center">🗂️ Gerenciamento de Logs & Planejamento</p></th>
    <th><p align="center">🔎 Busca Web & Aprendizado</p></th>
  </tr>
  <tr>
    <td align="center"><p align="center"><img src="assets/picoclaw_code.gif" width="240" height="180"></p></td>
    <td align="center"><p align="center"><img src="assets/picoclaw_memory.gif" width="240" height="180"></p></td>
    <td align="center"><p align="center"><img src="assets/picoclaw_search.gif" width="240" height="180"></p></td>
  </tr>
  <tr>
    <td align="center">Desenvolver • Implantar • Escalar</td>
    <td align="center">Agendar • Automatizar • Memorizar</td>
    <td align="center">Descobrir • Analisar • Tendências</td>
  </tr>
</table>

### 📱 Rode em celulares Android antigos

Dê uma segunda vida ao seu celular de dez anos atrás! Transforme-o em um assistente de IA inteligente com o PicoClaw. Início rápido:

1. **Instale o [Termux](https://github.com/termux/termux-app)** (Baixe em [GitHub Releases](https://github.com/termux/termux-app/releases), ou busque no F-Droid / Google Play).
2. **Execute os comandos**

```bash
# Baixe a versão mais recente em https://github.com/sipeed/picoclaw/releases
wget https://github.com/sipeed/picoclaw/releases/latest/download/picoclaw_Linux_arm64.tar.gz
tar xzf picoclaw_Linux_arm64.tar.gz
pkg install proot
termux-chroot ./picoclaw onboard   # chroot fornece um layout padrão do sistema de arquivos Linux
```

Depois siga as instruções na seção "Início Rápido" para completar a configuração!

<img src="assets/termux.jpg" alt="PicoClaw" width="512">

### 🐜 Implantação Inovadora com Baixo Consumo

O PicoClaw pode ser implantado em praticamente qualquer dispositivo Linux!

- $9.9 [LicheeRV-Nano](https://www.aliexpress.com/item/1005006519668532.html) versão E(Ethernet) ou W(WiFi6), para Assistente Doméstico Minimalista
- $30~50 [NanoKVM](https://www.aliexpress.com/item/1005007369816019.html), ou $100 [NanoKVM-Pro](https://www.aliexpress.com/item/1005010048471263.html) para Manutenção Automatizada de Servidores
- $50 [MaixCAM](https://www.aliexpress.com/item/1005008053333693.html) ou $100 [MaixCAM2](https://www.kickstarter.com/projects/zepan/maixcam2-build-your-next-gen-4k-ai-camera) para Monitoramento Inteligente

<https://private-user-images.githubusercontent.com/83055338/547056448-e7b031ff-d6f5-4468-bcca-5726b6fecb5c.mp4>

🌟 Mais cenários de implantação aguardam você!

## 📦 Instalação

### Baixar de picoclaw.io (Recomendado)

Visite **[picoclaw.io](https://picoclaw.io)** — o site oficial detecta automaticamente sua plataforma e oferece download com um clique. Sem necessidade de escolher manualmente a arquitetura.

### Baixar binário pré-compilado

Alternativamente, baixe o binário para sua plataforma na página de [GitHub Releases](https://github.com/sipeed/picoclaw/releases).

### Compilar a partir do código-fonte (para desenvolvimento)

```bash
git clone https://github.com/sipeed/picoclaw.git

cd picoclaw
make deps

# Build, sem necessidade de instalar
make build

# Build para múltiplas plataformas
make build-all

# Build para Raspberry Pi Zero 2 W (32-bit: make build-linux-arm; 64-bit: make build-linux-arm64)
make build-pi-zero

# Build e Instalar
make install
```

**Raspberry Pi Zero 2 W:** Use o binário correspondente ao seu SO: Raspberry Pi OS 32-bit → `make build-linux-arm`; 64-bit → `make build-linux-arm64`. Ou execute `make build-pi-zero` para compilar ambos.

## 📚 Documentação

Para guias detalhados, consulte a documentação abaixo. Este README cobre apenas o início rápido.

| Tópico | Descrição |
|--------|-----------|
| 🐳 [Docker & Início Rápido](docs/pt-br/docker.md) | Configuração Docker Compose, modos Launcher/Agent, configuração de Início Rápido |
| 💬 [Apps de Chat](docs/pt-br/chat-apps.md) | Telegram, Discord, WhatsApp, Matrix, QQ, Slack, IRC, DingTalk, LINE, Feishu, WeCom e mais |
| ⚙️ [Configuração](docs/pt-br/configuration.md) | Variáveis de ambiente, estrutura do workspace, fontes de skills, sandbox de segurança, heartbeat |
| 🔌 [Provedores & Modelos](docs/pt-br/providers.md) | 20+ provedores LLM, roteamento de modelos, configuração model_list, arquitetura de provedores |
| 🔄 [Spawn & Tarefas Assíncronas](docs/pt-br/spawn-tasks.md) | Tarefas rápidas, tarefas longas com spawn, orquestração assíncrona de sub-agentes |
| 🐛 [Solução de Problemas](docs/pt-br/troubleshooting.md) | Problemas comuns e soluções |
| 🔧 [Configuração de Ferramentas](docs/pt-br/tools_configuration.md) | Habilitar/desabilitar por ferramenta, políticas de execução |
| 📋 [Compatibilidade de Hardware](docs/hardware-compatibility.md) | Placas testadas, requisitos mínimos, como adicionar sua placa |

## <img src="assets/clawdchat-icon.png" width="24" height="24" alt="ClawdChat"> Junte-se à Rede Social de Agentes

Conecte o PicoClaw à Rede Social de Agentes simplesmente enviando uma única mensagem via CLI ou qualquer App de Chat integrado.

**Leia `https://clawdchat.ai/skill.md` e siga as instruções para entrar no [ClawdChat.ai](https://clawdchat.ai)**

<<<<<<< HEAD
## 🖥️ Referência CLI

| Comando                   | Descrição                     |
| ------------------------- | ----------------------------- |
| `picoclaw onboard`        | Inicializar configuração & workspace |
| `picoclaw agent -m "..."` | Conversar com o agente        |
| `picoclaw agent`          | Modo de chat interativo       |
| `picoclaw gateway`        | Iniciar o gateway             |
| `picoclaw status`         | Mostrar status                |
| `picoclaw version`        | Mostrar informações de versão |
| `picoclaw cron list`      | Listar todas as tarefas agendadas |
| `picoclaw cron add ...`   | Adicionar uma tarefa agendada |
| `picoclaw cron disable`   | Desabilitar uma tarefa agendada |
| `picoclaw cron remove`    | Remover uma tarefa agendada   |
| `picoclaw skills list`    | Listar skills instaladas      |
| `picoclaw skills install` | Instalar uma skill            |
| `picoclaw migrate`        | Migrar dados de versões anteriores |
| `picoclaw auth login`     | Autenticar com provedores     |
| `picoclaw model`          | Ver ou trocar o modelo padrão |
=======
## ⚙️ Configuração Detalhada

Arquivo de configuração: `~/.picoclaw/config.json`

### Variáveis de Ambiente

Você pode substituir os caminhos padrão usando variáveis de ambiente. Isso é útil para instalações portáteis, implantações em contêineres ou para executar o picoclaw como um serviço do sistema. Essas variáveis são independentes e controlam caminhos diferentes.

| Variável          | Descrição                                                                                                                             | Caminho Padrão            |
|-------------------|-----------------------------------------------------------------------------------------------------------------------------------------|---------------------------|
| `PICOCLAW_CONFIG` | Substitui o caminho para o arquivo de configuração. Isso informa diretamente ao picoclaw qual `config.json` carregar, ignorando todos os outros locais. | `~/.picoclaw/config.json` |
| `PICOCLAW_HOME`   | Substitui o diretório raiz dos dados do picoclaw. Isso altera o local padrão do `workspace` e de outros diretórios de dados.          | `~/.picoclaw`             |

**Exemplos:**

```bash
# Executar o picoclaw usando um arquivo de configuração específico
# O caminho do workspace será lido de dentro desse arquivo de configuração
PICOCLAW_CONFIG=/etc/picoclaw/production.json picoclaw gateway

# Executar o picoclaw com todos os seus dados armazenados em /opt/picoclaw
# A configuração será carregada do ~/.picoclaw/config.json padrão
# O workspace será criado em /opt/picoclaw/workspace
PICOCLAW_HOME=/opt/picoclaw picoclaw agent

# Use ambos para uma configuração totalmente personalizada
PICOCLAW_HOME=/srv/picoclaw PICOCLAW_CONFIG=/srv/picoclaw/main.json picoclaw gateway
```

### Estrutura do Workspace

O PicoClaw armazena dados no workspace configurado (padrão: `~/.picoclaw/workspace`):

```
~/.picoclaw/workspace/
├── sessions/          # Sessoes de conversa e historico
├── memory/            # Memoria de longo prazo (MEMORY.md)
├── state/             # Estado persistente (ultimo canal, etc.)
├── cron/              # Banco de dados de tarefas agendadas
├── skills/            # Skills personalizadas
├── AGENT.md           # Definicao estruturada do agente e prompt do sistema
├── HEARTBEAT.md       # Prompts de tarefas periodicas (verificado a cada 30 min)
├── SOUL.md            # Alma do Agente
└── ...
```

### 🔒 Sandbox de Segurança

O PicoClaw roda em um ambiente sandbox por padrão. O agente so pode acessar arquivos e executar comandos dentro do workspace configurado.

#### Configuração Padrão

```json
{
  "agents": {
    "defaults": {
      "workspace": "~/.picoclaw/workspace",
      "restrict_to_workspace": true
    }
  }
}
```

| Opção | Padrão | Descrição |
|-------|--------|-----------|
| `workspace` | `~/.picoclaw/workspace` | Diretório de trabalho do agente |
| `restrict_to_workspace` | `true` | Restringir acesso de arquivos/comandos ao workspace |

#### Ferramentas Protegidas

Quando `restrict_to_workspace: true`, as seguintes ferramentas são restritas ao sandbox:

| Ferramenta | Função | Restrição |
|------------|--------|-----------|
| `read_file` | Ler arquivos | Apenas arquivos dentro do workspace |
| `write_file` | Escrever arquivos | Apenas arquivos dentro do workspace |
| `list_dir` | Listar diretorios | Apenas diretorios dentro do workspace |
| `edit_file` | Editar arquivos | Apenas arquivos dentro do workspace |
| `append_file` | Adicionar a arquivos | Apenas arquivos dentro do workspace |
| `exec` | Executar comandos | Caminhos dos comandos devem estar dentro do workspace |

#### Proteção Adicional do Exec

Mesmo com `restrict_to_workspace: false`, a ferramenta `exec` bloqueia estes comandos perigosos:

* `rm -rf`, `del /f`, `rmdir /s` — Exclusão em massa
* `format`, `mkfs`, `diskpart` — Formatação de disco
* `dd if=` — Criação de imagem de disco
* Escrita em `/dev/sd[a-z]` — Escrita direta no disco
* `shutdown`, `reboot`, `poweroff` — Desligamento do sistema
* Fork bomb `:(){ :|:& };:`

#### Exemplos de Erro

```
[ERROR] tool: Tool execution failed
{tool=exec, error=Command blocked by safety guard (path outside working dir)}
```

```
[ERROR] tool: Tool execution failed
{tool=exec, error=Command blocked by safety guard (dangerous pattern detected)}
```

#### Desabilitar Restrições (Risco de Segurança)

Se você precisa que o agente acesse caminhos fora do workspace:

**Método 1: Arquivo de configuração**

```json
{
  "agents": {
    "defaults": {
      "restrict_to_workspace": false
    }
  }
}
```

**Método 2: Variável de ambiente**

```bash
export PICOCLAW_AGENTS_DEFAULTS_RESTRICT_TO_WORKSPACE=false
```

> ⚠️ **Aviso**: Desabilitar esta restrição permite que o agente acesse qualquer caminho no seu sistema. Use com cuidado apenas em ambientes controlados.

#### Consistência do Limite de Segurança

A configuração `restrict_to_workspace` se aplica consistentemente em todos os caminhos de execução:

| Caminho de Execução | Limite de Segurança |
|----------------------|---------------------|
| Agente Principal | `restrict_to_workspace` ✅ |
| Subagente / Spawn | Herda a mesma restrição ✅ |
| Tarefas Heartbeat | Herda a mesma restrição ✅ |

Todos os caminhos compartilham a mesma restrição de workspace — nao há como contornar o limite de segurança por meio de subagentes ou tarefas agendadas.

### Heartbeat (Tarefas Periódicas)

O PicoClaw pode executar tarefas periódicas automaticamente. Crie um arquivo `HEARTBEAT.md` no seu workspace:

```markdown
# Tarefas Periodicas

- Verificar meu email para mensagens importantes
- Revisar minha agenda para proximos eventos
- Verificar a previsao do tempo
```

O agente lerá este arquivo a cada 30 minutos (configurável) e executará as tarefas usando as ferramentas disponíveis.

#### Tarefas Assincronas com Spawn

Para tarefas de longa duração (busca web, chamadas de API), use a ferramenta `spawn` para criar um **subagente**:

```markdown
# Tarefas Periódicas

## Tarefas Rápidas (resposta direta)
- Informar hora atual

## Tarefas Longas (usar spawn para async)
- Buscar notícias de IA na web e resumir
- Verificar email e reportar mensagens importantes
```

**Comportamentos principais:**

| Funcionalidade | Descrição |
|----------------|-----------|
| **spawn** | Cria subagente assíncrono, não bloqueia o heartbeat |
| **Contexto independente** | Subagente tem seu próprio contexto, sem histórico de sessão |
| **Ferramenta message** | Subagente se comunica diretamente com o usuário via ferramenta message |
| **Não-bloqueante** | Após o spawn, o heartbeat continua para a próxima tarefa |

#### Como Funciona a Comunicação do Subagente

```
Heartbeat dispara
    ↓
Agente lê HEARTBEAT.md
    ↓
Para tarefa longa: spawn subagente
    ↓                           ↓
Continua próxima tarefa    Subagente trabalha independentemente
    ↓                           ↓
Todas tarefas concluídas   Subagente usa ferramenta "message"
    ↓                           ↓
Responde HEARTBEAT_OK      Usuário recebe resultado diretamente
```

O subagente tem acesso às ferramentas (message, web_search, etc.) e pode se comunicar com o usuário independentemente sem passar pelo agente principal.

**Configuração:**

```json
{
  "heartbeat": {
    "enabled": true,
    "interval": 30
  }
}
```

| Opção | Padrão | Descrição |
|-------|--------|-----------|
| `enabled` | `true` | Habilitar/desabilitar heartbeat |
| `interval` | `30` | Intervalo de verificação em minutos (min: 5) |

**Variáveis de ambiente:**

* `PICOCLAW_HEARTBEAT_ENABLED=false` para desabilitar
* `PICOCLAW_HEARTBEAT_INTERVAL=60` para alterar o intervalo

### Provedores

> [!NOTE]
> O Groq fornece transcrição de voz gratuita via Whisper. Se configurado, mensagens de áudio de qualquer canal serão automaticamente transcritas no nível do agente.

| Provedor | Finalidade | Obter API Key |
| --- | --- | --- |
| `gemini` | LLM (Gemini direto) | [aistudio.google.com](https://aistudio.google.com) |
| `zhipu` | LLM (Zhipu direto) | [bigmodel.cn](bigmodel.cn) |
| `volcengine`             | LLM(Volcengine direto)                   | [volcengine.com](https://www.volcengine.com/activity/codingplan?utm_campaign=PicoClaw&utm_content=PicoClaw&utm_medium=devrel&utm_source=OWO&utm_term=PicoClaw)           |
| `openrouter` (Em teste) | LLM (recomendado, acesso a todos os modelos) | [openrouter.ai](https://openrouter.ai) |
| `anthropic` (Em teste) | LLM (Claude direto) | [console.anthropic.com](https://console.anthropic.com) |
| `openai` (Em teste) | LLM (GPT direto) | [platform.openai.com](https://platform.openai.com) |
| `deepseek` (Em teste) | LLM (DeepSeek direto) | [platform.deepseek.com](https://platform.deepseek.com) |
| `qwen` | Alibaba Qwen | [dashscope.console.aliyun.com](https://dashscope.console.aliyun.com) |
| `cerebras` | Cerebras | [cerebras.ai](https://cerebras.ai) |
| `groq` | LLM + **Transcrição de voz** (Whisper) | [console.groq.com](https://console.groq.com) |

<details>
<summary><b>Configuração Zhipu</b></summary>

**1. Obter API key**

* Obtenha a [API key](https://bigmodel.cn/usercenter/proj-mgmt/apikeys)

**2. Configurar**

```json
{
  "agents": {
    "defaults": {
      "workspace": "~/.picoclaw/workspace",
      "model": "glm-4.7",
      "max_tokens": 8192,
      "temperature": 0.7,
      "max_tool_iterations": 20
    }
  },
  "providers": {
    "zhipu": {
      "api_key": "Sua API Key",
      "api_base": "https://open.bigmodel.cn/api/paas/v4"
    }
  }
}
```

**3. Executar**

```bash
picoclaw agent -m "Ola, como vai?"
```

</details>

<details>
<summary><b>Exemplo de configuraçao completa</b></summary>

```json
{
  "agents": {
    "defaults": {
      "model": "anthropic/claude-opus-4-5"
    }
  },
  "providers": {
    "openrouter": {
      "api_key": "sk-or-v1-xxx"
    },
    "groq": {
      "api_key": "gsk_xxx"
    }
  },
  "channels": {
    "telegram": {
      "enabled": true,
      "token": "123456:ABC...",
      "allow_from": ["123456789"]
    },
    "discord": {
      "enabled": true,
      "token": "",
      "allow_from": [""]
    },
    "whatsapp": {
      "enabled": false
    },
    "feishu": {
      "enabled": false,
      "app_id": "cli_xxx",
      "app_secret": "xxx",
      "encrypt_key": "",
      "verification_token": "",
      "allow_from": []
    },
    "qq": {
      "enabled": false,
      "app_id": "",
      "app_secret": "",
      "allow_from": []
    }
  },
  "tools": {
    "web": {
      "brave": {
        "enabled": false,
        "api_key": "BSA...",
        "max_results": 5
      },
      "duckduckgo": {
        "enabled": true,
        "max_results": 5
      }
    },
    "cron": {
      "exec_timeout_minutes": 5
    }
  },
  "heartbeat": {
    "enabled": true,
    "interval": 30
  }
}
```

</details>

### Configuração de Modelo (model_list)

> **Novidade!** PicoClaw agora usa uma abordagem de configuração **centrada no modelo**. Basta especificar o formato `fornecedor/modelo` (ex: `zhipu/glm-4.7`) para adicionar novos provedores—**nenhuma alteração de código necessária!**

Este design também possibilita o **suporte multi-agent** com seleção flexível de provedores:

- **Diferentes agentes, diferentes provedores** : Cada agente pode usar seu próprio provedor LLM
- **Modelos de fallback** : Configure modelos primários e de reserva para resiliência
- **Balanceamento de carga** : Distribua solicitações entre múltiplos endpoints
- **Configuração centralizada** : Gerencie todos os provedores em um só lugar

#### 📋 Todos os Fornecedores Suportados

| Fornecedor | Prefixo `model` | API Base Padrão | Protocolo | Chave API |
|-------------|-----------------|------------------|----------|-----------|
| **OpenAI** | `openai/` | `https://api.openai.com/v1` | OpenAI | [Obter Chave](https://platform.openai.com) |
| **Anthropic** | `anthropic/` | `https://api.anthropic.com/v1` | Anthropic | [Obter Chave](https://console.anthropic.com) |
| **Zhipu AI (GLM)** | `zhipu/` | `https://open.bigmodel.cn/api/paas/v4` | OpenAI | [Obter Chave](https://open.bigmodel.cn/usercenter/proj-mgmt/apikeys) |
| **DeepSeek** | `deepseek/` | `https://api.deepseek.com/v1` | OpenAI | [Obter Chave](https://platform.deepseek.com) |
| **Google Gemini** | `gemini/` | `https://generativelanguage.googleapis.com/v1beta` | OpenAI | [Obter Chave](https://aistudio.google.com/api-keys) |
| **Groq** | `groq/` | `https://api.groq.com/openai/v1` | OpenAI | [Obter Chave](https://console.groq.com) |
| **Moonshot** | `moonshot/` | `https://api.moonshot.cn/v1` | OpenAI | [Obter Chave](https://platform.moonshot.cn) |
| **Qwen (Alibaba)** | `qwen/` | `https://dashscope.aliyuncs.com/compatible-mode/v1` | OpenAI | [Obter Chave](https://dashscope.console.aliyun.com) |
| **NVIDIA** | `nvidia/` | `https://integrate.api.nvidia.com/v1` | OpenAI | [Obter Chave](https://build.nvidia.com) |
| **Ollama** | `ollama/` | `http://localhost:11434/v1` | OpenAI | Local (sem chave necessária) |
| **OpenRouter** | `openrouter/` | `https://openrouter.ai/api/v1` | OpenAI | [Obter Chave](https://openrouter.ai/keys) |
| **VLLM** | `vllm/` | `http://localhost:8000/v1` | OpenAI | Local |
| **Cerebras** | `cerebras/` | `https://api.cerebras.ai/v1` | OpenAI | [Obter Chave](https://cerebras.ai) |
| **VolcEngine (Doubao)** | `volcengine/` | `https://ark.cn-beijing.volces.com/api/v3` | OpenAI | [Obter Chave](https://www.volcengine.com/activity/codingplan?utm_campaign=PicoClaw&utm_content=PicoClaw&utm_medium=devrel&utm_source=OWO&utm_term=PicoClaw) |
| **ShengsuanYun** | `shengsuanyun/` | `https://router.shengsuanyun.com/api/v1` | OpenAI | - |
| **BytePlus**        | `byteplus/`       | `https://ark.ap-southeast.bytepluses.com/api/v3`    | OpenAI    | [Obter Chave](https://www.byteplus.com)                    |
| **LongCat**         | `longcat/`        | `https://api.longcat.chat/openai`                   | OpenAI    | [Obter Chave](https://longcat.chat/platform)                     |
| **ModelScope (魔搭)**| `modelscope/`    | `https://api-inference.modelscope.cn/v1`            | OpenAI    | [Obter Token](https://modelscope.cn/my/tokens)                   |
| **Antigravity** | `antigravity/` | Google Cloud | Custom | Apenas OAuth |
| **GitHub Copilot** | `github-copilot/` | `localhost:4321` | gRPC | - |

#### Configuração Básica

```json
{
  "model_list": [
    {
      "model_name": "ark-code-latest",
      "model": "volcengine/ark-code-latest",
      "api_key": "sk-your-api-key"
    },
    {
      "model_name": "gpt-5.4",
      "model": "openai/gpt-5.4",
      "api_key": "sk-your-openai-key"
    },
    {
      "model_name": "claude-sonnet-4.6",
      "model": "anthropic/claude-sonnet-4.6",
      "api_key": "sk-ant-your-key"
    },
    {
      "model_name": "glm-4.7",
      "model": "zhipu/glm-4.7",
      "api_key": "your-zhipu-key"
    }
  ],
  "agents": {
    "defaults": {
      "model": "gpt-5.4"
    }
  }
}
```

#### Exemplos por Fornecedor

**OpenAI**
```json
{
  "model_name": "gpt-5.4",
  "model": "openai/gpt-5.4",
  "api_key": "sk-..."
}
```

**VolcEngine (Doubao)**
```json
{
  "model_name": "ark-code-latest",
  "model": "volcengine/ark-code-latest",
  "api_key": "sk-..."
}
```

**Zhipu AI (GLM)**
```json
{
  "model_name": "glm-4.7",
  "model": "zhipu/glm-4.7",
  "api_key": "your-key"
}
```

**Anthropic (com OAuth)**
```json
{
  "model_name": "claude-sonnet-4.6",
  "model": "anthropic/claude-sonnet-4.6",
  "auth_method": "oauth"
}
```
> Execute `picoclaw auth login --provider anthropic` para configurar credenciais OAuth.

**Proxy/API personalizada**
```json
{
  "model_name": "my-custom-model",
  "model": "openai/custom-model",
  "api_base": "https://my-proxy.com/v1",
  "api_key": "sk-...",
  "request_timeout": 300
}
```

#### Balanceamento de Carga

Configure vários endpoints para o mesmo nome de modelo—PicoClaw fará round-robin automaticamente entre eles:

```json
{
  "model_list": [
    {
      "model_name": "gpt-5.4",
      "model": "openai/gpt-5.4",
      "api_base": "https://api1.example.com/v1",
      "api_key": "sk-key1"
    },
    {
      "model_name": "gpt-5.4",
      "model": "openai/gpt-5.4",
      "api_base": "https://api2.example.com/v1",
      "api_key": "sk-key2"
    }
  ]
}
```

#### Migração da Configuração Legada `providers`

A configuração antiga `providers` está **descontinuada** mas ainda é suportada para compatibilidade reversa.

**Configuração Antiga (descontinuada):**
```json
{
  "providers": {
    "zhipu": {
      "api_key": "your-key",
      "api_base": "https://open.bigmodel.cn/api/paas/v4"
    }
  },
  "agents": {
    "defaults": {
      "provider": "zhipu",
      "model": "glm-4.7"
    }
  }
}
```

**Nova Configuração (recomendada):**
```json
{
  "model_list": [
    {
      "model_name": "glm-4.7",
      "model": "zhipu/glm-4.7",
      "api_key": "your-key"
    }
  ],
  "agents": {
    "defaults": {
      "model": "glm-4.7"
    }
  }
}
```

Para o guia de migração detalhado, consulte [docs/migration/model-list-migration.md](docs/migration/model-list-migration.md).

## Referência CLI

| Comando | Descrição |
| --- | --- |
| `picoclaw onboard` | Inicializar configuração & workspace |
| `picoclaw agent -m "..."` | Conversar com o agente |
| `picoclaw agent` | Modo de chat interativo |
| `picoclaw gateway` | Iniciar o gateway (para bots de chat) |
| `picoclaw status` | Mostrar status |
| `picoclaw cron list` | Listar todas as tarefas agendadas |
| `picoclaw cron add ...` | Adicionar uma tarefa agendada |
>>>>>>> refactor/agent

### Tarefas Agendadas / Lembretes

O PicoClaw suporta lembretes agendados e tarefas recorrentes por meio da ferramenta `cron`:

* **Lembretes únicos**: "Me lembre em 10 minutos" → dispara uma vez após 10min
* **Tarefas recorrentes**: "Me lembre a cada 2 horas" → dispara a cada 2 horas
* **Expressões Cron**: "Me lembre às 9h todos os dias" → usa expressão cron

## 🤝 Contribuir & Roadmap

PRs são bem-vindos! O código-fonte é intencionalmente pequeno e legível. 🤗

Veja nosso [Roadmap da Comunidade](https://github.com/sipeed/picoclaw/blob/main/ROADMAP.md) completo.

Grupo de desenvolvedores em formação. Junte-se após seu primeiro PR com merge!

Grupos de usuários:

discord: <https://discord.gg/V4sAZ9XWpN>

<img src="assets/wechat.png" alt="PicoClaw" width="512">
