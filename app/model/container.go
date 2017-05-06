package model

//Container Creatrion struct
type Container struct {
	CreatedBy   string `gorm:"not null;"`
	Name        string `gorm:"primary_key"`
	DisplayName string `gorm:"not null;unique"`
	Status      string
}

//ContainerStat struct
type ContainerStat struct {
	Name           string
	IsRunning      bool
	MemUsed        string
	MemLimit       string
	KernelMemUsed  string
	KernelMemLimit string
	MemSwapUsed    string
	MemSwapLimit   string
	BulkIOUsage    string
	CPUTime        string
	CPUTimePerCPU  interface{}
	CPUStats       interface{}
	InterfaceStats interface{}
	IPAddress      []string
}

//ContainerBasicInfo for basic info abt container
type ContainerBasicInfo struct {
	Name        string
	DisplayName string
	Status      string
	CreatedBy   string
	IsRunning   bool
	MemUsed     string
	MemLimit    string
	IPAddress   interface{}
}
