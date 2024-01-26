package ice

import "zinx-zero/apps/gamex/proto/msg"

type ITemplateManager interface {
	AddTemplate(template ITemplate)
	GetTemplateByTemplateId(templateId int64) (template ITemplate, err error)
	GetTemplateByTemplateIdStr(templateIdStr string) (template ITemplate, err error)
	RemoveTemplate(template ITemplate)
}

type ITemplate interface {
	DoWriteLock(fn func())
	DoReadLock(fn func())
	InitTemplateByDbLock(dbTemplate *msg.DBTemplate)
	SetTemplateIdLock(templateId int64)
	GetTemplateIdLock() (templateId int64)
	GetTemplateIdStrLock() (templateIdStr string)
}
