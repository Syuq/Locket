# api/v2/activity_service.proto
## Version: version not set

---
## AuthService

### /api/v2/auth/signin

#### POST
##### Summary

SignIn signs in the user with the given username and password.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| username | query |  | No | string |
| password | query |  | No | string |
| neverExpire | query |  | No | boolean |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v2SignInResponse](#v2signinresponse) |
| default | An unexpected error response. | [googlerpcStatus](#googlerpcstatus) |

### /api/v2/auth/signin/sso

#### POST
##### Summary

SignInWithSSO signs in the user with the given SSO code.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| idpId | query |  | No | integer |
| code | query |  | No | string |
| redirectUri | query |  | No | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v2SignInWithSSOResponse](#v2signinwithssoresponse) |
| default | An unexpected error response. | [googlerpcStatus](#googlerpcstatus) |

### /api/v2/auth/signout

#### POST
##### Summary

SignOut signs out the user.

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v2SignOutResponse](#v2signoutresponse) |
| default | An unexpected error response. | [googlerpcStatus](#googlerpcstatus) |

### /api/v2/auth/signup

#### POST
##### Summary

SignUp signs up the user with the given username and password.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| username | query |  | No | string |
| password | query |  | No | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v2SignUpResponse](#v2signupresponse) |
| default | An unexpected error response. | [googlerpcStatus](#googlerpcstatus) |

### /api/v2/auth/status

#### POST
##### Summary

GetAuthStatus returns the current auth status of the user.

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v2GetAuthStatusResponse](#v2getauthstatusresponse) |
| default | An unexpected error response. | [googlerpcStatus](#googlerpcstatus) |

---
## InboxService

### /api/v2/inboxes

#### GET
##### Summary

ListInboxes lists inboxes for a user.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| user | query | Format: users/{username} | No | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v2ListInboxesResponse](#v2listinboxesresponse) |
| default | An unexpected error response. | [googlerpcStatus](#googlerpcstatus) |

### /api/v2/{inbox.name}

#### PATCH
##### Summary

UpdateInbox updates an inbox.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| inbox.name | path | The name of the inbox. Format: inboxes/{uid} | Yes | string |
| inbox | body |  | Yes | { **"sender"**: string, **"receiver"**: string, **"status"**: [v2InboxStatus](#v2inboxstatus), **"createTime"**: dateTime, **"type"**: [v2InboxType](#v2inboxtype), **"activityId"**: integer } |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v2UpdateInboxResponse](#v2updateinboxresponse) |
| default | An unexpected error response. | [googlerpcStatus](#googlerpcstatus) |

### /api/v2/{name_1}

#### GET
##### Summary

GetResource returns a resource by name.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| name_1 | path |  | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v2GetResourceResponse](#v2getresourceresponse) |
| default | An unexpected error response. | [googlerpcStatus](#googlerpcstatus) |

#### DELETE
##### Summary

DeleteInbox deletes an inbox.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| name_1 | path | The name of the inbox to delete. Format: inboxes/{uid} | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v2DeleteInboxResponse](#v2deleteinboxresponse) |
| default | An unexpected error response. | [googlerpcStatus](#googlerpcstatus) |

---
## LinkService

### /api/v2/link_metadata

#### GET
##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| link | query |  | No | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v2GetLinkMetadataResponse](#v2getlinkmetadataresponse) |
| default | An unexpected error response. | [googlerpcStatus](#googlerpcstatus) |

---
## LocketService

### /api/v2/lockets

#### GET
##### Summary

ListLockets lists lockets with pagination and filter.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| pageSize | query | The maximum number of lockets to return. | No | integer |
| pageToken | query | A page token, received from a previous `ListLockets` call. Provide this to retrieve the subsequent page. | No | string |
| filter | query | Filter is used to filter lockets returned in the list. Format: "creator == users/{uid} && visibilities == ['PUBLIC', 'PROTECTED']" | No | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v2ListLocketsResponse](#v2listlocketsresponse) |
| default | An unexpected error response. | [googlerpcStatus](#googlerpcstatus) |

#### POST
##### Summary

CreateLocket creates a locket.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| body | body |  | Yes | [v2CreateLocketRequest](#v2createlocketrequest) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v2CreateLocketResponse](#v2createlocketresponse) |
| default | An unexpected error response. | [googlerpcStatus](#googlerpcstatus) |

### /api/v2/lockets/stats

#### GET
##### Summary

GetUserLocketsStats gets stats of lockets for a user.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| name | query | name is the name of the user to get stats for. Format: users/{id} | No | string |
| timezone | query | timezone location Format: uses tz identifier https://en.wikipedia.org/wiki/List_of_tz_database_time_zones | No | string |
| filter | query | Same as ListLocketsRequest.filter | No | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v2GetUserLocketsStatsResponse](#v2getuserlocketsstatsresponse) |
| default | An unexpected error response. | [googlerpcStatus](#googlerpcstatus) |

### /api/v2/lockets:export

#### POST
##### Summary

ExportLockets exports lockets.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| filter | query | Same as ListLocketsRequest.filter | No | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v2ExportLocketsResponse](#v2exportlocketsresponse) |
| default | An unexpected error response. | [googlerpcStatus](#googlerpcstatus) |

### /api/v2/lockets:search

#### GET
##### Summary

SearchLockets searches lockets.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| filter | query | Filter is used to filter lockets returned. Format: "creator == users/{uid} && visibilities == ['PUBLIC', 'PROTECTED']" | No | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v2SearchLocketsResponse](#v2searchlocketsresponse) |
| default | An unexpected error response. | [googlerpcStatus](#googlerpcstatus) |

### /api/v2/{locket.name}

#### PATCH
##### Summary

UpdateLocket updates a locket.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| locket.name | path | The name of the locket. Format: lockets/{id} id is the system generated id. | Yes | string |
| locket | body |  | Yes | { **"uid"**: string, **"rowStatus"**: [apiv2RowStatus](#apiv2rowstatus), **"creator"**: string, **"createTime"**: dateTime, **"updateTime"**: dateTime, **"displayTime"**: dateTime, **"content"**: string, **"visibility"**: [v2Visibility](#v2visibility), **"pinned"**: boolean, **"parentId"**: integer, **"resources"**: [ [v2Resource](#v2resource) ], **"relations"**: [ [v2LocketRelation](#v2locketrelation) ], **"reactions"**: [ [apiv2Reaction](#apiv2reaction) ] } |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v2UpdateLocketResponse](#v2updatelocketresponse) |
| default | An unexpected error response. | [googlerpcStatus](#googlerpcstatus) |

### /api/v2/{name_2}

#### GET
##### Summary

GetLocket gets a locket.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| name_2 | path | The name of the locket. Format: lockets/{id} | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v2GetLocketResponse](#v2getlocketresponse) |
| default | An unexpected error response. | [googlerpcStatus](#googlerpcstatus) |

#### DELETE
##### Summary

DeleteResource deletes a resource by name.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| name_2 | path |  | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v2DeleteResourceResponse](#v2deleteresourceresponse) |
| default | An unexpected error response. | [googlerpcStatus](#googlerpcstatus) |

### /api/v2/{name_3}

#### DELETE
##### Summary

DeleteLocket deletes a locket.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| name_3 | path | The name of the locket. Format: lockets/{id} | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v2DeleteLocketResponse](#v2deletelocketresponse) |
| default | An unexpected error response. | [googlerpcStatus](#googlerpcstatus) |

### /api/v2/{name}/comments

#### GET
##### Summary

ListLocketComments lists comments for a locket.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| name | path | The name of the locket. Format: lockets/{id} | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v2ListLocketCommentsResponse](#v2listlocketcommentsresponse) |
| default | An unexpected error response. | [googlerpcStatus](#googlerpcstatus) |

#### POST
##### Summary

CreateLocketComment creates a comment for a locket.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| name | path | The name of the locket. Format: lockets/{id} | Yes | string |
| comment.content | query |  | No | string |
| comment.visibility | query |  | No | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v2CreateLocketCommentResponse](#v2createlocketcommentresponse) |
| default | An unexpected error response. | [googlerpcStatus](#googlerpcstatus) |

### /api/v2/{name}/reactions

#### GET
##### Summary

ListLocketReactions lists reactions for a locket.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| name | path | The name of the locket. Format: lockets/{id} | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v2ListLocketReactionsResponse](#v2listlocketreactionsresponse) |
| default | An unexpected error response. | [googlerpcStatus](#googlerpcstatus) |

#### POST
##### Summary

UpsertLocketReaction upserts a reaction for a locket.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| name | path | The name of the locket. Format: lockets/{id} | Yes | string |
| reaction.id | query |  | No | integer |
| reaction.creator | query | The name of the creator. Format: users/{id} | No | string |
| reaction.contentId | query |  | No | string |
| reaction.reactionType | query |  | No | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v2UpsertLocketReactionResponse](#v2upsertlocketreactionresponse) |
| default | An unexpected error response. | [googlerpcStatus](#googlerpcstatus) |

### /api/v2/{name}/reactions/{reactionId}

#### DELETE
##### Summary

DeleteLocketReaction deletes a reaction for a locket.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| name | path | The name of the locket. Format: lockets/{id} | Yes | string |
| reactionId | path |  | Yes | integer |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v2DeleteLocketReactionResponse](#v2deletelocketreactionresponse) |
| default | An unexpected error response. | [googlerpcStatus](#googlerpcstatus) |

### /api/v2/{name}/relations

#### GET
##### Summary

ListLocketRelations lists relations for a locket.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| name | path | The name of the locket. Format: lockets/{id} | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v2ListLocketRelationsResponse](#v2listlocketrelationsresponse) |
| default | An unexpected error response. | [googlerpcStatus](#googlerpcstatus) |

#### POST
##### Summary

SetLocketRelations sets relations for a locket.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| name | path | The name of the locket. Format: lockets/{id} | Yes | string |
| body | body |  | Yes | [LocketServiceSetLocketRelationsBody](#locketservicesetlocketrelationsbody) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v2SetLocketRelationsResponse](#v2setlocketrelationsresponse) |
| default | An unexpected error response. | [googlerpcStatus](#googlerpcstatus) |

### /api/v2/{name}/resources

#### GET
##### Summary

ListLocketResources lists resources for a locket.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| name | path | The name of the locket. Format: lockets/{id} | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v2ListLocketResourcesResponse](#v2listlocketresourcesresponse) |
| default | An unexpected error response. | [googlerpcStatus](#googlerpcstatus) |

#### POST
##### Summary

SetLocketResources sets resources for a locket.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| name | path | The name of the locket. Format: lockets/{id} | Yes | string |
| body | body |  | Yes | [LocketServiceSetLocketResourcesBody](#locketservicesetlocketresourcesbody) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v2SetLocketResourcesResponse](#v2setlocketresourcesresponse) |
| default | An unexpected error response. | [googlerpcStatus](#googlerpcstatus) |

---
## ResourceService

### /api/v2/resources

#### GET
##### Summary

ListResources lists all resources.

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v2ListResourcesResponse](#v2listresourcesresponse) |
| default | An unexpected error response. | [googlerpcStatus](#googlerpcstatus) |

#### POST
##### Summary

CreateResource creates a new resource.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| filename | query |  | No | string |
| externalLink | query |  | No | string |
| type | query |  | No | string |
| locketId | query |  | No | integer |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v2CreateResourceResponse](#v2createresourceresponse) |
| default | An unexpected error response. | [googlerpcStatus](#googlerpcstatus) |

### /api/v2/resources:search

#### GET
##### Summary

SearchResources searches lockets.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| filter | query |  | No | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v2SearchResourcesResponse](#v2searchresourcesresponse) |
| default | An unexpected error response. | [googlerpcStatus](#googlerpcstatus) |

### /api/v2/{name_1}

#### GET
##### Summary

GetResource returns a resource by name.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| name_1 | path |  | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v2GetResourceResponse](#v2getresourceresponse) |
| default | An unexpected error response. | [googlerpcStatus](#googlerpcstatus) |

#### DELETE
##### Summary

DeleteInbox deletes an inbox.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| name_1 | path | The name of the inbox to delete. Format: inboxes/{uid} | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v2DeleteInboxResponse](#v2deleteinboxresponse) |
| default | An unexpected error response. | [googlerpcStatus](#googlerpcstatus) |

### /api/v2/{name_2}

#### GET
##### Summary

GetLocket gets a locket.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| name_2 | path | The name of the locket. Format: lockets/{id} | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v2GetLocketResponse](#v2getlocketresponse) |
| default | An unexpected error response. | [googlerpcStatus](#googlerpcstatus) |

#### DELETE
##### Summary

DeleteResource deletes a resource by name.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| name_2 | path |  | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v2DeleteResourceResponse](#v2deleteresourceresponse) |
| default | An unexpected error response. | [googlerpcStatus](#googlerpcstatus) |

### /api/v2/{resource.name}

#### PATCH
##### Summary

UpdateResource updates a resource.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| resource.name | path | The name of the resource. Format: resources/{id} id is the system generated unique identifier. | Yes | string |
| resource | body |  | Yes | { **"uid"**: string, **"createTime"**: dateTime, **"filename"**: string, **"externalLink"**: string, **"type"**: string, **"size"**: string (int64), **"locketId"**: integer } |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v2UpdateResourceResponse](#v2updateresourceresponse) |
| default | An unexpected error response. | [googlerpcStatus](#googlerpcstatus) |

---
## TagService

### /api/v2/tags

#### GET
##### Summary

ListTags lists tags.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| user | query | The creator of tags. Format: users/{id} | No | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v2ListTagsResponse](#v2listtagsresponse) |
| default | An unexpected error response. | [googlerpcStatus](#googlerpcstatus) |

#### DELETE
##### Summary

DeleteTag deletes a tag.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| tag.name | query |  | No | string |
| tag.creator | query | The creator of tags. Format: users/{id} | No | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v2DeleteTagResponse](#v2deletetagresponse) |
| default | An unexpected error response. | [googlerpcStatus](#googlerpcstatus) |

#### POST
##### Summary

UpsertTag upserts a tag.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| name | query |  | No | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v2UpsertTagResponse](#v2upserttagresponse) |
| default | An unexpected error response. | [googlerpcStatus](#googlerpcstatus) |

### /api/v2/tags/suggestion

#### GET
##### Summary

GetTagSuggestions gets tag suggestions from the user's lockets.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| user | query | The creator of tags. Format: users/{id} | No | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v2GetTagSuggestionsResponse](#v2gettagsuggestionsresponse) |
| default | An unexpected error response. | [googlerpcStatus](#googlerpcstatus) |

### /api/v2/tags:batchUpsert

#### POST
##### Summary

BatchUpsertTag upserts multiple tags.

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v2BatchUpsertTagResponse](#v2batchupserttagresponse) |
| default | An unexpected error response. | [googlerpcStatus](#googlerpcstatus) |

### /api/v2/tags:rename

#### PATCH
##### Summary

RenameTag renames a tag.
All related lockets will be updated.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| user | query | The creator of tags. Format: users/{id} | No | string |
| oldName | query |  | No | string |
| newName | query |  | No | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v2RenameTagResponse](#v2renametagresponse) |
| default | An unexpected error response. | [googlerpcStatus](#googlerpcstatus) |

---
## UserService

### /api/v2/users

#### GET
##### Summary

ListUsers returns a list of users.

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v2ListUsersResponse](#v2listusersresponse) |
| default | An unexpected error response. | [googlerpcStatus](#googlerpcstatus) |

#### POST
##### Summary

CreateUser creates a new user.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| user | body |  | Yes | [v2User](#v2user) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v2CreateUserResponse](#v2createuserresponse) |
| default | An unexpected error response. | [googlerpcStatus](#googlerpcstatus) |

### /api/v2/users:search

#### GET
##### Summary

SearchUsers searches users by filter.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| filter | query |  | No | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v2SearchUsersResponse](#v2searchusersresponse) |
| default | An unexpected error response. | [googlerpcStatus](#googlerpcstatus) |

### /api/v2/{name}

#### GET
##### Summary

GetUser gets a user by name.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| name | path | The name of the user. Format: users/{id} | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v2GetUserResponse](#v2getuserresponse) |
| default | An unexpected error response. | [googlerpcStatus](#googlerpcstatus) |

#### DELETE
##### Summary

DeleteUser deletes a user.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| name | path | The name of the user. Format: users/{id} | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v2DeleteUserResponse](#v2deleteuserresponse) |
| default | An unexpected error response. | [googlerpcStatus](#googlerpcstatus) |

### /api/v2/{name}/access_tokens

#### GET
##### Summary

ListUserAccessTokens returns a list of access tokens for a user.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| name | path | The name of the user. Format: users/{id} | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v2ListUserAccessTokensResponse](#v2listuseraccesstokensresponse) |
| default | An unexpected error response. | [googlerpcStatus](#googlerpcstatus) |

#### POST
##### Summary

CreateUserAccessToken creates a new access token for a user.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| name | path | The name of the user. Format: users/{id} | Yes | string |
| body | body |  | Yes | [UserServiceCreateUserAccessTokenBody](#userservicecreateuseraccesstokenbody) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v2CreateUserAccessTokenResponse](#v2createuseraccesstokenresponse) |
| default | An unexpected error response. | [googlerpcStatus](#googlerpcstatus) |

### /api/v2/{name}/access_tokens/{accessToken}

#### DELETE
##### Summary

DeleteUserAccessToken deletes an access token for a user.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| name | path | The name of the user. Format: users/{id} | Yes | string |
| accessToken | path | access_token is the access token to delete. | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v2DeleteUserAccessTokenResponse](#v2deleteuseraccesstokenresponse) |
| default | An unexpected error response. | [googlerpcStatus](#googlerpcstatus) |

### /api/v2/{name}/setting

#### GET
##### Summary

GetUserSetting gets the setting of a user.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| name | path | The name of the user. Format: users/{id} | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v2GetUserSettingResponse](#v2getusersettingresponse) |
| default | An unexpected error response. | [googlerpcStatus](#googlerpcstatus) |

### /api/v2/{setting.name}

#### PATCH
##### Summary

UpdateUserSetting updates the setting of a user.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| setting.name | path | The name of the user. Format: users/{id} | Yes | string |
| setting | body |  | Yes | { **"locale"**: string, **"appearance"**: string, **"locketVisibility"**: string, **"telegramUserId"**: string } |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v2UpdateUserSettingResponse](#v2updateusersettingresponse) |
| default | An unexpected error response. | [googlerpcStatus](#googlerpcstatus) |

### /api/v2/{user.name}

#### PATCH
##### Summary

UpdateUser updates a user.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| user.name | path | The name of the user. Format: users/{id} | Yes | string |
| user | body |  | Yes | { **"id"**: integer, **"role"**: [UserRole](#userrole), **"username"**: string, **"email"**: string, **"nickname"**: string, **"avatarUrl"**: string, **"description"**: string, **"password"**: string, **"rowStatus"**: [apiv2RowStatus](#apiv2rowstatus), **"createTime"**: dateTime, **"updateTime"**: dateTime } |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v2UpdateUserResponse](#v2updateuserresponse) |
| default | An unexpected error response. | [googlerpcStatus](#googlerpcstatus) |

---
## WebhookService

### /api/v2/webhooks

#### GET
##### Summary

ListWebhooks returns a list of webhooks.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| creatorId | query |  | No | integer |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v2ListWebhooksResponse](#v2listwebhooksresponse) |
| default | An unexpected error response. | [googlerpcStatus](#googlerpcstatus) |

#### POST
##### Summary

CreateWebhook creates a new webhook.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| body | body |  | Yes | [v2CreateWebhookRequest](#v2createwebhookrequest) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v2CreateWebhookResponse](#v2createwebhookresponse) |
| default | An unexpected error response. | [googlerpcStatus](#googlerpcstatus) |

### /api/v2/webhooks/{id}

#### GET
##### Summary

GetWebhook returns a webhook by id.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| id | path |  | Yes | integer |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v2GetWebhookResponse](#v2getwebhookresponse) |
| default | An unexpected error response. | [googlerpcStatus](#googlerpcstatus) |

#### DELETE
##### Summary

DeleteWebhook deletes a webhook by id.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| id | path |  | Yes | integer |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v2DeleteWebhookResponse](#v2deletewebhookresponse) |
| default | An unexpected error response. | [googlerpcStatus](#googlerpcstatus) |

### /api/v2/webhooks/{webhook.id}

#### PATCH
##### Summary

UpdateWebhook updates a webhook.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| webhook.id | path |  | Yes | integer |
| webhook | body |  | Yes | { **"creatorId"**: integer, **"createdTime"**: dateTime, **"updatedTime"**: dateTime, **"rowStatus"**: [apiv2RowStatus](#apiv2rowstatus), **"name"**: string, **"url"**: string } |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v2UpdateWebhookResponse](#v2updatewebhookresponse) |
| default | An unexpected error response. | [googlerpcStatus](#googlerpcstatus) |

---
## WorkspaceService

### /api/v2/workspace/profile

#### GET
##### Summary

GetWorkspaceProfile returns the workspace profile.

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v2GetWorkspaceProfileResponse](#v2getworkspaceprofileresponse) |
| default | An unexpected error response. | [googlerpcStatus](#googlerpcstatus) |

---
## WorkspaceSettingService

### /api/v2/workspace/{name}

#### GET
##### Summary

GetWorkspaceSetting returns the setting by name.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| name | path | The resource name of the workspace setting. Format: settings/{setting} | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v2GetWorkspaceSettingResponse](#v2getworkspacesettingresponse) |
| default | An unexpected error response. | [googlerpcStatus](#googlerpcstatus) |

### /api/v2/workspace/{setting.name}

#### PATCH
##### Summary

SetWorkspaceSetting updates the setting.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| setting.name | path | name is the name of the setting. Format: settings/{setting} | Yes | string |
| setting | body | setting is the setting to update. | Yes | { **"generalSetting"**: [apiv2WorkspaceGeneralSetting](#apiv2workspacegeneralsetting) } |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v2SetWorkspaceSettingResponse](#v2setworkspacesettingresponse) |
| default | An unexpected error response. | [googlerpcStatus](#googlerpcstatus) |

---
## ActivityService

### /v2/activities/{id}

#### GET
##### Summary

GetActivity returns the activity with the given id.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| id | path |  | Yes | integer |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v2GetActivityResponse](#v2getactivityresponse) |
| default | An unexpected error response. | [googlerpcStatus](#googlerpcstatus) |

---
### Models

#### LocketServiceSetLocketRelationsBody

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| relations | [ [v2LocketRelation](#v2locketrelation) ] |  | No |

#### LocketServiceSetLocketResourcesBody

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| resources | [ [v2Resource](#v2resource) ] |  | No |

#### UserRole

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| UserRole | string |  |  |

#### UserServiceCreateUserAccessTokenBody

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| description | string |  | No |
| expiresAt | dateTime |  | No |

#### apiv2ActivityLocketCommentPayload

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| locketId | integer |  | No |
| relatedLocketId | integer |  | No |

#### apiv2ActivityPayload

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| locketComment | [apiv2ActivityLocketCommentPayload](#apiv2activitylocketcommentpayload) |  | No |
| versionUpdate | [apiv2ActivityVersionUpdatePayload](#apiv2activityversionupdatepayload) |  | No |

#### apiv2ActivityVersionUpdatePayload

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| version | string |  | No |

#### apiv2Reaction

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | integer |  | No |
| creator | string |  | No |
| contentId | string |  | No |
| reactionType | [apiv2ReactionType](#apiv2reactiontype) |  | No |

#### apiv2ReactionType

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| apiv2ReactionType | string |  |  |

#### apiv2RowStatus

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| apiv2RowStatus | string |  |  |

#### apiv2UserSetting

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| name | string |  | No |
| locale | string | The preferred locale of the user. | No |
| appearance | string | The preferred appearance of the user. | No |
| locketVisibility | string | The default visibility of the locket. | No |
| telegramUserId | string | The telegram user id of the user. | No |

#### apiv2Webhook

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | integer |  | No |
| creatorId | integer |  | No |
| createdTime | dateTime |  | No |
| updatedTime | dateTime |  | No |
| rowStatus | [apiv2RowStatus](#apiv2rowstatus) |  | No |
| name | string |  | No |
| url | string |  | No |

#### apiv2WorkspaceGeneralSetting

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| instanceUrl | string | instance_url is the instance URL. | No |
| disallowSignup | boolean | disallow_signup is the flag to disallow signup. | No |
| disallowPasswordLogin | boolean | disallow_password_login is the flag to disallow password login. | No |
| additionalScript | string | additional_script is the additional script. | No |
| additionalStyle | string | additional_style is the additional style. | No |

#### apiv2WorkspaceSetting

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| name | string |  | No |
| generalSetting | [apiv2WorkspaceGeneralSetting](#apiv2workspacegeneralsetting) | general_setting is the general setting of workspace. | No |

#### googlerpcStatus

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| code | integer |  | No |
| message | string |  | No |
| details | [ [protobufAny](#protobufany) ] |  | No |

#### protobufAny

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| @type | string |  | No |

#### v2Activity

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | integer |  | No |
| creatorId | integer |  | No |
| type | string |  | No |
| level | string |  | No |
| createTime | dateTime |  | No |
| payload | [apiv2ActivityPayload](#apiv2activitypayload) |  | No |

#### v2BatchUpsertTagResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| v2BatchUpsertTagResponse | object |  |  |

#### v2CreateLocketCommentResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| locket | [v2Locket](#v2locket) |  | No |

#### v2CreateLocketRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| content | string |  | No |
| visibility | [v2Visibility](#v2visibility) |  | No |

#### v2CreateLocketResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| locket | [v2Locket](#v2locket) |  | No |

#### v2CreateResourceResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| resource | [v2Resource](#v2resource) |  | No |

#### v2CreateUserAccessTokenResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| accessToken | [v2UserAccessToken](#v2useraccesstoken) |  | No |

#### v2CreateUserResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| user | [v2User](#v2user) |  | No |

#### v2CreateWebhookRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| name | string |  | No |
| url | string |  | No |

#### v2CreateWebhookResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| webhook | [apiv2Webhook](#apiv2webhook) |  | No |

#### v2DeleteInboxResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| v2DeleteInboxResponse | object |  |  |

#### v2DeleteLocketReactionResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| v2DeleteLocketReactionResponse | object |  |  |

#### v2DeleteLocketResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| v2DeleteLocketResponse | object |  |  |

#### v2DeleteResourceResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| v2DeleteResourceResponse | object |  |  |

#### v2DeleteTagResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| v2DeleteTagResponse | object |  |  |

#### v2DeleteUserAccessTokenResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| v2DeleteUserAccessTokenResponse | object |  |  |

#### v2DeleteUserResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| v2DeleteUserResponse | object |  |  |

#### v2DeleteWebhookResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| v2DeleteWebhookResponse | object |  |  |

#### v2ExportLocketsResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| content | byte |  | No |

#### v2GetActivityResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| activity | [v2Activity](#v2activity) |  | No |

#### v2GetAuthStatusResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| user | [v2User](#v2user) |  | No |

#### v2GetLinkMetadataResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| linkMetadata | [v2LinkMetadata](#v2linkmetadata) |  | No |

#### v2GetLocketResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| locket | [v2Locket](#v2locket) |  | No |

#### v2GetResourceResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| resource | [v2Resource](#v2resource) |  | No |

#### v2GetTagSuggestionsResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| tags | [ string ] |  | No |

#### v2GetUserLocketsStatsResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| stats | object | stats is the stats of locket creating/updating activities. key is the year-month-day string. e.g. "2020-01-01". | No |

#### v2GetUserResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| user | [v2User](#v2user) |  | No |

#### v2GetUserSettingResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| setting | [apiv2UserSetting](#apiv2usersetting) |  | No |

#### v2GetWebhookResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| webhook | [apiv2Webhook](#apiv2webhook) |  | No |

#### v2GetWorkspaceProfileResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| workspaceProfile | [v2WorkspaceProfile](#v2workspaceprofile) |  | No |

#### v2GetWorkspaceSettingResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| setting | [apiv2WorkspaceSetting](#apiv2workspacesetting) |  | No |

#### v2Inbox

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| name | string |  | No |
| sender | string |  | No |
| receiver | string |  | No |
| status | [v2InboxStatus](#v2inboxstatus) |  | No |
| createTime | dateTime |  | No |
| type | [v2InboxType](#v2inboxtype) |  | No |
| activityId | integer |  | No |

#### v2InboxStatus

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| v2InboxStatus | string |  |  |

#### v2InboxType

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| v2InboxType | string |  |  |

#### v2LinkMetadata

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| title | string |  | No |
| description | string |  | No |
| image | string |  | No |

#### v2ListInboxesResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| inboxes | [ [v2Inbox](#v2inbox) ] |  | No |

#### v2ListLocketCommentsResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| lockets | [ [v2Locket](#v2locket) ] |  | No |

#### v2ListLocketReactionsResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| reactions | [ [apiv2Reaction](#apiv2reaction) ] |  | No |

#### v2ListLocketRelationsResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| relations | [ [v2LocketRelation](#v2locketrelation) ] |  | No |

#### v2ListLocketResourcesResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| resources | [ [v2Resource](#v2resource) ] |  | No |

#### v2ListLocketsResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| lockets | [ [v2Locket](#v2locket) ] |  | No |
| nextPageToken | string | A token, which can be sent as `page_token` to retrieve the next page. If this field is omitted, there are no subsequent pages. | No |

#### v2ListResourcesResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| resources | [ [v2Resource](#v2resource) ] |  | No |

#### v2ListTagsResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| tags | [ [v2Tag](#v2tag) ] |  | No |

#### v2ListUserAccessTokensResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| accessTokens | [ [v2UserAccessToken](#v2useraccesstoken) ] |  | No |

#### v2ListUsersResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| users | [ [v2User](#v2user) ] |  | No |

#### v2ListWebhooksResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| webhooks | [ [apiv2Webhook](#apiv2webhook) ] |  | No |

#### v2Locket

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| name | string | The name of the locket. Format: lockets/{id} id is the system generated id. | No |
| uid | string | The user defined id of the locket. | No |
| rowStatus | [apiv2RowStatus](#apiv2rowstatus) |  | No |
| creator | string |  | No |
| createTime | dateTime |  | No |
| updateTime | dateTime |  | No |
| displayTime | dateTime |  | No |
| content | string |  | No |
| visibility | [v2Visibility](#v2visibility) |  | No |
| pinned | boolean |  | No |
| parentId | integer |  | No |
| resources | [ [v2Resource](#v2resource) ] |  | No |
| relations | [ [v2LocketRelation](#v2locketrelation) ] |  | No |
| reactions | [ [apiv2Reaction](#apiv2reaction) ] |  | No |

#### v2LocketRelation

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| locket | string |  | No |
| relatedLocket | string |  | No |
| type | [v2LocketRelationType](#v2locketrelationtype) |  | No |

#### v2LocketRelationType

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| v2LocketRelationType | string |  |  |

#### v2RenameTagResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| tag | [v2Tag](#v2tag) |  | No |

#### v2Resource

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| name | string | The name of the resource. Format: resources/{id} id is the system generated unique identifier. | No |
| uid | string | The user defined id of the resource. | No |
| createTime | dateTime |  | No |
| filename | string |  | No |
| externalLink | string |  | No |
| type | string |  | No |
| size | string (int64) |  | No |
| locketId | integer |  | No |

#### v2SearchLocketsResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| lockets | [ [v2Locket](#v2locket) ] |  | No |

#### v2SearchResourcesResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| resources | [ [v2Resource](#v2resource) ] |  | No |

#### v2SearchUsersResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| users | [ [v2User](#v2user) ] |  | No |

#### v2SetLocketRelationsResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| v2SetLocketRelationsResponse | object |  |  |

#### v2SetLocketResourcesResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| v2SetLocketResourcesResponse | object |  |  |

#### v2SetWorkspaceSettingResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| setting | [apiv2WorkspaceSetting](#apiv2workspacesetting) |  | No |

#### v2SignInResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| user | [v2User](#v2user) |  | No |

#### v2SignInWithSSOResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| user | [v2User](#v2user) |  | No |

#### v2SignOutResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| v2SignOutResponse | object |  |  |

#### v2SignUpResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| user | [v2User](#v2user) |  | No |

#### v2Tag

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| name | string |  | No |
| creator | string |  | No |

#### v2UpdateInboxResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| inbox | [v2Inbox](#v2inbox) |  | No |

#### v2UpdateLocketResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| locket | [v2Locket](#v2locket) |  | No |

#### v2UpdateResourceResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| resource | [v2Resource](#v2resource) |  | No |

#### v2UpdateUserResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| user | [v2User](#v2user) |  | No |

#### v2UpdateUserSettingResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| setting | [apiv2UserSetting](#apiv2usersetting) |  | No |

#### v2UpdateWebhookResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| webhook | [apiv2Webhook](#apiv2webhook) |  | No |

#### v2UpsertLocketReactionResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| reaction | [apiv2Reaction](#apiv2reaction) |  | No |

#### v2UpsertTagRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| name | string |  | No |

#### v2UpsertTagResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| tag | [v2Tag](#v2tag) |  | No |

#### v2User

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| name | string |  | No |
| id | integer | The system generated uid of the user. | No |
| role | [UserRole](#userrole) |  | No |
| username | string |  | No |
| email | string |  | No |
| nickname | string |  | No |
| avatarUrl | string |  | No |
| description | string |  | No |
| password | string |  | No |
| rowStatus | [apiv2RowStatus](#apiv2rowstatus) |  | No |
| createTime | dateTime |  | No |
| updateTime | dateTime |  | No |

#### v2UserAccessToken

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| accessToken | string |  | No |
| description | string |  | No |
| issuedAt | dateTime |  | No |
| expiresAt | dateTime |  | No |

#### v2Visibility

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| v2Visibility | string |  |  |

#### v2WorkspaceProfile

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| owner | string |  | No |
| version | string |  | No |
| mode | string | mode is the instance mode (e.g. "prod", "dev" or "demo"). | No |
| disallowSignup | boolean | disallow_signup is whether the signup is disallowed. | No |
| disablePasswordLogin | boolean | disable_password_login is whether the password login is disabled. | No |
| additionalScript | string | additional_script is the additional script. | No |
| additionalStyle | string | additional_style is the additional style. | No |
