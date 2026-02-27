package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strings"

	"github.com/pratice/golang/bytedance/apm/f_apm_performance_montior/model"
)

// 统计结果结构
type PathStats struct {
	Count     int       // 请求次数
	TotalTime float64   // 总响应时间（秒）
	Times     []float64 // 所有响应时间（秒）
	AvgTime   float64   // 平均响应时间（秒）
	MaxTime   float64   // 最大响应时间（秒）
	MinTime   float64   // 最小响应时间（秒）
	P90Time   float64   // P90 响应时间（秒）
}

func main() {
	// harFile := "./bytedance/apm/f_apm_performance_montior/har_apm_insight.json" // 替换为实际 HAR 文件路径
	harFile := "./bytedance/apm/f_apm_performance_montior/har_apm_imis.json" // 替换为实际 HAR 文件路径

	// 读取 HAR 文件
	data, err := os.ReadFile(harFile)
	if err != nil {
		log.Fatalf("Error reading HAR file: %v", err)
	}

	// 解析 HAR 文件
	var har model.HAR
	if err := json.Unmarshal(data, &har); err != nil {
		log.Fatalf("Error parsing HAR file: %v", err)
	}

	// 按 Path 统计请求
	statsByPath := make(map[string]*PathStats)

	for _, entry := range har.Log.Entries {
		// 忽略非 2xx/3xx 的请求（可选）
		if entry.Response.Status < 200 || entry.Response.Status >= 400 {
			continue
		}

		// 提取 Path（去掉查询参数和域名）
		// path := extractPath(entry.Request.URL)
		path := extractPathImis(entry.Request)

		// 初始化统计
		if _, ok := statsByPath[path]; !ok {
			statsByPath[path] = &PathStats{
				MinTime: math.MaxFloat64,
			}
		}

		// 更新统计
		timeSec := entry.Time / 1000 // 毫秒 -> 秒
		stats := statsByPath[path]
		stats.Count++
		stats.Times = append(stats.Times, timeSec)
		stats.TotalTime += timeSec
		if timeSec > stats.MaxTime {
			stats.MaxTime = timeSec
		}
		if timeSec < stats.MinTime {
			stats.MinTime = timeSec
		}
	}

	// 计算平均时间和 P90
	for _, stats := range statsByPath {
		stats.AvgTime = stats.TotalTime / float64(stats.Count)
		stats.P90Time = calculateP90(stats.Times)
	}

	// 打印结果
	printStats(statsByPath)
}

func extractPathImis(request model.Request) string {
	t := request.PostData.Text
	test := &model.Text{}
	if err := json.Unmarshal([]byte(t), test); err != nil {
		log.Fatalf("Error parsing HAR file: %v", err)
	}
	return extractPath(test.Config.URL)
}

// 从 URL 中提取 Path（去掉域名和查询参数）
func extractPath(url string) string {
	// 去掉协议和域名
	parts := strings.SplitN(url, "/", 4)
	if len(parts) < 4 {
		return url
	}
	path := "/" + parts[3]

	// 去掉查询参数
	if idx := strings.Index(path, "?"); idx != -1 {
		path = path[:idx]
	}

	return path
}

// 计算 P90 响应时间
func calculateP90(times []float64) float64 {
	if len(times) == 0 {
		return 0
	}

	// 排序
	sort.Float64s(times)

	// 计算 P90 的索引
	index := int(math.Ceil(float64(len(times)) * 0.9))
	if index >= len(times) {
		index = len(times) - 1
	}

	return times[index]
}

// 打印统计结果
func printStats(statsByPath map[string]*PathStats) {
	// 按请求次数排序
	type pathStat struct {
		Path  string
		Stats *PathStats
	}
	var sortedStats []pathStat
	for path, stats := range statsByPath {
		sortedStats = append(sortedStats, pathStat{path, stats})
	}
	sort.Slice(sortedStats, func(i, j int) bool {
		return sortedStats[i].Stats.Count > sortedStats[j].Stats.Count
	})

	// 打印表头
	fmt.Printf("%-70s %8s %10s %10s %10s %10s\n",
		"Path", "Count", "Avg(sec)", "Min(sec)", "Max(sec)", "P90(sec)")
	fmt.Println(strings.Repeat("-", 100))

	// 打印每个 Path 的统计
	for _, item := range sortedStats {
		stats := item.Stats
		fmt.Printf("%-70s %8d %10.2f %10.2f %10.2f %10.2f\n",
			truncatePath(item.Path, 70),
			stats.Count,
			stats.AvgTime,
			stats.MinTime,
			stats.MaxTime,
			stats.P90Time)
	}
}

// 截断过长的 Path
func truncatePath(path string, maxLen int) string {
	if len(path) <= maxLen {
		return path
	}
	return "..." + path[len(path)-maxLen+3:]
}
