package build

import (
	"github.com/peractio/gdk/pkg/syncx"
	"github.com/peractio/gdk/pkg/tags"
)

var (
	md          *Metadata
	onceInfo    syncx.Once
	onceInfoRes map[string]string
)

// Metadata is information set on binary build.
type Metadata struct {
	// Hash contains the information of the last git commit hash.
	Hash string
	// User contains the username that built the binary.
	User string
	// Time contains the information when the binary is being built.
	Time string
}

// New creates the metadata.
func New(hash, user, time string) *Metadata {
	if hash == "" && user == "" && time == "" {
		return nil
	}

	md = &Metadata{
		Hash: hash,
		User: user,
		Time: time,
	}
	return md
}

// Info returns current build info.
func Info() map[string]string {
	onceInfo.Do(func() {
		if md == nil {
			return
		}

		info := map[string]string{}

		if md.Hash != "" {
			info[tags.Hash] = md.Hash
		}
		if md.User != "" {
			info[tags.User] = md.User
		}
		if md.Time != "" {
			info[tags.Time] = md.Time
		}

		onceInfoRes = info
	})

	return onceInfoRes
}
