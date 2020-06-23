package hlib

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
)

// Try 指定した回数 関数を実行するための構造体
type Try struct {
	next bool
}

// ForcedTermination TryFuncの指定回数を無視して強制終了させる
func (try *Try) ForcedTermination() {
	try.next = false
}

// Func numTryで指定した回数fnを実行する。fnが成功したらその時点でfnを実行した回数とnilをリターンする
// Try.ForcedTermination()を呼び出すことで強制終了させることができる
func (try *Try) Func(numTry int, fn func(count int) error) (int, error) {
	var err error
	var c int
	// nextフィールドがfalseになっている可能性があるので必ず最初にtrueにしておく
	try.next = true

	for try.next {
		c++
		err = fn(c)
		if !try.next {
			break
		}
		// ループ回数がnumTryと一致した場合ブレークする
		if c == numTry {
			break
		}
		// fn()の結果がエラーなら次のループへ
		if err != nil {
			continue
		}

		// fnに成功した場合の処理。前のループでエラーが出てる可能性があるのでエラーをここで潰してブレークする
		err = nil
		break
	}

	return c, err
}

// TryFunc Try.Funcのショートカット。強制終了が必要なければこちらを使う
func TryFunc(numTry int, fn func(count int) error) (int, error) {
	try := &Try{}
	return try.Func(numTry, fn)
}

// JSONUnmarshalFromFile ファイルを読み込み引数vへunmarshalする
func JSONUnmarshalFromFile(path string, v interface{}) error {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, v)
}

// WriteFileJSONPretty jsonを整形してファイルに書き込む
func WriteFileJSONPretty(v interface{}, path string, perm os.FileMode) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	if err = json.Indent(&buf, b, "", "  "); err != nil {
		return err
	}

	return ioutil.WriteFile(path, buf.Bytes(), perm)
}
