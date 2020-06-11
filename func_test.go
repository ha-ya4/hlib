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

func TestTryFuncCount(t *testing.T) {
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
	c, err = TryFuncCount(numTry, fn)
	assert.NoError(t, err)
	assert.True(t, c == tryCount)

	// 関数の処理が成功する前に指定回数に到達した場合
	numTry = 8
	tryCount = 8
	c, err = TryFuncCount(numTry, fn)
	assert.Error(t, err)
	assert.True(t, c == tryCount)
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

	// 指定回数に到達する前に関数が処理に成功した場合
	numTry := 15
	assert.NoError(t, TryFunc(numTry, fn))

	// 関数の処理が成功する前に指定回数に到達した場合
	numTry = 8
	assert.Error(t, TryFunc(numTry, fn))
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

func TestFileLoad(t *testing.T) {
	_, err := FileLoad("./testhelper/test.json")
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
