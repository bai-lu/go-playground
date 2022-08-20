package define

type InspectResult interface {
	Save()
	Load(PID int)
	Count(ttype string) int // 返回巡检Metric的数量
}
