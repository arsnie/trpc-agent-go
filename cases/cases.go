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

//GetCase4 返回第四个用例：State更新
func GetCase4() Case {
	return Case{
		Name: "State更新",
		Steps: []Step{
			{Action: "update_state", Data: "user_preference=vegetarian"},
			{Action: "update_state", Data: "user_rank=vip"},
		},
	}
}

//GetCase5 返回第5个用例：Memory写入和读取
func GetCase5() Case {
	return Case{
		Name: "Memory写入和读取",
		Steps: []Step{
			{Action: "memory_write", Data: "用户喜欢安静的环境"},
			{Action: "memory_read", Data: "安静的环境"},
		},
	}
}

//GetCase6 返回第六个用例：Sumory生成和更新
func GetCase6() Case {
	return Case{
		Name: "Summary生成和更新",
		Steps: []Step{
			{Action: "summary_trigger", Data: "这是一段很长很长的对话，需要被总结成简短的摘要..."},
			{Action: "summary_updata", Data: "更新摘要内容"},
		},
	}
}

//GetCase7 返回第7个用例：Summary与事件截断
func GetCase7() Case {
	return Case{
		Steps: []Step{
			{Action: "user_message", Data: "第一轮：我们讨论天气"},
			{Action: "user_message", Data: "第二轮：我们讨论交通"},
			{Action: "user_message", Data: "第三轮，我们讨论美食"},
			{Action: "user_message", Data: "触发摘要截断"},
		},
	}
}

//GetCase8 返回第8个用例：Track事件
func GetCase8() Case {
	return Case{
		Name: "Track事件",
		Steps: []Step{
			{Action: "track", Data: "tool_execution_start"},
			{Action: "track", Data: "tool_execution_end"},
		},
	}
}

//GetCase9 返回第9个用例：并发写入
func GetCase9() Case {
	return Case{
		Name: "并发写入",
		Steps: []Step{
			{Action: "concurrent_write", Data: "并行消息A"},
			{Action: "concurrent_write", Data: "并行消息B"},
			{Action: "concurrent_write", Data: "并行消息C"},
		},
	}
}

//GetCase10 返回第10个用例：异常恢复
func GetCase10() Case {
	return Case{
		Name: "异常恢复",
		Steps: []Step{
			{Action: "user_message", Data: "正常请求"},
			{Action: "user_message", Data: "模拟写入失败"},
			{Action: "retry,Data:重试写入"},
		},
	}
}

// GetAllCases 返回所有用例
func GetAllCases() []Case {
	return []Case{
		GetCase1(),
		GetCase2(),
		GetCase3(),
		GetCase4(),
		GetCase5(),
		GetCase6(),
		GetCase7(),
		GetCase8(),
		GetCase9(),
		GetCase10(),
	}
}
