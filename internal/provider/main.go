package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	nexus "github.com/williamt1997/go-nexus-client/nexus3"
	"github.com/williamt1997/go-nexus-client/nexus3/pkg/client"
	"github.com/williamt1997/terraform-provider-nexus/internal/services/repository"
	"github.com/williamt1997/terraform-provider-nexus/internal/services/security"
)

// Provider returns a terraform.Provider
func Provider() *schema.Provider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{},
		ResourcesMap: map[string]*schema.Resource{
			"nexus_repository_apt_hosted":                 repository.ResourceRepositoryAptHosted(),
			"nexus_repository_bower_group":                repository.ResourceRepositoryBowerGroup(),
			"nexus_repository_bower_hosted":               repository.ResourceRepositoryBowerHosted(),
			"nexus_repository_cargo_group":                repository.ResourceRepositoryCargoGroup(),
			"nexus_repository_cargo_hosted":               repository.ResourceRepositoryCargoHosted(),
			"nexus_repository_docker_group":               repository.ResourceRepositoryDockerGroup(),
			"nexus_repository_docker_hosted":              repository.ResourceRepositoryDockerHosted(),
			"nexus_repository_gitlfs_hosted":              repository.ResourceRepositoryGitlfsHosted(),
			"nexus_repository_go_group":                   repository.ResourceRepositoryGoGroup(),
			"nexus_repository_helm_hosted":                repository.ResourceRepositoryHelmHosted(),
			"nexus_repository_maven_group":                repository.ResourceRepositoryMavenGroup(),
			"nexus_repository_maven_hosted":               repository.ResourceRepositoryMavenHosted(),
			"nexus_repository_maven_proxy":                repository.ResourceRepositoryMavenProxy(),
			"nexus_repository_npm_group":                  repository.ResourceRepositoryNpmGroup(),
			"nexus_repository_npm_hosted":                 repository.ResourceRepositoryNpmHosted(),
			"nexus_repository_nuget_group":                repository.ResourceRepositoryNugetGroup(),
			"nexus_repository_nuget_hosted":               repository.ResourceRepositoryNugetHosted(),
			"nexus_repository_pypi_group":                 repository.ResourceRepositoryPypiGroup(),
			"nexus_repository_pypi_hosted":                repository.ResourceRepositoryPypiHosted(),
			"nexus_repository_r_group":                    repository.ResourceRepositoryRGroup(),
			"nexus_repository_r_hosted":                   repository.ResourceRepositoryRHosted(),
			"nexus_repository_raw_group":                  repository.ResourceRepositoryRawGroup(),
			"nexus_repository_raw_hosted":                 repository.ResourceRepositoryRawHosted(),
			"nexus_repository_rubygems_group":             repository.ResourceRepositoryRubygemsGroup(),
			"nexus_repository_rubygems_hosted":            repository.ResourceRepositoryRubygemsHosted(),
			"nexus_repository_yum_group":                  repository.ResourceRepositoryYumGroup(),
			"nexus_repository_yum_hosted":                 repository.ResourceRepositoryYumHosted(),
			"nexus_security_content_selector":             security.ResourceSecurityContentSelector(),
			"nexus_security_role":                         security.ResourceSecurityRole(),
			"nexus_security_user":                         security.ResourceSecurityUser(),
			"nexus_privilege_repository_content_selector": security.ResourceSecurityPrivilegeRepositoryContentSelector(),
		},
		Schema: map[string]*schema.Schema{
			"insecure": {
				Description: "Boolean to specify wether insecure SSL connections are allowed or not. Reading environment variable NEXUS_INSECURE_SKIP_VERIFY. Default:`true`",
				Default:     false,
				DefaultFunc: schema.EnvDefaultFunc("NEXUS_INSECURE_SKIP_VERIFY", "true"),
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"password": {
				Description: "Password of user to connect to API. Reading environment variable NEXUS_PASSWORD. Default:`admin123`",
				DefaultFunc: schema.EnvDefaultFunc("NEXUS_PASSWORD", "admin123"),
				Optional:    true,
				Type:        schema.TypeString,
			},
			"url": {
				Description: "URL of Nexus to reach API. Reading environment variable NEXUS_URL. Default:`http://127.0.0.1:8080`",
				DefaultFunc: schema.EnvDefaultFunc("NEXUS_URL", "http://127.0.0.1:8080"),
				Optional:    true,
				Type:        schema.TypeString,
			},
			"username": {
				Description: "Username used to connect to API. Reading environment variable NEXUS_USERNAME. Default:`admin`",
				DefaultFunc: schema.EnvDefaultFunc("NEXUS_USERNAME", "admin"),
				Optional:    true,
				Type:        schema.TypeString,
			},
			"timeout": {
				Description: "Timeout in seconds to connect to API. Reading environment variable NEXUS_TIMEOUT. Default:`30`",
				DefaultFunc: schema.EnvDefaultFunc("NEXUS_TIMEOUT", 30),
				Optional:    true,
				Type:        schema.TypeInt,
			},
			"client_cert_path": {
				Description: "Path to a client PEM certificate to load for mTLS. Reading environment variable NEXUS_CLIENT_CERT_PATH. Default:``",
				DefaultFunc: schema.EnvDefaultFunc("NEXUS_CLIENT_CERT_PATH", ""),
				Optional:    true,
				Type:        schema.TypeString,
			},
			"client_key_path": {
				Description: "Path to a client PEM key to load for mTLS. Reading environment variable NEXUS_CLIENT_KEY_PATH. Default:``",
				DefaultFunc: schema.EnvDefaultFunc("NEXUS_CLIENT_KEY_PATH", ""),
				Optional:    true,
				Type:        schema.TypeString,
			},
			"root_ca_path": {
				Description: "Path to a root CA certificate to load for mTLS. Reading environment variable NEXUS_ROOT_CA_PATH. Default:``",
				DefaultFunc: schema.EnvDefaultFunc("NEXUS_ROOT_CA_PATH", ""),
				Optional:    true,
				Type:        schema.TypeString,
			},
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	timeout := d.Get("timeout").(int)
	clientCertPath := d.Get("client_cert_path").(string)
	clientKeyPath := d.Get("client_key_path").(string)
	rootCaPath := d.Get("root_ca_path").(string)
	config := client.Config{
		Insecure:              d.Get("insecure").(bool),
		Password:              d.Get("password").(string),
		URL:                   d.Get("url").(string),
		Username:              d.Get("username").(string),
		Timeout:               &timeout,
		ClientCertificatePath: &clientCertPath,
		ClientKeyPath:         &clientKeyPath,
		RootCAPath:            &rootCaPath,
	}

	return nexus.NewClient(config), nil
}
