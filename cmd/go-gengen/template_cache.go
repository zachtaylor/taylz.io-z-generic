package main

// CacheTemplate is template for cache generation
const CacheTemplate = `// Code generated by go-gengen(v0.0.0) DO NOT EDIT.

package {{P}}

// Cache is an observable concurrent in-memory datastore
type Cache struct {
	dat   map[{{TK}}]{{TV}}
	mu    sync.Mutex
	obs   []Func
}

// Storer is an abstraction of in-memory datastore
type Storer interface {
	Get({{TK}}) {{TV}}
	Set({{TK}}, {{TV}})
	Each(Func)
	Sync(func(map[{{TK}}]{{TV}}))
	Keys() []{{TK}}
	Observe(Func)
	Remove({{TK}})
}

// Func is a callback func
type Func = func({{TK}}, {{TV}})

// NewCache returns a new Cache
func NewCache() *Cache {
	return &Cache{
		dat: make(map[{{TK}}]{{TV}}),
		obs:   make([]Func, 0),
	}
}

// Get returns the {{TV}} for a {{TK}}
func (c *Cache) Get(k {{TK}}) {{TV}} { return c.dat[k] }

// Set saves a {{TV}} for a {{TK}}
func (c *Cache) Set(k {{TK}}, v {{TV}}) {
	c.mu.Lock()
	if v != nil {
		c.dat[k] = v
	} else {
		delete(c.dat, k)
	}
	c.mu.Unlock()
	for _, f := range c.obs {
		f(k, v)
	}
}

// Each calls the func for each {{TK}},{{TV}} in this Cache
func (c *Cache) Each(f Func) {
	c.mu.Lock()
	for k, v := range c.dat {
		f(k, v)
	}
	c.mu.Unlock()
}

// Sync calls the func within the cache lock state
func (c *Cache) Sync(f func(map[{{TK}}]{{TV}})) {
	c.mu.Lock()
	f(c.dat)
	c.mu.Unlock()
}

// Keys returns a new slice with all the {{TK}} keys
func (c *Cache) Keys() []{{TK}} {
	c.mu.Lock()
	keys := make([]{{TK}}, 0, len(c.dat))
	for k := range c.dat {
		keys = append(keys, k)
	}
	c.mu.Unlock()
	return keys
}

// Observe adds a func to be called when a {{TV}} is explicitly set
func (c *Cache) Observe(f Func) { c.obs = append(c.obs, f) }

// Remove deletes a {{TK}},{{TV}}
func (c *Cache) Remove(k {{TK}}) { c.Set(k, nil) }
`
