package hlib

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	code := m.Run()

	err := os.Remove("./testhelper/test.json")
	if err != nil {
		fmt.Println("err" + err.Error())
	}

	os.Exit(code)
}

func TestTryFunc(t *testing.T) {
	successCount := 10
	// successCountと引数countが一致したときにnil（成功）を返す
	fn := func(count int) error {
		if successCount == count {
			return nil
		}
		return errors.New("")
	}

	try := Try{}
	var err error
	var c int

	// 指定回数に到達する前に関数が処理に成功した場合
	numTry := 15
	tryCount := 10
	c, err = try.Func(numTry, fn)
	assert.NoError(t, err)
	assert.True(t, c == tryCount)

	// 関数の処理が成功する前に指定回数に到達した場合
	numTry = 8
	tryCount = 8
	c, err = try.Func(numTry, fn)
	assert.Error(t, err)
	assert.True(t, c == tryCount)
}

// Try.Funcが強制終了するか
func TestForcedTermination(t *testing.T) {
	successCount := 10
	numTry := 10
	tryCount := 2
	try := Try{}
	e := errors.New("ForcedTermination")

	c, err := try.Func(numTry, func(count int) error {
		if count == 2 {
			try.ForcedTermination()
			return e
		}
		if successCount == count {
			return nil
		}
		return errors.New("")
	})
	assert.True(t, err.Error() == e.Error())
	assert.True(t, c == tryCount)
}

func TestTryFuncShortCut(t *testing.T) {
	successCount := 10
	// successCountと引数countが一致したときにnil（成功）を返す
	fn := func(count int) error {
		if successCount == count {
			return nil
		}
		return errors.New("")
	}

	var err error
	var c int

	// 指定回数に到達する前に関数が処理に成功した場合
	numTry := 15
	tryCount := 10
	c, err = TryFunc(numTry, fn)
	assert.NoError(t, err)
	assert.True(t, c == tryCount)

	// 関数の処理が成功する前に指定回数に到達した場合
	numTry = 8
	tryCount = 8
	c, err = TryFunc(numTry, fn)
	assert.Error(t, err)
	assert.True(t, c == tryCount)
}

var testSt = struct {
	Text string   `json:"text"`
	Num  int      `json:"num"`
	List []string `json:"list"`
}{
	Text: "hello!",
	Num:  100,
	List: []string{"go", "js", "rust"},
}

func TestWriteFileJSONPretty(t *testing.T) {
	err := WriteFileJSONPretty(testSt, "./testhelper/test.json", 0777)
	assert.NoError(t, err)
}

func TestJSONUnmarshalFromFile(t *testing.T) {
	st := struct {
		Text string   `json:"text"`
		Num  int      `json:"num"`
		List []string `json:"list"`
	}{}
	err := JSONUnmarshalFromFile("./testhelper/test.json", &st)
	assert.NoError(t, err)
	assert.Exactly(t, testSt, st)
}
