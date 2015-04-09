package rand

import (
    "math/rand"
    "strings"
    "time"
)

var (
    Capital   string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    Lowercase string = "abcdefghijklmnopqrstuvwxyz"
    Number    string = "0123456789"
    SpeStr    string = "~!@#$%^&*()/-_+-=|{}[]<>?"
)

func Random(baseStr []string, req int) (res string) {
    baseLen := len(baseStr)
    if baseLen == 0 {
        return
    }
    random := rand.New(rand.NewSource(time.Now().UnixNano()))
    for flag := 0; flag < req; flag++ {
        ran := random.Intn(baseLen)
        res += baseStr[ran]
    }
    return
}

func StringRand(str string, req int) string {
    if req <= 0 {
        return ""
    }
    baseStr := strings.Split(str, "")
    return Random(baseStr, req)
}

func CapitalOnly(req int) string {
    if req <= 0 {
        return ""
    }
    baseStr := strings.Split(Capital, "")
    return Random(baseStr, req)
}

func LowercaseOnly(req int) string {
    if req <= 0 {
        return ""
    }
    baseStr := strings.Split(Lowercase, "")
    return Random(baseStr, req)
}

func NumberOnly(req int) string {
    if req <= 0 {
        return ""
    }
    baseStr := strings.Split(Number, "")
    return Random(baseStr, req)
}

func SpeStrOnly(req int) string {
    if req <= 0 {
        return ""
    }
    baseStr := strings.Split(SpeStr, "")
    return Random(baseStr, req)
}

func NoSpeStr(req int) string {
    if req <= 0 {
        return ""
    }
    baseStr := strings.Split(Capital+Number+Lowercase, "")
    return Random(baseStr, req)
}
