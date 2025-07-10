# oauth

------------------------------------------------------------------------------------------

#### Create/Generate Auth Token

<details>
 <summary><code>POST</code> <code><b>/</b></code> <code>api/oauth/access-token/v1/</code></summary>

##### URL

> ```javascript
>  https://oauth.bongo.chat/api/oauth/access-token/v1/
> ```

##### Parameters

> | name            |  type     | data type         | description                                                           |
> |-----------------|-----------|-------------------|-----------------------------------------------------------------------|
> | grant_type      |  required | string            | N/A  |
> | phone_number    |  required | string            | N/A  |
> | password        |  required | string            | N/A  |


##### Responses

> | http code     | content-type                      | response                                                            |
> |---------------|-----------------------------------|---------------------------------------------------------------------|
> | `201`         | `application/json`                | `Token created successfully`                                |
> | `401`         | `application/json`                | `{"message": "Wrong password","display_message": "Wrong password","status": 401,"error": "unauthorized","causes": null}`                            |
> | `405`         | `text/html;charset=utf-8`         | None                                                                |

##### Example cURL

> ```javascript
>  curl -X POST -H "Content-Type: application/json" --data @post.json https://oauth.bongochat.app/api/oauth/access-token/v1/
> ```

</details>

------------------------------------------------------------------------------------------

#### Verify Auth Token

<details>
 <summary><code>GET</code> <code><b>/</b></code> <code>api/oauth/1/verify-token/v1/</code></summary>

##### URL

> ```javascript
>  https://oauth.bongo.chat/api/oauth/<user_id>/verify-token/v1/
> ```


##### Responses

> | http code     | content-type                      | response                                                            |
> |---------------|-----------------------------------|---------------------------------------------------------------------|
> | `200`         | `application/json`                | `{"result": {"access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9","user_id": 5,"date_created": "2024-0115T10:16:24.084Z"},"status": 200}`         |
> | `401`         | `application/json`                | `{"message": "Wrong password","display_message": "Wrong password","status": 401,"error": "unauthorized","causes": null}`                            |
> | `405`         | `text/html;charset=utf-8`         | None                                                                |

##### Example cURL

> ```javascript
>  curl -X POST -H "Content-Type: application/json" --data @post.json https://oauth.bongo.chat/api/oauth/<user_id>/verify-token/v1/
> ```

</details>

------------------------------------------------------------------------------------------
