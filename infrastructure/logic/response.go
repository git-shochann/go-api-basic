package logic

import (
	"ddd/domain/model"
	"encoding/json"
	"fmt"
	"net/http"
)

type ResponseLogic interface {
	SendResponse(w http.ResponseWriter, response []byte, code int) error
	SendErrorResponse(w http.ResponseWriter, errorMessage string, code int) error
	SendAuthResponse(w http.ResponseWriter, user *model.User, code int) error
}

type responseLogic struct {
	jwtLogic JwtLogic
}

func NewResponseLogic(jwtLogic JwtLogic) ResponseLogic {
	return &responseLogic{
		jwtLogic: jwtLogic,
	}
}

// ステータスコード200の場合のレスポンス
func (rl *responseLogic) SendResponse(w http.ResponseWriter, response []byte, code int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err := w.Write(response)
	if err != nil {
		return err
	}
	return nil
}

// ステータスコード200以外のレスポンスで使用
// message: err.Error() とする
func (rl *responseLogic) SendErrorResponse(w http.ResponseWriter, errorMessage string, code int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	response := map[string]string{
		"message": errorMessage,
	}
	// jsonに変換する
	responseBody, err := json.Marshal(response)
	if err != nil {
		return err
	}
	_, err = w.Write(responseBody)
	if err != nil {
		return err
	}
	return nil
}

// 新規登録とログイン時のレスポンスとしてJWTトークンとUser構造体を返却する
func (rl *responseLogic) SendAuthResponse(w http.ResponseWriter, user *model.User, code int) error {

	fmt.Println("SendAuthResponse!")

	// これはどう呼び出せばいい？
	jwtToken, err := rl.jwtLogic.CreateJWTToken(user)
	if err != nil {
		return err
	}

	// レスポンス
	response := model.UserAuthResponse{
		User:     *user, // デリファレンスする
		JwtToken: jwtToken,
	}

	// 構造体をjsonに変換
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return err
	}
	fmt.Printf("jsonResponse: %v\n", string(jsonResponse))

	if err := rl.SendResponse(w, jsonResponse, code); err != nil {
		return err
	}

	return nil

}
