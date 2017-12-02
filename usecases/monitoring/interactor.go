package monitoring

//Interactor : the monitoring interactor
type Interactor struct {
	MonitoringRW MonitoringRW
	Version      string
}

// NewMonitoringInteractor : Admin Interactor constructor
func NewMonitoringInteractor(version string, MONRW MonitoringRW) Interactor {
	return Interactor{MONRW, version}
}
