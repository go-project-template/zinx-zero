package templateManager

import (
	"sync"
	"zinx-zero/apps/gamex/internal/ice"
	"zinx-zero/apps/gamex/proto/msg"

	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"
)

// Check interface implementation.
var _ ice.ITemplate = (*Template)(nil)

func NewTemplate(id int64) ice.ITemplate {
	template := &Template{}
	template.SetTemplateIdLock(id)
	return template
}

type Template struct {
	*msg.DBTemplate
	sync.RWMutex

	templateIdStr string
}

// InitTemplateLock implements ice.ITemplate.
func (a *Template) InitTemplateByDbLock(dbTemplate *msg.DBTemplate) {
	if dbTemplate == nil {
		logx.Errorf("InitTemplateByDbLock dbTemplate is nil")
		return
	}
	a.DBTemplate = dbTemplate
}

// GetTemplateIdLock implements ice.ITemplate.
func (a *Template) GetTemplateIdLock() (templateId int64) {
	return a.GetTemplateIdLock()
}

// GetTemplateIdStrLock implements ice.ITemplate.
func (a *Template) GetTemplateIdStrLock() (templateIdStr string) {
	return a.templateIdStr
}

// SetTemplateIdLock implements ice.ITemplate.
func (a *Template) SetTemplateIdLock(templateId int64) {
	a.TemplateId = templateId
	a.templateIdStr = cast.ToString(templateId)
}

func (a *Template) DoWriteLock(fn func()) {
	a.Lock()
	defer a.Unlock()
	fn()
}

func (a *Template) DoReadLock(fn func()) {
	a.RLock()
	defer a.RUnlock()
	fn()
}
