package conf
import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"demo1/model"
	"strings"
	"demo1/util"
	"github.com/joho/godotenv"
)
var Dictinary *map[interface{}]interface{}
func LoadLocales(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		return err
	}

	Dictinary = &m

	return nil
}

func T(key string) string {
	dic := *Dictinary
	keys := strings.Split(key, ".")
	for index, path := range keys {
		// 如果到达了最后一层，寻找目标翻译
		if len(keys) == (index + 1) {
			for k, v := range dic {
				if k, ok := k.(string); ok {
					if k == path {
						if value, ok := v.(string); ok {
							return value
						}
					}
				}
			}
			return path
		}
		// 如果还有下一层，继续寻找
		for k, v := range dic {
			if ks, ok := k.(string); ok {
				if ks == path {
					if dic, ok = v.(map[interface{}]interface{}); ok == false {
						return path
					}
				}
			} else {
				return ""
			}
		}
	}

	return ""
}


func Init()  {
	godotenv.Load()
	util.BuildLogger(os.Getenv("LOG_LEVEL"))

	if err := LoadLocales("conf/locales/zh-cn.yaml"); err!=nil{
		util.Log().Panic("翻译文件加载失败", err)
	}
	model.Database(os.Getenv("MYSQL_DSN"))
}
