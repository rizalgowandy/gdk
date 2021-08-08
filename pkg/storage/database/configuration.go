package database

type PostgreConfiguration struct {
	// Address is database connection address.
	// e.g. user=unicorn_user password=magical_password dbname=application_name host=127.0.0.1 port=5432 sslmode=disable
	Address string
	// MinConnection is the minimum size of the pool.
	// The health check will increase the number of connections to this
	// amount if it had dropped below.
	MinConnection int32
	// MaxConnection is the maximum size of the pool.
	MaxConnection int32
	// MaxConnectionLifetime is the duration since creation after which a connection will be automatically closed.
	// Time: second.
	MaxConnectionLifetime int64
	// MaxConnectionIdleTime is the duration after which an idle connection will be automatically closed by the health check.
	// Time: second.
	MaxConnectionIdleTime int64
}
