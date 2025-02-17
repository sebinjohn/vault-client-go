module github.com/sebinjohn/vault-client-go

go 1.21

require (
	github.com/hashicorp/go-cleanhttp v0.5.2
	github.com/hashicorp/go-retryablehttp v0.7.1
	github.com/hashicorp/go-rootcerts v1.0.2
	github.com/hashicorp/go-secure-stdlib/strutil v0.1.2
	github.com/hashicorp/vault-client-go v0.4.1
	github.com/stretchr/testify v1.8.0
	golang.org/x/sys v0.4.0
	golang.org/x/time v0.0.0-20220922220347-f3bd1da661af
)

replace (
	github.com/hashicorp/vault-client-go => github.com/sebinjohn/vault-client-go v1.1.3
	github.com/hashicorp/vault-client-go/schema => github.com/sebinjohn/vault-client-go/schema v1.1.3
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/ryanuber/go-glob v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
