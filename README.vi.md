<div align="center">
  <img src="assets/logo.webp" alt="PicoClaw" width="512">

  <h1>PicoClaw: Trợ lý AI Siêu Nhẹ viết bằng Go</h1>

  <h3>Phần cứng $10 · <10MB RAM · Khởi động <1 giây · Nào, xuất phát!</h3>
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

[中文](README.zh.md) | [日本語](README.ja.md) | [Português](README.pt-br.md) | **Tiếng Việt** | [Français](README.fr.md) | [Italiano](README.it.md) | [Bahasa Indonesia](README.id.md) | [English](README.md)

</div>

---

> **PicoClaw** là dự án mã nguồn mở độc lập được khởi xướng bởi [Sipeed](https://sipeed.com). Được viết hoàn toàn bằng **Go** — không phải là bản fork của OpenClaw, NanoBot hay bất kỳ dự án nào khác.

🦐 PicoClaw là trợ lý AI cá nhân siêu nhẹ, lấy cảm hứng từ [NanoBot](https://github.com/HKUDS/nanobot), được viết lại hoàn toàn bằng Go thông qua quá trình "tự khởi tạo" (self-bootstrapping) — nơi chính AI Agent đã tự dẫn dắt toàn bộ quá trình chuyển đổi kiến trúc và tối ưu hóa mã nguồn.

⚡️ Chạy trên phần cứng chỉ $10 với RAM <10MB: Tiết kiệm 99% bộ nhớ so với OpenClaw và rẻ hơn 98% so với Mac mini!

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
> **🚨 TUYÊN BỐ BẢO MẬT & KÊNH CHÍNH THỨC**
>
> * **KHÔNG CÓ CRYPTO:** PicoClaw **KHÔNG** có bất kỳ token/coin chính thức nào. Mọi thông tin trên `pump.fun` hoặc các sàn giao dịch khác đều là **LỪA ĐẢO**.
>
> * **DOMAIN CHÍNH THỨC:** Website chính thức **DUY NHẤT** là **[picoclaw.io](https://picoclaw.io)**, website công ty là **[sipeed.com](https://sipeed.com)**
> * **Cảnh báo:** Nhiều tên miền `.ai/.org/.com/.net/...` đã bị bên thứ ba đăng ký.
> * **Cảnh báo:** PicoClaw đang trong giai đoạn phát triển sớm và có thể còn các vấn đề bảo mật mạng chưa được giải quyết. Không nên triển khai lên môi trường production trước phiên bản v1.0.
> * **Lưu ý:** PicoClaw gần đây đã merge nhiều PR, dẫn đến bộ nhớ sử dụng có thể lớn hơn (10–20MB) ở các phiên bản mới nhất. Chúng tôi sẽ ưu tiên tối ưu tài nguyên khi bộ tính năng đã ổn định.

## 📢 Tin tức

2026-03-17 🚀 **v0.2.3 Phát hành!** Giao diện khay hệ thống (Windows & Linux), theo dõi trạng thái sub-agent (`spawn_status`), hot-reload gateway thử nghiệm, cổng bảo mật cron và 2 bản vá bảo mật. PicoClaw đạt **25K ⭐**!

2026-03-09 🎉 **v0.2.1 — Bản cập nhật lớn nhất!** Hỗ trợ giao thức MCP, 4 kênh mới (Matrix/IRC/WeCom/Discord Proxy), 3 nhà cung cấp mới (Kimi/Minimax/Avian), pipeline xử lý hình ảnh, bộ nhớ JSONL và định tuyến mô hình.

2026-02-28 📦 **v0.2.0** phát hành với hỗ trợ Docker Compose và launcher Web UI.

2026-02-26 🎉 PicoClaw đạt **20K stars** chỉ trong 17 ngày! Tự động điều phối kênh và giao diện năng lực đã được triển khai.

<details>
<summary>Tin tức cũ hơn...</summary>

2026-02-16 🎉 PicoClaw đạt 12K stars chỉ trong một tuần! Vai trò maintainer cộng đồng và [roadmap](ROADMAP.md) đã được công bố chính thức.

2026-02-13 🎉 PicoClaw đạt 5000 stars trong 4 ngày! Lộ trình dự án và Nhóm phát triển đang được thiết lập.

2026-02-09 🎉 **PicoClaw chính thức ra mắt!** Được xây dựng trong 1 ngày để mang AI Agent đến phần cứng $10 với RAM <10MB. 🦐 PicoClaw, Lên Đường!

</details>

## ✨ Tính năng nổi bật

🪶 **Siêu nhẹ**: Bộ nhớ sử dụng <10MB — nhỏ hơn 99% so với OpenClaw (chức năng cốt lõi).*

💰 **Chi phí tối thiểu**: Đủ hiệu quả để chạy trên phần cứng $10 — rẻ hơn 98% so với Mac mini.

⚡️ **Khởi động siêu nhanh**: Nhanh gấp 400 lần, khởi động trong <1 giây ngay cả trên CPU đơn nhân 0.6GHz.

🌍 **Di động thực sự**: Một file binary duy nhất chạy trên RISC-V, ARM, MIPS và x86. Một click là chạy!

🤖 **AI tự xây dựng**: Triển khai Go-native tự động — 95% mã nguồn cốt lõi được Agent tạo ra, với sự tinh chỉnh của con người.

🔌 **Hỗ trợ MCP**: Tích hợp [Model Context Protocol](https://modelcontextprotocol.io/) gốc — kết nối bất kỳ máy chủ MCP nào để mở rộng khả năng của agent.

👁️ **Pipeline Xử lý Hình ảnh**: Gửi hình ảnh và tệp trực tiếp cho agent — tự động mã hóa base64 cho các LLM đa phương thức.

🧠 **Định tuyến Thông minh**: Định tuyến mô hình dựa trên quy tắc — truy vấn đơn giản chuyển đến mô hình nhẹ, tiết kiệm chi phí API.

_*Các phiên bản gần đây có thể sử dụng 10–20MB do merge tính năng nhanh chóng. Tối ưu tài nguyên đang được lên kế hoạch. So sánh thời gian khởi động dựa trên benchmark đơn nhân 0.8GHz (xem bảng bên dưới)._

|                               | OpenClaw      | NanoBot                  | **PicoClaw**                              |
| ----------------------------- | ------------- | ------------------------ | ----------------------------------------- |
| **Ngôn ngữ**                  | TypeScript    | Python                   | **Go**                                    |
| **RAM**                       | >1GB          | >100MB                   | **< 10MB***                               |
| **Thời gian khởi động**</br>(CPU 0.8GHz) | >500s         | >30s                     | **<1s**                                   |
| **Chi phí**                   | Mac Mini $599 | Hầu hết SBC Linux ~$50  | **Mọi bo mạch Linux**</br>**Chỉ từ $10** |

<img src="assets/compare.jpg" alt="PicoClaw" width="512">

> 📋 **[Danh Sách Tương Thích Phần Cứng](docs/hardware-compatibility.md)** — Xem tất cả các board đã được kiểm tra, từ RISC-V $5 đến Raspberry Pi và điện thoại Android. Board của bạn chưa có? Gửi PR!

## 🦾 Demo

### 🛠️ Quy trình trợ lý tiêu chuẩn

<table align="center">
  <tr align="center">
    <th><p align="center">🧩 Lập trình Full-Stack</p></th>
    <th><p align="center">🗂️ Quản lý Nhật ký & Kế hoạch</p></th>
    <th><p align="center">🔎 Tìm kiếm Web & Học hỏi</p></th>
  </tr>
  <tr>
    <td align="center"><p align="center"><img src="assets/picoclaw_code.gif" width="240" height="180"></p></td>
    <td align="center"><p align="center"><img src="assets/picoclaw_memory.gif" width="240" height="180"></p></td>
    <td align="center"><p align="center"><img src="assets/picoclaw_search.gif" width="240" height="180"></p></td>
  </tr>
  <tr>
    <td align="center">Phát triển • Triển khai • Mở rộng</td>
    <td align="center">Lên lịch • Tự động hóa • Ghi nhớ</td>
    <td align="center">Khám phá • Phân tích • Xu hướng</td>
  </tr>
</table>

### 📱 Chạy trên điện thoại Android cũ

Hãy cho chiếc điện thoại cũ một cuộc sống mới! Biến nó thành trợ lý AI thông minh với PicoClaw. Bắt đầu nhanh:

1. **Cài đặt [Termux](https://github.com/termux/termux-app)** (Tải từ [GitHub Releases](https://github.com/termux/termux-app/releases), hoặc tìm trên F-Droid / Google Play).
2. **Chạy các lệnh**

```bash
# Tải phiên bản mới nhất từ https://github.com/sipeed/picoclaw/releases
wget https://github.com/sipeed/picoclaw/releases/latest/download/picoclaw_Linux_arm64.tar.gz
tar xzf picoclaw_Linux_arm64.tar.gz
pkg install proot
termux-chroot ./picoclaw onboard   # chroot cung cấp bố cục hệ thống tệp Linux tiêu chuẩn
```

Sau đó làm theo hướng dẫn trong phần "Bắt đầu nhanh" để hoàn tất cấu hình!

<img src="assets/termux.jpg" alt="PicoClaw" width="512">

### 🐜 Triển khai sáng tạo trên phần cứng tối thiểu

PicoClaw có thể triển khai trên hầu hết mọi thiết bị Linux!

- $9.9 [LicheeRV-Nano](https://www.aliexpress.com/item/1005006519668532.html) phiên bản E(Ethernet) hoặc W(WiFi6), dùng làm Trợ lý Gia đình tối giản
- $30~50 [NanoKVM](https://www.aliexpress.com/item/1005007369816019.html), hoặc $100 [NanoKVM-Pro](https://www.aliexpress.com/item/1005010048471263.html) dùng cho quản trị Server tự động
- $50 [MaixCAM](https://www.aliexpress.com/item/1005008053333693.html) hoặc $100 [MaixCAM2](https://www.kickstarter.com/projects/zepan/maixcam2-build-your-next-gen-4k-ai-camera) dùng cho Giám sát thông minh

<https://private-user-images.githubusercontent.com/83055338/547056448-e7b031ff-d6f5-4468-bcca-5726b6fecb5c.mp4>

🌟 Nhiều hình thức triển khai hơn đang chờ bạn khám phá!

## 📦 Cài đặt

### Tải từ picoclaw.io (Khuyến nghị)

Truy cập **[picoclaw.io](https://picoclaw.io)** — trang web chính thức tự động phát hiện nền tảng của bạn và cung cấp tải xuống một cú nhấp. Không cần chọn kiến trúc thủ công.

### Tải binary đã biên dịch sẵn

Hoặc tải binary cho nền tảng của bạn từ trang [GitHub Releases](https://github.com/sipeed/picoclaw/releases).

### Biên dịch từ mã nguồn (cho phát triển)

```bash
git clone https://github.com/sipeed/picoclaw.git

cd picoclaw
make deps

# Build (không cần cài đặt)
make build

# Build cho nhiều nền tảng
make build-all

# Build cho Raspberry Pi Zero 2 W (32-bit: make build-linux-arm; 64-bit: make build-linux-arm64)
make build-pi-zero

# Build và cài đặt
make install
```

**Raspberry Pi Zero 2 W:** Sử dụng binary phù hợp với hệ điều hành: Raspberry Pi OS 32-bit → `make build-linux-arm`; 64-bit → `make build-linux-arm64`. Hoặc chạy `make build-pi-zero` để build cả hai.

## 📚 Tài liệu

Để xem hướng dẫn chi tiết, tham khảo tài liệu bên dưới. README này chỉ bao gồm phần bắt đầu nhanh.

| Chủ đề | Mô tả |
|--------|-------|
| 🐳 [Docker & Bắt đầu nhanh](docs/vi/docker.md) | Thiết lập Docker Compose, chế độ Launcher/Agent, cấu hình Bắt đầu nhanh |
| 💬 [Ứng dụng Chat](docs/vi/chat-apps.md) | Telegram, Discord, WhatsApp, Matrix, QQ, Slack, IRC, DingTalk, LINE, Feishu, WeCom và nhiều hơn |
| ⚙️ [Cấu hình](docs/vi/configuration.md) | Biến môi trường, cấu trúc workspace, nguồn skill, sandbox bảo mật, heartbeat |
| 🔌 [Nhà cung cấp & Mô hình](docs/vi/providers.md) | 20+ nhà cung cấp LLM, định tuyến mô hình, cấu hình model_list, kiến trúc nhà cung cấp |
| 🔄 [Spawn & Tác vụ bất đồng bộ](docs/vi/spawn-tasks.md) | Tác vụ nhanh, tác vụ dài với spawn, điều phối sub-agent bất đồng bộ |
| 🐛 [Xử lý sự cố](docs/vi/troubleshooting.md) | Các vấn đề thường gặp và giải pháp |
| 🔧 [Cấu hình Công cụ](docs/vi/tools_configuration.md) | Bật/tắt từng công cụ, chính sách thực thi |
| 📋 [Tương Thích Phần Cứng](docs/hardware-compatibility.md) | Các board đã kiểm tra, yêu cầu tối thiểu, cách thêm board |

## <img src="assets/clawdchat-icon.png" width="24" height="24" alt="ClawdChat"> Tham gia Mạng xã hội Agent

Kết nối PicoClaw với Mạng xã hội Agent chỉ bằng cách gửi một tin nhắn qua CLI hoặc bất kỳ ứng dụng Chat nào đã tích hợp.

**Đọc `https://clawdchat.ai/skill.md` và làm theo hướng dẫn để tham gia [ClawdChat.ai](https://clawdchat.ai)**

<<<<<<< HEAD
## 🖥️ Tham chiếu CLI

| Lệnh                      | Mô tả                         |
| -------------------------- | ------------------------------ |
| `picoclaw onboard`         | Khởi tạo cấu hình & workspace |
| `picoclaw agent -m "..."`  | Trò chuyện với agent           |
| `picoclaw agent`           | Chế độ chat tương tác          |
| `picoclaw gateway`         | Khởi động gateway              |
| `picoclaw status`          | Hiển thị trạng thái            |
| `picoclaw version`         | Hiển thị thông tin phiên bản   |
| `picoclaw cron list`       | Liệt kê tất cả tác vụ định kỳ |
| `picoclaw cron add ...`    | Thêm tác vụ định kỳ           |
| `picoclaw cron disable`    | Tắt tác vụ định kỳ            |
| `picoclaw cron remove`     | Xóa tác vụ định kỳ            |
| `picoclaw skills list`     | Liệt kê các skill đã cài      |
| `picoclaw skills install`  | Cài đặt một skill              |
| `picoclaw migrate`         | Di chuyển dữ liệu từ phiên bản cũ |
| `picoclaw auth login`      | Xác thực với nhà cung cấp     |
| `picoclaw model`           | Xem hoặc chuyển đổi model mặc định |
=======
## ⚙️ Cấu hình chi tiết

File cấu hình: `~/.picoclaw/config.json`

### Biến môi trường

Bạn có thể ghi đè các đường dẫn mặc định bằng cách sử dụng các biến môi trường. Điều này hữu ích cho việc cài đặt di động, triển khai container hóa hoặc chạy picoclaw như một dịch vụ hệ thống. Các biến này độc lập và kiểm soát các đường dẫn khác nhau.

| Biến              | Mô tả                                                                                                                             | Đường dẫn mặc định        |
|-------------------|-----------------------------------------------------------------------------------------------------------------------------------------|---------------------------|
| `PICOCLAW_CONFIG` | Ghi đè đường dẫn đến file cấu hình. Điều này trực tiếp yêu cầu picoclaw tải file `config.json` nào, bỏ qua tất cả các vị trí khác. | `~/.picoclaw/config.json` |
| `PICOCLAW_HOME`   | Ghi đè thư mục gốc cho dữ liệu picoclaw. Điều này thay đổi vị trí mặc định của `workspace` và các thư mục dữ liệu khác.          | `~/.picoclaw`             |

**Ví dụ:**

```bash
# Chạy picoclaw bằng một file cấu hình cụ thể
# Đường dẫn workspace sẽ được đọc từ trong file cấu hình đó
PICOCLAW_CONFIG=/etc/picoclaw/production.json picoclaw gateway

# Chạy picoclaw với tất cả dữ liệu được lưu trữ trong /opt/picoclaw
# Cấu hình sẽ được tải từ ~/.picoclaw/config.json mặc định
# Workspace sẽ được tạo tại /opt/picoclaw/workspace
PICOCLAW_HOME=/opt/picoclaw picoclaw agent

# Sử dụng cả hai để có thiết lập tùy chỉnh hoàn toàn
PICOCLAW_HOME=/srv/picoclaw PICOCLAW_CONFIG=/srv/picoclaw/main.json picoclaw gateway
```

### Cấu trúc Workspace

PicoClaw lưu trữ dữ liệu trong workspace đã cấu hình (mặc định: `~/.picoclaw/workspace`):

```
~/.picoclaw/workspace/
├── sessions/          # Phiên hội thoại và lịch sử
├── memory/           # Bộ nhớ dài hạn (MEMORY.md)
├── state/            # Trạng thái lưu trữ (kênh cuối cùng, v.v.)
├── cron/             # Cơ sở dữ liệu tác vụ định kỳ
├── skills/           # Kỹ năng tùy chỉnh
├── AGENT.md          # Định nghĩa agent có cấu trúc và system prompt
├── HEARTBEAT.md      # Prompt tác vụ định kỳ (kiểm tra mỗi 30 phút)
├── SOUL.md           # Tâm hồn/Tính cách Agent
└── ...
```

### 🔒 Hộp cát bảo mật (Security Sandbox)

PicoClaw chạy trong môi trường sandbox theo mặc định. Agent chỉ có thể truy cập file và thực thi lệnh trong phạm vi workspace.

#### Cấu hình mặc định

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

| Tùy chọn | Mặc định | Mô tả |
|----------|---------|-------|
| `workspace` | `~/.picoclaw/workspace` | Thư mục làm việc của agent |
| `restrict_to_workspace` | `true` | Giới hạn truy cập file/lệnh trong workspace |

#### Công cụ được bảo vệ

Khi `restrict_to_workspace: true`, các công cụ sau bị giới hạn trong sandbox:

| Công cụ | Chức năng | Giới hạn |
|---------|----------|---------|
| `read_file` | Đọc file | Chỉ file trong workspace |
| `write_file` | Ghi file | Chỉ file trong workspace |
| `list_dir` | Liệt kê thư mục | Chỉ thư mục trong workspace |
| `edit_file` | Sửa file | Chỉ file trong workspace |
| `append_file` | Thêm vào file | Chỉ file trong workspace |
| `exec` | Thực thi lệnh | Đường dẫn lệnh phải trong workspace |

#### Bảo vệ bổ sung cho Exec

Ngay cả khi `restrict_to_workspace: false`, công cụ `exec` vẫn chặn các lệnh nguy hiểm sau:

* `rm -rf`, `del /f`, `rmdir /s` — Xóa hàng loạt
* `format`, `mkfs`, `diskpart` — Định dạng ổ đĩa
* `dd if=` — Tạo ảnh đĩa
* Ghi vào `/dev/sd[a-z]` — Ghi trực tiếp lên đĩa
* `shutdown`, `reboot`, `poweroff` — Tắt/khởi động lại hệ thống
* Fork bomb `:(){ :|:& };:`

#### Ví dụ lỗi

```
[ERROR] tool: Tool execution failed
{tool=exec, error=Command blocked by safety guard (path outside working dir)}
```

```
[ERROR] tool: Tool execution failed
{tool=exec, error=Command blocked by safety guard (dangerous pattern detected)}
```

#### Tắt giới hạn (Rủi ro bảo mật)

Nếu bạn cần agent truy cập đường dẫn ngoài workspace:

**Cách 1: File cấu hình**

```json
{
  "agents": {
    "defaults": {
      "restrict_to_workspace": false
    }
  }
}
```

**Cách 2: Biến môi trường**

```bash
export PICOCLAW_AGENTS_DEFAULTS_RESTRICT_TO_WORKSPACE=false
```

> ⚠️ **Cảnh báo**: Tắt giới hạn này cho phép agent truy cập mọi đường dẫn trên hệ thống. Chỉ sử dụng cẩn thận trong môi trường được kiểm soát.

#### Tính nhất quán của ranh giới bảo mật

Cài đặt `restrict_to_workspace` áp dụng nhất quán trên mọi đường thực thi:

| Đường thực thi | Ranh giới bảo mật |
|----------------|-------------------|
| Agent chính | `restrict_to_workspace` ✅ |
| Subagent / Spawn | Kế thừa cùng giới hạn ✅ |
| Tác vụ Heartbeat | Kế thừa cùng giới hạn ✅ |

Tất cả đường thực thi chia sẻ cùng giới hạn workspace — không có cách nào vượt qua ranh giới bảo mật thông qua subagent hoặc tác vụ định kỳ.

### Heartbeat (Tác vụ định kỳ)

PicoClaw có thể tự động thực hiện các tác vụ định kỳ. Tạo file `HEARTBEAT.md` trong workspace:

```markdown
# Tác vụ định kỳ

- Kiểm tra email xem có tin nhắn quan trọng không
- Xem lại lịch cho các sự kiện sắp tới
- Kiểm tra dự báo thời tiết
```

Agent sẽ đọc file này mỗi 30 phút (có thể cấu hình) và thực hiện các tác vụ bằng công cụ có sẵn.

#### Tác vụ bất đồng bộ với Spawn

Đối với các tác vụ chạy lâu (tìm kiếm web, gọi API), sử dụng công cụ `spawn` để tạo **subagent**:

```markdown
# Tác vụ định kỳ

## Tác vụ nhanh (trả lời trực tiếp)
- Báo cáo thời gian hiện tại

## Tác vụ lâu (dùng spawn cho async)
- Tìm kiếm tin tức AI trên web và tóm tắt
- Kiểm tra email và báo cáo tin nhắn quan trọng
```

**Hành vi chính:**

| Tính năng | Mô tả |
|-----------|-------|
| **spawn** | Tạo subagent bất đồng bộ, không chặn heartbeat |
| **Context độc lập** | Subagent có context riêng, không có lịch sử phiên |
| **message tool** | Subagent giao tiếp trực tiếp với người dùng qua công cụ message |
| **Không chặn** | Sau khi spawn, heartbeat tiếp tục tác vụ tiếp theo |

#### Cách Subagent giao tiếp

```
Heartbeat kích hoạt
    ↓
Agent đọc HEARTBEAT.md
    ↓
Tác vụ lâu: spawn subagent
    ↓                           ↓
Tiếp tục tác vụ tiếp theo   Subagent làm việc độc lập
    ↓                           ↓
Tất cả tác vụ hoàn thành    Subagent dùng công cụ "message"
    ↓                           ↓
Phản hồi HEARTBEAT_OK       Người dùng nhận kết quả trực tiếp
```

Subagent có quyền truy cập các công cụ (message, web_search, v.v.) và có thể giao tiếp với người dùng một cách độc lập mà không cần thông qua agent chính.

**Cấu hình:**

```json
{
  "heartbeat": {
    "enabled": true,
    "interval": 30
  }
}
```

| Tùy chọn | Mặc định | Mô tả |
|----------|---------|-------|
| `enabled` | `true` | Bật/tắt heartbeat |
| `interval` | `30` | Khoảng thời gian kiểm tra (phút, tối thiểu: 5) |

**Biến môi trường:**

* `PICOCLAW_HEARTBEAT_ENABLED=false` để tắt
* `PICOCLAW_HEARTBEAT_INTERVAL=60` để thay đổi khoảng thời gian

### Nhà cung cấp (Providers)

> [!NOTE]
> Groq cung cấp dịch vụ chuyển giọng nói thành văn bản miễn phí qua Whisper. Nếu đã cấu hình Groq, tin nhắn âm thanh từ bất kỳ kênh nào sẽ được tự động chuyển thành văn bản ở cấp độ agent.

| Nhà cung cấp | Mục đích | Lấy API Key |
| --- | --- | --- |
| `gemini` | LLM (Gemini trực tiếp) | [aistudio.google.com](https://aistudio.google.com) |
| `zhipu` | LLM (Zhipu trực tiếp) | [bigmodel.cn](bigmodel.cn) |
| `volcengine`             | LLM(Volcengine trực tiếp)                   | [volcengine.com](https://www.volcengine.com/activity/codingplan?utm_campaign=PicoClaw&utm_content=PicoClaw&utm_medium=devrel&utm_source=OWO&utm_term=PicoClaw)           |
| `openrouter` (Đang thử nghiệm) | LLM (khuyên dùng, truy cập mọi model) | [openrouter.ai](https://openrouter.ai) |
| `anthropic` (Đang thử nghiệm) | LLM (Claude trực tiếp) | [console.anthropic.com](https://console.anthropic.com) |
| `openai` (Đang thử nghiệm) | LLM (GPT trực tiếp) | [platform.openai.com](https://platform.openai.com) |
| `deepseek` (Đang thử nghiệm) | LLM (DeepSeek trực tiếp) | [platform.deepseek.com](https://platform.deepseek.com) |
| `groq` | LLM + **Chuyển giọng nói** (Whisper) | [console.groq.com](https://console.groq.com) |
| `qwen` | LLM (Qwen trực tiếp) | [dashscope.console.aliyun.com](https://dashscope.console.aliyun.com) |
| `cerebras` | LLM (Cerebras trực tiếp) | [cerebras.ai](https://cerebras.ai) |

<details>
<summary><b>Cấu hình Zhipu</b></summary>

**1. Lấy API key**

* Lấy [API key](https://bigmodel.cn/usercenter/proj-mgmt/apikeys)

**2. Cấu hình**

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

**3. Chạy**

```bash
picoclaw agent -m "Xin chào"
```

</details>

<details>
<summary><b>Ví dụ cấu hình đầy đủ</b></summary>

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
    }
  },
  "heartbeat": {
    "enabled": true,
    "interval": 30
  }
}
```

</details>

### Cấu hình Mô hình (model_list)

> **Tính năng mới!** PicoClaw hiện sử dụng phương pháp cấu hình **đặt mô hình vào trung tâm**. Chỉ cần chỉ định dạng `nhà cung cấp/mô hình` (ví dụ: `zhipu/glm-4.7`) để thêm nhà cung cấp mới—**không cần thay đổi mã!**

Thiết kế này cũng cho phép **hỗ trợ đa tác nhân** với lựa chọn nhà cung cấp linh hoạt:

- **Tác nhân khác nhau, nhà cung cấp khác nhau** : Mỗi tác nhân có thể sử dụng nhà cung cấp LLM riêng
- **Mô hình dự phòng** : Cấu hình mô hình chính và dự phòng để tăng độ tin cậy
- **Cân bằng tải** : Phân phối yêu cầu trên nhiều endpoint khác nhau
- **Cấu hình tập trung** : Quản lý tất cả nhà cung cấp ở một nơi

#### 📋 Tất cả Nhà cung cấp được Hỗ trợ

| Nhà cung cấp | Prefix `model` | API Base Mặc định | Giao thức | Khóa API |
|-------------|----------------|-------------------|-----------|----------|
| **OpenAI** | `openai/` | `https://api.openai.com/v1` | OpenAI | [Lấy Khóa](https://platform.openai.com) |
| **Anthropic** | `anthropic/` | `https://api.anthropic.com/v1` | Anthropic | [Lấy Khóa](https://console.anthropic.com) |
| **Zhipu AI (GLM)** | `zhipu/` | `https://open.bigmodel.cn/api/paas/v4` | OpenAI | [Lấy Khóa](https://open.bigmodel.cn/usercenter/proj-mgmt/apikeys) |
| **DeepSeek** | `deepseek/` | `https://api.deepseek.com/v1` | OpenAI | [Lấy Khóa](https://platform.deepseek.com) |
| **Google Gemini** | `gemini/` | `https://generativelanguage.googleapis.com/v1beta` | OpenAI | [Lấy Khóa](https://aistudio.google.com/api-keys) |
| **Groq** | `groq/` | `https://api.groq.com/openai/v1` | OpenAI | [Lấy Khóa](https://console.groq.com) |
| **Moonshot** | `moonshot/` | `https://api.moonshot.cn/v1` | OpenAI | [Lấy Khóa](https://platform.moonshot.cn) |
| **Qwen (Alibaba)** | `qwen/` | `https://dashscope.aliyuncs.com/compatible-mode/v1` | OpenAI | [Lấy Khóa](https://dashscope.console.aliyun.com) |
| **NVIDIA** | `nvidia/` | `https://integrate.api.nvidia.com/v1` | OpenAI | [Lấy Khóa](https://build.nvidia.com) |
| **Ollama** | `ollama/` | `http://localhost:11434/v1` | OpenAI | Local (không cần khóa) |
| **OpenRouter** | `openrouter/` | `https://openrouter.ai/api/v1` | OpenAI | [Lấy Khóa](https://openrouter.ai/keys) |
| **VLLM** | `vllm/` | `http://localhost:8000/v1` | OpenAI | Local |
| **Cerebras** | `cerebras/` | `https://api.cerebras.ai/v1` | OpenAI | [Lấy Khóa](https://cerebras.ai) |
| **VolcEngine (Doubao)** | `volcengine/` | `https://ark.cn-beijing.volces.com/api/v3` | OpenAI | [Lấy Khóa](https://www.volcengine.com/activity/codingplan?utm_campaign=PicoClaw&utm_content=PicoClaw&utm_medium=devrel&utm_source=OWO&utm_term=PicoClaw) |
| **ShengsuanYun** | `shengsuanyun/` | `https://router.shengsuanyun.com/api/v1` | OpenAI | - |
| **BytePlus**        | `byteplus/`       | `https://ark.ap-southeast.bytepluses.com/api/v3`    | OpenAI    | [Lấy Khóa](https://www.byteplus.com)                      |
| **LongCat**         | `longcat/`        | `https://api.longcat.chat/openai`                   | OpenAI    | [Lấy Key](https://longcat.chat/platform)                        |
| **ModelScope (魔搭)**| `modelscope/`    | `https://api-inference.modelscope.cn/v1`            | OpenAI    | [Lấy Token](https://modelscope.cn/my/tokens)                    |
| **Antigravity** | `antigravity/` | Google Cloud | Tùy chỉnh | Chỉ OAuth |
| **GitHub Copilot** | `github-copilot/` | `localhost:4321` | gRPC | - |

#### Cấu hình Cơ bản

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

#### Ví dụ theo Nhà cung cấp

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

**Anthropic (với OAuth)**
```json
{
  "model_name": "claude-sonnet-4.6",
  "model": "anthropic/claude-sonnet-4.6",
  "auth_method": "oauth"
}
```
> Chạy `picoclaw auth login --provider anthropic` để thiết lập thông tin xác thực OAuth.

**Proxy/API tùy chỉnh**
```json
{
  "model_name": "my-custom-model",
  "model": "openai/custom-model",
  "api_base": "https://my-proxy.com/v1",
  "api_key": "sk-...",
  "request_timeout": 300
}
```

#### Cân bằng Tải tải

Định cấu hình nhiều endpoint cho cùng một tên mô hình—PicoClaw sẽ tự động phân phối round-robin giữa chúng:

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

#### Chuyển đổi từ Cấu hình `providers` Cũ

Cấu hình `providers` cũ đã **ngừng sử dụng** nhưng vẫn được hỗ trợ để tương thích ngược.

**Cấu hình Cũ (đã ngừng sử dụng):**
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

**Cấu hình Mới (khuyến nghị):**
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

Xem hướng dẫn chuyển đổi chi tiết tại [docs/migration/model-list-migration.md](docs/migration/model-list-migration.md).

## Tham chiếu CLI

| Lệnh | Mô tả |
| --- | --- |
| `picoclaw onboard` | Khởi tạo cấu hình & workspace |
| `picoclaw agent -m "..."` | Trò chuyện với agent |
| `picoclaw agent` | Chế độ chat tương tác |
| `picoclaw gateway` | Khởi động gateway (cho bot chat) |
| `picoclaw status` | Hiển thị trạng thái |
| `picoclaw cron list` | Liệt kê tất cả tác vụ định kỳ |
| `picoclaw cron add ...` | Thêm tác vụ định kỳ |
>>>>>>> refactor/agent

### Tác vụ định kỳ / Nhắc nhở

PicoClaw hỗ trợ nhắc nhở theo lịch và tác vụ lặp lại thông qua công cụ `cron`:

* **Nhắc nhở một lần**: "Nhắc tôi sau 10 phút" → kích hoạt một lần sau 10 phút
* **Tác vụ lặp lại**: "Nhắc tôi mỗi 2 giờ" → kích hoạt mỗi 2 giờ
* **Biểu thức Cron**: "Nhắc tôi lúc 9 giờ sáng mỗi ngày" → sử dụng biểu thức cron

## 🤝 Đóng góp & Lộ trình

Chào đón mọi PR! Mã nguồn được thiết kế nhỏ gọn và dễ đọc. 🤗

Xem [Lộ trình Cộng đồng](https://github.com/sipeed/picoclaw/blob/main/ROADMAP.md) đầy đủ.

Nhóm phát triển đang được xây dựng. Tham gia sau khi có PR đầu tiên được merge!

Nhóm người dùng:

discord: <https://discord.gg/V4sAZ9XWpN>

<img src="assets/wechat.png" alt="PicoClaw" width="512">
