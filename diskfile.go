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

func getLabelsFromFirst5125Bytes(r io.Reader) (io.Reader, []string) {
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

type myFile struct {
	passedName string
	absPath    string
	name       string
	size       int64
	mode       os.FileMode
	modTime    time.Time
	types      []string
}

var redisPool *redis.Pool

func main() {
	logger.LogOn = true
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		filename := scanner.Text()
		if len(filename) < 1 {
			logger.Logerror.Printf("filename empty [%s]\n", filename)
			continue
		}
		f, err := os.OpenFile(filename, os.O_RDONLY, 0666)
		if err != nil {
			logger.Logerror.Println(err)
			continue
		}
		logger.Loginfo.Printf("file [%s] opened", filename)
		_, labels := getLabelsFromFirst5125Bytes(f)
		// newReader, labels := getLabelsFromFirst5125Bytes(f)
		// b, err := ioutil.ReadAll(newReader)
		// k := fmt.Sprintf("%x", md5.Sum(b))
		// client := redisPool.Get()
		// if _, err := redis.String(client.Do("SET", "f:"+k, b)); err != nil {
		// 	continue
		// }

		// b, err := ioutil.ReadAll(newReader)
		// if err != nil {
		// 	f.Close()
		// 	logger.Logerror.Println(err)
		// 	continue
		// }
		// ioutil.WriteFile("/dev/shm/thanks-frog.ben", b, 0644)
		stat, err := f.Stat()
		if err != nil {
			f.Close()
			logger.Logerror.Println(err)
			continue
		}
		absPath, err := filepath.Abs(filename)
		aFile := &myFile{
			passedName: filename,
			name:       stat.Name(),
			absPath:    absPath,
			size:       stat.Size(),
			mode:       stat.Mode(),
			modTime:    stat.ModTime(),
			types:      labels,
		}
		f.Close()
		logger.Loginfo.Printf("name [%s] [%s] [%s] [%d] modtime [%v]\n",
			aFile.name, aFile.passedName, aFile.absPath, aFile.size, aFile.modTime)
		for _, v := range aFile.types {
			logger.Loginfo.Printf("label [%s]\n", v)
		}
	}
}
