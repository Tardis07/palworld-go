// status包
package status

var memoryIssueDetected bool = false
var successReadGameWorldSettings bool = false

// SetMemoryIssueDetected 设置内存问题检测标志
func SetMemoryIssueDetected(flag bool) {
	memoryIssueDetected = flag
}

// GetMemoryIssueDetected 获取内存问题检测标志的当前值
func GetMemoryIssueDetected() bool {
	return memoryIssueDetected
}

// SetMemoryIssueDetected 设置内存问题检测标志
func SetsuccessReadGameWorldSettings(flag bool) {
	successReadGameWorldSettings = flag
}

// GetMemoryIssueDetected 获取内存问题检测标志的当前值
func GetsuccessReadGameWorldSettings() bool {
	return successReadGameWorldSettings
}
