# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [store/activity.proto](#store_activity-proto)
    - [ActivityLocketCommentPayload](#lockets-store-ActivityLocketCommentPayload)
    - [ActivityPayload](#lockets-store-ActivityPayload)
    - [ActivityVersionUpdatePayload](#lockets-store-ActivityVersionUpdatePayload)
  
- [store/common.proto](#store_common-proto)
    - [RowStatus](#lockets-store-RowStatus)
  
- [store/idp.proto](#store_idp-proto)
    - [IdentityProviderConfig](#lockets-store-IdentityProviderConfig)
    - [IdentityProviderConfig.FieldMapping](#lockets-store-IdentityProviderConfig-FieldMapping)
    - [IdentityProviderConfig.OAuth2](#lockets-store-IdentityProviderConfig-OAuth2)
  
- [store/inbox.proto](#store_inbox-proto)
    - [InboxMessage](#lockets-store-InboxMessage)
  
    - [InboxMessage.Type](#lockets-store-InboxMessage-Type)
  
- [store/reaction.proto](#store_reaction-proto)
    - [Reaction](#lockets-store-Reaction)
  
    - [Reaction.Type](#lockets-store-Reaction-Type)
  
- [store/user_setting.proto](#store_user_setting-proto)
    - [AccessTokensUserSetting](#lockets-store-AccessTokensUserSetting)
    - [AccessTokensUserSetting.AccessToken](#lockets-store-AccessTokensUserSetting-AccessToken)
    - [UserSetting](#lockets-store-UserSetting)
  
    - [UserSettingKey](#lockets-store-UserSettingKey)
  
- [store/webhook.proto](#store_webhook-proto)
    - [Webhook](#lockets-store-Webhook)
  
- [store/workspace_setting.proto](#store_workspace_setting-proto)
    - [WorkspaceGeneralSetting](#lockets-store-WorkspaceGeneralSetting)
    - [WorkspaceSetting](#lockets-store-WorkspaceSetting)
  
    - [WorkspaceSettingKey](#lockets-store-WorkspaceSettingKey)
  
- [Scalar Value Types](#scalar-value-types)



<a name="store_activity-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## store/activity.proto



<a name="lockets-store-ActivityLocketCommentPayload"></a>

### ActivityLocketCommentPayload



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| locket_id | [int32](#int32) |  |  |
| related_locket_id | [int32](#int32) |  |  |






<a name="lockets-store-ActivityPayload"></a>

### ActivityPayload



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| locket_comment | [ActivityLocketCommentPayload](#lockets-store-ActivityLocketCommentPayload) |  |  |
| version_update | [ActivityVersionUpdatePayload](#lockets-store-ActivityVersionUpdatePayload) |  |  |






<a name="lockets-store-ActivityVersionUpdatePayload"></a>

### ActivityVersionUpdatePayload



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| version | [string](#string) |  |  |





 

 

 

 



<a name="store_common-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## store/common.proto


 


<a name="lockets-store-RowStatus"></a>

### RowStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| ROW_STATUS_UNSPECIFIED | 0 |  |
| NORMAL | 1 |  |
| ARCHIVED | 2 |  |


 

 

 



<a name="store_idp-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## store/idp.proto



<a name="lockets-store-IdentityProviderConfig"></a>

### IdentityProviderConfig



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| oauth2 | [IdentityProviderConfig.OAuth2](#lockets-store-IdentityProviderConfig-OAuth2) |  |  |






<a name="lockets-store-IdentityProviderConfig-FieldMapping"></a>

### IdentityProviderConfig.FieldMapping



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| identifier | [string](#string) |  |  |
| display_name | [string](#string) |  |  |
| email | [string](#string) |  |  |






<a name="lockets-store-IdentityProviderConfig-OAuth2"></a>

### IdentityProviderConfig.OAuth2



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| client_id | [string](#string) |  |  |
| client_secret | [string](#string) |  |  |
| auth_url | [string](#string) |  |  |
| token_url | [string](#string) |  |  |
| user_info_url | [string](#string) |  |  |
| scopes | [string](#string) | repeated |  |
| field_mapping | [IdentityProviderConfig.FieldMapping](#lockets-store-IdentityProviderConfig-FieldMapping) |  |  |





 

 

 

 



<a name="store_inbox-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## store/inbox.proto



<a name="lockets-store-InboxMessage"></a>

### InboxMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [InboxMessage.Type](#lockets-store-InboxMessage-Type) |  |  |
| activity_id | [int32](#int32) | optional |  |





 


<a name="lockets-store-InboxMessage-Type"></a>

### InboxMessage.Type


| Name | Number | Description |
| ---- | ------ | ----------- |
| TYPE_UNSPECIFIED | 0 |  |
| TYPE_LOCKET_COMMENT | 1 |  |
| TYPE_VERSION_UPDATE | 2 |  |


 

 

 



<a name="store_reaction-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## store/reaction.proto



<a name="lockets-store-Reaction"></a>

### Reaction



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| created_ts | [int64](#int64) |  |  |
| creator_id | [int32](#int32) |  |  |
| content_id | [string](#string) |  | content_id is the id of the content that the reaction is for. This can be a locket. e.g. lockets/101 |
| reaction_type | [Reaction.Type](#lockets-store-Reaction-Type) |  |  |





 


<a name="lockets-store-Reaction-Type"></a>

### Reaction.Type


| Name | Number | Description |
| ---- | ------ | ----------- |
| TYPE_UNSPECIFIED | 0 |  |
| THUMBS_UP | 1 |  |
| THUMBS_DOWN | 2 |  |
| HEART | 3 |  |
| FIRE | 4 |  |
| CLAPPING_HANDS | 5 |  |
| LAUGH | 6 |  |
| OK_HAND | 7 |  |
| ROCKET | 8 |  |
| EYES | 9 |  |
| THINKING_FACE | 10 |  |
| CLOWN_FACE | 11 |  |
| QUESTION_MARK | 12 |  |


 

 

 



<a name="store_user_setting-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## store/user_setting.proto



<a name="lockets-store-AccessTokensUserSetting"></a>

### AccessTokensUserSetting



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| access_tokens | [AccessTokensUserSetting.AccessToken](#lockets-store-AccessTokensUserSetting-AccessToken) | repeated |  |






<a name="lockets-store-AccessTokensUserSetting-AccessToken"></a>

### AccessTokensUserSetting.AccessToken



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| access_token | [string](#string) |  | The access token is a JWT token. Including expiration time, issuer, etc. |
| description | [string](#string) |  | A description for the access token. |






<a name="lockets-store-UserSetting"></a>

### UserSetting



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [int32](#int32) |  |  |
| key | [UserSettingKey](#lockets-store-UserSettingKey) |  |  |
| access_tokens | [AccessTokensUserSetting](#lockets-store-AccessTokensUserSetting) |  |  |
| locale | [string](#string) |  |  |
| appearance | [string](#string) |  |  |
| locket_visibility | [string](#string) |  |  |
| telegram_user_id | [string](#string) |  |  |





 


<a name="lockets-store-UserSettingKey"></a>

### UserSettingKey


| Name | Number | Description |
| ---- | ------ | ----------- |
| USER_SETTING_KEY_UNSPECIFIED | 0 |  |
| USER_SETTING_ACCESS_TOKENS | 1 | Access tokens for the user. |
| USER_SETTING_LOCALE | 2 | The locale of the user. |
| USER_SETTING_APPEARANCE | 3 | The appearance of the user. |
| USER_SETTING_LOCKET_VISIBILITY | 4 | The visibility of the locket. |
| USER_SETTING_TELEGRAM_USER_ID | 5 | The telegram user id of the user. |


 

 

 



<a name="store_webhook-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## store/webhook.proto



<a name="lockets-store-Webhook"></a>

### Webhook



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| created_ts | [int64](#int64) |  |  |
| updated_ts | [int64](#int64) |  |  |
| creator_id | [int32](#int32) |  |  |
| row_status | [RowStatus](#lockets-store-RowStatus) |  |  |
| name | [string](#string) |  |  |
| url | [string](#string) |  |  |





 

 

 

 



<a name="store_workspace_setting-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## store/workspace_setting.proto



<a name="lockets-store-WorkspaceGeneralSetting"></a>

### WorkspaceGeneralSetting



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| instance_url | [string](#string) |  | instance_url is the instance URL. |
| disallow_signup | [bool](#bool) |  | disallow_signup is the flag to disallow signup. |
| disallow_password_login | [bool](#bool) |  | disallow_password_login is the flag to disallow password login. |
| additional_script | [string](#string) |  | additional_script is the additional script. |
| additional_style | [string](#string) |  | additional_style is the additional style. |






<a name="lockets-store-WorkspaceSetting"></a>

### WorkspaceSetting



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [WorkspaceSettingKey](#lockets-store-WorkspaceSettingKey) |  |  |
| general | [WorkspaceGeneralSetting](#lockets-store-WorkspaceGeneralSetting) |  |  |





 


<a name="lockets-store-WorkspaceSettingKey"></a>

### WorkspaceSettingKey


| Name | Number | Description |
| ---- | ------ | ----------- |
| WORKSPACE_SETTING_KEY_UNSPECIFIED | 0 |  |
| WORKSPACE_SETTING_GENERAL | 1 | WORKSPACE_SETTING_GENERAL is the key for general settings. |


 

 

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

