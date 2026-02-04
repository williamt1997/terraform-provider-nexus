package acceptance

const (
	TemplateStringRepositoryNpmHosted = `
resource "nexus_repository_npm_hosted" "acceptance" {
` + TemplateStringHostedRepository

	TemplateStringRepositoryNpmGroup = `
resource "nexus_repository_npm_group" "acceptance" {
	depends_on = [
		nexus_repository_npm_hosted.acceptance
	]
` + TemplateStringGroupDeployRepository
)
