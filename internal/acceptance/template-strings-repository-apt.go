package acceptance

const (
	TemplateStringRepositoryAptHosted = `
resource "nexus_repository_apt_hosted" "acceptance" {
	distribution = "{{ .Apt.Distribution }}"
	signing {
		keypair = "{{ .AptSigning.Keypair }}"
{{- if .AptSigning.Passphrase }}
		passphrase = "{{ .AptSigning.Passphrase }}"
{{- end }}
	}
` + TemplateStringHostedRepository
)
