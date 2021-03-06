// +build !ignore_autogenerated

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/openshift/certman-operator/pkg/apis/certman/v1alpha1.CertificateRequest":       schema_pkg_apis_certman_v1alpha1_CertificateRequest(ref),
		"github.com/openshift/certman-operator/pkg/apis/certman/v1alpha1.CertificateRequestSpec":   schema_pkg_apis_certman_v1alpha1_CertificateRequestSpec(ref),
		"github.com/openshift/certman-operator/pkg/apis/certman/v1alpha1.CertificateRequestStatus": schema_pkg_apis_certman_v1alpha1_CertificateRequestStatus(ref),
	}
}

func schema_pkg_apis_certman_v1alpha1_CertificateRequest(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CertificateRequest is the Schema for the certificaterequests API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openshift/certman-operator/pkg/apis/certman/v1alpha1.CertificateRequestSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openshift/certman-operator/pkg/apis/certman/v1alpha1.CertificateRequestStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/openshift/certman-operator/pkg/apis/certman/v1alpha1.CertificateRequestSpec", "github.com/openshift/certman-operator/pkg/apis/certman/v1alpha1.CertificateRequestStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_certman_v1alpha1_CertificateRequestSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CertificateRequestSpec defines the desired state of CertificateRequest",
				Properties: map[string]spec.Schema{
					"acmeDNSDomain": {
						SchemaProps: spec.SchemaProps{
							Description: "ACMEDNSDomain is the DNS zone that will house the TXT records needed for the certificate to be created. In Route53 this would be the public Route53 hosted zone (the Domain Name not the ZoneID)",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"certificateSecret": {
						SchemaProps: spec.SchemaProps{
							Description: "CertificateSecret is the reference to the secret where certificates are stored.",
							Ref:         ref("k8s.io/api/core/v1.ObjectReference"),
						},
					},
					"platformSecrets": {
						SchemaProps: spec.SchemaProps{
							Description: "PlatformSecrets contains the credentials and secrets for the cluster infrastructure.",
							Ref:         ref("github.com/openshift/certman-operator/pkg/apis/certman/v1alpha1.PlatformSecrets"),
						},
					},
					"dnsNames": {
						SchemaProps: spec.SchemaProps{
							Description: "DNSNames is a list of subject alt names to be used on the Certificate.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"email": {
						SchemaProps: spec.SchemaProps{
							Description: "Let's Encrypt will use this to contact you about expiring certificates, and issues related to your account.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"renewBeforeDays": {
						SchemaProps: spec.SchemaProps{
							Description: "Certificate renew before expiration duration in days.",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
				},
				Required: []string{"acmeDNSDomain", "certificateSecret", "platformSecrets", "dnsNames", "email"},
			},
		},
		Dependencies: []string{
			"github.com/openshift/certman-operator/pkg/apis/certman/v1alpha1.PlatformSecrets", "k8s.io/api/core/v1.ObjectReference"},
	}
}

func schema_pkg_apis_certman_v1alpha1_CertificateRequestStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CertificateRequestStatus defines the observed state of CertificateRequest",
				Properties: map[string]spec.Schema{
					"issued": {
						SchemaProps: spec.SchemaProps{
							Description: "Issued is true once certificates have been issued.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Description: "Status",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"notAfter": {
						SchemaProps: spec.SchemaProps{
							Description: "The expiration time of the certificate stored in the secret named by this resource in spec.secretName.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"notBefore": {
						SchemaProps: spec.SchemaProps{
							Description: "The earliest time and date on which the certificate stored in the secret named by this resource in spec.secretName is valid.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"issuerName": {
						SchemaProps: spec.SchemaProps{
							Description: "The entity that verified the information and signed the certificate.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"serialNumber": {
						SchemaProps: spec.SchemaProps{
							Description: "The serial number of the certificate stored in the secret named by this resource in spec.secretName.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"conditions": {
						SchemaProps: spec.SchemaProps{
							Description: "Conditions includes more detailed status for the Certificate Request",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/openshift/certman-operator/pkg/apis/certman/v1alpha1.CertificateRequestCondition"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/openshift/certman-operator/pkg/apis/certman/v1alpha1.CertificateRequestCondition"},
	}
}
