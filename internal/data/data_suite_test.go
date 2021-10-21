package data_test

import (
	"adolesce/internal/biz"
	"adolesce/internal/conf"
	"adolesce/internal/data"
	"adolesce/pkg/log"
	"fmt"
	"github.com/fatih/color"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	db       *data.Data
	infoRepo biz.CallbackInfoRepo
	logRepo  biz.CallbackLogRepo
)

func TestData(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Data Suite")
}

var _ = BeforeSuite(func() {
	var err error
	file := filepath.Join("../../configs", "config.yaml")
	db, err = NewMockData(file)
	Expect(err).NotTo(HaveOccurred())

	infoRepo = data.NewCallbackInfoRepo(db, nil)
	logRepo = data.NewCallbackLogRepo(db, nil)
})

var _ = AfterSuite(func() {
	Expect(db.Close()).NotTo(HaveOccurred())
})

func NewMockData(file string) (*data.Data, error) {
	var config conf.Configs
	// 读取yaml配置
	configFile, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("NewMockData->ReadFile:%v", err)
		return nil, err
	}
	if err := yaml.Unmarshal(configFile, &config); err != nil {
		fmt.Printf("NewMockData->yaml.Unmarshal:%v", err)
		return nil, err
	}
	config.Log.OutputPath = filepath.Join("../../storage/logs", "test.inf.log")
	config.Log.ErrorOutputPath = filepath.Join("../../storage/logs", "test.err.log")
	// initLogger
	logger, err := log.NewLogger(&config)
	if err != nil {
		fmt.Printf("NewMockData->NewLogger:%v %v\n", color.RedString("Error:"), err)
		return nil, err
	}
	d, _, err := data.NewData(&config, logger)
	return d, err
}
