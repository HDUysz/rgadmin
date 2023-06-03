// Code generated by goctl. DO NOT EDIT.
package types

type Response struct {
	Code uint32 `json:"code"`
	Msg  string `json:"msg"`
}

type UserLoginReq struct {
	StuffCode string `json:"stuffCode"`
	Password  string `json:"password"`
}

type UserLoginResp struct {
	Response
	Role        string `json:"role"`
	AccessToken string `json:"accessToken"`
}

type UserLogoutReq struct {
}

type UserLogoutResp struct {
	Response
}

type UserInfoReq struct {
}

type UserInfoResp struct {
	Response
	Id        string `json:"id"`
	StuffCode string `json:"stuffCode"`
	Name      string `json:"name"`
	Role      string `json:"role"`
}

type LoadInfoReq struct {
	UserInfoTable []byte `json:"userInfoTable"`
}

type LoadInfoResp struct {
	Response
	OutputTable []byte `json:"outputTable"`
}

type CodeError struct {
	Code uint32
	Msg  string
}

// 返回给前端的错误码
func (e *CodeError) GetErrCode() uint32 {
	return e.Code
}

// 返回给前端显示端错误信息
func (e *CodeError) GetErrMsg() string {
	return e.Msg
}

func (e *CodeError) Error() string {
	return fmt.Sprintf("ErrCode:%d，ErrMsg:%s", e.Code, e.Msg)
}
func NewResponse(code uint32, msg string) Response {
	return Response{code, msg}
}
func NewErrCodeMsg(errCode uint32, errMsg string) *CodeError {
	return &CodeError{Code: errCode, Msg: errMsg}
}

const (
	EXPIRE        = 6 * time.Hour
	issuer string = "hduer"
)

type Claims struct {
	ID   string
	Name string
	jwtgo.StandardClaims
}

func NewClaim() Claims {
	return Claims{}
}
func Messagedigest5(s, salt string) string {
	if (s + salt) == "" {
		s = "123456"
	}
	data := md5.Sum([]byte(s + salt))
	return fmt.Sprintf("%x", data)
}
func (c *Claims) GeneraterJwt(id, name, salt string, duration time.Duration) (string, error) {
	now := time.Now().In(time.Local)
	c.ID = id
	c.Name = name
	c.Issuer = issuer + name
	c.ExpiresAt = now.Add(duration).Unix()
	tokenclaim := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, c)
	jwt, err := tokenclaim.SignedString([]byte(salt))
	return jwt, err
}
