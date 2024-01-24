package ice

type ITemplate interface {
	SetTemplateId(templateId int64)
	GetTemplateId() (templateId int64)
	GetTemplateIdStr() (templateIdStr string)
}

type ITemplateManager interface {
	NewTemplate(id int64) (template ITemplate)
	AddTemplate(template ITemplate)
	GetTemplateByTemplateId(templateId int64) (template ITemplate, err error)
	GetTemplateByTemplateIdStr(templateIdStr string) (template ITemplate, err error)
	RemoveTemplate(template ITemplate)
}
