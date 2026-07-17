package cases

// Step 定义每一步的具体操作
type Step struct {
	Action string // "user_message" 或 "tool_call"
	Data   string // 具体内容
}

// Case 定义一个完整的测试剧本
type Case struct {
	Name  string
	Steps []Step
}

// GetCase1 返回第1个用例：单轮对话
func GetCase1() Case {
	return Case{
		Name: "单轮对话",
		Steps: []Step{
			{Action: "user_message", Data: "你好，请介绍一下你自己"}, // ← 注意这里结尾有逗号
		},
	}
}

// GetCase2 返回第2个用例：多轮对话（占位）
func GetCase2() Case {
	return Case{
		Name: "多轮对话",
		Steps: []Step{
			{Action: "user_message", Data: "你好，我想订一张去上海的机票"},
			{Action: "user_message", Data: "明天下午3点左右的"},
			{Action: "user_message", Data: "好的帮我确认一下订单"},
		},
	}
}

//GetCase3 返回第3个用例：工具调用对话（模拟）
func GetCase3() Case {
	return Case{
		Name: "工具调用对话",
		Steps: []Step{
			{Action: "user_message", Data: "帮我查一下今天的天气"},
			//这里暂时只存用户消息，暂时不存真正的tool_call
		},
	}
}

// GetAllCases 返回所有用例
func GetAllCases() []Case {
	return []Case{
		GetCase1(),
		GetCase2(),
		GetCase3(),
	}
}
