package iam

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
)

func DataSourceAwsIamOpenIDConnectProvider() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAwsIamOpenIDConnectProviderRead,

		Schema: map[string]*schema.Schema{
			"arn": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: verify.ValidARN,
			},
			"url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"client_id_list": {
				Elem:     &schema.Schema{Type: schema.TypeString},
				Type:     schema.TypeList,
				Computed: true,
			},
			"thumbprint_list": {
				Elem:     &schema.Schema{Type: schema.TypeString},
				Type:     schema.TypeList,
				Computed: true,
			},
		},
	}
}

func dataSourceAwsIamOpenIDConnectProviderRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*AWSClient).IAMConn

	arn := d.Get("arn").(string)
	input := &iam.GetOpenIDConnectProviderInput{
		OpenIDConnectProviderArn: aws.String(arn),
	}
	out, err := iamconn.GetOpenIDConnectProvider(input)
	if err != nil {
		return fmt.Errorf("error reading IAM OIDC Provider (%s): %w", arn, err)
	}

	d.SetId(arn)
	d.Set("url", out.Url)
	d.Set("client_id_list", flattenStringList(out.ClientIDList))
	d.Set("thumbprint_list", flattenStringList(out.ThumbprintList))

	return nil
}
