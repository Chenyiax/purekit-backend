package service

import (
	"strings"
	"unicode/utf8"
)

type TextService interface {
	Process(text string, action string) (string, interface{}, error)
}

type textService struct{}

func NewTextService() TextService {
	return &textService{}
}

func (s *textService) Process(text string, action string) (string, interface{}, error) {
	var result string

	// 计算统计信息
	stats := map[string]int{
		"charCount": utf8.RuneCountInString(text),
		"wordCount": len(strings.Fields(text)),
		"lineCount": len(strings.Split(text, "\n")),
	}

	switch action {
	case "upper":
		result = strings.ToUpper(text)
	case "lower":
		result = strings.ToLower(text)
	case "reverse":
		runes := []rune(text)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		result = string(runes)
	case "trim":
		result = strings.TrimSpace(text)
	case "collapse":
		// 将多个空格/制表符等连续空白替换为单个空格
		words := strings.Fields(text)
		result = strings.Join(words, " ")
	case "cnToEn":
		result = convertSymbols(text, true)
	case "enToCn":
		result = convertSymbols(text, false)
	default:
		result = text
	}

	return result, stats, nil
}

func convertSymbols(text string, toEn bool) string {
	cnSymbols := []string{"，", "。", "！", "？", "：", "；", "“", "”", "‘", "’", "（", "）", "【", "】", "—", "…"}
	enSymbols := []string{",", ".", "!", "?", ":", ";", "\"", "\"", "'", "'", "(", ")", "[", "]", "-", "..."}

	result := text
	if toEn {
		for i := 0; i < len(cnSymbols); i++ {
			result = strings.ReplaceAll(result, cnSymbols[i], enSymbols[i])
		}
	} else {
		for i := 0; i < len(enSymbols); i++ {
			// 对于双引号和单引号，中文区分左和右，这里做一个简单的映射
			result = strings.ReplaceAll(result, enSymbols[i], cnSymbols[i])
		}
	}
	return result
}
