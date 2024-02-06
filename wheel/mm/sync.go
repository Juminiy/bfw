package mm

import "sync"

type RPool struct {
	sync.Pool
}
