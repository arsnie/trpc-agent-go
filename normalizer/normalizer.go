package normalizer

import (
	"encoding/json"
	"fmt"
	"strings"
)

// NormalizeJSON 接收 JSON 字符串，返回归一化后的 JSON 字符串
func NormalizeJSON(raw string) (string, error) {
	var data interface{}
	err := json.Unmarshal([]byte(raw), &data)
	if err != nil {
		return "", fmt.Errorf("解析 JSON 失败: %v", err)
	}
	cleanData := cleanValue(data)
	// 使用 json.MarshalIndent 并设置 SetEscapeHTML(false)
	buffer := &strings.Builder{}
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent("", "  ")
	encoder.SetEscapeHTML(false) // 关闭 HTML 转义，让 <IGNORED> 正常显示
	err = encoder.Encode(cleanData)
	if err != nil {
		return "", fmt.Errorf("生成 JSON 失败: %v", err)
	}
	result := strings.TrimSpace(buffer.String())
	if err != nil {
		return "", fmt.Errorf("生成 JSON 失败: %v", err)
	}
	return string(result), nil
}

// cleanValue 递归清洗数据
func cleanValue(v interface{}) interface{} {
	switch val := v.(type) {
	case map[string]interface{}:
		newMap := make(map[string]interface{})
		for key, value := range val {
			if isIgnoredKey(key) {
				newMap[key] = "<IGNORED>"
			} else {
				newMap[key] = cleanValue(value)
			}
		}
		return newMap
	case []interface{}:
		newSlice := make([]interface{}, len(val))
		for i, item := range val {
			newSlice[i] = cleanValue(item) // 递归调用自身
		}
		return newSlice
	default:
		return v
	}
}

// isIgnoredKey 判断字段是否应该被忽略
func isIgnoredKey(key string) bool {
	ignoredKeys := map[string]bool{
		"id":           true,
		"timestamp":    true,
		"createdAt":    true,
		"updatedAt":    true,
		"invocationId": true,
		"response":     true, // 因为 response 里也有 timestamp，整体忽略
	}
	return ignoredKeys[key]
}
