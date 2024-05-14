# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [api/v2/activity_service.proto](#api_v2_activity_service-proto)
    - [Activity](#lockets-api-v2-Activity)
    - [ActivityLocketCommentPayload](#lockets-api-v2-ActivityLocketCommentPayload)
    - [ActivityPayload](#lockets-api-v2-ActivityPayload)
    - [ActivityVersionUpdatePayload](#lockets-api-v2-ActivityVersionUpdatePayload)
    - [GetActivityRequest](#lockets-api-v2-GetActivityRequest)
    - [GetActivityResponse](#lockets-api-v2-GetActivityResponse)
  
    - [ActivityService](#lockets-api-v2-ActivityService)
  
- [api/v2/common.proto](#api_v2_common-proto)
    - [PageToken](#lockets-api-v2-PageToken)
  
    - [RowStatus](#lockets-api-v2-RowStatus)
  
- [api/v2/user_service.proto](#api_v2_user_service-proto)
    - [CreateUserAccessTokenRequest](#lockets-api-v2-CreateUserAccessTokenRequest)
    - [CreateUserAccessTokenResponse](#lockets-api-v2-CreateUserAccessTokenResponse)
    - [CreateUserRequest](#lockets-api-v2-CreateUserRequest)
    - [CreateUserResponse](#lockets-api-v2-CreateUserResponse)
    - [DeleteUserAccessTokenRequest](#lockets-api-v2-DeleteUserAccessTokenRequest)
    - [DeleteUserAccessTokenResponse](#lockets-api-v2-DeleteUserAccessTokenResponse)
    - [DeleteUserRequest](#lockets-api-v2-DeleteUserRequest)
    - [DeleteUserResponse](#lockets-api-v2-DeleteUserResponse)
    - [GetUserRequest](#lockets-api-v2-GetUserRequest)
    - [GetUserResponse](#lockets-api-v2-GetUserResponse)
    - [GetUserSettingRequest](#lockets-api-v2-GetUserSettingRequest)
    - [GetUserSettingResponse](#lockets-api-v2-GetUserSettingResponse)
    - [ListUserAccessTokensRequest](#lockets-api-v2-ListUserAccessTokensRequest)
    - [ListUserAccessTokensResponse](#lockets-api-v2-ListUserAccessTokensResponse)
    - [ListUsersRequest](#lockets-api-v2-ListUsersRequest)
    - [ListUsersResponse](#lockets-api-v2-ListUsersResponse)
    - [SearchUsersRequest](#lockets-api-v2-SearchUsersRequest)
    - [SearchUsersResponse](#lockets-api-v2-SearchUsersResponse)
    - [UpdateUserRequest](#lockets-api-v2-UpdateUserRequest)
    - [UpdateUserResponse](#lockets-api-v2-UpdateUserResponse)
    - [UpdateUserSettingRequest](#lockets-api-v2-UpdateUserSettingRequest)
    - [UpdateUserSettingResponse](#lockets-api-v2-UpdateUserSettingResponse)
    - [User](#lockets-api-v2-User)
    - [UserAccessToken](#lockets-api-v2-UserAccessToken)
    - [UserSetting](#lockets-api-v2-UserSetting)
  
    - [User.Role](#lockets-api-v2-User-Role)
  
    - [UserService](#lockets-api-v2-UserService)
  
- [api/v2/auth_service.proto](#api_v2_auth_service-proto)
    - [GetAuthStatusRequest](#lockets-api-v2-GetAuthStatusRequest)
    - [GetAuthStatusResponse](#lockets-api-v2-GetAuthStatusResponse)
    - [SignInRequest](#lockets-api-v2-SignInRequest)
    - [SignInResponse](#lockets-api-v2-SignInResponse)
    - [SignInWithSSORequest](#lockets-api-v2-SignInWithSSORequest)
    - [SignInWithSSOResponse](#lockets-api-v2-SignInWithSSOResponse)
    - [SignOutRequest](#lockets-api-v2-SignOutRequest)
    - [SignOutResponse](#lockets-api-v2-SignOutResponse)
    - [SignUpRequest](#lockets-api-v2-SignUpRequest)
    - [SignUpResponse](#lockets-api-v2-SignUpResponse)
  
    - [AuthService](#lockets-api-v2-AuthService)
  
- [api/v2/idp_service.proto](#api_v2_idp_service-proto)
    - [CreateIdentityProviderRequest](#lockets-api-v2-CreateIdentityProviderRequest)
    - [CreateIdentityProviderResponse](#lockets-api-v2-CreateIdentityProviderResponse)
    - [DeleteIdentityProviderRequest](#lockets-api-v2-DeleteIdentityProviderRequest)
    - [DeleteIdentityProviderResponse](#lockets-api-v2-DeleteIdentityProviderResponse)
    - [GetIdentityProviderRequest](#lockets-api-v2-GetIdentityProviderRequest)
    - [GetIdentityProviderResponse](#lockets-api-v2-GetIdentityProviderResponse)
    - [IdentityProvider](#lockets-api-v2-IdentityProvider)
    - [IdentityProvider.Config](#lockets-api-v2-IdentityProvider-Config)
    - [IdentityProvider.Config.FieldMapping](#lockets-api-v2-IdentityProvider-Config-FieldMapping)
    - [IdentityProvider.Config.OAuth2](#lockets-api-v2-IdentityProvider-Config-OAuth2)
    - [ListIdentityProvidersRequest](#lockets-api-v2-ListIdentityProvidersRequest)
    - [ListIdentityProvidersResponse](#lockets-api-v2-ListIdentityProvidersResponse)
    - [UpdateIdentityProviderRequest](#lockets-api-v2-UpdateIdentityProviderRequest)
    - [UpdateIdentityProviderResponse](#lockets-api-v2-UpdateIdentityProviderResponse)
  
    - [IdentityProvider.Type](#lockets-api-v2-IdentityProvider-Type)
  
    - [IdentityProviderService](#lockets-api-v2-IdentityProviderService)
  
- [api/v2/inbox_service.proto](#api_v2_inbox_service-proto)
    - [DeleteInboxRequest](#lockets-api-v2-DeleteInboxRequest)
    - [DeleteInboxResponse](#lockets-api-v2-DeleteInboxResponse)
    - [Inbox](#lockets-api-v2-Inbox)
    - [ListInboxesRequest](#lockets-api-v2-ListInboxesRequest)
    - [ListInboxesResponse](#lockets-api-v2-ListInboxesResponse)
    - [UpdateInboxRequest](#lockets-api-v2-UpdateInboxRequest)
    - [UpdateInboxResponse](#lockets-api-v2-UpdateInboxResponse)
  
    - [Inbox.Status](#lockets-api-v2-Inbox-Status)
    - [Inbox.Type](#lockets-api-v2-Inbox-Type)
  
    - [InboxService](#lockets-api-v2-InboxService)
  
- [api/v2/link_service.proto](#api_v2_link_service-proto)
    - [GetLinkMetadataRequest](#lockets-api-v2-GetLinkMetadataRequest)
    - [GetLinkMetadataResponse](#lockets-api-v2-GetLinkMetadataResponse)
    - [LinkMetadata](#lockets-api-v2-LinkMetadata)
  
    - [LinkService](#lockets-api-v2-LinkService)
  
- [api/v2/locket_relation_service.proto](#api_v2_locket_relation_service-proto)
    - [LocketRelation](#lockets-api-v2-LocketRelation)
  
    - [LocketRelation.Type](#lockets-api-v2-LocketRelation-Type)
  
- [api/v2/reaction_service.proto](#api_v2_reaction_service-proto)
    - [Reaction](#lockets-api-v2-Reaction)
  
    - [Reaction.Type](#lockets-api-v2-Reaction-Type)
  
- [api/v2/resource_service.proto](#api_v2_resource_service-proto)
    - [CreateResourceRequest](#lockets-api-v2-CreateResourceRequest)
    - [CreateResourceResponse](#lockets-api-v2-CreateResourceResponse)
    - [DeleteResourceRequest](#lockets-api-v2-DeleteResourceRequest)
    - [DeleteResourceResponse](#lockets-api-v2-DeleteResourceResponse)
    - [GetResourceRequest](#lockets-api-v2-GetResourceRequest)
    - [GetResourceResponse](#lockets-api-v2-GetResourceResponse)
    - [ListResourcesRequest](#lockets-api-v2-ListResourcesRequest)
    - [ListResourcesResponse](#lockets-api-v2-ListResourcesResponse)
    - [Resource](#lockets-api-v2-Resource)
    - [SearchResourcesRequest](#lockets-api-v2-SearchResourcesRequest)
    - [SearchResourcesResponse](#lockets-api-v2-SearchResourcesResponse)
    - [UpdateResourceRequest](#lockets-api-v2-UpdateResourceRequest)
    - [UpdateResourceResponse](#lockets-api-v2-UpdateResourceResponse)
  
    - [ResourceService](#lockets-api-v2-ResourceService)
  
- [api/v2/locket_service.proto](#api_v2_locket_service-proto)
    - [CreateLocketCommentRequest](#lockets-api-v2-CreateLocketCommentRequest)
    - [CreateLocketCommentResponse](#lockets-api-v2-CreateLocketCommentResponse)
    - [CreateLocketRequest](#lockets-api-v2-CreateLocketRequest)
    - [CreateLocketResponse](#lockets-api-v2-CreateLocketResponse)
    - [DeleteLocketReactionRequest](#lockets-api-v2-DeleteLocketReactionRequest)
    - [DeleteLocketReactionResponse](#lockets-api-v2-DeleteLocketReactionResponse)
    - [DeleteLocketRequest](#lockets-api-v2-DeleteLocketRequest)
    - [DeleteLocketResponse](#lockets-api-v2-DeleteLocketResponse)
    - [ExportLocketsRequest](#lockets-api-v2-ExportLocketsRequest)
    - [ExportLocketsResponse](#lockets-api-v2-ExportLocketsResponse)
    - [GetLocketRequest](#lockets-api-v2-GetLocketRequest)
    - [GetLocketResponse](#lockets-api-v2-GetLocketResponse)
    - [GetUserLocketsStatsRequest](#lockets-api-v2-GetUserLocketsStatsRequest)
    - [GetUserLocketsStatsResponse](#lockets-api-v2-GetUserLocketsStatsResponse)
    - [GetUserLocketsStatsResponse.StatsEntry](#lockets-api-v2-GetUserLocketsStatsResponse-StatsEntry)
    - [ListLocketCommentsRequest](#lockets-api-v2-ListLocketCommentsRequest)
    - [ListLocketCommentsResponse](#lockets-api-v2-ListLocketCommentsResponse)
    - [ListLocketReactionsRequest](#lockets-api-v2-ListLocketReactionsRequest)
    - [ListLocketReactionsResponse](#lockets-api-v2-ListLocketReactionsResponse)
    - [ListLocketRelationsRequest](#lockets-api-v2-ListLocketRelationsRequest)
    - [ListLocketRelationsResponse](#lockets-api-v2-ListLocketRelationsResponse)
    - [ListLocketResourcesRequest](#lockets-api-v2-ListLocketResourcesRequest)
    - [ListLocketResourcesResponse](#lockets-api-v2-ListLocketResourcesResponse)
    - [ListLocketsRequest](#lockets-api-v2-ListLocketsRequest)
    - [ListLocketsResponse](#lockets-api-v2-ListLocketsResponse)
    - [Locket](#lockets-api-v2-Locket)
    - [SearchLocketsRequest](#lockets-api-v2-SearchLocketsRequest)
    - [SearchLocketsResponse](#lockets-api-v2-SearchLocketsResponse)
    - [SetLocketRelationsRequest](#lockets-api-v2-SetLocketRelationsRequest)
    - [SetLocketRelationsResponse](#lockets-api-v2-SetLocketRelationsResponse)
    - [SetLocketResourcesRequest](#lockets-api-v2-SetLocketResourcesRequest)
    - [SetLocketResourcesResponse](#lockets-api-v2-SetLocketResourcesResponse)
    - [UpdateLocketRequest](#lockets-api-v2-UpdateLocketRequest)
    - [UpdateLocketResponse](#lockets-api-v2-UpdateLocketResponse)
    - [UpsertLocketReactionRequest](#lockets-api-v2-UpsertLocketReactionRequest)
    - [UpsertLocketReactionResponse](#lockets-api-v2-UpsertLocketReactionResponse)
  
    - [Visibility](#lockets-api-v2-Visibility)
  
    - [LocketService](#lockets-api-v2-LocketService)
  
- [api/v2/tag_service.proto](#api_v2_tag_service-proto)
    - [BatchUpsertTagRequest](#lockets-api-v2-BatchUpsertTagRequest)
    - [BatchUpsertTagResponse](#lockets-api-v2-BatchUpsertTagResponse)
    - [DeleteTagRequest](#lockets-api-v2-DeleteTagRequest)
    - [DeleteTagResponse](#lockets-api-v2-DeleteTagResponse)
    - [GetTagSuggestionsRequest](#lockets-api-v2-GetTagSuggestionsRequest)
    - [GetTagSuggestionsResponse](#lockets-api-v2-GetTagSuggestionsResponse)
    - [ListTagsRequest](#lockets-api-v2-ListTagsRequest)
    - [ListTagsResponse](#lockets-api-v2-ListTagsResponse)
    - [RenameTagRequest](#lockets-api-v2-RenameTagRequest)
    - [RenameTagResponse](#lockets-api-v2-RenameTagResponse)
    - [Tag](#lockets-api-v2-Tag)
    - [UpsertTagRequest](#lockets-api-v2-UpsertTagRequest)
    - [UpsertTagResponse](#lockets-api-v2-UpsertTagResponse)
  
    - [TagService](#lockets-api-v2-TagService)
  
- [api/v2/webhook_service.proto](#api_v2_webhook_service-proto)
    - [CreateWebhookRequest](#lockets-api-v2-CreateWebhookRequest)
    - [CreateWebhookResponse](#lockets-api-v2-CreateWebhookResponse)
    - [DeleteWebhookRequest](#lockets-api-v2-DeleteWebhookRequest)
    - [DeleteWebhookResponse](#lockets-api-v2-DeleteWebhookResponse)
    - [GetWebhookRequest](#lockets-api-v2-GetWebhookRequest)
    - [GetWebhookResponse](#lockets-api-v2-GetWebhookResponse)
    - [ListWebhooksRequest](#lockets-api-v2-ListWebhooksRequest)
    - [ListWebhooksResponse](#lockets-api-v2-ListWebhooksResponse)
    - [UpdateWebhookRequest](#lockets-api-v2-UpdateWebhookRequest)
    - [UpdateWebhookResponse](#lockets-api-v2-UpdateWebhookResponse)
    - [Webhook](#lockets-api-v2-Webhook)
  
    - [WebhookService](#lockets-api-v2-WebhookService)
  
- [api/v2/workspace_service.proto](#api_v2_workspace_service-proto)
    - [GetWorkspaceProfileRequest](#lockets-api-v2-GetWorkspaceProfileRequest)
    - [GetWorkspaceProfileResponse](#lockets-api-v2-GetWorkspaceProfileResponse)
    - [WorkspaceProfile](#lockets-api-v2-WorkspaceProfile)
  
    - [WorkspaceService](#lockets-api-v2-WorkspaceService)
  
- [api/v2/workspace_setting_service.proto](#api_v2_workspace_setting_service-proto)
    - [GetWorkspaceSettingRequest](#lockets-api-v2-GetWorkspaceSettingRequest)
    - [GetWorkspaceSettingResponse](#lockets-api-v2-GetWorkspaceSettingResponse)
    - [SetWorkspaceSettingRequest](#lockets-api-v2-SetWorkspaceSettingRequest)
    - [SetWorkspaceSettingResponse](#lockets-api-v2-SetWorkspaceSettingResponse)
    - [WorkspaceGeneralSetting](#lockets-api-v2-WorkspaceGeneralSetting)
    - [WorkspaceSetting](#lockets-api-v2-WorkspaceSetting)
  
    - [WorkspaceSettingService](#lockets-api-v2-WorkspaceSettingService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="api_v2_activity_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v2/activity_service.proto



<a name="lockets-api-v2-Activity"></a>

### Activity



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| creator_id | [int32](#int32) |  |  |
| type | [string](#string) |  |  |
| level | [string](#string) |  |  |
| create_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| payload | [ActivityPayload](#lockets-api-v2-ActivityPayload) |  |  |






<a name="lockets-api-v2-ActivityLocketCommentPayload"></a>

### ActivityLocketCommentPayload



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| locket_id | [int32](#int32) |  |  |
| related_locket_id | [int32](#int32) |  |  |






<a name="lockets-api-v2-ActivityPayload"></a>

### ActivityPayload



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| locket_comment | [ActivityLocketCommentPayload](#lockets-api-v2-ActivityLocketCommentPayload) |  |  |
| version_update | [ActivityVersionUpdatePayload](#lockets-api-v2-ActivityVersionUpdatePayload) |  |  |






<a name="lockets-api-v2-ActivityVersionUpdatePayload"></a>

### ActivityVersionUpdatePayload



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| version | [string](#string) |  |  |






<a name="lockets-api-v2-GetActivityRequest"></a>

### GetActivityRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |






<a name="lockets-api-v2-GetActivityResponse"></a>

### GetActivityResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| activity | [Activity](#lockets-api-v2-Activity) |  |  |





 

 

 


<a name="lockets-api-v2-ActivityService"></a>

### ActivityService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetActivity | [GetActivityRequest](#lockets-api-v2-GetActivityRequest) | [GetActivityResponse](#lockets-api-v2-GetActivityResponse) | GetActivity returns the activity with the given id. |

 



<a name="api_v2_common-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v2/common.proto



<a name="lockets-api-v2-PageToken"></a>

### PageToken
Used internally for obfuscating the page token.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| limit | [int32](#int32) |  |  |
| offset | [int32](#int32) |  |  |





 


<a name="lockets-api-v2-RowStatus"></a>

### RowStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| ROW_STATUS_UNSPECIFIED | 0 |  |
| ACTIVE | 1 |  |
| ARCHIVED | 2 |  |


 

 

 



<a name="api_v2_user_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v2/user_service.proto



<a name="lockets-api-v2-CreateUserAccessTokenRequest"></a>

### CreateUserAccessTokenRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the user. Format: users/{id} |
| description | [string](#string) |  |  |
| expires_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) | optional |  |






<a name="lockets-api-v2-CreateUserAccessTokenResponse"></a>

### CreateUserAccessTokenResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| access_token | [UserAccessToken](#lockets-api-v2-UserAccessToken) |  |  |






<a name="lockets-api-v2-CreateUserRequest"></a>

### CreateUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#lockets-api-v2-User) |  |  |






<a name="lockets-api-v2-CreateUserResponse"></a>

### CreateUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#lockets-api-v2-User) |  |  |






<a name="lockets-api-v2-DeleteUserAccessTokenRequest"></a>

### DeleteUserAccessTokenRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the user. Format: users/{id} |
| access_token | [string](#string) |  | access_token is the access token to delete. |






<a name="lockets-api-v2-DeleteUserAccessTokenResponse"></a>

### DeleteUserAccessTokenResponse







<a name="lockets-api-v2-DeleteUserRequest"></a>

### DeleteUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the user. Format: users/{id} |






<a name="lockets-api-v2-DeleteUserResponse"></a>

### DeleteUserResponse







<a name="lockets-api-v2-GetUserRequest"></a>

### GetUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the user. Format: users/{id} |






<a name="lockets-api-v2-GetUserResponse"></a>

### GetUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#lockets-api-v2-User) |  |  |






<a name="lockets-api-v2-GetUserSettingRequest"></a>

### GetUserSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the user. Format: users/{id} |






<a name="lockets-api-v2-GetUserSettingResponse"></a>

### GetUserSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| setting | [UserSetting](#lockets-api-v2-UserSetting) |  |  |






<a name="lockets-api-v2-ListUserAccessTokensRequest"></a>

### ListUserAccessTokensRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the user. Format: users/{id} |






<a name="lockets-api-v2-ListUserAccessTokensResponse"></a>

### ListUserAccessTokensResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| access_tokens | [UserAccessToken](#lockets-api-v2-UserAccessToken) | repeated |  |






<a name="lockets-api-v2-ListUsersRequest"></a>

### ListUsersRequest







<a name="lockets-api-v2-ListUsersResponse"></a>

### ListUsersResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| users | [User](#lockets-api-v2-User) | repeated |  |






<a name="lockets-api-v2-SearchUsersRequest"></a>

### SearchUsersRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| filter | [string](#string) |  | Filter is used to filter users returned in the list. Format: &#34;username == frank&#34; |






<a name="lockets-api-v2-SearchUsersResponse"></a>

### SearchUsersResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| users | [User](#lockets-api-v2-User) | repeated |  |






<a name="lockets-api-v2-UpdateUserRequest"></a>

### UpdateUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#lockets-api-v2-User) |  |  |
| update_mask | [google.protobuf.FieldMask](#google-protobuf-FieldMask) |  |  |






<a name="lockets-api-v2-UpdateUserResponse"></a>

### UpdateUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#lockets-api-v2-User) |  |  |






<a name="lockets-api-v2-UpdateUserSettingRequest"></a>

### UpdateUserSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| setting | [UserSetting](#lockets-api-v2-UserSetting) |  |  |
| update_mask | [google.protobuf.FieldMask](#google-protobuf-FieldMask) |  |  |






<a name="lockets-api-v2-UpdateUserSettingResponse"></a>

### UpdateUserSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| setting | [UserSetting](#lockets-api-v2-UserSetting) |  |  |






<a name="lockets-api-v2-User"></a>

### User



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the user. Format: users/{id} |
| id | [int32](#int32) |  | The system generated uid of the user. |
| role | [User.Role](#lockets-api-v2-User-Role) |  |  |
| username | [string](#string) |  |  |
| email | [string](#string) |  |  |
| nickname | [string](#string) |  |  |
| avatar_url | [string](#string) |  |  |
| description | [string](#string) |  |  |
| password | [string](#string) |  |  |
| row_status | [RowStatus](#lockets-api-v2-RowStatus) |  |  |
| create_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| update_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






<a name="lockets-api-v2-UserAccessToken"></a>

### UserAccessToken



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| access_token | [string](#string) |  |  |
| description | [string](#string) |  |  |
| issued_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| expires_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






<a name="lockets-api-v2-UserSetting"></a>

### UserSetting



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the user. Format: users/{id} |
| locale | [string](#string) |  | The preferred locale of the user. |
| appearance | [string](#string) |  | The preferred appearance of the user. |
| locket_visibility | [string](#string) |  | The default visibility of the locket. |
| telegram_user_id | [string](#string) |  | The telegram user id of the user. |





 


<a name="lockets-api-v2-User-Role"></a>

### User.Role


| Name | Number | Description |
| ---- | ------ | ----------- |
| ROLE_UNSPECIFIED | 0 |  |
| HOST | 1 |  |
| ADMIN | 2 |  |
| USER | 3 |  |


 

 


<a name="lockets-api-v2-UserService"></a>

### UserService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| ListUsers | [ListUsersRequest](#lockets-api-v2-ListUsersRequest) | [ListUsersResponse](#lockets-api-v2-ListUsersResponse) | ListUsers returns a list of users. |
| SearchUsers | [SearchUsersRequest](#lockets-api-v2-SearchUsersRequest) | [SearchUsersResponse](#lockets-api-v2-SearchUsersResponse) | SearchUsers searches users by filter. |
| GetUser | [GetUserRequest](#lockets-api-v2-GetUserRequest) | [GetUserResponse](#lockets-api-v2-GetUserResponse) | GetUser gets a user by name. |
| CreateUser | [CreateUserRequest](#lockets-api-v2-CreateUserRequest) | [CreateUserResponse](#lockets-api-v2-CreateUserResponse) | CreateUser creates a new user. |
| UpdateUser | [UpdateUserRequest](#lockets-api-v2-UpdateUserRequest) | [UpdateUserResponse](#lockets-api-v2-UpdateUserResponse) | UpdateUser updates a user. |
| DeleteUser | [DeleteUserRequest](#lockets-api-v2-DeleteUserRequest) | [DeleteUserResponse](#lockets-api-v2-DeleteUserResponse) | DeleteUser deletes a user. |
| GetUserSetting | [GetUserSettingRequest](#lockets-api-v2-GetUserSettingRequest) | [GetUserSettingResponse](#lockets-api-v2-GetUserSettingResponse) | GetUserSetting gets the setting of a user. |
| UpdateUserSetting | [UpdateUserSettingRequest](#lockets-api-v2-UpdateUserSettingRequest) | [UpdateUserSettingResponse](#lockets-api-v2-UpdateUserSettingResponse) | UpdateUserSetting updates the setting of a user. |
| ListUserAccessTokens | [ListUserAccessTokensRequest](#lockets-api-v2-ListUserAccessTokensRequest) | [ListUserAccessTokensResponse](#lockets-api-v2-ListUserAccessTokensResponse) | ListUserAccessTokens returns a list of access tokens for a user. |
| CreateUserAccessToken | [CreateUserAccessTokenRequest](#lockets-api-v2-CreateUserAccessTokenRequest) | [CreateUserAccessTokenResponse](#lockets-api-v2-CreateUserAccessTokenResponse) | CreateUserAccessToken creates a new access token for a user. |
| DeleteUserAccessToken | [DeleteUserAccessTokenRequest](#lockets-api-v2-DeleteUserAccessTokenRequest) | [DeleteUserAccessTokenResponse](#lockets-api-v2-DeleteUserAccessTokenResponse) | DeleteUserAccessToken deletes an access token for a user. |

 



<a name="api_v2_auth_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v2/auth_service.proto



<a name="lockets-api-v2-GetAuthStatusRequest"></a>

### GetAuthStatusRequest







<a name="lockets-api-v2-GetAuthStatusResponse"></a>

### GetAuthStatusResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#lockets-api-v2-User) |  |  |






<a name="lockets-api-v2-SignInRequest"></a>

### SignInRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| username | [string](#string) |  |  |
| password | [string](#string) |  |  |
| never_expire | [bool](#bool) |  |  |






<a name="lockets-api-v2-SignInResponse"></a>

### SignInResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#lockets-api-v2-User) |  |  |






<a name="lockets-api-v2-SignInWithSSORequest"></a>

### SignInWithSSORequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| idp_id | [int32](#int32) |  |  |
| code | [string](#string) |  |  |
| redirect_uri | [string](#string) |  |  |






<a name="lockets-api-v2-SignInWithSSOResponse"></a>

### SignInWithSSOResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#lockets-api-v2-User) |  |  |






<a name="lockets-api-v2-SignOutRequest"></a>

### SignOutRequest







<a name="lockets-api-v2-SignOutResponse"></a>

### SignOutResponse







<a name="lockets-api-v2-SignUpRequest"></a>

### SignUpRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| username | [string](#string) |  |  |
| password | [string](#string) |  |  |






<a name="lockets-api-v2-SignUpResponse"></a>

### SignUpResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#lockets-api-v2-User) |  |  |





 

 

 


<a name="lockets-api-v2-AuthService"></a>

### AuthService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetAuthStatus | [GetAuthStatusRequest](#lockets-api-v2-GetAuthStatusRequest) | [GetAuthStatusResponse](#lockets-api-v2-GetAuthStatusResponse) | GetAuthStatus returns the current auth status of the user. |
| SignIn | [SignInRequest](#lockets-api-v2-SignInRequest) | [SignInResponse](#lockets-api-v2-SignInResponse) | SignIn signs in the user with the given username and password. |
| SignInWithSSO | [SignInWithSSORequest](#lockets-api-v2-SignInWithSSORequest) | [SignInWithSSOResponse](#lockets-api-v2-SignInWithSSOResponse) | SignInWithSSO signs in the user with the given SSO code. |
| SignUp | [SignUpRequest](#lockets-api-v2-SignUpRequest) | [SignUpResponse](#lockets-api-v2-SignUpResponse) | SignUp signs up the user with the given username and password. |
| SignOut | [SignOutRequest](#lockets-api-v2-SignOutRequest) | [SignOutResponse](#lockets-api-v2-SignOutResponse) | SignOut signs out the user. |

 



<a name="api_v2_idp_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v2/idp_service.proto



<a name="lockets-api-v2-CreateIdentityProviderRequest"></a>

### CreateIdentityProviderRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| identity_provider | [IdentityProvider](#lockets-api-v2-IdentityProvider) |  | The identityProvider to create. |






<a name="lockets-api-v2-CreateIdentityProviderResponse"></a>

### CreateIdentityProviderResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| identity_provider | [IdentityProvider](#lockets-api-v2-IdentityProvider) |  | The created identityProvider. |






<a name="lockets-api-v2-DeleteIdentityProviderRequest"></a>

### DeleteIdentityProviderRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the identityProvider to delete. Format: identityProviders/{id} |






<a name="lockets-api-v2-DeleteIdentityProviderResponse"></a>

### DeleteIdentityProviderResponse







<a name="lockets-api-v2-GetIdentityProviderRequest"></a>

### GetIdentityProviderRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the identityProvider to get. Format: identityProviders/{id} |






<a name="lockets-api-v2-GetIdentityProviderResponse"></a>

### GetIdentityProviderResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| identity_provider | [IdentityProvider](#lockets-api-v2-IdentityProvider) |  | The identityProvider. |






<a name="lockets-api-v2-IdentityProvider"></a>

### IdentityProvider



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the identityProvider. Format: identityProviders/{id} |
| type | [IdentityProvider.Type](#lockets-api-v2-IdentityProvider-Type) |  |  |
| title | [string](#string) |  |  |
| identifier_filter | [string](#string) |  |  |
| config | [IdentityProvider.Config](#lockets-api-v2-IdentityProvider-Config) |  |  |






<a name="lockets-api-v2-IdentityProvider-Config"></a>

### IdentityProvider.Config



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| oauth2 | [IdentityProvider.Config.OAuth2](#lockets-api-v2-IdentityProvider-Config-OAuth2) |  |  |






<a name="lockets-api-v2-IdentityProvider-Config-FieldMapping"></a>

### IdentityProvider.Config.FieldMapping



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| identifier | [string](#string) |  |  |
| display_name | [string](#string) |  |  |
| email | [string](#string) |  |  |






<a name="lockets-api-v2-IdentityProvider-Config-OAuth2"></a>

### IdentityProvider.Config.OAuth2



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| client_id | [string](#string) |  |  |
| client_secret | [string](#string) |  |  |
| auth_url | [string](#string) |  |  |
| token_url | [string](#string) |  |  |
| user_info_url | [string](#string) |  |  |
| scopes | [string](#string) | repeated |  |
| field_mapping | [IdentityProvider.Config.FieldMapping](#lockets-api-v2-IdentityProvider-Config-FieldMapping) |  |  |






<a name="lockets-api-v2-ListIdentityProvidersRequest"></a>

### ListIdentityProvidersRequest







<a name="lockets-api-v2-ListIdentityProvidersResponse"></a>

### ListIdentityProvidersResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| identity_providers | [IdentityProvider](#lockets-api-v2-IdentityProvider) | repeated |  |






<a name="lockets-api-v2-UpdateIdentityProviderRequest"></a>

### UpdateIdentityProviderRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| identity_provider | [IdentityProvider](#lockets-api-v2-IdentityProvider) |  | The identityProvider to update. |
| update_mask | [google.protobuf.FieldMask](#google-protobuf-FieldMask) |  | The update mask applies to the resource. Only the top level fields of IdentityProvider are supported. |






<a name="lockets-api-v2-UpdateIdentityProviderResponse"></a>

### UpdateIdentityProviderResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| identity_provider | [IdentityProvider](#lockets-api-v2-IdentityProvider) |  | The updated identityProvider. |





 


<a name="lockets-api-v2-IdentityProvider-Type"></a>

### IdentityProvider.Type


| Name | Number | Description |
| ---- | ------ | ----------- |
| TYPE_UNSPECIFIED | 0 |  |
| OAUTH2 | 1 |  |


 

 


<a name="lockets-api-v2-IdentityProviderService"></a>

### IdentityProviderService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| ListIdentityProviders | [ListIdentityProvidersRequest](#lockets-api-v2-ListIdentityProvidersRequest) | [ListIdentityProvidersResponse](#lockets-api-v2-ListIdentityProvidersResponse) |  |
| GetIdentityProvider | [GetIdentityProviderRequest](#lockets-api-v2-GetIdentityProviderRequest) | [GetIdentityProviderResponse](#lockets-api-v2-GetIdentityProviderResponse) |  |
| CreateIdentityProvider | [CreateIdentityProviderRequest](#lockets-api-v2-CreateIdentityProviderRequest) | [CreateIdentityProviderResponse](#lockets-api-v2-CreateIdentityProviderResponse) |  |
| UpdateIdentityProvider | [UpdateIdentityProviderRequest](#lockets-api-v2-UpdateIdentityProviderRequest) | [UpdateIdentityProviderResponse](#lockets-api-v2-UpdateIdentityProviderResponse) | UpdateIdentityProvider updates an identity provider. |
| DeleteIdentityProvider | [DeleteIdentityProviderRequest](#lockets-api-v2-DeleteIdentityProviderRequest) | [DeleteIdentityProviderResponse](#lockets-api-v2-DeleteIdentityProviderResponse) | DeleteIdentityProvider deletes an identity provider. |

 



<a name="api_v2_inbox_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v2/inbox_service.proto



<a name="lockets-api-v2-DeleteInboxRequest"></a>

### DeleteInboxRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the inbox to delete. Format: inboxes/{id} |






<a name="lockets-api-v2-DeleteInboxResponse"></a>

### DeleteInboxResponse







<a name="lockets-api-v2-Inbox"></a>

### Inbox



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the inbox. Format: inboxes/{id} |
| sender | [string](#string) |  | Format: users/{id} |
| receiver | [string](#string) |  | Format: users/{id} |
| status | [Inbox.Status](#lockets-api-v2-Inbox-Status) |  |  |
| create_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| type | [Inbox.Type](#lockets-api-v2-Inbox-Type) |  |  |
| activity_id | [int32](#int32) | optional |  |






<a name="lockets-api-v2-ListInboxesRequest"></a>

### ListInboxesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [string](#string) |  | Format: users/{id} |






<a name="lockets-api-v2-ListInboxesResponse"></a>

### ListInboxesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| inboxes | [Inbox](#lockets-api-v2-Inbox) | repeated |  |






<a name="lockets-api-v2-UpdateInboxRequest"></a>

### UpdateInboxRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| inbox | [Inbox](#lockets-api-v2-Inbox) |  |  |
| update_mask | [google.protobuf.FieldMask](#google-protobuf-FieldMask) |  |  |






<a name="lockets-api-v2-UpdateInboxResponse"></a>

### UpdateInboxResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| inbox | [Inbox](#lockets-api-v2-Inbox) |  |  |





 


<a name="lockets-api-v2-Inbox-Status"></a>

### Inbox.Status


| Name | Number | Description |
| ---- | ------ | ----------- |
| STATUS_UNSPECIFIED | 0 |  |
| UNREAD | 1 |  |
| ARCHIVED | 2 |  |



<a name="lockets-api-v2-Inbox-Type"></a>

### Inbox.Type


| Name | Number | Description |
| ---- | ------ | ----------- |
| TYPE_UNSPECIFIED | 0 |  |
| TYPE_LOCKET_COMMENT | 1 |  |
| TYPE_VERSION_UPDATE | 2 |  |


 

 


<a name="lockets-api-v2-InboxService"></a>

### InboxService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| ListInboxes | [ListInboxesRequest](#lockets-api-v2-ListInboxesRequest) | [ListInboxesResponse](#lockets-api-v2-ListInboxesResponse) | ListInboxes lists inboxes for a user. |
| UpdateInbox | [UpdateInboxRequest](#lockets-api-v2-UpdateInboxRequest) | [UpdateInboxResponse](#lockets-api-v2-UpdateInboxResponse) | UpdateInbox updates an inbox. |
| DeleteInbox | [DeleteInboxRequest](#lockets-api-v2-DeleteInboxRequest) | [DeleteInboxResponse](#lockets-api-v2-DeleteInboxResponse) | DeleteInbox deletes an inbox. |

 



<a name="api_v2_link_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v2/link_service.proto



<a name="lockets-api-v2-GetLinkMetadataRequest"></a>

### GetLinkMetadataRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| link | [string](#string) |  |  |






<a name="lockets-api-v2-GetLinkMetadataResponse"></a>

### GetLinkMetadataResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| link_metadata | [LinkMetadata](#lockets-api-v2-LinkMetadata) |  |  |






<a name="lockets-api-v2-LinkMetadata"></a>

### LinkMetadata



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| title | [string](#string) |  |  |
| description | [string](#string) |  |  |
| image | [string](#string) |  |  |





 

 

 


<a name="lockets-api-v2-LinkService"></a>

### LinkService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetLinkMetadata | [GetLinkMetadataRequest](#lockets-api-v2-GetLinkMetadataRequest) | [GetLinkMetadataResponse](#lockets-api-v2-GetLinkMetadataResponse) |  |

 



<a name="api_v2_locket_relation_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v2/locket_relation_service.proto



<a name="lockets-api-v2-LocketRelation"></a>

### LocketRelation



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| locket | [string](#string) |  | The name of locket. Format: &#34;lockets/{uid}&#34; |
| related_locket | [string](#string) |  | The name of related locket. Format: &#34;lockets/{uid}&#34; |
| type | [LocketRelation.Type](#lockets-api-v2-LocketRelation-Type) |  |  |





 


<a name="lockets-api-v2-LocketRelation-Type"></a>

### LocketRelation.Type


| Name | Number | Description |
| ---- | ------ | ----------- |
| TYPE_UNSPECIFIED | 0 |  |
| REFERENCE | 1 |  |
| COMMENT | 2 |  |


 

 

 



<a name="api_v2_reaction_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v2/reaction_service.proto



<a name="lockets-api-v2-Reaction"></a>

### Reaction



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| creator | [string](#string) |  | The name of the creator. Format: users/{id} |
| content_id | [string](#string) |  |  |
| reaction_type | [Reaction.Type](#lockets-api-v2-Reaction-Type) |  |  |





 


<a name="lockets-api-v2-Reaction-Type"></a>

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


 

 

 



<a name="api_v2_resource_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v2/resource_service.proto



<a name="lockets-api-v2-CreateResourceRequest"></a>

### CreateResourceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| filename | [string](#string) |  |  |
| external_link | [string](#string) |  |  |
| type | [string](#string) |  |  |
| locket | [string](#string) | optional | Format: lockets/{id} |






<a name="lockets-api-v2-CreateResourceResponse"></a>

### CreateResourceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| resource | [Resource](#lockets-api-v2-Resource) |  |  |






<a name="lockets-api-v2-DeleteResourceRequest"></a>

### DeleteResourceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |






<a name="lockets-api-v2-DeleteResourceResponse"></a>

### DeleteResourceResponse







<a name="lockets-api-v2-GetResourceRequest"></a>

### GetResourceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |






<a name="lockets-api-v2-GetResourceResponse"></a>

### GetResourceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| resource | [Resource](#lockets-api-v2-Resource) |  |  |






<a name="lockets-api-v2-ListResourcesRequest"></a>

### ListResourcesRequest







<a name="lockets-api-v2-ListResourcesResponse"></a>

### ListResourcesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| resources | [Resource](#lockets-api-v2-Resource) | repeated |  |






<a name="lockets-api-v2-Resource"></a>

### Resource



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the resource. Format: resources/{id} id is the system generated unique identifier. |
| uid | [string](#string) |  | The user defined id of the resource. |
| create_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| filename | [string](#string) |  |  |
| external_link | [string](#string) |  |  |
| type | [string](#string) |  |  |
| size | [int64](#int64) |  |  |
| locket | [string](#string) | optional | Format: lockets/{id} |






<a name="lockets-api-v2-SearchResourcesRequest"></a>

### SearchResourcesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| filter | [string](#string) |  |  |






<a name="lockets-api-v2-SearchResourcesResponse"></a>

### SearchResourcesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| resources | [Resource](#lockets-api-v2-Resource) | repeated |  |






<a name="lockets-api-v2-UpdateResourceRequest"></a>

### UpdateResourceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| resource | [Resource](#lockets-api-v2-Resource) |  |  |
| update_mask | [google.protobuf.FieldMask](#google-protobuf-FieldMask) |  |  |






<a name="lockets-api-v2-UpdateResourceResponse"></a>

### UpdateResourceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| resource | [Resource](#lockets-api-v2-Resource) |  |  |





 

 

 


<a name="lockets-api-v2-ResourceService"></a>

### ResourceService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateResource | [CreateResourceRequest](#lockets-api-v2-CreateResourceRequest) | [CreateResourceResponse](#lockets-api-v2-CreateResourceResponse) | CreateResource creates a new resource. |
| ListResources | [ListResourcesRequest](#lockets-api-v2-ListResourcesRequest) | [ListResourcesResponse](#lockets-api-v2-ListResourcesResponse) | ListResources lists all resources. |
| SearchResources | [SearchResourcesRequest](#lockets-api-v2-SearchResourcesRequest) | [SearchResourcesResponse](#lockets-api-v2-SearchResourcesResponse) | SearchResources searches lockets. |
| GetResource | [GetResourceRequest](#lockets-api-v2-GetResourceRequest) | [GetResourceResponse](#lockets-api-v2-GetResourceResponse) | GetResource returns a resource by name. |
| UpdateResource | [UpdateResourceRequest](#lockets-api-v2-UpdateResourceRequest) | [UpdateResourceResponse](#lockets-api-v2-UpdateResourceResponse) | UpdateResource updates a resource. |
| DeleteResource | [DeleteResourceRequest](#lockets-api-v2-DeleteResourceRequest) | [DeleteResourceResponse](#lockets-api-v2-DeleteResourceResponse) | DeleteResource deletes a resource by name. |

 



<a name="api_v2_locket_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v2/locket_service.proto



<a name="lockets-api-v2-CreateLocketCommentRequest"></a>

### CreateLocketCommentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the locket. Format: lockets/{id} |
| comment | [CreateLocketRequest](#lockets-api-v2-CreateLocketRequest) |  |  |






<a name="lockets-api-v2-CreateLocketCommentResponse"></a>

### CreateLocketCommentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| locket | [Locket](#lockets-api-v2-Locket) |  |  |






<a name="lockets-api-v2-CreateLocketRequest"></a>

### CreateLocketRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| content | [string](#string) |  |  |
| visibility | [Visibility](#lockets-api-v2-Visibility) |  |  |






<a name="lockets-api-v2-CreateLocketResponse"></a>

### CreateLocketResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| locket | [Locket](#lockets-api-v2-Locket) |  |  |






<a name="lockets-api-v2-DeleteLocketReactionRequest"></a>

### DeleteLocketReactionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| reaction_id | [int32](#int32) |  |  |






<a name="lockets-api-v2-DeleteLocketReactionResponse"></a>

### DeleteLocketReactionResponse







<a name="lockets-api-v2-DeleteLocketRequest"></a>

### DeleteLocketRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the locket. Format: lockets/{id} |






<a name="lockets-api-v2-DeleteLocketResponse"></a>

### DeleteLocketResponse







<a name="lockets-api-v2-ExportLocketsRequest"></a>

### ExportLocketsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| filter | [string](#string) |  | Same as ListLocketsRequest.filter |






<a name="lockets-api-v2-ExportLocketsResponse"></a>

### ExportLocketsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| content | [bytes](#bytes) |  |  |






<a name="lockets-api-v2-GetLocketRequest"></a>

### GetLocketRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the locket. Format: lockets/{id} |






<a name="lockets-api-v2-GetLocketResponse"></a>

### GetLocketResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| locket | [Locket](#lockets-api-v2-Locket) |  |  |






<a name="lockets-api-v2-GetUserLocketsStatsRequest"></a>

### GetUserLocketsStatsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | name is the name of the user to get stats for. Format: users/{id} |
| timezone | [string](#string) |  | timezone location Format: uses tz identifier https://en.wikipedia.org/wiki/List_of_tz_database_time_zones |
| filter | [string](#string) |  | Same as ListLocketsRequest.filter |






<a name="lockets-api-v2-GetUserLocketsStatsResponse"></a>

### GetUserLocketsStatsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| stats | [GetUserLocketsStatsResponse.StatsEntry](#lockets-api-v2-GetUserLocketsStatsResponse-StatsEntry) | repeated | stats is the stats of locket creating/updating activities. key is the year-month-day string. e.g. &#34;2020-01-01&#34;. |






<a name="lockets-api-v2-GetUserLocketsStatsResponse-StatsEntry"></a>

### GetUserLocketsStatsResponse.StatsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [int32](#int32) |  |  |






<a name="lockets-api-v2-ListLocketCommentsRequest"></a>

### ListLocketCommentsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the locket. Format: lockets/{id} |






<a name="lockets-api-v2-ListLocketCommentsResponse"></a>

### ListLocketCommentsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| lockets | [Locket](#lockets-api-v2-Locket) | repeated |  |






<a name="lockets-api-v2-ListLocketReactionsRequest"></a>

### ListLocketReactionsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the locket. Format: lockets/{id} |






<a name="lockets-api-v2-ListLocketReactionsResponse"></a>

### ListLocketReactionsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| reactions | [Reaction](#lockets-api-v2-Reaction) | repeated |  |






<a name="lockets-api-v2-ListLocketRelationsRequest"></a>

### ListLocketRelationsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the locket. Format: lockets/{id} |






<a name="lockets-api-v2-ListLocketRelationsResponse"></a>

### ListLocketRelationsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| relations | [LocketRelation](#lockets-api-v2-LocketRelation) | repeated |  |






<a name="lockets-api-v2-ListLocketResourcesRequest"></a>

### ListLocketResourcesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the locket. Format: lockets/{id} |






<a name="lockets-api-v2-ListLocketResourcesResponse"></a>

### ListLocketResourcesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| resources | [Resource](#lockets-api-v2-Resource) | repeated |  |






<a name="lockets-api-v2-ListLocketsRequest"></a>

### ListLocketsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page_size | [int32](#int32) |  | The maximum number of lockets to return. |
| page_token | [string](#string) |  | A page token, received from a previous `ListLockets` call. Provide this to retrieve the subsequent page. |
| filter | [string](#string) |  | Filter is used to filter lockets returned in the list. Format: &#34;creator == users/{uid} &amp;&amp; visibilities == [&#39;PUBLIC&#39;, &#39;PROTECTED&#39;]&#34; |






<a name="lockets-api-v2-ListLocketsResponse"></a>

### ListLocketsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| lockets | [Locket](#lockets-api-v2-Locket) | repeated |  |
| next_page_token | [string](#string) |  | A token, which can be sent as `page_token` to retrieve the next page. If this field is omitted, there are no subsequent pages. |






<a name="lockets-api-v2-Locket"></a>

### Locket



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the locket. Format: lockets/{id} id is the system generated id. |
| uid | [string](#string) |  | The user defined id of the locket. |
| row_status | [RowStatus](#lockets-api-v2-RowStatus) |  |  |
| creator | [string](#string) |  | The name of the creator. Format: users/{id} |
| create_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| update_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| display_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| content | [string](#string) |  |  |
| visibility | [Visibility](#lockets-api-v2-Visibility) |  |  |
| pinned | [bool](#bool) |  |  |
| parent_id | [int32](#int32) | optional |  |
| resources | [Resource](#lockets-api-v2-Resource) | repeated |  |
| relations | [LocketRelation](#lockets-api-v2-LocketRelation) | repeated |  |
| reactions | [Reaction](#lockets-api-v2-Reaction) | repeated |  |






<a name="lockets-api-v2-SearchLocketsRequest"></a>

### SearchLocketsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| filter | [string](#string) |  | Filter is used to filter lockets returned. Format: &#34;creator == users/{uid} &amp;&amp; visibilities == [&#39;PUBLIC&#39;, &#39;PROTECTED&#39;]&#34; |






<a name="lockets-api-v2-SearchLocketsResponse"></a>

### SearchLocketsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| lockets | [Locket](#lockets-api-v2-Locket) | repeated |  |






<a name="lockets-api-v2-SetLocketRelationsRequest"></a>

### SetLocketRelationsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the locket. Format: lockets/{id} |
| relations | [LocketRelation](#lockets-api-v2-LocketRelation) | repeated |  |






<a name="lockets-api-v2-SetLocketRelationsResponse"></a>

### SetLocketRelationsResponse







<a name="lockets-api-v2-SetLocketResourcesRequest"></a>

### SetLocketResourcesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the locket. Format: lockets/{id} |
| resources | [Resource](#lockets-api-v2-Resource) | repeated |  |






<a name="lockets-api-v2-SetLocketResourcesResponse"></a>

### SetLocketResourcesResponse







<a name="lockets-api-v2-UpdateLocketRequest"></a>

### UpdateLocketRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| locket | [Locket](#lockets-api-v2-Locket) |  |  |
| update_mask | [google.protobuf.FieldMask](#google-protobuf-FieldMask) |  |  |






<a name="lockets-api-v2-UpdateLocketResponse"></a>

### UpdateLocketResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| locket | [Locket](#lockets-api-v2-Locket) |  |  |






<a name="lockets-api-v2-UpsertLocketReactionRequest"></a>

### UpsertLocketReactionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the locket. Format: lockets/{id} |
| reaction | [Reaction](#lockets-api-v2-Reaction) |  |  |






<a name="lockets-api-v2-UpsertLocketReactionResponse"></a>

### UpsertLocketReactionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| reaction | [Reaction](#lockets-api-v2-Reaction) |  |  |





 


<a name="lockets-api-v2-Visibility"></a>

### Visibility


| Name | Number | Description |
| ---- | ------ | ----------- |
| VISIBILITY_UNSPECIFIED | 0 |  |
| PRIVATE | 1 |  |
| PROTECTED | 2 |  |
| PUBLIC | 3 |  |


 

 


<a name="lockets-api-v2-LocketService"></a>

### LocketService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateLocket | [CreateLocketRequest](#lockets-api-v2-CreateLocketRequest) | [CreateLocketResponse](#lockets-api-v2-CreateLocketResponse) | CreateLocket creates a locket. |
| ListLockets | [ListLocketsRequest](#lockets-api-v2-ListLocketsRequest) | [ListLocketsResponse](#lockets-api-v2-ListLocketsResponse) | ListLockets lists lockets with pagination and filter. |
| SearchLockets | [SearchLocketsRequest](#lockets-api-v2-SearchLocketsRequest) | [SearchLocketsResponse](#lockets-api-v2-SearchLocketsResponse) | SearchLockets searches lockets. |
| GetLocket | [GetLocketRequest](#lockets-api-v2-GetLocketRequest) | [GetLocketResponse](#lockets-api-v2-GetLocketResponse) | GetLocket gets a locket. |
| UpdateLocket | [UpdateLocketRequest](#lockets-api-v2-UpdateLocketRequest) | [UpdateLocketResponse](#lockets-api-v2-UpdateLocketResponse) | UpdateLocket updates a locket. |
| DeleteLocket | [DeleteLocketRequest](#lockets-api-v2-DeleteLocketRequest) | [DeleteLocketResponse](#lockets-api-v2-DeleteLocketResponse) | DeleteLocket deletes a locket. |
| ExportLockets | [ExportLocketsRequest](#lockets-api-v2-ExportLocketsRequest) | [ExportLocketsResponse](#lockets-api-v2-ExportLocketsResponse) | ExportLockets exports lockets. |
| SetLocketResources | [SetLocketResourcesRequest](#lockets-api-v2-SetLocketResourcesRequest) | [SetLocketResourcesResponse](#lockets-api-v2-SetLocketResourcesResponse) | SetLocketResources sets resources for a locket. |
| ListLocketResources | [ListLocketResourcesRequest](#lockets-api-v2-ListLocketResourcesRequest) | [ListLocketResourcesResponse](#lockets-api-v2-ListLocketResourcesResponse) | ListLocketResources lists resources for a locket. |
| SetLocketRelations | [SetLocketRelationsRequest](#lockets-api-v2-SetLocketRelationsRequest) | [SetLocketRelationsResponse](#lockets-api-v2-SetLocketRelationsResponse) | SetLocketRelations sets relations for a locket. |
| ListLocketRelations | [ListLocketRelationsRequest](#lockets-api-v2-ListLocketRelationsRequest) | [ListLocketRelationsResponse](#lockets-api-v2-ListLocketRelationsResponse) | ListLocketRelations lists relations for a locket. |
| CreateLocketComment | [CreateLocketCommentRequest](#lockets-api-v2-CreateLocketCommentRequest) | [CreateLocketCommentResponse](#lockets-api-v2-CreateLocketCommentResponse) | CreateLocketComment creates a comment for a locket. |
| ListLocketComments | [ListLocketCommentsRequest](#lockets-api-v2-ListLocketCommentsRequest) | [ListLocketCommentsResponse](#lockets-api-v2-ListLocketCommentsResponse) | ListLocketComments lists comments for a locket. |
| GetUserLocketsStats | [GetUserLocketsStatsRequest](#lockets-api-v2-GetUserLocketsStatsRequest) | [GetUserLocketsStatsResponse](#lockets-api-v2-GetUserLocketsStatsResponse) | GetUserLocketsStats gets stats of lockets for a user. |
| ListLocketReactions | [ListLocketReactionsRequest](#lockets-api-v2-ListLocketReactionsRequest) | [ListLocketReactionsResponse](#lockets-api-v2-ListLocketReactionsResponse) | ListLocketReactions lists reactions for a locket. |
| UpsertLocketReaction | [UpsertLocketReactionRequest](#lockets-api-v2-UpsertLocketReactionRequest) | [UpsertLocketReactionResponse](#lockets-api-v2-UpsertLocketReactionResponse) | UpsertLocketReaction upserts a reaction for a locket. |
| DeleteLocketReaction | [DeleteLocketReactionRequest](#lockets-api-v2-DeleteLocketReactionRequest) | [DeleteLocketReactionResponse](#lockets-api-v2-DeleteLocketReactionResponse) | DeleteLocketReaction deletes a reaction for a locket. |

 



<a name="api_v2_tag_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v2/tag_service.proto



<a name="lockets-api-v2-BatchUpsertTagRequest"></a>

### BatchUpsertTagRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| requests | [UpsertTagRequest](#lockets-api-v2-UpsertTagRequest) | repeated |  |






<a name="lockets-api-v2-BatchUpsertTagResponse"></a>

### BatchUpsertTagResponse







<a name="lockets-api-v2-DeleteTagRequest"></a>

### DeleteTagRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tag | [Tag](#lockets-api-v2-Tag) |  |  |






<a name="lockets-api-v2-DeleteTagResponse"></a>

### DeleteTagResponse







<a name="lockets-api-v2-GetTagSuggestionsRequest"></a>

### GetTagSuggestionsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [string](#string) |  | The creator of tags. Format: users/{id} |






<a name="lockets-api-v2-GetTagSuggestionsResponse"></a>

### GetTagSuggestionsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tags | [string](#string) | repeated |  |






<a name="lockets-api-v2-ListTagsRequest"></a>

### ListTagsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [string](#string) |  | The creator of tags. Format: users/{id} |






<a name="lockets-api-v2-ListTagsResponse"></a>

### ListTagsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tags | [Tag](#lockets-api-v2-Tag) | repeated |  |






<a name="lockets-api-v2-RenameTagRequest"></a>

### RenameTagRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [string](#string) |  | The creator of tags. Format: users/{id} |
| old_name | [string](#string) |  |  |
| new_name | [string](#string) |  |  |






<a name="lockets-api-v2-RenameTagResponse"></a>

### RenameTagResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tag | [Tag](#lockets-api-v2-Tag) |  |  |






<a name="lockets-api-v2-Tag"></a>

### Tag



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| creator | [string](#string) |  | The creator of tags. Format: users/{id} |






<a name="lockets-api-v2-UpsertTagRequest"></a>

### UpsertTagRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |






<a name="lockets-api-v2-UpsertTagResponse"></a>

### UpsertTagResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tag | [Tag](#lockets-api-v2-Tag) |  |  |





 

 

 


<a name="lockets-api-v2-TagService"></a>

### TagService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| UpsertTag | [UpsertTagRequest](#lockets-api-v2-UpsertTagRequest) | [UpsertTagResponse](#lockets-api-v2-UpsertTagResponse) | UpsertTag upserts a tag. |
| BatchUpsertTag | [BatchUpsertTagRequest](#lockets-api-v2-BatchUpsertTagRequest) | [BatchUpsertTagResponse](#lockets-api-v2-BatchUpsertTagResponse) | BatchUpsertTag upserts multiple tags. |
| ListTags | [ListTagsRequest](#lockets-api-v2-ListTagsRequest) | [ListTagsResponse](#lockets-api-v2-ListTagsResponse) | ListTags lists tags. |
| RenameTag | [RenameTagRequest](#lockets-api-v2-RenameTagRequest) | [RenameTagResponse](#lockets-api-v2-RenameTagResponse) | RenameTag renames a tag. All related lockets will be updated. |
| DeleteTag | [DeleteTagRequest](#lockets-api-v2-DeleteTagRequest) | [DeleteTagResponse](#lockets-api-v2-DeleteTagResponse) | DeleteTag deletes a tag. |
| GetTagSuggestions | [GetTagSuggestionsRequest](#lockets-api-v2-GetTagSuggestionsRequest) | [GetTagSuggestionsResponse](#lockets-api-v2-GetTagSuggestionsResponse) | GetTagSuggestions gets tag suggestions from the user&#39;s lockets. |

 



<a name="api_v2_webhook_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v2/webhook_service.proto



<a name="lockets-api-v2-CreateWebhookRequest"></a>

### CreateWebhookRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| url | [string](#string) |  |  |






<a name="lockets-api-v2-CreateWebhookResponse"></a>

### CreateWebhookResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| webhook | [Webhook](#lockets-api-v2-Webhook) |  |  |






<a name="lockets-api-v2-DeleteWebhookRequest"></a>

### DeleteWebhookRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |






<a name="lockets-api-v2-DeleteWebhookResponse"></a>

### DeleteWebhookResponse







<a name="lockets-api-v2-GetWebhookRequest"></a>

### GetWebhookRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |






<a name="lockets-api-v2-GetWebhookResponse"></a>

### GetWebhookResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| webhook | [Webhook](#lockets-api-v2-Webhook) |  |  |






<a name="lockets-api-v2-ListWebhooksRequest"></a>

### ListWebhooksRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| creator_id | [int32](#int32) |  |  |






<a name="lockets-api-v2-ListWebhooksResponse"></a>

### ListWebhooksResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| webhooks | [Webhook](#lockets-api-v2-Webhook) | repeated |  |






<a name="lockets-api-v2-UpdateWebhookRequest"></a>

### UpdateWebhookRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| webhook | [Webhook](#lockets-api-v2-Webhook) |  |  |
| update_mask | [google.protobuf.FieldMask](#google-protobuf-FieldMask) |  |  |






<a name="lockets-api-v2-UpdateWebhookResponse"></a>

### UpdateWebhookResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| webhook | [Webhook](#lockets-api-v2-Webhook) |  |  |






<a name="lockets-api-v2-Webhook"></a>

### Webhook



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| creator_id | [int32](#int32) |  |  |
| created_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| updated_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| row_status | [RowStatus](#lockets-api-v2-RowStatus) |  |  |
| name | [string](#string) |  |  |
| url | [string](#string) |  |  |





 

 

 


<a name="lockets-api-v2-WebhookService"></a>

### WebhookService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateWebhook | [CreateWebhookRequest](#lockets-api-v2-CreateWebhookRequest) | [CreateWebhookResponse](#lockets-api-v2-CreateWebhookResponse) | CreateWebhook creates a new webhook. |
| GetWebhook | [GetWebhookRequest](#lockets-api-v2-GetWebhookRequest) | [GetWebhookResponse](#lockets-api-v2-GetWebhookResponse) | GetWebhook returns a webhook by id. |
| ListWebhooks | [ListWebhooksRequest](#lockets-api-v2-ListWebhooksRequest) | [ListWebhooksResponse](#lockets-api-v2-ListWebhooksResponse) | ListWebhooks returns a list of webhooks. |
| UpdateWebhook | [UpdateWebhookRequest](#lockets-api-v2-UpdateWebhookRequest) | [UpdateWebhookResponse](#lockets-api-v2-UpdateWebhookResponse) | UpdateWebhook updates a webhook. |
| DeleteWebhook | [DeleteWebhookRequest](#lockets-api-v2-DeleteWebhookRequest) | [DeleteWebhookResponse](#lockets-api-v2-DeleteWebhookResponse) | DeleteWebhook deletes a webhook by id. |

 



<a name="api_v2_workspace_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v2/workspace_service.proto



<a name="lockets-api-v2-GetWorkspaceProfileRequest"></a>

### GetWorkspaceProfileRequest







<a name="lockets-api-v2-GetWorkspaceProfileResponse"></a>

### GetWorkspaceProfileResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| workspace_profile | [WorkspaceProfile](#lockets-api-v2-WorkspaceProfile) |  |  |






<a name="lockets-api-v2-WorkspaceProfile"></a>

### WorkspaceProfile



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| owner | [string](#string) |  | The name of intance owner. Format: &#34;users/{id}&#34; |
| version | [string](#string) |  | version is the current version of instance |
| mode | [string](#string) |  | mode is the instance mode (e.g. &#34;prod&#34;, &#34;dev&#34; or &#34;demo&#34;). |
| disallow_signup | [bool](#bool) |  | disallow_signup is whether the signup is disallowed. |
| disable_password_login | [bool](#bool) |  | disable_password_login is whether the password login is disabled. |
| additional_script | [string](#string) |  | additional_script is the additional script. |
| additional_style | [string](#string) |  | additional_style is the additional style. |





 

 

 


<a name="lockets-api-v2-WorkspaceService"></a>

### WorkspaceService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetWorkspaceProfile | [GetWorkspaceProfileRequest](#lockets-api-v2-GetWorkspaceProfileRequest) | [GetWorkspaceProfileResponse](#lockets-api-v2-GetWorkspaceProfileResponse) | GetWorkspaceProfile returns the workspace profile. |

 



<a name="api_v2_workspace_setting_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v2/workspace_setting_service.proto



<a name="lockets-api-v2-GetWorkspaceSettingRequest"></a>

### GetWorkspaceSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The resource name of the workspace setting. Format: settings/{setting} |






<a name="lockets-api-v2-GetWorkspaceSettingResponse"></a>

### GetWorkspaceSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| setting | [WorkspaceSetting](#lockets-api-v2-WorkspaceSetting) |  |  |






<a name="lockets-api-v2-SetWorkspaceSettingRequest"></a>

### SetWorkspaceSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| setting | [WorkspaceSetting](#lockets-api-v2-WorkspaceSetting) |  | setting is the setting to update. |






<a name="lockets-api-v2-SetWorkspaceSettingResponse"></a>

### SetWorkspaceSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| setting | [WorkspaceSetting](#lockets-api-v2-WorkspaceSetting) |  |  |






<a name="lockets-api-v2-WorkspaceGeneralSetting"></a>

### WorkspaceGeneralSetting



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| instance_url | [string](#string) |  | instance_url is the instance URL. |
| disallow_signup | [bool](#bool) |  | disallow_signup is the flag to disallow signup. |
| disallow_password_login | [bool](#bool) |  | disallow_password_login is the flag to disallow password login. |
| additional_script | [string](#string) |  | additional_script is the additional script. |
| additional_style | [string](#string) |  | additional_style is the additional style. |






<a name="lockets-api-v2-WorkspaceSetting"></a>

### WorkspaceSetting



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | name is the name of the setting. Format: settings/{setting} |
| general_setting | [WorkspaceGeneralSetting](#lockets-api-v2-WorkspaceGeneralSetting) |  | general_setting is the general setting of workspace. |





 

 

 


<a name="lockets-api-v2-WorkspaceSettingService"></a>

### WorkspaceSettingService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetWorkspaceSetting | [GetWorkspaceSettingRequest](#lockets-api-v2-GetWorkspaceSettingRequest) | [GetWorkspaceSettingResponse](#lockets-api-v2-GetWorkspaceSettingResponse) | GetWorkspaceSetting returns the setting by name. |
| SetWorkspaceSetting | [SetWorkspaceSettingRequest](#lockets-api-v2-SetWorkspaceSettingRequest) | [SetWorkspaceSettingResponse](#lockets-api-v2-SetWorkspaceSettingResponse) | SetWorkspaceSetting updates the setting. |

 



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

