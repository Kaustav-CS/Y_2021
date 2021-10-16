func isMatch(s string, p string) bool {
    dp := make([][]int, len(s) + 1)
    for i := 0; i < len(dp); i++ {
        dp[i] = make([]int, len(p) + 1)
    }
    return helper(s, p, &dp)
}

func helper(s string, p string, dp *[][]int) bool {
    if (*dp)[len(s)][len(p)] == notMatched {
        return false
    } else if (*dp)[len(s)][len(p)] == matched {
        return true
    }
    
    if len(s) == 0 && len(p) == 0 {
        (*dp)[len(s)][len(p)] = matched
        return true
    }
    
    if len(p) == 0 {
        (*dp)[len(s)][len(p)] = notMatched
        return false
    }
    
    if len(s) == 0 {
        if p[0] == '*' {
            if helper(s, p[1:], dp) {
                (*dp)[len(s)][len(p)] = matched
                return true
            } else {
                (*dp)[len(s)][len(p)] = notMatched
                return false
            }
        } else {
            (*dp)[len(s)][len(p)] = notMatched
            return false
        }
    }
    
    if p[0] == '?' || p[0] == s[0] {
        if helper(s[1:], p[1:], dp) {
            (*dp)[len(s)][len(p)] = matched
            return true
        } else {
            (*dp)[len(s)][len(p)] = notMatched
            return false
        }
    }
    
    if p[0] == '*' {
        for i := 0; i <= len(s); i++ {
            if helper(s[i:], p[1:], dp) {
                (*dp)[len(s)][len(p)] = matched
                return true
            }
        }
    }
    (*dp)[len(s)][len(p)] = notMatched
    return false
}
const (
    untouched = iota
    notMatched
    matched
)
