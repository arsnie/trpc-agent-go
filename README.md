# Session Replay Consistency Test Framework


## 项目介绍

本项目是一个针对 trpc-agent-go Session / Memory 后端的一致性回放测试框架。


项目目标：

使用同一组标准化的 Agent 执行轨迹，
分别运行在不同 Session 后端，
并比较不同后端保存的数据是否一致。


## 项目背景

在 Agent 系统中，一个 Session 不仅包含聊天记录，
还包含：

- Event（事件）
- State（状态）
- Memory（长期记忆）
- Summary（摘要）
- Track（执行轨迹）


开发阶段通常使用：

- InMemory


生产环境可能使用：

- SQLite
- Redis
- PostgreSQL
- MySQL
- ClickHouse


如果不同 Backend 数据不一致，可能导致：

- 对话历史丢失
- 上下文恢复错误
- Memory污染
- Summary异常


## 当前进度


### 已完成

- [x] 基础项目结构
- [x] InMemory Session接入
- [x] SQLite Session接入
- [x] Replay Case初版
- [x] JSON Normalizer初版


### 正在进行

- [ ] 重构Replay Case数据模型
- [ ] Backend Adapter设计
- [ ] Session Snapshot设计
- [ ] Comparator实现
- [ ] Diff Report生成


## 学习目标

通过这个项目学习：

- Go语言工程开发
- Agent Session架构
- 后端抽象设计
- 测试框架设计
- 数据一致性验证


## 技术栈

- Go
- trpc-agent-go
- SQLite
- JSON
- Git/GitHub


## 项目状态

开发中。
