// Code generated by goctl. DO NOT EDIT.
package types

import (
	"crypto/md5"
	"fmt"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Response struct {
	Code uint32 `json:"code"`
	Msg  string `json:"msg"`
}

type UserLoginReq struct {
	StaffCode string `json:"staffCode"`
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
	StaffCode string `json:"staffCode"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	BaseWage  string `json:"baseWage"`
	ElseFee   string `json:"elseFee"`
}

type CalwageReq struct {
	WorkTime float64 `json:"workTime"`
}

type CalwageResp struct {
	Response
	Id            string  `json:"id"`
	StaffCode     string  `json:"staffCode"`
	Name          string  `json:"name"`
	Year          int     `json:"year"`
	Month         int     `json:"month"`
	WageBeforeTax float64 `json:"wageBeforeTax"`
	Tax           float64 `json:"tax"`
	ActualWage    float64 `json:"actualWage"`
}

type LookupResp struct {
	Response
	CurrentCount string `json:"currentCount"`
	AllCount     string `json:"allCount"`
}

type QuerywageReq struct {
	Year  int `json:"year"`
	Month int `json:"month"`
}

type QuerywageResp struct {
	Response
	WageBeforeTax float64 `json:"wageBeforeTax"`
	Tax           float64 `json:"tax"`
	ActualWage    float64 `json:"actualWage"`
}

type WageExcelReq struct {
	Year  int `json:"year"`
	Month int `json:"month"`
}

type WageExcelResp struct {
	Response
	FileName    string `json:"fileName"`
	FileType    string `json:"fileType"`
	DownloadUrl string `json:"downloadUrl"`
}

type UploadInfoReq struct {
	Hash string `json:"hash,optional"`
	Name string `json:"name,optional"`
	Ext  string `json:"ext,optional"`
	Size int64  `json:"size,optional"`
	Path string `json:"path,optional"`
}

type UploadInfoResp struct {
	Response
	Identity string `json:"identity"`
	Ext      string `json:"ext"`
	Name     string `json:"name"`
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
func GetJwtToken(secretKey string, iat, seconds int64, payload string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["payload"] = payload
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
