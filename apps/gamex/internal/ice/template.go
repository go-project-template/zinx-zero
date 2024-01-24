package ice

import "zinx-zero/apps/gamex/msg"

type ITemplate interface {
	Init(dbTemplate *msg.DBTemplate)
	SetTemplateId(templateId int64)
	GetTemplateId() (templateId int64)
	GetTemplateIdStr() (templateIdStr string)
}

type ITemplateManager interface {
	AddTemplate(template ITemplate)
	GetTemplateByTemplateId(templateId int64) (template ITemplate, err error)
	GetTemplateByTemplateIdStr(templateIdStr string) (template ITemplate, err error)
	RemoveTemplate(template ITemplate)
}
