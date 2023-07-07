package manager

// Configurer represents the config manager interface.
type Configurer interface {
	// Configure the service config based on user specification
	Configure()
	// Fill populates the service config
	Fill([]byte) error
}
