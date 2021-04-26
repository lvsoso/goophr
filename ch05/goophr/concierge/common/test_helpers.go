package common


import "testing"


const checkMark = "\u2713"
const ballotX = "\u2717"
const prefix = "\t\t"

// 测试成功的时候
func OnTestSuccess(t *testing.T, msg string){
	t.Log(prefix+msg, checkMark)
}


// 测试失败的时候
func OnTestError(t *testing.T, msg string){
	t.Error(prefix+msg, ballotX)
}

// 测试不可预测错误导致失败
func OnTestUnexpectedError(t *testing.T, err error) {
        OnTestError(t, "Unexpected Error:\n"+err.Error())
}

