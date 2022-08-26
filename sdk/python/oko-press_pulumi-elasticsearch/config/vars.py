# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

import types

__config__ = pulumi.Config('elasticsearch')


class _ExportableConfig(types.ModuleType):
    @property
    def aws_access_key(self) -> Optional[str]:
        """
        The access key for use with AWS Elasticsearch Service domains
        """
        return __config__.get('awsAccessKey')

    @property
    def aws_assume_role_arn(self) -> Optional[str]:
        """
        Amazon Resource Name of an IAM Role to assume prior to making AWS API calls.
        """
        return __config__.get('awsAssumeRoleArn')

    @property
    def aws_profile(self) -> Optional[str]:
        """
        The AWS profile for use with AWS Elasticsearch Service domains
        """
        return __config__.get('awsProfile')

    @property
    def aws_region(self) -> Optional[str]:
        """
        The AWS region for use in signing of AWS elasticsearch requests. Must be specified in order to use AWS URL signing with
        AWS ElasticSearch endpoint exposed on a custom DNS domain.
        """
        return __config__.get('awsRegion')

    @property
    def aws_secret_key(self) -> Optional[str]:
        """
        The secret key for use with AWS Elasticsearch Service domains
        """
        return __config__.get('awsSecretKey')

    @property
    def aws_signature_service(self) -> Optional[str]:
        """
        AWS service name used in the credential scope of signed requests to ElasticSearch.
        """
        return __config__.get('awsSignatureService')

    @property
    def aws_token(self) -> Optional[str]:
        """
        The session token for use with AWS Elasticsearch Service domains
        """
        return __config__.get('awsToken')

    @property
    def cacert_file(self) -> Optional[str]:
        """
        A Custom CA certificate
        """
        return __config__.get('cacertFile')

    @property
    def client_cert_path(self) -> Optional[str]:
        """
        A X509 certificate to connect to elasticsearch
        """
        return __config__.get('clientCertPath')

    @property
    def client_key_path(self) -> Optional[str]:
        """
        A X509 key to connect to elasticsearch
        """
        return __config__.get('clientKeyPath')

    @property
    def elasticsearch_version(self) -> Optional[str]:
        """
        ElasticSearch Version
        """
        return __config__.get('elasticsearchVersion')

    @property
    def healthcheck(self) -> Optional[bool]:
        """
        Set the client healthcheck option for the elastic client. Healthchecking is designed for direct access to the cluster.
        """
        return __config__.get_bool('healthcheck')

    @property
    def host_override(self) -> Optional[str]:
        """
        If provided, sets the 'Host' header of requests and the 'ServerName' for certificate validation to this value. See the
        documentation on connecting to Elasticsearch via an SSH tunnel.
        """
        return __config__.get('hostOverride')

    @property
    def insecure(self) -> Optional[bool]:
        """
        Disable SSL verification of API calls
        """
        return __config__.get_bool('insecure')

    @property
    def kibana_url(self) -> Optional[str]:
        """
        URL to reach the Kibana API
        """
        return __config__.get('kibanaUrl')

    @property
    def password(self) -> Optional[str]:
        """
        Password to use to connect to elasticsearch using basic auth
        """
        return __config__.get('password')

    @property
    def sign_aws_requests(self) -> Optional[bool]:
        """
        Enable signing of AWS elasticsearch requests. The `url` must refer to AWS ES domain (`*.<region>.es.amazonaws.com`), or
        `aws_region` must be specified explicitly.
        """
        return __config__.get_bool('signAwsRequests')

    @property
    def sniff(self) -> Optional[bool]:
        """
        Set the node sniffing option for the elastic client. Client won't work with sniffing if nodes are not routable.
        """
        return __config__.get_bool('sniff')

    @property
    def token(self) -> Optional[str]:
        """
        A bearer token or ApiKey for an Authorization header, e.g. Active Directory API key.
        """
        return __config__.get('token')

    @property
    def token_name(self) -> Optional[str]:
        """
        The type of token, usually ApiKey or Bearer
        """
        return __config__.get('tokenName')

    @property
    def url(self) -> Optional[str]:
        """
        Elasticsearch URL
        """
        return __config__.get('url')

    @property
    def username(self) -> Optional[str]:
        """
        Username to use to connect to elasticsearch using basic auth
        """
        return __config__.get('username')

    @property
    def version_ping_timeout(self) -> Optional[int]:
        """
        Version ping timeout in seconds
        """
        return __config__.get_int('versionPingTimeout')
