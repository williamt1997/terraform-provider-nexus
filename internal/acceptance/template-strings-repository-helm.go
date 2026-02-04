package acceptance

const (
	TemplateStringRepositoryHelmHosted = `
resource "nexus_repository_helm_hosted" "acceptance" {
` + TemplateStringHostedRepository
)
