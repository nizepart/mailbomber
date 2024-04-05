package store

type Store interface {
	User() UserRepository
	EmailTemplate() EmailTemplateRepository
}
