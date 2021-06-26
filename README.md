# Sign In Project

一個簡單的第三方驗證整合專案，以golang實現，依照各大平台Oauth2驗證流程，取得使用者的會員ID、使用名稱、email等資料，並回傳固定格式。

## 專案架構

使用微服務架構，以traefik做proxy server，串聯所有容器：

* frontEnd: 登入按鈕，將使用者導向平台授權，並取得授權碼。

  * 所需環境變數：

    * SIGN_IN_FRONT_END_HOST: 該容器專屬host，以traefik轉導。
    * API_BASE_URL: 驗證授權碼並取得使用者資料的server端位置，可以是gateway容器的host。
    * LINE_CLIENT_ID： line client id
    * LINE_REDIECT_URL: line導回URL，應該會與frontEnd容器位置一樣。
    * FB_APP_ID: fb app id
    * FB_REDIECT_URL: fb導回URL，應該會與frontEnd容器位置一樣。
    * GOOGLE_CLIENT_ID: google client id
    * GOOGLE_REDIECT_URL: google導回URL，應該會與frontEnd容器位置一樣。
    <br><br/>

* gateway: 接收授權碼，並傳送給對應的容器驗證、取得使用者資料，提供http和grpc接口。

  * 所需環境變數：

    * CONNECT_MODE: 控制使用http或grpc接口，預設http。
    * SIGN_IN_GATEWAY_HOST: 該容器專屬host，以traefik轉導。
    <br><br/>

* line: line平台驗證。

  * 所需環境變數：

    * REDIRECT_URL: line導回URL，應該會與frontEnd容器位置一樣。
    * CLIENT_ID: line client id
    * CLIENT_SECRET: line client secret
    <br><br/>

  * 取得使用者資料：

    * accessToken: access token
    * accessTokenExpireIn: access token 過期時間
    * refreshToken: refresh token
    * userId: 會員ID
    * name： 名稱
    * picture： 顯示照片
    * email： 信箱
    * statusMessage： 狀態資訊
    <br><br/>

* fb: fb平台驗證。

  * 所需環境變數：

    * REDIRECT_URL: fb導回URL，應該會與frontEnd容器位置一樣。
    * CLIENT_ID(app id): fb client id
    * CLIENT_SECRET: fb client secret
    <br><br/>

  * 取得使用者資料：

    * accessToken: access token
    * accessTokenExpireIn: access token 過期時間
    * userId: 會員ID
    * name： 名稱
    * picture： 顯示照片
    * email： 信箱
    * birthday： 生日
    <br><br/>

* google: google平台驗證。

  * 所需環境變數：

    * REDIRECT_URL: google導回URL，應該會與frontEnd容器位置一樣。
    * CLIENT_ID: google client id
    * CLIENT_SECRET: google client secret
    <br><br/>

  * 取得使用者資料：

    * accessToken: access token
    * accessTokenExpireIn: access token 過期時間
    * refreshToken: refresh token
    * userId: 會員ID
    * name： 名稱
    * picture： 顯示照片
    * email： 信箱

## 容器架構

參考[bxcodec](https://github.com/bxcodec)的[go-clean-arch](https://github.com/bxcodec/go-clean-arch)