package collectGroup

import "sync"

type Group struct {
	cancel func()
	wg     sync.WaitGroup
	once   sync.Once
	Errs   map[string]error
}

func (g *Group) Go(id string, fn func() error) {
	g.wg.Add(1)
	go func() {
		id := id
		defer g.wg.Done()
		if err := fn(); err != nil {
			g.Errs[id] = err
			g.once.Do(func() {
				if g.cancel != nil {
					g.cancel()
				}
			})
		}
	}()
}

func (g *Group) Wait() bool {
	g.wg.Wait()
	if g.cancel != nil {
		g.cancel()
	}
	return len(g.Errs) > 0
}

// refer: https://github.com/higker/collgroup