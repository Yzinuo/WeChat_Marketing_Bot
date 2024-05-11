package conf
import(
	"fmt"
	"os"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type Conf struct{
	APP App `json:"app" yaml:"app"`
	Keys Keys `json:"keys" yaml :"keys"`
	MysqlConf MysqlConf `json:"mysql" yaml:"mysql"`
}

type MysqlConf struct{
	IP 		string `json:"ip" yaml:"ip"`
	Port 	string	`json:"port" yaml:"port"`
	Passwd	string	`json:"passwd" yaml:"passwd"`
}

type App struct{
	Env string `json:"env" yaml:"env"`
}

type Keys struct{
	Botname string `json:"bot_name" yaml:"bot_name"`
}

func GetConf(cfg string)(conf *Conf,err error){
	var(
		yamlFile = make([]byte,0)
	)

	filepath := fmt.Sprintf("%s",cfg)
	logrus.Infof("filepath: %s",filepath)
	yamlFile, err = os.ReadFile(filepath)
	if err != nil{
		err = errors.Wrapf(err,"ReadFile error")
		logrus.Errorf(err.Error())
		return conf, err
	}

	err = yaml.Unmarshal(yamlFile,&conf)
	if err != nil{
		err = errors.Wrapf(err,"yaml.Unmarshal error")
		logrus.Errorf(err.Error())
		return conf,err
	}

	return conf,nil
}
