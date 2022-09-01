package redis

import "time"

type DialConfig struct {
	ConnectTimeout  time.Duration
	ReadTimeout     time.Duration
	Host            string
	Port            int
	Database        int
	Password        string
	MaxIdle         int
	MaxActive       int
	IdleTimeout     int
	Wait            bool
	MaxConnLifetime int
}
