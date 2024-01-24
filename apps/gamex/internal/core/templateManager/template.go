package templateManager

import (
	"sync"
	"zinx-zero/apps/gamex/internal/ice"

	"github.com/spf13/cast"
)

// Check interface implementation.
var _ ice.ITemplate = (*Template)(nil)

type Template struct {
	sync.RWMutex

	templateId    int64
	templateIdStr string
}

// GetTemplateId implements ice.ITemplate.
func (a *Template) GetTemplateId() (templateId int64) {
	a.doRead(func() {
		templateId = a.templateId
	})
	return templateId
}

// GetTemplateIdStr implements ice.ITemplate.
func (a *Template) GetTemplateIdStr() (templateIdStr string) {
	a.doRead(func() {
		templateIdStr = a.templateIdStr
	})
	return templateIdStr
}

// SetTemplateId implements ice.ITemplate.
func (a *Template) SetTemplateId(templateId int64) {
	a.doWrite(func() {
		a.templateId = templateId
		a.templateIdStr = cast.ToString(templateId)
	})
}

func (a *Template) doWrite(fn func()) {
	a.Lock()
	defer a.Unlock()
	fn()
}

func (a *Template) doRead(fn func()) {
	a.RLock()
	defer a.RUnlock()
	fn()
}
