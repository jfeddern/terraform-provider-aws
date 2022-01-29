package iam_test

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
)

func TestAccAWSDataSourceIamOpenIDConnectProvider_basic(t *testing.T) {
	rString := acctest.RandString(5)
	url := "accounts.testle.com/" + rString
	resourceName := "data.aws_iam_openid_connect_provider.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:   func() { acctest.PreCheck(t) },
		ErrorCheck: acctest.ErrorCheck(t, iam.EndpointsID),
		Providers:  acctest.Providers,
		Steps: []resource.TestStep{
			{
				Config: testAccDatasourceAwsIamOpenIDConnectProviderConfig(url),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(resourceName, "arn", "aws_iam_openid_connect_provider.test", "arn"),
					resource.TestCheckResourceAttr(resourceName, "url", url),
					resource.TestCheckResourceAttr(resourceName, "client_id_list.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "client_id_list.0",
						"266362248691-re108qaeld573ia0l6clj2i5ac7r7291.apps.testleusercontent.com"),
					resource.TestCheckResourceAttr(resourceName, "thumbprint_list.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "thumbprint_list.0", "cf23df2207d99a74fbe169e3eba035e633b65d94"),
					resource.TestCheckResourceAttr(resourceName, "thumbprint_list.1", "c784713d6f9cb67b55dd84f4e4af7832d42b8f55"),
				),
			},
		},
	})
}

func testAccDatasourceAwsIamOpenIDConnectProviderConfig(url string) string {
	return fmt.Sprintf(`
resource "aws_iam_openid_connect_provider" "test" {
  url = "https://accounts.testle.com/%s"
  
  client_id_list = [
    "266362248691-re108qaeld573ia0l6clj2i5ac7r7291.apps.testleusercontent.com",
  ]
  
  thumbprint_list = ["cf23df2207d99a74fbe169e3eba035e633b65d94", "c784713d6f9cb67b55dd84f4e4af7832d42b8f55"]
}
data "aws_iam_openid_connect_provider" "test" {
  arn = aws_iam_openid_connect_provider.test.arn
}
`, url)
}
