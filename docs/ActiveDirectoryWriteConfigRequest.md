# ActiveDirectoryWriteConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnonymousGroupSearch** | Pointer to **bool** | Use anonymous binds when performing LDAP group searches (if true the initial credentials will still be used for the initial connection test). | [optional] [default to false]
**Binddn** | Pointer to **string** | LDAP DN for searching for the user DN (optional) | [optional] 
**Bindpass** | Pointer to **string** | LDAP password for searching for the user DN (optional) | [optional] 
**CaseSensitiveNames** | Pointer to **bool** | If true, case sensitivity will be used when comparing usernames and groups for matching policies. | [optional] 
**Certificate** | Pointer to **string** | CA certificate to use when verifying LDAP server certificate, must be x509 PEM encoded (optional) | [optional] 
**ClientTlsCert** | Pointer to **string** | Client certificate to provide to the LDAP server, must be x509 PEM encoded (optional) | [optional] 
**ClientTlsKey** | Pointer to **string** | Client certificate key to provide to the LDAP server, must be x509 PEM encoded (optional) | [optional] 
**DenyNullBind** | Pointer to **bool** | Denies an unauthenticated LDAP bind request if the user&#39;s password is empty; defaults to true | [optional] [default to true]
**Discoverdn** | Pointer to **bool** | Use anonymous bind to discover the bind DN of a user (optional) | [optional] 
**Formatter** | Pointer to **string** | Text to insert the password into, ex. \&quot;customPrefix{{PASSWORD}}customSuffix\&quot;. | [optional] 
**Groupattr** | Pointer to **string** | LDAP attribute to follow on objects returned by &lt;groupfilter&gt; in order to enumerate user group membership. Examples: \&quot;cn\&quot; or \&quot;memberOf\&quot;, etc. Default: cn | [optional] [default to "cn"]
**Groupdn** | Pointer to **string** | LDAP search base to use for group membership search (eg: ou&#x3D;Groups,dc&#x3D;example,dc&#x3D;org) | [optional] 
**Groupfilter** | Pointer to **string** | Go template for querying group membership of user (optional) The template can access the following context variables: UserDN, Username Example: (&amp;(objectClass&#x3D;group)(member:1.2.840.113556.1.4.1941:&#x3D;{{.UserDN}})) Default: (|(memberUid&#x3D;{{.Username}})(member&#x3D;{{.UserDN}})(uniqueMember&#x3D;{{.UserDN}})) | [optional] [default to "(|(memberUid={{.Username}})(member={{.UserDN}})(uniqueMember={{.UserDN}}))"]
**InsecureTls** | Pointer to **bool** | Skip LDAP server SSL Certificate verification - VERY insecure (optional) | [optional] 
**LastRotationTolerance** | Pointer to **int32** | The number of seconds after a Vault rotation where, if Active Directory shows a later rotation, it should be considered out-of-band. | [optional] [default to 5]
**Length** | Pointer to **int32** | The desired length of passwords that Vault generates. | [optional] [default to 64]
**MaxTtl** | Pointer to **int32** | In seconds, the maximum password time-to-live. | [optional] 
**PasswordPolicy** | Pointer to **string** | Name of the password policy to use to generate passwords. | [optional] 
**RequestTimeout** | Pointer to **int32** | Timeout, in seconds, for the connection when making requests against the server before returning back an error. | [optional] 
**Starttls** | Pointer to **bool** | Issue a StartTLS command after establishing unencrypted connection (optional) | [optional] 
**TlsMaxVersion** | Pointer to **string** | Maximum TLS version to use. Accepted values are &#39;tls10&#39;, &#39;tls11&#39;, &#39;tls12&#39; or &#39;tls13&#39;. Defaults to &#39;tls12&#39; | [optional] [default to "tls12"]
**TlsMinVersion** | Pointer to **string** | Minimum TLS version to use. Accepted values are &#39;tls10&#39;, &#39;tls11&#39;, &#39;tls12&#39; or &#39;tls13&#39;. Defaults to &#39;tls12&#39; | [optional] [default to "tls12"]
**Ttl** | Pointer to **int32** | In seconds, the default password time-to-live. | [optional] 
**Upndomain** | Pointer to **string** | Enables userPrincipalDomain login with [username]@UPNDomain (optional) | [optional] 
**Url** | Pointer to **string** | LDAP URL to connect to (default: ldap://127.0.0.1). Multiple URLs can be specified by concatenating them with commas; they will be tried in-order. | [optional] [default to "ldap://127.0.0.1"]
**UsePre111GroupCnBehavior** | Pointer to **bool** | In Vault 1.1.1 a fix for handling group CN values of different cases unfortunately introduced a regression that could cause previously defined groups to not be found due to a change in the resulting name. If set true, the pre-1.1.1 behavior for matching group CNs will be used. This is only needed in some upgrade scenarios for backwards compatibility. It is enabled by default if the config is upgraded but disabled by default on new configurations. | [optional] 
**UseTokenGroups** | Pointer to **bool** | If true, use the Active Directory tokenGroups constructed attribute of the user to find the group memberships. This will find all security groups including nested ones. | [optional] [default to false]
**Userattr** | Pointer to **string** | Attribute used for users (default: cn) | [optional] [default to "cn"]
**Userdn** | Pointer to **string** | LDAP domain to use for users (eg: ou&#x3D;People,dc&#x3D;example,dc&#x3D;org) | [optional] 
**Userfilter** | Pointer to **string** | Go template for LDAP user search filer (optional) The template can access the following context variables: UserAttr, Username Default: ({{.UserAttr}}&#x3D;{{.Username}}) | [optional] [default to "({{.UserAttr}}={{.Username}})"]
**UsernameAsAlias** | Pointer to **bool** | If true, sets the alias name to the username | [optional] [default to false]

## Methods

### NewActiveDirectoryWriteConfigRequest

`func NewActiveDirectoryWriteConfigRequest() *ActiveDirectoryWriteConfigRequest`

NewActiveDirectoryWriteConfigRequest instantiates a new ActiveDirectoryWriteConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActiveDirectoryWriteConfigRequestWithDefaults

`func NewActiveDirectoryWriteConfigRequestWithDefaults() *ActiveDirectoryWriteConfigRequest`

NewActiveDirectoryWriteConfigRequestWithDefaults instantiates a new ActiveDirectoryWriteConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnonymousGroupSearch

`func (o *ActiveDirectoryWriteConfigRequest) GetAnonymousGroupSearch() bool`

GetAnonymousGroupSearch returns the AnonymousGroupSearch field if non-nil, zero value otherwise.

### GetAnonymousGroupSearchOk

`func (o *ActiveDirectoryWriteConfigRequest) GetAnonymousGroupSearchOk() (*bool, bool)`

GetAnonymousGroupSearchOk returns a tuple with the AnonymousGroupSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymousGroupSearch

`func (o *ActiveDirectoryWriteConfigRequest) SetAnonymousGroupSearch(v bool)`

SetAnonymousGroupSearch sets AnonymousGroupSearch field to given value.

### HasAnonymousGroupSearch

`func (o *ActiveDirectoryWriteConfigRequest) HasAnonymousGroupSearch() bool`

HasAnonymousGroupSearch returns a boolean if a field has been set.

### GetBinddn

`func (o *ActiveDirectoryWriteConfigRequest) GetBinddn() string`

GetBinddn returns the Binddn field if non-nil, zero value otherwise.

### GetBinddnOk

`func (o *ActiveDirectoryWriteConfigRequest) GetBinddnOk() (*string, bool)`

GetBinddnOk returns a tuple with the Binddn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinddn

`func (o *ActiveDirectoryWriteConfigRequest) SetBinddn(v string)`

SetBinddn sets Binddn field to given value.

### HasBinddn

`func (o *ActiveDirectoryWriteConfigRequest) HasBinddn() bool`

HasBinddn returns a boolean if a field has been set.

### GetBindpass

`func (o *ActiveDirectoryWriteConfigRequest) GetBindpass() string`

GetBindpass returns the Bindpass field if non-nil, zero value otherwise.

### GetBindpassOk

`func (o *ActiveDirectoryWriteConfigRequest) GetBindpassOk() (*string, bool)`

GetBindpassOk returns a tuple with the Bindpass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindpass

`func (o *ActiveDirectoryWriteConfigRequest) SetBindpass(v string)`

SetBindpass sets Bindpass field to given value.

### HasBindpass

`func (o *ActiveDirectoryWriteConfigRequest) HasBindpass() bool`

HasBindpass returns a boolean if a field has been set.

### GetCaseSensitiveNames

`func (o *ActiveDirectoryWriteConfigRequest) GetCaseSensitiveNames() bool`

GetCaseSensitiveNames returns the CaseSensitiveNames field if non-nil, zero value otherwise.

### GetCaseSensitiveNamesOk

`func (o *ActiveDirectoryWriteConfigRequest) GetCaseSensitiveNamesOk() (*bool, bool)`

GetCaseSensitiveNamesOk returns a tuple with the CaseSensitiveNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseSensitiveNames

`func (o *ActiveDirectoryWriteConfigRequest) SetCaseSensitiveNames(v bool)`

SetCaseSensitiveNames sets CaseSensitiveNames field to given value.

### HasCaseSensitiveNames

`func (o *ActiveDirectoryWriteConfigRequest) HasCaseSensitiveNames() bool`

HasCaseSensitiveNames returns a boolean if a field has been set.

### GetCertificate

`func (o *ActiveDirectoryWriteConfigRequest) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *ActiveDirectoryWriteConfigRequest) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *ActiveDirectoryWriteConfigRequest) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *ActiveDirectoryWriteConfigRequest) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetClientTlsCert

`func (o *ActiveDirectoryWriteConfigRequest) GetClientTlsCert() string`

GetClientTlsCert returns the ClientTlsCert field if non-nil, zero value otherwise.

### GetClientTlsCertOk

`func (o *ActiveDirectoryWriteConfigRequest) GetClientTlsCertOk() (*string, bool)`

GetClientTlsCertOk returns a tuple with the ClientTlsCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientTlsCert

`func (o *ActiveDirectoryWriteConfigRequest) SetClientTlsCert(v string)`

SetClientTlsCert sets ClientTlsCert field to given value.

### HasClientTlsCert

`func (o *ActiveDirectoryWriteConfigRequest) HasClientTlsCert() bool`

HasClientTlsCert returns a boolean if a field has been set.

### GetClientTlsKey

`func (o *ActiveDirectoryWriteConfigRequest) GetClientTlsKey() string`

GetClientTlsKey returns the ClientTlsKey field if non-nil, zero value otherwise.

### GetClientTlsKeyOk

`func (o *ActiveDirectoryWriteConfigRequest) GetClientTlsKeyOk() (*string, bool)`

GetClientTlsKeyOk returns a tuple with the ClientTlsKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientTlsKey

`func (o *ActiveDirectoryWriteConfigRequest) SetClientTlsKey(v string)`

SetClientTlsKey sets ClientTlsKey field to given value.

### HasClientTlsKey

`func (o *ActiveDirectoryWriteConfigRequest) HasClientTlsKey() bool`

HasClientTlsKey returns a boolean if a field has been set.

### GetDenyNullBind

`func (o *ActiveDirectoryWriteConfigRequest) GetDenyNullBind() bool`

GetDenyNullBind returns the DenyNullBind field if non-nil, zero value otherwise.

### GetDenyNullBindOk

`func (o *ActiveDirectoryWriteConfigRequest) GetDenyNullBindOk() (*bool, bool)`

GetDenyNullBindOk returns a tuple with the DenyNullBind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyNullBind

`func (o *ActiveDirectoryWriteConfigRequest) SetDenyNullBind(v bool)`

SetDenyNullBind sets DenyNullBind field to given value.

### HasDenyNullBind

`func (o *ActiveDirectoryWriteConfigRequest) HasDenyNullBind() bool`

HasDenyNullBind returns a boolean if a field has been set.

### GetDiscoverdn

`func (o *ActiveDirectoryWriteConfigRequest) GetDiscoverdn() bool`

GetDiscoverdn returns the Discoverdn field if non-nil, zero value otherwise.

### GetDiscoverdnOk

`func (o *ActiveDirectoryWriteConfigRequest) GetDiscoverdnOk() (*bool, bool)`

GetDiscoverdnOk returns a tuple with the Discoverdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverdn

`func (o *ActiveDirectoryWriteConfigRequest) SetDiscoverdn(v bool)`

SetDiscoverdn sets Discoverdn field to given value.

### HasDiscoverdn

`func (o *ActiveDirectoryWriteConfigRequest) HasDiscoverdn() bool`

HasDiscoverdn returns a boolean if a field has been set.

### GetFormatter

`func (o *ActiveDirectoryWriteConfigRequest) GetFormatter() string`

GetFormatter returns the Formatter field if non-nil, zero value otherwise.

### GetFormatterOk

`func (o *ActiveDirectoryWriteConfigRequest) GetFormatterOk() (*string, bool)`

GetFormatterOk returns a tuple with the Formatter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatter

`func (o *ActiveDirectoryWriteConfigRequest) SetFormatter(v string)`

SetFormatter sets Formatter field to given value.

### HasFormatter

`func (o *ActiveDirectoryWriteConfigRequest) HasFormatter() bool`

HasFormatter returns a boolean if a field has been set.

### GetGroupattr

`func (o *ActiveDirectoryWriteConfigRequest) GetGroupattr() string`

GetGroupattr returns the Groupattr field if non-nil, zero value otherwise.

### GetGroupattrOk

`func (o *ActiveDirectoryWriteConfigRequest) GetGroupattrOk() (*string, bool)`

GetGroupattrOk returns a tuple with the Groupattr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupattr

`func (o *ActiveDirectoryWriteConfigRequest) SetGroupattr(v string)`

SetGroupattr sets Groupattr field to given value.

### HasGroupattr

`func (o *ActiveDirectoryWriteConfigRequest) HasGroupattr() bool`

HasGroupattr returns a boolean if a field has been set.

### GetGroupdn

`func (o *ActiveDirectoryWriteConfigRequest) GetGroupdn() string`

GetGroupdn returns the Groupdn field if non-nil, zero value otherwise.

### GetGroupdnOk

`func (o *ActiveDirectoryWriteConfigRequest) GetGroupdnOk() (*string, bool)`

GetGroupdnOk returns a tuple with the Groupdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupdn

`func (o *ActiveDirectoryWriteConfigRequest) SetGroupdn(v string)`

SetGroupdn sets Groupdn field to given value.

### HasGroupdn

`func (o *ActiveDirectoryWriteConfigRequest) HasGroupdn() bool`

HasGroupdn returns a boolean if a field has been set.

### GetGroupfilter

`func (o *ActiveDirectoryWriteConfigRequest) GetGroupfilter() string`

GetGroupfilter returns the Groupfilter field if non-nil, zero value otherwise.

### GetGroupfilterOk

`func (o *ActiveDirectoryWriteConfigRequest) GetGroupfilterOk() (*string, bool)`

GetGroupfilterOk returns a tuple with the Groupfilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupfilter

`func (o *ActiveDirectoryWriteConfigRequest) SetGroupfilter(v string)`

SetGroupfilter sets Groupfilter field to given value.

### HasGroupfilter

`func (o *ActiveDirectoryWriteConfigRequest) HasGroupfilter() bool`

HasGroupfilter returns a boolean if a field has been set.

### GetInsecureTls

`func (o *ActiveDirectoryWriteConfigRequest) GetInsecureTls() bool`

GetInsecureTls returns the InsecureTls field if non-nil, zero value otherwise.

### GetInsecureTlsOk

`func (o *ActiveDirectoryWriteConfigRequest) GetInsecureTlsOk() (*bool, bool)`

GetInsecureTlsOk returns a tuple with the InsecureTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecureTls

`func (o *ActiveDirectoryWriteConfigRequest) SetInsecureTls(v bool)`

SetInsecureTls sets InsecureTls field to given value.

### HasInsecureTls

`func (o *ActiveDirectoryWriteConfigRequest) HasInsecureTls() bool`

HasInsecureTls returns a boolean if a field has been set.

### GetLastRotationTolerance

`func (o *ActiveDirectoryWriteConfigRequest) GetLastRotationTolerance() int32`

GetLastRotationTolerance returns the LastRotationTolerance field if non-nil, zero value otherwise.

### GetLastRotationToleranceOk

`func (o *ActiveDirectoryWriteConfigRequest) GetLastRotationToleranceOk() (*int32, bool)`

GetLastRotationToleranceOk returns a tuple with the LastRotationTolerance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRotationTolerance

`func (o *ActiveDirectoryWriteConfigRequest) SetLastRotationTolerance(v int32)`

SetLastRotationTolerance sets LastRotationTolerance field to given value.

### HasLastRotationTolerance

`func (o *ActiveDirectoryWriteConfigRequest) HasLastRotationTolerance() bool`

HasLastRotationTolerance returns a boolean if a field has been set.

### GetLength

`func (o *ActiveDirectoryWriteConfigRequest) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *ActiveDirectoryWriteConfigRequest) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *ActiveDirectoryWriteConfigRequest) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *ActiveDirectoryWriteConfigRequest) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetMaxTtl

`func (o *ActiveDirectoryWriteConfigRequest) GetMaxTtl() int32`

GetMaxTtl returns the MaxTtl field if non-nil, zero value otherwise.

### GetMaxTtlOk

`func (o *ActiveDirectoryWriteConfigRequest) GetMaxTtlOk() (*int32, bool)`

GetMaxTtlOk returns a tuple with the MaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTtl

`func (o *ActiveDirectoryWriteConfigRequest) SetMaxTtl(v int32)`

SetMaxTtl sets MaxTtl field to given value.

### HasMaxTtl

`func (o *ActiveDirectoryWriteConfigRequest) HasMaxTtl() bool`

HasMaxTtl returns a boolean if a field has been set.

### GetPasswordPolicy

`func (o *ActiveDirectoryWriteConfigRequest) GetPasswordPolicy() string`

GetPasswordPolicy returns the PasswordPolicy field if non-nil, zero value otherwise.

### GetPasswordPolicyOk

`func (o *ActiveDirectoryWriteConfigRequest) GetPasswordPolicyOk() (*string, bool)`

GetPasswordPolicyOk returns a tuple with the PasswordPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordPolicy

`func (o *ActiveDirectoryWriteConfigRequest) SetPasswordPolicy(v string)`

SetPasswordPolicy sets PasswordPolicy field to given value.

### HasPasswordPolicy

`func (o *ActiveDirectoryWriteConfigRequest) HasPasswordPolicy() bool`

HasPasswordPolicy returns a boolean if a field has been set.

### GetRequestTimeout

`func (o *ActiveDirectoryWriteConfigRequest) GetRequestTimeout() int32`

GetRequestTimeout returns the RequestTimeout field if non-nil, zero value otherwise.

### GetRequestTimeoutOk

`func (o *ActiveDirectoryWriteConfigRequest) GetRequestTimeoutOk() (*int32, bool)`

GetRequestTimeoutOk returns a tuple with the RequestTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTimeout

`func (o *ActiveDirectoryWriteConfigRequest) SetRequestTimeout(v int32)`

SetRequestTimeout sets RequestTimeout field to given value.

### HasRequestTimeout

`func (o *ActiveDirectoryWriteConfigRequest) HasRequestTimeout() bool`

HasRequestTimeout returns a boolean if a field has been set.

### GetStarttls

`func (o *ActiveDirectoryWriteConfigRequest) GetStarttls() bool`

GetStarttls returns the Starttls field if non-nil, zero value otherwise.

### GetStarttlsOk

`func (o *ActiveDirectoryWriteConfigRequest) GetStarttlsOk() (*bool, bool)`

GetStarttlsOk returns a tuple with the Starttls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarttls

`func (o *ActiveDirectoryWriteConfigRequest) SetStarttls(v bool)`

SetStarttls sets Starttls field to given value.

### HasStarttls

`func (o *ActiveDirectoryWriteConfigRequest) HasStarttls() bool`

HasStarttls returns a boolean if a field has been set.

### GetTlsMaxVersion

`func (o *ActiveDirectoryWriteConfigRequest) GetTlsMaxVersion() string`

GetTlsMaxVersion returns the TlsMaxVersion field if non-nil, zero value otherwise.

### GetTlsMaxVersionOk

`func (o *ActiveDirectoryWriteConfigRequest) GetTlsMaxVersionOk() (*string, bool)`

GetTlsMaxVersionOk returns a tuple with the TlsMaxVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsMaxVersion

`func (o *ActiveDirectoryWriteConfigRequest) SetTlsMaxVersion(v string)`

SetTlsMaxVersion sets TlsMaxVersion field to given value.

### HasTlsMaxVersion

`func (o *ActiveDirectoryWriteConfigRequest) HasTlsMaxVersion() bool`

HasTlsMaxVersion returns a boolean if a field has been set.

### GetTlsMinVersion

`func (o *ActiveDirectoryWriteConfigRequest) GetTlsMinVersion() string`

GetTlsMinVersion returns the TlsMinVersion field if non-nil, zero value otherwise.

### GetTlsMinVersionOk

`func (o *ActiveDirectoryWriteConfigRequest) GetTlsMinVersionOk() (*string, bool)`

GetTlsMinVersionOk returns a tuple with the TlsMinVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsMinVersion

`func (o *ActiveDirectoryWriteConfigRequest) SetTlsMinVersion(v string)`

SetTlsMinVersion sets TlsMinVersion field to given value.

### HasTlsMinVersion

`func (o *ActiveDirectoryWriteConfigRequest) HasTlsMinVersion() bool`

HasTlsMinVersion returns a boolean if a field has been set.

### GetTtl

`func (o *ActiveDirectoryWriteConfigRequest) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *ActiveDirectoryWriteConfigRequest) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *ActiveDirectoryWriteConfigRequest) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *ActiveDirectoryWriteConfigRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUpndomain

`func (o *ActiveDirectoryWriteConfigRequest) GetUpndomain() string`

GetUpndomain returns the Upndomain field if non-nil, zero value otherwise.

### GetUpndomainOk

`func (o *ActiveDirectoryWriteConfigRequest) GetUpndomainOk() (*string, bool)`

GetUpndomainOk returns a tuple with the Upndomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpndomain

`func (o *ActiveDirectoryWriteConfigRequest) SetUpndomain(v string)`

SetUpndomain sets Upndomain field to given value.

### HasUpndomain

`func (o *ActiveDirectoryWriteConfigRequest) HasUpndomain() bool`

HasUpndomain returns a boolean if a field has been set.

### GetUrl

`func (o *ActiveDirectoryWriteConfigRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ActiveDirectoryWriteConfigRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ActiveDirectoryWriteConfigRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ActiveDirectoryWriteConfigRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUsePre111GroupCnBehavior

`func (o *ActiveDirectoryWriteConfigRequest) GetUsePre111GroupCnBehavior() bool`

GetUsePre111GroupCnBehavior returns the UsePre111GroupCnBehavior field if non-nil, zero value otherwise.

### GetUsePre111GroupCnBehaviorOk

`func (o *ActiveDirectoryWriteConfigRequest) GetUsePre111GroupCnBehaviorOk() (*bool, bool)`

GetUsePre111GroupCnBehaviorOk returns a tuple with the UsePre111GroupCnBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePre111GroupCnBehavior

`func (o *ActiveDirectoryWriteConfigRequest) SetUsePre111GroupCnBehavior(v bool)`

SetUsePre111GroupCnBehavior sets UsePre111GroupCnBehavior field to given value.

### HasUsePre111GroupCnBehavior

`func (o *ActiveDirectoryWriteConfigRequest) HasUsePre111GroupCnBehavior() bool`

HasUsePre111GroupCnBehavior returns a boolean if a field has been set.

### GetUseTokenGroups

`func (o *ActiveDirectoryWriteConfigRequest) GetUseTokenGroups() bool`

GetUseTokenGroups returns the UseTokenGroups field if non-nil, zero value otherwise.

### GetUseTokenGroupsOk

`func (o *ActiveDirectoryWriteConfigRequest) GetUseTokenGroupsOk() (*bool, bool)`

GetUseTokenGroupsOk returns a tuple with the UseTokenGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTokenGroups

`func (o *ActiveDirectoryWriteConfigRequest) SetUseTokenGroups(v bool)`

SetUseTokenGroups sets UseTokenGroups field to given value.

### HasUseTokenGroups

`func (o *ActiveDirectoryWriteConfigRequest) HasUseTokenGroups() bool`

HasUseTokenGroups returns a boolean if a field has been set.

### GetUserattr

`func (o *ActiveDirectoryWriteConfigRequest) GetUserattr() string`

GetUserattr returns the Userattr field if non-nil, zero value otherwise.

### GetUserattrOk

`func (o *ActiveDirectoryWriteConfigRequest) GetUserattrOk() (*string, bool)`

GetUserattrOk returns a tuple with the Userattr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserattr

`func (o *ActiveDirectoryWriteConfigRequest) SetUserattr(v string)`

SetUserattr sets Userattr field to given value.

### HasUserattr

`func (o *ActiveDirectoryWriteConfigRequest) HasUserattr() bool`

HasUserattr returns a boolean if a field has been set.

### GetUserdn

`func (o *ActiveDirectoryWriteConfigRequest) GetUserdn() string`

GetUserdn returns the Userdn field if non-nil, zero value otherwise.

### GetUserdnOk

`func (o *ActiveDirectoryWriteConfigRequest) GetUserdnOk() (*string, bool)`

GetUserdnOk returns a tuple with the Userdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserdn

`func (o *ActiveDirectoryWriteConfigRequest) SetUserdn(v string)`

SetUserdn sets Userdn field to given value.

### HasUserdn

`func (o *ActiveDirectoryWriteConfigRequest) HasUserdn() bool`

HasUserdn returns a boolean if a field has been set.

### GetUserfilter

`func (o *ActiveDirectoryWriteConfigRequest) GetUserfilter() string`

GetUserfilter returns the Userfilter field if non-nil, zero value otherwise.

### GetUserfilterOk

`func (o *ActiveDirectoryWriteConfigRequest) GetUserfilterOk() (*string, bool)`

GetUserfilterOk returns a tuple with the Userfilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserfilter

`func (o *ActiveDirectoryWriteConfigRequest) SetUserfilter(v string)`

SetUserfilter sets Userfilter field to given value.

### HasUserfilter

`func (o *ActiveDirectoryWriteConfigRequest) HasUserfilter() bool`

HasUserfilter returns a boolean if a field has been set.

### GetUsernameAsAlias

`func (o *ActiveDirectoryWriteConfigRequest) GetUsernameAsAlias() bool`

GetUsernameAsAlias returns the UsernameAsAlias field if non-nil, zero value otherwise.

### GetUsernameAsAliasOk

`func (o *ActiveDirectoryWriteConfigRequest) GetUsernameAsAliasOk() (*bool, bool)`

GetUsernameAsAliasOk returns a tuple with the UsernameAsAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameAsAlias

`func (o *ActiveDirectoryWriteConfigRequest) SetUsernameAsAlias(v bool)`

SetUsernameAsAlias sets UsernameAsAlias field to given value.

### HasUsernameAsAlias

`func (o *ActiveDirectoryWriteConfigRequest) HasUsernameAsAlias() bool`

HasUsernameAsAlias returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

