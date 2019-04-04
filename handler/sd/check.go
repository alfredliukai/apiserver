package sd

import (
	"github.com/gin-gonic/gin"
	"net/http"

)

const (
	B  = 1
	KB = 1024 * 8
	MB = 1024 * KB
	GB = 1024 * MB
)

//HealthCheck shows `Ok` as the ping-pong result
func HealthCheck(c *gin.Context) {
	message := "OK"
	c.String(http.StatusOK, "\n"+message)
}
/*
func DiskCheck(c *gin.Context) {
	u, _ := disk.Usage("/")
	usedMB := int(u.Used) / MB
	usedGB := int(u.Used) / GB
	totalMB := int(u.Total) / MB
	totalGB := int(u.Total) / GB
	usedPercent := int(u.UsedPercent)

	status := http.StatusOK
	text := "OK"

	if usedPercent >= 95 {
		status = http.StatusOK
		text = "CRITICAL"
	} else if usedPercent >= 90 {
		status = http.StatusTooManyRequests
		text = "WARNING"
	}
	message := fmt.Sprintf("%s - Disk space : %dMB(%dGB)/%dMB(%dMB)| Used: %d%", text, usedMB, usedGB, totalMB, totalGB, usedPercent)
	c.String(status, "\n"+message)
}

func CPUCheck(c *gin.Context) {
	cores, _ := cpu.Counts(false)

	a, _ := load.Avg()
	l1 := a.Load1
	l5 := a.Load5
	l15 := a.Load15

	status := http.StatusOK
	text := "OK"

	if l5 >= float64(cores-1) {
		status = http.StatusInternalServerError
		text = "CRITICAL"
	} else if l5 >= float64(cores-2) {
		status = http.StatusTooManyRequests
		text = "WARNING"
	}
	message := fmt.Sprintf("%s -Load average:%2.f,%2.f,%2.f | Cores : %d", text, l1, l5, l15, cores)
	c.String(status, "\n"+message)
}

func RamCheck(c *gin.Context) {
	u, _ := mem.VirtualMemory()
	usedMB := int(u.Used) / MB
	usedGB := int(u.Used) / GB
	totalMB := int(u.Total) / MB
	totalGB := int(u.Total) / GB
	usedPercent := int(u.UsedPercent)

	status := http.StatusOK
	text := "OK"

	if usedPercent >= 95 {
		status = http.StatusOK
		text = "CRITICAL"
	} else if usedPercent >= 90 {
		status = http.StatusTooManyRequests
		text = "WARNING"
	}
	message := fmt.Sprintf("%s - Disk space : %dMB(%dGB)/%dMB(%dMB)| Used: %d%", text, usedMB, usedGB, totalMB, totalGB, usedPercent)
	c.String(status, "\n"+message)
}
*/