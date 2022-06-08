// Code generated by goctl. DO NOT EDIT.
package types

type AddTestDataReq struct {
	Username string `json:"username"`
	Type     string `json:"type"`
}

type AddTestDataResp struct {
	Msg string `json:"msg"`
}

type Data struct {
	Username  string `json:"username"`
	StudentId string `json:"student_id"`
	Type      string `json:"type"`
	Time      string `json:"time"`
}

type GetTestDataListReq struct {
}

type GetTestDataListResp struct {
	List []Data `json:"data_list"`
}

type GetLatestDataReq struct {
}

type GetLatestDataResp struct {
	Type string `json:"type"`
}
