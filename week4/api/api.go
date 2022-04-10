package api

import (
	"fmt"
	"golang_weekend/week4/internal/biz"
	"golang_weekend/week4/internal/pkg/mysql"
	"net/http"
)

// func UnorderHandler(info *mysql.DBModel) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		restr := biz.UnOrderedString(info, 11111)
// 		fmt.Fprintln(w, restr)
// 	})
// }

func UnorderHandler(info *mysql.DBModel) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		restr := biz.UnOrderedString(info, 11111)
		fmt.Fprintln(w, restr)
	})
}
