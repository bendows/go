package diskfile

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	logger "github.com/bendows/go"
	"github.com/gomodule/redigo/redis"
)

func GetLabelsFromFirst5125Bytes(r io.Reader) (io.Reader, []string) {
	b := make([]byte, 512)
	b, err := ioutil.ReadAll(io.LimitReader(r, 512))
	if err != nil {
		logger.Logerror.Println(err)
		return r, []string{}
	}
	nr := io.MultiReader(bytes.NewReader(b), r)
	// nr := io.MultiReader(bytes.NewReader([]byte{'g', 'o', 'l', 'a', 'n', 'g'}), r)
	// b[1:4] == []byte{'o', 'l', 'a'}, sharing the same storage as b
	label := http.DetectContentType(b)
	s := make(map[string]struct{})
	s[label] = struct{}{}
	for _, v := range strings.Split(label, " ") {
		v = strings.TrimSuffix(v, ";")
		s[v] = struct{}{}
		for _, v2 := range strings.Split(v, " ") {
			v2 = strings.TrimSuffix(v2, ";")
			for _, part := range strings.Split(v2, "/") {
				s[part] = struct{}{}
			}
			for _, part := range strings.Split(v2, "=") {
				s[part] = struct{}{}
			}
		}
	}
	labels := []string{}
	for key := range s {
		labels = append(labels, key)
	}
	return nr, labels
}
