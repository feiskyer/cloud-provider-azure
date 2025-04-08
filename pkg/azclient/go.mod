module sigs.k8s.io/cloud-provider-azure/pkg/azclient

go 1.23.1

require (
	github.com/Azure/azure-sdk-for-go/sdk/azcore v1.18.0
	github.com/Azure/azure-sdk-for-go/sdk/azidentity v1.8.2
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v2 v2.2.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v6 v6.4.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry v1.2.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v6 v6.5.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault v1.5.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/msi/armmsi v1.2.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v6 v6.2.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/privatedns/armprivatedns v1.3.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources v1.2.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage v1.7.0
	github.com/Azure/azure-sdk-for-go/sdk/security/keyvault/azsecrets v1.3.1
	github.com/google/uuid v1.6.0
	github.com/onsi/ginkgo/v2 v2.23.4
	github.com/onsi/gomega v1.36.3
	go.opentelemetry.io/otel v1.35.0
	go.opentelemetry.io/otel/metric v1.35.0
	go.uber.org/mock v0.5.1
	golang.org/x/crypto v0.36.0
	golang.org/x/sync v0.12.0
	golang.org/x/time v0.11.0
	gopkg.in/dnaeon/go-vcr.v3 v3.2.0
	k8s.io/utils v0.0.0-20230505201702-9f6742963106
)

require (
	github.com/Azure/azure-sdk-for-go/sdk/internal v1.11.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/security/keyvault/internal v1.1.1 // indirect
	github.com/AzureAD/microsoft-authentication-library-for-go v1.4.2 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-task/slim-sprig/v3 v3.0.0 // indirect
	github.com/golang-jwt/jwt/v5 v5.2.2 // indirect
	github.com/google/go-cmp v0.7.0 // indirect
	github.com/google/pprof v0.0.0-20250403155104-27863c87afa6 // indirect
	github.com/kylelemons/godebug v1.1.0 // indirect
	github.com/pkg/browser v0.0.0-20240102092130-5ac0b6a4141c // indirect
	github.com/rogpeppe/go-internal v1.13.1 // indirect
	go.uber.org/automaxprocs v1.6.0 // indirect
	golang.org/x/net v0.38.0 // indirect
	golang.org/x/sys v0.32.0 // indirect
	golang.org/x/text v0.23.0 // indirect
	golang.org/x/tools v0.31.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
