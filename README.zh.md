<div align="center">
<img src="assets/logo.webp" alt="PicoClaw" width="512">

<h1>PicoClaw: 基于Go语言的超高效 AI 助手</h1>

<h3>$10 硬件 · <10MB 内存 · <1s 启动 · 皮皮虾，我们走！</h3>
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

**中文** | [日本語](README.ja.md) | [Português](README.pt-br.md) | [Tiếng Việt](README.vi.md) | [Français](README.fr.md) | [Italiano](README.it.md) | [Bahasa Indonesia](README.id.md) | [English](README.md)

</div>

---

> **PicoClaw** 是由 [矽速科技 (Sipeed)](https://sipeed.com) 发起的独立开源项目，完全使用 **Go 语言**从零编写——不是 OpenClaw、NanoBot 或其他项目的分支。

🦐 **PicoClaw** 是一个受 [NanoBot](https://github.com/HKUDS/nanobot) 启发的超轻量级个人 AI 助手。它采用 **Go 语言** 从零重构，经历了一个"自举"过程——即由 AI Agent 自身驱动了整个架构迁移和代码优化。

⚡️ **极致轻量**：可在 **10 美元** 的硬件上运行，内存占用 **<10MB**。这意味着比 OpenClaw 节省 99% 的内存，比 Mac mini 便宜 98%！

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
> **🚨 安全声明**
>
> - **无加密货币 (NO CRYPTO):** PicoClaw **没有** 发行任何官方代币、Token 或虚拟货币。所有在 `pump.fun` 或其他交易平台上的相关声称均为 **诈骗**。
> - **官方域名:** 唯一的官方网站是 **[picoclaw.io](https://picoclaw.io)**，公司官网是 **[sipeed.com](https://sipeed.com)**。
> - **警惕:** 许多 `.ai/.org/.com/.net/...` 后缀的域名被第三方抢注，请勿轻信。
> - **注意:** PicoClaw 正在初期的快速功能开发阶段，可能有尚未修复的网络安全问题，在 1.0 正式版发布前，请不要将其部署到生产环境中。
> - **注意:** PicoClaw 最近合并了大量 PR，近期版本可能内存占用较大 (10~20MB)，我们将在功能较为收敛后进行资源占用优化。

## 📢 新闻

2026-03-17 🚀 **v0.2.3 发布！** 系统托盘 UI（Windows & Linux）、子 Agent 状态查询 (`spawn_status`)、实验性 Gateway 热重载、Cron 安全门控，以及 2 项安全修复。PicoClaw 已达 **25K ⭐**！

2026-03-09 🎉 **v0.2.1 — 史上最大更新！** MCP 协议支持、4 个新频道 (Matrix/IRC/WeCom/Discord Proxy)、3 个新 Provider (Kimi/Minimax/Avian)、视觉管线、JSONL 记忆存储、模型路由。

2026-02-28 📦 **v0.2.0** 发布，支持 Docker Compose 和 Web UI 启动器。

2026-02-26 🎉 PicoClaw 仅 17 天突破 **20K Stars**！频道自动编排和能力接口上线。

<details>
<summary>更早的新闻...</summary>

2026-02-16 🎉 PicoClaw 一周内突破 12K Stars！社区维护者角色和 [路线图](ROADMAP.md) 正式发布。

2026-02-13 🎉 PicoClaw 4 天内突破 5000 Stars！项目路线图和开发者群组筹建中。

2026-02-09 🎉 **PicoClaw 正式发布！** 仅用 1 天构建，将 AI Agent 带入 $10 硬件与 <10MB 内存的世界。🦐 皮皮虾，我们走！

</details>

## ✨ 特性

🪶 **超轻量级**: 核心功能内存占用 <10MB — 比 OpenClaw 小 99%。*

💰 **极低成本**: 高效到足以在 $10 的硬件上运行 — 比 Mac mini 便宜 98%。

⚡️ **闪电启动**: 启动速度快 400 倍，即使在 0.6GHz 单核处理器上也能在 1 秒内启动。

🌍 **真正可移植**: 跨 RISC-V、ARM、MIPS 和 x86 架构的单二进制文件，一键运行！

🤖 **AI 自举**: 纯 Go 语言原生实现 — 95% 的核心代码由 Agent 生成，并经由"人机回环"微调。

🔌 **MCP 支持**: 原生 [Model Context Protocol](https://modelcontextprotocol.io/) 集成 — 连接任意 MCP 服务器扩展 Agent 能力。

👁️ **视觉管线**: 直接向 Agent 发送图片和文件 — 自动 base64 编码对接多模态 LLM。

🧠 **智能路由**: 基于规则的模型路由 — 简单查询走轻量模型，节省 API 成本。

_*近期版本因快速合并 PR 可能占用 10–20MB，资源优化已列入计划。启动速度对比基于 0.8GHz 单核实测（见下方对比表）。_

|                                | OpenClaw      | NanoBot                  | **PicoClaw**                           |
| ------------------------------ | ------------- | ------------------------ | -------------------------------------- |
| **语言**                       | TypeScript    | Python                   | **Go**                                 |
| **RAM**                        | >1GB          | >100MB                   | **< 10MB***                            |
| **启动时间**</br>(0.8GHz core) | >500s         | >30s                     | **<1s**                                |
| **成本**                       | Mac Mini $599 | 大多数 Linux 开发板 ~$50 | **任意 Linux 开发板**</br>**低至 $10** |

<img src="assets/compare.jpg" alt="PicoClaw" width="512">

> 📋 **[硬件兼容列表](docs/hardware-compatibility.md)** — 查看所有已测试的板卡，从 $5 RISC-V 到树莓派到安卓手机。你的板卡没在列表中？欢迎提交 PR！

## 🦾 演示

### 🛠️ 标准助手工作流

<table align="center">
<tr align="center">
<th><p align="center">🧩 全栈工程师模式</p></th>
<th><p align="center">🗂️ 日志与规划管理</p></th>
<th><p align="center">🔎 网络搜索与学习</p></th>
</tr>
<tr>
<td align="center"><p align="center"><img src="assets/picoclaw_code.gif" width="240" height="180"></p></td>
<td align="center"><p align="center"><img src="assets/picoclaw_memory.gif" width="240" height="180"></p></td>
<td align="center"><p align="center"><img src="assets/picoclaw_search.gif" width="240" height="180"></p></td>
</tr>
<tr>
<td align="center">开发 • 部署 • 扩展</td>
<td align="center">日程 • 自动化 • 记忆</td>
<td align="center">发现 • 洞察 • 趋势</td>
</tr>
</table>

### 📱 在手机上轻松运行

PicoClaw 可以将你 10 年前的老旧手机废物利用，变身成为你的 AI 助理！快速指南：

1. 安装 [Termux](https://github.com/termux/termux-app)（可从 [GitHub Releases](https://github.com/termux/termux-app/releases) 下载，或在 F-Droid 等应用商店搜索）
2. 打开后执行指令

```bash
# 从 Release 页面下载最新版本
wget https://github.com/sipeed/picoclaw/releases/latest/download/picoclaw_Linux_arm64.tar.gz
tar xzf picoclaw_Linux_arm64.tar.gz
pkg install proot
termux-chroot ./picoclaw onboard   # chroot 提供标准 Linux 文件系统布局
```

然后跟随下面的"快速开始"章节继续配置 PicoClaw 即可使用！

<img src="assets/termux.jpg" alt="PicoClaw" width="512">

### 🐜 创新的低占用部署

PicoClaw 几乎可以部署在任何 Linux 设备上！

- $9.9 [LicheeRV-Nano](https://www.aliexpress.com/item/1005006519668532.html) E(网口) 或 W(WiFi6) 版本，用于极简家庭助手
- $30~50 [NanoKVM](https://www.aliexpress.com/item/1005007369816019.html)，或 $100 [NanoKVM-Pro](https://www.aliexpress.com/item/1005010048471263.html)，用于自动化服务器运维
- $50 [MaixCAM](https://www.aliexpress.com/item/1005008053333693.html) 或 $100 [MaixCAM2](https://www.kickstarter.com/projects/zepan/maixcam2-build-your-next-gen-4k-ai-camera)，用于智能监控

<https://private-user-images.githubusercontent.com/83055338/547056448-e7b031ff-d6f5-4468-bcca-5726b6fecb5c.mp4>

🌟 更多部署案例敬请期待！

## 📦 安装

### 从 picoclaw.io 下载（推荐）

访问 **[picoclaw.io](https://picoclaw.io)** — 官网自动检测你的平台，提供一键下载，无需手动选择架构。

### 下载预编译二进制文件

也可以从 [GitHub Releases](https://github.com/sipeed/picoclaw/releases) 页面手动下载对应平台的二进制文件。

### 从源码构建（开发用）

```bash
git clone https://github.com/sipeed/picoclaw.git

cd picoclaw
make deps

# 构建（无需安装）
make build

# 为多平台构建
make build-all

# 为 Raspberry Pi Zero 2 W 构建（32位: make build-linux-arm; 64位: make build-linux-arm64）
make build-pi-zero

# 构建并安装
make install
```

**Raspberry Pi Zero 2 W:** 请使用与系统匹配的二进制文件：32 位 Raspberry Pi OS → `make build-linux-arm`；64 位 → `make build-linux-arm64`。或运行 `make build-pi-zero` 同时构建两者。

## 📚 文档

详细指南请参阅以下文档，README 仅涵盖快速入门。

| 主题 | 说明 |
|------|------|
| 🐳 [Docker 与快速开始](docs/zh/docker.md) | Docker Compose 配置、Launcher/Agent 模式、快速开始 |
| 💬 [聊天应用配置](docs/zh/chat-apps.md) | Telegram、Discord、WhatsApp、Matrix、QQ、Slack、IRC、钉钉、LINE、飞书、企业微信等 |
| ⚙️ [配置指南](docs/zh/configuration.md) | 环境变量、工作区布局、技能来源、安全沙箱、心跳任务 |
| 🔌 [提供商与模型配置](docs/zh/providers.md) | 20+ LLM 提供商、模型路由、model_list 配置、Provider 架构 |
| 🔄 [异步任务与 Spawn](docs/zh/spawn-tasks.md) | 快速任务、长任务与 Spawn、异步子 Agent 编排 |
| 🐛 [疑难解答](docs/zh/troubleshooting.md) | 常见问题与解决方案 |
| 🔧 [工具配置](docs/zh/tools_configuration.md) | 工具启用/禁用、执行策略 |
| 📋 [硬件兼容列表](docs/hardware-compatibility.md) | 已测试板卡、最低要求、如何添加你的板卡 |

## <img src="assets/clawdchat-icon.png" width="24" height="24" alt="ClawdChat"> 加入 Agent 社交网络

<<<<<<< HEAD
通过 CLI 或任何已集成的聊天应用发送一条消息，即可将 PicoClaw 连接到 Agent 社交网络。

**阅读 `https://clawdchat.ai/skill.md` 并按照说明加入 [ClawdChat.ai](https://clawdchat.ai)**

## 🖥️ CLI 命令行参考

| 命令                       | 说明                   |
| ------------------------- | ---------------------- |
| `picoclaw onboard`        | 初始化配置与工作区       |
| `picoclaw agent -m "..."` | 与 Agent 对话           |
| `picoclaw agent`          | 交互式对话模式           |
| `picoclaw gateway`        | 启动网关                |
| `picoclaw status`         | 查看状态                |
| `picoclaw version`        | 查看版本信息             |
| `picoclaw cron list`      | 列出所有定时任务         |
| `picoclaw cron add ...`   | 添加定时任务             |
| `picoclaw cron disable`   | 禁用定时任务             |
| `picoclaw cron remove`    | 删除定时任务             |
| `picoclaw skills list`    | 列出已安装技能           |
| `picoclaw skills install` | 安装技能                |
| `picoclaw migrate`        | 从旧版本迁移数据         |
| `picoclaw auth login`     | 认证提供商               |
| `picoclaw model`          | 查看或切换默认模型 |

### 定时任务 / 提醒
=======
只需通过 CLI 或任何集成的聊天应用发送一条消息，即可将 PicoClaw 连接到 Agent 社交网络。

\*\*阅读 `https://clawdchat.ai/skill.md` 并按照说明加入 [ClawdChat.ai](https://clawdchat.ai)

## ⚙️ 配置详解

配置文件路径: `~/.picoclaw/config.json`

### 环境变量

你可以使用环境变量覆盖默认路径。这对于便携安装、容器化部署或将 picoclaw 作为系统服务运行非常有用。这些变量是独立的，控制不同的路径。

| 变量              | 描述                                                                                                                             | 默认路径                  |
|-------------------|-----------------------------------------------------------------------------------------------------------------------------------------|---------------------------|
| `PICOCLAW_CONFIG` | 覆盖配置文件的路径。这直接告诉 picoclaw 加载哪个 `config.json`，忽略所有其他位置。 | `~/.picoclaw/config.json` |
| `PICOCLAW_HOME`   | 覆盖 picoclaw 数据根目录。这会更改 `workspace` 和其他数据目录的默认位置。          | `~/.picoclaw`             |

**示例：**

```bash
# 使用特定的配置文件运行 picoclaw
# 工作区路径将从该配置文件中读取
PICOCLAW_CONFIG=/etc/picoclaw/production.json picoclaw gateway

# 在 /opt/picoclaw 中存储所有数据运行 picoclaw
# 配置将从默认的 ~/.picoclaw/config.json 加载
# 工作区将在 /opt/picoclaw/workspace 创建
PICOCLAW_HOME=/opt/picoclaw picoclaw agent

# 同时使用两者进行完全自定义设置
PICOCLAW_HOME=/srv/picoclaw PICOCLAW_CONFIG=/srv/picoclaw/main.json picoclaw gateway
```

### 工作区布局 (Workspace Layout)

PicoClaw 将数据存储在您配置的工作区中（默认：`~/.picoclaw/workspace`）：

```
~/.picoclaw/workspace/
├── sessions/          # 对话会话和历史
├── memory/            # 长期记忆 (MEMORY.md)
├── state/             # 持久化状态 (最后一次频道等)
├── cron/              # 定时任务数据库
├── skills/            # 工作区级技能
├── AGENT.md           # 结构化 Agent 定义与系统提示词
├── SOUL.md            # Agent 灵魂/性格
├── USER.md            # 当前工作区的用户资料与偏好
├── HEARTBEAT.md       # 周期性任务提示词 (每 30 分钟检查一次)
└── ...

```

### 技能来源 (Skill Sources)

默认情况下，技能会按以下顺序加载：

1. `~/.picoclaw/workspace/skills`（工作区）
2. `~/.picoclaw/skills`（全局）
3. `<current-working-directory>/skills`（内置）

在高级/测试场景下，可通过以下环境变量覆盖内置技能目录：

```bash
export PICOCLAW_BUILTIN_SKILLS=/path/to/skills
```

### 统一命令执行策略

- 通用斜杠命令通过 `pkg/agent/loop.go` 中的 `commands.Executor` 统一执行。
- Channel 适配器不再在本地消费通用命令；它们只负责把入站文本转发到 bus/agent 路径。Telegram 仍会在启动时自动注册其支持的命令菜单。
- 未注册的斜杠命令（例如 `/foo`）会透传给 LLM 按普通输入处理。
- 已注册但当前 channel 不支持的命令（例如 WhatsApp 上的 `/show`）会返回明确的用户可见错误，并停止后续处理。
### 心跳 / 周期性任务 (Heartbeat)

PicoClaw 可以自动执行周期性任务。在工作区创建 `HEARTBEAT.md` 文件：

```markdown
# Periodic Tasks

- Check my email for important messages
- Review my calendar for upcoming events
- Check the weather forecast
```

Agent 将每隔 30 分钟（可配置）读取此文件，并使用可用工具执行任务。

#### 使用 Spawn 的异步任务

对于耗时较长的任务（网络搜索、API 调用），使用 `spawn` 工具创建一个 **子 Agent (subagent)**：

```markdown
# Periodic Tasks

## Quick Tasks (respond directly)

- Report current time

## Long Tasks (use spawn for async)

- Search the web for AI news and summarize
- Check email and report important messages
```

**关键行为：**

| 特性             | 描述                                     |
| ---------------- | ---------------------------------------- |
| **spawn**        | 创建异步子 Agent，不阻塞主心跳进程       |
| **独立上下文**   | 子 Agent 拥有独立上下文，无会话历史      |
| **message tool** | 子 Agent 通过 message 工具直接与用户通信 |
| **非阻塞**       | spawn 后，心跳继续处理下一个任务         |

#### 子 Agent 通信原理

```
心跳触发 (Heartbeat triggers)
    ↓
Agent 读取 HEARTBEAT.md
    ↓
对于长任务: spawn 子 Agent
    ↓                           ↓
继续下一个任务               子 Agent 独立工作
    ↓                           ↓
所有任务完成                 子 Agent 使用 "message" 工具
    ↓                           ↓
响应 HEARTBEAT_OK            用户直接收到结果

```

子 Agent 可以访问工具（message, web_search 等），并且无需通过主 Agent 即可独立与用户通信。

**配置：**

```json
{
  "heartbeat": {
    "enabled": true,
    "interval": 30
  }
}
```

| 选项       | 默认值 | 描述                         |
| ---------- | ------ | ---------------------------- |
| `enabled`  | `true` | 启用/禁用心跳                |
| `interval` | `30`   | 检查间隔，单位分钟 (最小: 5) |

**环境变量:**

- `PICOCLAW_HEARTBEAT_ENABLED=false` 禁用
- `PICOCLAW_HEARTBEAT_INTERVAL=60` 更改间隔

### 提供商 (Providers)

> [!NOTE]
> Groq 通过 Whisper 提供免费的语音转录。如果配置了 Groq，任意渠道的音频消息都将在 Agent 层面自动转录为文字。

| 提供商               | 用途                         | 获取 API Key                                                         |
| -------------------- | ---------------------------- | -------------------------------------------------------------------- |
| `gemini`             | LLM (Gemini 直连)            | [aistudio.google.com](https://aistudio.google.com)                   |
| `zhipu`              | LLM (智谱直连)               | [bigmodel.cn](bigmodel.cn)                                           |
| `volcengine` | LLM (火山引擎直连)                  | [volcengine.com](https://www.volcengine.com/activity/codingplan?utm_campaign=PicoClaw&utm_content=PicoClaw&utm_medium=devrel&utm_source=OWO&utm_term=PicoClaw)                 |
| `openrouter` | LLM (推荐，可访问所有模型)   | [openrouter.ai](https://openrouter.ai)                               |
| `anthropic`  | LLM (Claude 直连)            | [console.anthropic.com](https://console.anthropic.com)               |
| `openai`     | LLM (GPT 直连)               | [platform.openai.com](https://platform.openai.com)                   |
| `deepseek`   | LLM (DeepSeek 直连)          | [platform.deepseek.com](https://platform.deepseek.com)               |
| `qwen`               | LLM (通义千问)               | [dashscope.console.aliyun.com](https://dashscope.console.aliyun.com) |
| `groq`               | LLM + **语音转录** (Whisper) | [console.groq.com](https://console.groq.com)                         |
| `cerebras`           | LLM (Cerebras 直连)          | [cerebras.ai](https://cerebras.ai)                                   |

### 模型配置 (model_list)

> **新功能！** PicoClaw 现在采用**以模型为中心**的配置方式。只需使用 `厂商/模型` 格式（如 `zhipu/glm-4.7`）即可添加新的 provider——**无需修改任何代码！**

该设计同时支持**多 Agent 场景**，提供灵活的 Provider 选择：

- **不同 Agent 使用不同 Provider**：每个 Agent 可以使用自己的 LLM provider
- **模型回退（Fallback）**：配置主模型和备用模型，提高可靠性
- **负载均衡**：在多个 API 端点之间分配请求
- **集中化配置**：在一个地方管理所有 provider

#### 📋 所有支持的厂商

| 厂商                | `model` 前缀      | 默认 API Base                                       | 协议      | 获取 API Key                                                      |
| ------------------- | ----------------- | --------------------------------------------------- | --------- | ----------------------------------------------------------------- |
| **OpenAI**          | `openai/`         | `https://api.openai.com/v1`                         | OpenAI    | [获取密钥](https://platform.openai.com)                           |
| **Anthropic**       | `anthropic/`      | `https://api.anthropic.com/v1`                      | Anthropic | [获取密钥](https://console.anthropic.com)                         |
| **智谱 AI (GLM)**   | `zhipu/`          | `https://open.bigmodel.cn/api/paas/v4`              | OpenAI    | [获取密钥](https://open.bigmodel.cn/usercenter/proj-mgmt/apikeys) |
| **DeepSeek**        | `deepseek/`       | `https://api.deepseek.com/v1`                       | OpenAI    | [获取密钥](https://platform.deepseek.com)                         |
| **Google Gemini**   | `gemini/`         | `https://generativelanguage.googleapis.com/v1beta`  | OpenAI    | [获取密钥](https://aistudio.google.com/api-keys)                  |
| **Groq**            | `groq/`           | `https://api.groq.com/openai/v1`                    | OpenAI    | [获取密钥](https://console.groq.com)                              |
| **Moonshot**        | `moonshot/`       | `https://api.moonshot.cn/v1`                        | OpenAI    | [获取密钥](https://platform.moonshot.cn)                          |
| **通义千问 (Qwen)** | `qwen/`           | `https://dashscope.aliyuncs.com/compatible-mode/v1` | OpenAI    | [获取密钥](https://dashscope.console.aliyun.com)                  |
| **NVIDIA**          | `nvidia/`         | `https://integrate.api.nvidia.com/v1`               | OpenAI    | [获取密钥](https://build.nvidia.com)                              |
| **Ollama**          | `ollama/`         | `http://localhost:11434/v1`                         | OpenAI    | 本地（无需密钥）                                                  |
| **OpenRouter**      | `openrouter/`     | `https://openrouter.ai/api/v1`                      | OpenAI    | [获取密钥](https://openrouter.ai/keys)                            |
| **VLLM**            | `vllm/`           | `http://localhost:8000/v1`                          | OpenAI    | 本地                                                              |
| **Cerebras**        | `cerebras/`       | `https://api.cerebras.ai/v1`                        | OpenAI    | [获取密钥](https://cerebras.ai)                                   |
| **火山引擎（Doubao）**        | `volcengine/`     | `https://ark.cn-beijing.volces.com/api/v3`          | OpenAI    | [获取密钥](https://www.volcengine.com/activity/codingplan?utm_campaign=PicoClaw&utm_content=PicoClaw&utm_medium=devrel&utm_source=OWO&utm_term=PicoClaw)                        |
| **神算云**          | `shengsuanyun/`   | `https://router.shengsuanyun.com/api/v1`            | OpenAI    | -                                                                 |
| **BytePlus**        | `byteplus/`       | `https://ark.ap-southeast.bytepluses.com/api/v3`    | OpenAI    | [获取密钥](https://www.byteplus.com)                        |
| **LongCat**         | `longcat/`        | `https://api.longcat.chat/openai`                   | OpenAI    | [获取密钥](https://longcat.chat/platform)                        |
| **ModelScope (魔搭)**| `modelscope/`    | `https://api-inference.modelscope.cn/v1`            | OpenAI    | [获取 Token](https://modelscope.cn/my/tokens)                    |
| **Antigravity**     | `antigravity/`    | Google Cloud                                        | 自定义    | 仅 OAuth                                                          |
| **GitHub Copilot**  | `github-copilot/` | `localhost:4321`                                    | gRPC      | -                                                                 |

#### 基础配置示例

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

#### 各厂商配置示例

**OpenAI**

```json
{
  "model_name": "gpt-5.4",
  "model": "openai/gpt-5.4",
  "api_key": "sk-..."
}
```

**火山引擎（Doubao）**

```json
{
  "model_name": "ark-code-latest",
  "model": "volcengine/ark-code-latest",
  "api_key": "sk-..."
}
```

**智谱 AI (GLM)**

```json
{
  "model_name": "glm-4.7",
  "model": "zhipu/glm-4.7",
  "api_key": "your-key"
}
```

**DeepSeek**

```json
{
  "model_name": "deepseek-chat",
  "model": "deepseek/deepseek-chat",
  "api_key": "sk-..."
}
```

**Anthropic (使用 OAuth)**

```json
{
  "model_name": "claude-sonnet-4.6",
  "model": "anthropic/claude-sonnet-4.6",
  "auth_method": "oauth"
}
```

> 运行 `picoclaw auth login --provider anthropic` 来设置 OAuth 凭证。

**Anthropic Messages API（原生格式）**

用于直接访问 Anthropic API 或仅支持 Anthropic 原生消息格式的自定义端点：

```json
{
  "model_name": "claude-opus-4-6",
  "model": "anthropic-messages/claude-opus-4-6",
  "api_key": "sk-ant-your-key",
  "api_base": "https://api.anthropic.com"
}
```

> 使用 `anthropic-messages` 协议的场景：
> - 使用仅支持 Anthropic 原生 `/v1/messages` 端点的第三方代理（不支持 OpenAI 兼容的 `/v1/chat/completions`）
> - 连接到 MiniMax、Synthetic 等需要 Anthropic 原生消息格式的服务
> - 现有的 `anthropic` 协议返回 404 错误（说明端点不支持 OpenAI 兼容格式）
>
> **注意：** `anthropic` 协议使用 OpenAI 兼容格式（`/v1/chat/completions`），而 `anthropic-messages` 使用 Anthropic 原生格式（`/v1/messages`）。请根据端点支持的格式选择。

**Ollama (本地)**

```json
{
  "model_name": "llama3",
  "model": "ollama/llama3"
}
```

**自定义代理/API**

```json
{
  "model_name": "my-custom-model",
  "model": "openai/custom-model",
  "api_base": "https://my-proxy.com/v1",
  "api_key": "sk-...",
  "request_timeout": 300
}
```

#### 负载均衡

为同一个模型名称配置多个端点——PicoClaw 会自动在它们之间轮询：

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

#### 从旧的 `providers` 配置迁移

旧的 `providers` 配置格式**已弃用**，但为向后兼容仍支持。

**旧配置（已弃用）：**

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

**新配置（推荐）：**

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

详细的迁移指南请参考 [docs/migration/model-list-migration.md](docs/migration/model-list-migration.md)。

<details>
<summary><b>智谱 (Zhipu) 配置示例</b></summary>

**1. 获取 API key 和 base URL**

- 获取 [API key](https://bigmodel.cn/usercenter/proj-mgmt/apikeys)

**2. 配置**

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
      "api_key": "Your API Key",
      "api_base": "https://open.bigmodel.cn/api/paas/v4"
    }
  }
}
```

**3. 运行**

```bash
picoclaw agent -m "你好"

```

</details>

<details>
<summary><b>完整配置示例</b></summary>

```json
{
  "agents": {
    "defaults": {
      "model": "anthropic/claude-opus-4-5"
    }
  },
  "session": {
    "dm_scope": "per-channel-peer",
    "backlog_limit": 20
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
        "api_key": "YOUR_BRAVE_API_KEY",
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

## CLI 命令行参考

| 命令                      | 描述               |
| ------------------------- | ------------------ |
| `picoclaw onboard`        | 初始化配置和工作区 |
| `picoclaw agent -m "..."` | 与 Agent 对话      |
| `picoclaw agent`          | 交互式聊天模式     |
| `picoclaw gateway`        | 启动网关 (Gateway) |
| `picoclaw status`         | 显示状态           |
| `picoclaw cron list`      | 列出所有定时任务   |
| `picoclaw cron add ...`   | 添加定时任务       |

### 定时任务 / 提醒 (Scheduled Tasks)
>>>>>>> refactor/agent

PicoClaw 通过 `cron` 工具支持定时提醒和重复任务：

* **一次性提醒**: "10分钟后提醒我" → 10分钟后触发一次
* **重复任务**: "每2小时提醒我" → 每2小时触发
* **Cron 表达式**: "每天上午9点提醒我" → 使用 cron 表达式

## 🤝 贡献与路线图

欢迎提交 PR！代码库刻意保持小巧和可读。🤗

查看完整的 [社区路线图](https://github.com/sipeed/picoclaw/blob/main/ROADMAP.md)。

开发者群组正在组建中，入群门槛：至少合并过 1 个 PR。

用户群组：

Discord: <https://discord.gg/V4sAZ9XWpN>

<img src="assets/wechat.png" alt="PicoClaw" width="512">
