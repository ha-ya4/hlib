package hlib

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
)

// TryFuncCount numTryで指定した回数fnを実行する。fnが成功したらその時点でfnを実行した回数とnilをリターンする
func TryFuncCount(numTry int, fn func(count int) error) (int, error) {
	var err error
	var c int

	for {
		c++
		err = fn(c)
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

// TryFunc TryFuncCountを呼び出し返り値のfnが実行された回数を潰してエラーのみ返す。fnが実行された回数が必要ない場合に使う
func TryFunc(numTry int, fn func(count int) error) error {
	_, err := TryFuncCount(numTry, fn)
	return err
}

// FileLoad 引数で指定したファイルを読み込んで[]byteで返す
func FileLoad(path string) (b []byte, err error) {
	file, err := os.Open(path)
	if err != nil {
		return []byte{}, err
	}
	defer file.Close()
	return ioutil.ReadAll(file)
}

// JSONUnmarshalFromFile ファイルを読み込み引数vへunmarshalする
func JSONUnmarshalFromFile(path string, v interface{}) error {
	b, err := FileLoad(path)
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
