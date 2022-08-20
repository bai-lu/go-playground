package define

type Inspector interface {
	Run() error
	Stop() error // cancel the inspecter
	Progress() (uint, error)
}
