package ch9
/**
*	Go圣经 JSON数据转换为实体
*/

import (
	"net/http"
	"log"
)

// IssueURL 问题URL
const IssueURL = "https://api.github.com/v3/search/issues"
// Q4 JSON转换为实体
func Q4() {
	res, err := http.Get("")	
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()


}