package randomx

import (
	"fmt"
	"math/rand"
	"time"
)

// GenerateOrderNo 生成唯一订单号
// 格式: PAY + 时间戳(14位) + 随机数(6位)
// 示例: PAY20260204153045123456
func GenerateOrderNo() string {
	// 时间戳部分: YYYYMMDDHHmmss
	timestamp := time.Now().Format("20060102150405")

	// 随机数部分: 6位随机数
	random := rand.Intn(1000000)

	return fmt.Sprintf("PAY%s%06d", timestamp, random)
}

// GenerateTransactionNo 生成唯一流水号
// 格式: TXN + 时间戳(14位) + 随机数(6位)
// 示例: TXN20260204153045123456
func GenerateTransactionNo() string {
	// 时间戳部分: YYYYMMDDHHmmss
	timestamp := time.Now().Format("20060102150405")

	// 随机数部分: 6位随机数
	random := rand.Intn(1000000)

	return fmt.Sprintf("TXN%s%06d", timestamp, random)
}

func init() {
	// 初始化随机数种子
	rand.Seed(time.Now().UnixNano())
}
