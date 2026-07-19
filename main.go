package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"

	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/session"
	"trpc.group/trpc-go/trpc-agent-go/session/inmemory"
	"trpc.group/trpc-go/trpc-agent-go/session/sqlite"

	"trpc-replay-test/cases"
	"trpc-replay-test/normalizer"
)

// buildUserEvent 把一句用户话术变成标准的 Event 结构体
func buildUserEvent(content string) *event.Event {
	return &event.Event{
		Response: &model.Response{
			Choices: []model.Choice{
				{
					Message: model.Message{
						Role:    model.RoleUser,
						Content: content,
					},
				},
			},
		},
		ID:        uuid.New().String(),
		Timestamp: time.Now(),
	}
}

func main() {

	_ = os.Remove("session.db")
	ctx := context.Background()

	// 1. 初始化后端
	inmem := inmemory.NewSessionService()
	fmt.Println("✅ InMemory 后端初始化成功")

	db, err := sql.Open("sqlite3", "./session.db")
	if err != nil {
		log.Fatalf("打开 SQLite 失败: %v", err)
	}
	sqlDB, err := sqlite.NewService(db)
	if err != nil {
		log.Fatalf("SQLite 初始化失败: %v", err)
	}
	fmt.Println("✅ SQLite 后端初始化成功")

	// 2. 加载用例
	allCases := cases.GetAllCases()
	fmt.Printf("✅ 共加载 %d 个测试用例\n", len(allCases))

	// 3. 遍历每个 Case
	for _, tc := range allCases {
		key := session.Key{
			AppName:   "replay-test",
			UserID:    "tester",
			SessionID: tc.Name,
		}

		// 3.1 先在两个后端创建 Session
		_, err := inmem.CreateSession(ctx, key, session.StateMap{})
		if err != nil {
			log.Printf("⚠️ InMemory 创建失败 [%s]: %v", tc.Name, err)
			continue
		}
		_, err = sqlDB.CreateSession(ctx, key, session.StateMap{})
		if err != nil {
			log.Printf("⚠️ SQLite 创建失败 [%s]: %v", tc.Name, err)
			continue
		}

		// 3.2 针对这个 Case 里的每一步，写入事件
		for _, step := range tc.Steps {
			switch step.Action {
			case "user_message":
				evt := buildUserEvent(step.Data)

				// ---- 写入 InMemory ----
				inmemSess, err := inmem.GetSession(ctx, key)
				if err != nil {
					log.Printf("⚠️ InMemory 获取 Session 失败: %v", err)
					break
				}
				err = inmem.AppendEvent(ctx, inmemSess, evt)
				if err != nil {
					log.Printf("⚠️ InMemory 追加事件失败: %v", err)
				}

				// ---- 写入 SQLite ----
				sqlSess, err := sqlDB.GetSession(ctx, key)
				if err != nil {
					log.Printf("⚠️ SQLite 获取 Session 失败: %v", err)
					continue
				}
				err = sqlDB.AppendEvent(ctx, sqlSess, evt)
				if err != nil {
					log.Printf("⚠️ SQLite 追加事件失败: %v", err)
				}
			case "update_state":
				// step.Data 的格式是 "key=value"，比如 "user_preference=vegetarian"
				parts := strings.SplitN(step.Data, "=", 2)
				if len(parts) != 2 {
					log.Printf("⚠️ 无效的 State 数据格式: %s", step.Data)
					break
				}
				statekey := parts[0]
				value := parts[1]

				// 构建 StateMap
				stateMap := session.StateMap{
					statekey: []byte(value),
				}

				// 写入 InMemory
				err := inmem.UpdateSessionState(ctx, key, stateMap)
				if err != nil {
					log.Printf("⚠️ InMemory 更新 State 失败: %v", err)
				}

				// 写入 SQLite
				err = sqlDB.UpdateSessionState(ctx, key, stateMap)
				if err != nil {
					log.Printf("⚠️ SQLite 更新 State 失败: %v", err)
				}

			default:
				log.Printf("⚠️ 未知的 Action 类型: %s (Case: %s)", step.Action, tc.Name)
			}
		}

		// 3.3 写入完成后，从两个后端读取完整的 Session 数据
		inmemSess, err := inmem.GetSession(ctx, key)
		if err != nil {
			log.Printf("⚠️ InMemory 读取失败: %v", err)
			continue
		}
		sqlSess, err := sqlDB.GetSession(ctx, key)
		if err != nil {
			log.Printf("⚠️ SQLite 读取失败: %v", err)
			continue
		}
		// 3.4 把两个后端读出的数据转成 JSON，并归一化
		fmt.Printf("\n📝 Case: %s\n", tc.Name)

		// InMemory：先转成 JSON，再归一化
		inmemRaw, _ := json.MarshalIndent(inmemSess, "", "  ")
		inmemNormalized, err := normalizer.NormalizeJSON(string(inmemRaw))
		if err != nil {
			log.Printf("⚠️ InMemory 归一化失败: %v", err)
		}
		fmt.Printf("  [InMemory 归一化后]\n  %s\n", inmemNormalized)

		// SQLite：同理
		sqlRaw, _ := json.MarshalIndent(sqlSess, "", "  ")
		sqlNormalized, err := normalizer.NormalizeJSON(string(sqlRaw))
		if err != nil {
			log.Printf("⚠️ SQLite 归一化失败: %v", err)
		}
		fmt.Printf("  [SQLite 归一化后]\n  %s\n", sqlNormalized)

	}

	fmt.Println("\n🎉 数据写入并读取完成！请观察上面两个 JSON 的差异（重点关注 ID 和时间戳）。")
}
