package gosdk

import (
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strconv"
)

type server struct {
	token      *jwt.Token
	tokenExist bool
}

var serverInstance = &server{tokenExist: false}

var tokenData = map[string]interface{}{}

func GetServerInstance(header http.Header) (*server, *CommError) {
	token1 := GetBearerToken(header)
	if token1 != "" {
		serverInstance.token, _ = jwt.Parse(token1, func(token *jwt.Token) (i interface{}, e error) {
			return token, nil
		})
		if _, ok := serverInstance.token.Claims.(jwt.MapClaims); ok {
			serverInstance.tokenExist = true
		}else{
			return nil,&CommError{204,"token format claim error"}
		}
	}
	return serverInstance,nil
}

func (server *server) GetTokenData() (map[string]interface{},*CommError) {
	if server.token == nil {
		return nil,&CommError{204,"token is empty"}
	}

	tokenData = make(map[string]interface{})
	claims, err := server.token.Claims.(jwt.MapClaims)
	if err {
		for key, value := range claims {
			tokenData[key] = value
		}
	}else{
		return nil,&CommError{204,"token format claim error"}
	}

	return tokenData,nil
}

func (server *server) GetAppkey() string {
	if server.token != nil {
		appkey,err:=server.token.Claims.(jwt.MapClaims)[TO_APPKEY_KEY].(string)
		if err{
			return appkey
		}
	}
	return ""
}

func (server *server) GetChannel() string {
	if server.token != nil {
		channel,err:=server.token.Claims.(jwt.MapClaims)[TO_CHANNEL].(string)
		if err{
			return channel
		}
		channelFloat,err:=server.token.Claims.(jwt.MapClaims)[TO_CHANNEL].(float64)
		if err{
			channel=strconv.FormatFloat(channelFloat, 'f', 0, 64)
			return channel
		}
	}
	return ""
}

func (server *server) GetAccountId() string {
	if server.token != nil {
		accountId,err:=server.token.Claims.(jwt.MapClaims)[ACCOUNT_ID_KEY].(string)
		if err{
			return accountId
		}
	}
	return ""
}

func (server *server) GetSubOrgKey() string {
	if server.token != nil {
		subOrgKey,err:=server.token.Claims.(jwt.MapClaims)[SUB_ORG_KEY_KEY].(string)
		if err{
			return subOrgKey
		}
	}
	return ""
}

func (server *server) GetUserInfo() map[string]string {
	if server.token != nil {
		userInfo,err:= server.token.Claims.(jwt.MapClaims)[USER_INFO_KEY].(map[string]string)
		if err{
			return userInfo
		}
	}
	return nil
}

func (server *server) GetFromAppkey() string {
	if server.token != nil {
		fromAppkey,err:=server.token.Claims.(jwt.MapClaims)[FROM_APPKEY_KEY].(string)
		if err{
			return fromAppkey
		}
	}
	return ""
}
func (server *server) GetFromChannel() string {
	if server.token != nil {
		fromChannel,err:=server.token.Claims.(jwt.MapClaims)[TO_CHANNEL].(string)
		if err{
			return fromChannel
		}
		channelFloat,err:=server.token.Claims.(jwt.MapClaims)[TO_CHANNEL].(float64)
		if err{
			fromChannel=strconv.FormatFloat(channelFloat, 'f', 0, 64)
			return fromChannel
		}
	}
	return ""
}
func (server *server) GetFromAppid() string {
	if server.token != nil{
		fromAppid,err:=server.token.Claims.(jwt.MapClaims)[FROM_APPID_KEY].(string)
		if err{
			return fromAppid
		}
	}
	return ""
}
func (server *server) GetCallStack() []map[string]string {
	if server.token != nil {
		callStack,err:=server.token.Claims.(jwt.MapClaims)[CALL_STACK_KEY].([]map[string]string)
		if err{
			return callStack
		}
	}
	return nil
}
