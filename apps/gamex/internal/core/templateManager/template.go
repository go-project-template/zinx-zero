package templateManager

import (
	"sync"
	"zinx-zero/apps/gamex/internal/ice"
	"zinx-zero/apps/gamex/msg"

	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"
)

// Check interface implementation.
var _ ice.ITemplate = (*Template)(nil)

func NewTemplate(id int64) (template ice.ITemplate) {
	template = &Template{}
	template.SetTemplateId(id)
	return template
}

type Template struct {
	*msg.DBTemplate
	sync.RWMutex

	templateIdStr string
}

// Init implements ice.ITemplate.
func (a *Template) Init(dbTemplate *msg.DBTemplate) {
	if dbTemplate == nil {
		logx.Errorf("Init Template dbTemplate is nil")
		return
	}
	a.DBTemplate = dbTemplate
}

// GetTemplateId implements ice.ITemplate.
func (a *Template) GetTemplateId() (templateId int64) {
	a.doRead(func() {
		templateId = a.GetTemplateId()
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
		a.TemplateId = templateId
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
