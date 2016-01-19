package config

type Config struct {
	BindAddress string     `yaml:"BindAddress"`
	Database    string     `yaml:"Database" `
	Devices     []Device   `yaml:"Devices"`
	Scheduler   []Schedule `yaml:"Scheduler"`
	Triggers    []Trigger  `yaml:"Triggers"`
}

type Device struct {
	Name        string    `yaml:"Name" json:"name"`
	Description string    `yaml:"Description" json:"description"`
	Tags        []string  `yaml:"Tags" json:"tags"`
	Enabled     bool      `yaml:"Enabled" json:"enabled"`
	DeviceType  string    `yaml:"DeviceType" json:"deviceType"`
	Connection  string    `yaml:"Connection" json:"connection"`
	Channels    []Channel `yaml:"Channels" json:"channels"`
}

type Channel struct {
	Name        string   `yaml:"Name" json:"name"`
	Description string   `yaml:"Description" json:"description"`
	Tags        []string `yaml:"Tags" json:"tags"`
	ChannelType string   `yaml:"ChannelType" json:"channelType"`
	Address     string   `yaml:"Address" json:"address"`
	Enabled     bool     `yaml:"Enabled" json:"enabled"`
	Watt        float64  `yaml:"Watt" json:"watt"`
}

type Schedule struct {
	Name    string   `yaml:"Name" json:"name"`
	When    string   `yaml:"When" json:"when"`
	Command []string `yaml:"Command" json:"command"`
}

type Trigger struct {
	Name      string   `yaml:"Name" json:"name"`
	EventName string   `yaml:"EventName" json:"eventName"`
	Command   []string `yaml:"Command" json:"command"`
}
