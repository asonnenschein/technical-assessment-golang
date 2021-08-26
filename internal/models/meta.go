package models

import "net/url"

type Meta struct {
	OriginUrl url.URL
}

func (m *Meta) SetOriginUrl(originUrl url.URL) {
	m.OriginUrl = originUrl
}

func (m *Meta) GetOriginUrl() url.URL {
	return m.OriginUrl
}
