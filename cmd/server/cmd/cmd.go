package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go-project/pkg/config"
	"go-project/server/controller"
	"path"
	"runtime"
)

var cmdReady bool

var rootCmd = &cobra.Command{
	Use:   "server",
	Short: "Go Project Server",
	Run: func(cmd *cobra.Command, args []string) {
		cmdReady = true
	},
	Example: "go-project -c config.yaml",
}

var configFile string

const (
	logLevel   = 4
	serverBind = "0.0.0.0:8080"
)

func init() {
	rootCmd.Flags().StringVarP(&configFile, "config", "c", "", "config file")
	rootCmd.Flags().IntP("log-level", "l", logLevel, "log level")
	rootCmd.Flags().String("server-bind", serverBind, "server bind addr")
}

func initConfig() error {
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			fileName := path.Base(frame.File)
			return frame.Function, fileName
		},
	})

	if len(configFile) > 0 {
		viper.SetConfigFile(configFile)
		err := viper.ReadInConfig()
		if err != nil {
			return err
		}
	}
	_ = viper.BindPFlag("log_level", rootCmd.Flags().Lookup("log-level"))
	_ = viper.BindPFlag("server_bind", rootCmd.Flags().Lookup("server-bind"))

	err := viper.Unmarshal(&config.Conf)
	if err != nil {
		return err
	}
	logrus.SetLevel(logrus.Level(config.Conf.LogLevel))
	gin.SetMode(gin.ReleaseMode)
	if logrus.Level(config.Conf.LogLevel) >= logrus.DebugLevel {
		gin.SetMode(gin.DebugMode)
		logrus.SetReportCaller(true)
	}
	return nil
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Fatal(err)
	}
	if cmdReady {
		err := initConfig()
		if err != nil {
			logrus.Fatal(err)
		}

		logrus.Debugf("config %+v", config.Conf)

		initConfigPost()
	}
}

func initConfigPost() {
	controller.RunServer(":8080")
}
