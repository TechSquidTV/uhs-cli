package service

type Ports struct {
	Http   *int `yaml:"http,omitempty"`
	Https  *int `yaml:"https,omitempty"`
	Udp    *int `yaml:"udp,omitempty"`
	P2p    *int `yaml:"p2p,omitempty"`
	P2pudp *int `yaml:"p2pudp,omitempty"`
}

type Image struct {
	Repository string `yaml:"repository"`
	Tag        string `yaml:"tag"`
	PullPolicy string `yaml:"pullPolicy"`
}

type Service struct {
	Enabled      bool  `yaml:"enabled"`
	ReplicaCount int   `yaml:"replicaCount"`
	Image        Image `yaml:"image"`
	Ports        Ports `yaml:"ports,omitempty"`
}
