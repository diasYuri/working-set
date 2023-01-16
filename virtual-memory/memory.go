package virtual_memory

import "time"

const pageSize = 1

type Page struct {
	data       *[]byte
	id         int
	referenced bool
	lastAccess time.Time
}

func NewPage(pageId int) *Page {
	data := make([]byte, pageSize)
	return &Page{
		id:         pageId,
		data:       &data,
		referenced: true,
		lastAccess: time.Now(),
	}
}

func (pg *Page) Access() *[]byte {
	pg.lastAccess = time.Now()
	return pg.data
}

type Memory struct {
	pages []*Page
}

func NewMemory(size int) *Memory {
	m := &Memory{
		pages: make([]*Page, size),
	}
	for i := 0; i < size; i++ {
		m.pages[i] = nil
	}
	return m
}
