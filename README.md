# GoHosts

<p align="center">
  <img src="build/appicon.png" alt="GoHosts Logo" width="128" height="128">
</p>

<p align="center">
  <strong>一款轻量、常驻、可自动更新的 Hosts 管理工具</strong>
</p>

<p align="center">
  <a href="#-主要功能">主要功能</a> •
  <a href="#-下载安装">下载安装</a> •
  <a href="#-使用指南">使用指南</a> •
  <a href="#-常见问题">常见问题</a>
</p>

---

## 🎯 主要功能

- **轻量小巧**：安装包仅 15MB 左右，运行时占用内存极低，不拖慢电脑。
- **常驻后台**：关闭窗口后仍可在系统托盘运行，定时自动更新 Hosts，无需手动操作。
- **多方案一键切换**：可创建开发、测试、生产等多个 Hosts 方案，点击即可快速切换。
- **远程自动同步**：支持从 URL 或 Git 仓库拉取远程 Hosts，可设置定时更新间隔。
- **智能备份**：每次修改前自动备份系统原始 Hosts 文件，出问题可一键恢复。
- **跨平台通用**：完美支持 Windows、macOS 和 Linux 系统。

---

## 📥 下载安装

访问 [Releases 页面](https://github.com/yourusername/gohosts/releases) 下载最新版本。

| 操作系统 | 安装包 |
| :--- | :--- |
| **Windows** | `GoHosts-Setup-x64.exe` |
| **macOS (Intel)** | `GoHosts-darwin-amd64.dmg` |
| **macOS (Apple Silicon)** | `GoHosts-darwin-arm64.dmg` |
| **Linux** | `GoHosts-linux-amd64.AppImage` 或 `.deb` / `.rpm` |

### Windows 用户注意

安装和运行时可能需要**管理员权限**才能修改系统 Hosts 文件。首次运行时会自动提示授权。

### macOS / Linux 用户注意

由于系统 Hosts 文件位于 `/etc/hosts`，应用在写入时会请求管理员密码，请放心授权。

---

## 📖 使用指南

### 1️⃣ 首次启动

- 启动后，应用会出现在**系统托盘**中（屏幕右下角或右上角）。
- 右键点击托盘图标，选择 **“打开主面板”** 进入管理界面。

### 2️⃣ 查看当前 Hosts

- 主面板首页会显示当前系统生效的 Hosts 内容。
- 点击 **“刷新”** 按钮可重新读取最新内容。

### 3️⃣ 创建与管理方案

- 进入 **“方案管理”** 页面，可创建多个配置方案（如“开发环境”、“测试环境”）。
- 每个方案可填写自定义的 Hosts 内容，或配置一个远程 URL 自动获取。
- 支持编辑、删除、启用/禁用方案。

### 4️⃣ 切换方案

- 在方案列表中，点击方案右侧的 **“应用”** 按钮即可立即切换。
- 托盘右键菜单中也提供快捷切换列表，无需打开主窗口。

### 5️⃣ 配置自动更新（远程方案）

- 创建或编辑方案时，选择 **“远程”** 类型。
- 填写远程 Hosts 文件的 URL（支持 HTTP/HTTPS），并设置更新频率（如每小时）。
- 开启 **“启用定时更新”** 后，后台服务会按设定时间自动拉取并合并到系统 Hosts。

### 6️⃣ 恢复备份

- 每次修改系统 Hosts 前，应用都会自动备份。
- 备份文件保存在 `用户目录/.gohosts/backups/` 下，需要时可手动复制恢复。

---

## ❓ 常见问题

### Q1：为什么修改后浏览器仍访问旧地址？

部分系统和浏览器会缓存 DNS 解析结果。可尝试以下方法刷新：
- **Windows**：打开 CMD 执行 `ipconfig /flushdns`
- **macOS**：打开终端执行 `sudo dscacheutil -flushcache; sudo killall -HUP mDNSResponder`
- **Linux**：重启网络服务或执行 `sudo systemctl restart systemd-resolved`

GoHosts 会在切换方案时自动尝试刷新 DNS 缓存（部分系统可能需要授权）。

### Q2：远程 URL 拉取失败怎么办？

- 请检查 URL 是否可直接在浏览器中访问。
- 如果 URL 需要身份验证（如私有 Git 仓库），目前版本暂不支持，后续会添加 Token 配置功能。
- 检查网络代理设置，确保应用能正常访问外网。

### Q3：如何彻底退出后台服务？

右键点击托盘图标，选择 **“退出”** 即可完全关闭应用。

### Q4：配置文件存储在哪里？

所有方案配置保存在 `用户目录/.gohosts/config.json` 中，需要备份或迁移时可复制该文件。

### Q5：如何卸载？

- **Windows**：通过“控制面板 → 程序和功能”卸载，或直接删除安装目录。
- **macOS**：将 `GoHosts.app` 拖入废纸篓。
- **Linux**：使用包管理器卸载，或删除 `AppImage` 文件。

卸载不会删除备份文件和配置文件，如需彻底清除可手动删除 `用户目录/.gohosts/` 文件夹。

---

## 📝 反馈与建议

如果遇到 Bug 或有功能建议，欢迎到 [Issues](https://github.com/yourusername/gohosts/issues) 页面提交。

---

## 📄 许可证

本项目基于 [MIT License](LICENSE) 开源。

---

**⭐ 如果觉得有用，不妨给个 Star 支持一下！**