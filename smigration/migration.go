package smigration

// Migration defines an interface that all migrations must obey.
type Migration interface {
	Version() uint64
	Up() error
	Down() error
}
