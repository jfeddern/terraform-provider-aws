// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package appsync

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []func(context.Context) (datasource.DataSourceWithConfigure, error) {
	return []func(context.Context) (datasource.DataSourceWithConfigure, error){}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []func(context.Context) (resource.ResourceWithConfigure, error) {
	return []func(context.Context) (resource.ResourceWithConfigure, error){}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) map[string]func() *schema.Resource {
	return map[string]func() *schema.Resource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) map[string]func() *schema.Resource {
	return map[string]func() *schema.Resource{
		"aws_appsync_api_cache":                   ResourceAPICache,
		"aws_appsync_api_key":                     ResourceAPIKey,
		"aws_appsync_datasource":                  ResourceDataSource,
		"aws_appsync_domain_name":                 ResourceDomainName,
		"aws_appsync_domain_name_api_association": ResourceDomainNameAPIAssociation,
		"aws_appsync_function":                    ResourceFunction,
		"aws_appsync_graphql_api":                 ResourceGraphQLAPI,
		"aws_appsync_resolver":                    ResourceResolver,
		"aws_appsync_type":                        ResourceType,
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.AppSync
}

var ServicePackage = &servicePackage{}
