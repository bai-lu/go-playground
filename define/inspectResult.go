package define

type InspectResult interface {
	Save()
	Load(PID int)
	Count(Type string) int // 返回巡检Metric的数量
}
