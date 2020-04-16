package radio

// Connector defines connector, suitable to perform operations
type Connector interface {
	Invoke(Context) error
}
