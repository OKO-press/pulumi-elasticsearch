// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

// The access key for use with AWS Elasticsearch Service domains
func GetAwsAccessKey(ctx *pulumi.Context) string {
	return config.Get(ctx, "elasticsearch:awsAccessKey")
}

// Amazon Resource Name of an IAM Role to assume prior to making AWS API calls.
func GetAwsAssumeRoleArn(ctx *pulumi.Context) string {
	return config.Get(ctx, "elasticsearch:awsAssumeRoleArn")
}

// The AWS profile for use with AWS Elasticsearch Service domains
func GetAwsProfile(ctx *pulumi.Context) string {
	return config.Get(ctx, "elasticsearch:awsProfile")
}

// The AWS region for use in signing of AWS elasticsearch requests. Must be specified in order to use AWS URL signing with
// AWS ElasticSearch endpoint exposed on a custom DNS domain.
func GetAwsRegion(ctx *pulumi.Context) string {
	return config.Get(ctx, "elasticsearch:awsRegion")
}

// The secret key for use with AWS Elasticsearch Service domains
func GetAwsSecretKey(ctx *pulumi.Context) string {
	return config.Get(ctx, "elasticsearch:awsSecretKey")
}

// AWS service name used in the credential scope of signed requests to ElasticSearch.
func GetAwsSignatureService(ctx *pulumi.Context) string {
	return config.Get(ctx, "elasticsearch:awsSignatureService")
}

// The session token for use with AWS Elasticsearch Service domains
func GetAwsToken(ctx *pulumi.Context) string {
	return config.Get(ctx, "elasticsearch:awsToken")
}

// A Custom CA certificate
func GetCacertFile(ctx *pulumi.Context) string {
	return config.Get(ctx, "elasticsearch:cacertFile")
}

// A X509 certificate to connect to elasticsearch
func GetClientCertPath(ctx *pulumi.Context) string {
	return config.Get(ctx, "elasticsearch:clientCertPath")
}

// A X509 key to connect to elasticsearch
func GetClientKeyPath(ctx *pulumi.Context) string {
	return config.Get(ctx, "elasticsearch:clientKeyPath")
}

// ElasticSearch Version
func GetElasticsearchVersion(ctx *pulumi.Context) string {
	return config.Get(ctx, "elasticsearch:elasticsearchVersion")
}

// Set the client healthcheck option for the elastic client. Healthchecking is designed for direct access to the cluster.
func GetHealthcheck(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "elasticsearch:healthcheck")
}

// If provided, sets the 'Host' header of requests and the 'ServerName' for certificate validation to this value. See the
// documentation on connecting to Elasticsearch via an SSH tunnel.
func GetHostOverride(ctx *pulumi.Context) string {
	return config.Get(ctx, "elasticsearch:hostOverride")
}

// Disable SSL verification of API calls
func GetInsecure(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "elasticsearch:insecure")
}

// URL to reach the Kibana API
func GetKibanaUrl(ctx *pulumi.Context) string {
	return config.Get(ctx, "elasticsearch:kibanaUrl")
}

// Password to use to connect to elasticsearch using basic auth
func GetPassword(ctx *pulumi.Context) string {
	return config.Get(ctx, "elasticsearch:password")
}

// Enable signing of AWS elasticsearch requests. The `url` must refer to AWS ES domain (`*.<region>.es.amazonaws.com`), or
// `aws_region` must be specified explicitly.
func GetSignAwsRequests(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "elasticsearch:signAwsRequests")
}

// Set the node sniffing option for the elastic client. Client won't work with sniffing if nodes are not routable.
func GetSniff(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "elasticsearch:sniff")
}

// A bearer token or ApiKey for an Authorization header, e.g. Active Directory API key.
func GetToken(ctx *pulumi.Context) string {
	return config.Get(ctx, "elasticsearch:token")
}

// The type of token, usually ApiKey or Bearer
func GetTokenName(ctx *pulumi.Context) string {
	return config.Get(ctx, "elasticsearch:tokenName")
}

// Elasticsearch URL
func GetUrl(ctx *pulumi.Context) string {
	return config.Get(ctx, "elasticsearch:url")
}

// Username to use to connect to elasticsearch using basic auth
func GetUsername(ctx *pulumi.Context) string {
	return config.Get(ctx, "elasticsearch:username")
}

// Version ping timeout in seconds
func GetVersionPingTimeout(ctx *pulumi.Context) int {
	return config.GetInt(ctx, "elasticsearch:versionPingTimeout")
}
