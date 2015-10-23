package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var rePhone, reEmail *regexp.Regexp

func init() {
	rePhone = regexp.MustCompile("^1\\d{10}$")
	reEmail = regexp.MustCompile("^[^@]+@[^\\.]+(\\.[^\\.]+)+$")
}

func MatchPhone(value string) bool {
	return rePhone.MatchString(value)
}

func MatchEmail(value string) bool {
	return reEmail.MatchString(value)
}

//生成32位md5字串
func GeneralMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func StrToInt(s string) (int, error) {
	id64, err := strconv.ParseInt(s, 10, 32)
	if nil != err {
		return 0, err
	}
	return int(id64), nil
}

/*
	根据父级结点的Id和Path生成子级结点的Path
*/
func GetChildPath(fatherId int, fatherPath string) string {
	return fmt.Sprintf("%s%d/", fatherPath, fatherId)
}

/*
	分割字符串为id切片
*/
func SplitToIdsComma(idsStr string, ignoreBlank bool) ([]int, error) {
	return SplitToIds(idsStr, ",", ignoreBlank)
}

func SplitToIdsSlice(idsStr string, ignoreBlank bool) ([]int, error) {
	return SplitToIds(idsStr, "/", ignoreBlank)
}

func SplitToIds(idsStr, sep string, ignoreBlank bool) (ids []int, err error) {
	for _, idStr := range strings.Split(idsStr, sep) {
		if idStr == "" {
			continue
		}
		if id_, err := strconv.Atoi(idStr); err != nil {
			if ignoreBlank {
				continue
			}
			return nil, err
		} else {
			ids = append(ids, id_)
		}
	}
	return ids, nil
}

func ImplodeStrSplitToString(strSplit []string, sep string) string {
	result := ""
	for _, s := range strSplit {
		result = result + s + sep
	}
	if "" != result {
		rs := []rune(result)
		resultLen := len(rs)
		sepLen := len([]rune(sep))
		return string(rs[0:(resultLen - sepLen)])
	}
	return ""

}
