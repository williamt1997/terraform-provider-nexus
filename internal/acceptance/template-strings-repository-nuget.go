package acceptance

const (
	TemplateStringRepositoryNugetHosted = `
resource "nexus_repository_nuget_hosted" "acceptance" {
` + TemplateStringHostedRepository

	TemplateStringRepositoryNugetGroup = `
resource "nexus_repository_nuget_group" "acceptance" {
	depends_on = [
		nexus_repository_nuget_hosted.acceptance
	]
` + TemplateStringGroupRepository
)
