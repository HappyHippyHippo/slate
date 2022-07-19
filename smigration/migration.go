package smigration

// IMigration defines an interface that all migrations must obey.
type IMigration interface {
	Version() uint64
	Up() error
	Down() error
}
