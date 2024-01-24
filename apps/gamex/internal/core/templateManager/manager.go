package templateManager

import (
	"errors"
	"zinx-zero/apps/gamex/internal/ice"

	"github.com/aceld/zinx/zutils"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"
)

// Check interface implementation.
var _ ice.ITemplateManager = (*TemplateManager)(nil)

var templateManagerObj = newTemplateManager()

func newTemplateManager() *TemplateManager {
	return &TemplateManager{
		templateMap: zutils.NewShardLockMaps(),
	}
}

func GetTemplateManager() ice.ITemplateManager {
	return templateManagerObj
}

type TemplateManager struct {
	templateMap zutils.ShardLockMaps
}

// NewTemplate implements ice.ITemplateManager.
func (*TemplateManager) NewTemplate(id int64) (template ice.ITemplate) {
	template = &Template{}
	template.SetTemplateId(id)
	return template
}

func (a *TemplateManager) AddTemplate(template ice.ITemplate) {
	a.templateMap.Set(template.GetTemplateIdStr(), template)
	logx.Infof("AddTemplate success. %d", template.GetTemplateId())
}

func (a *TemplateManager) GetTemplateByTemplateId(templateId int64) (template ice.ITemplate, err error) {
	return a.GetTemplateByTemplateIdStr(cast.ToString(templateId))
}

func (a *TemplateManager) GetTemplateByTemplateIdStr(templateIdStr string) (template ice.ITemplate, err error) {
	if conn, ok := a.templateMap.Get(templateIdStr); ok {
		return conn.(ice.ITemplate), nil
	}
	return nil, errors.New("template not found")
}

func (a *TemplateManager) RemoveTemplate(template ice.ITemplate) {
	a.templateMap.Remove(template.GetTemplateIdStr())
	logx.Infof("RemoveTemplate fail. templateId=%d", template.GetTemplateId())
}
