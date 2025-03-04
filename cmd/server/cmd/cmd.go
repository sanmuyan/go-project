package cmd

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go-project/pkg/config"
	"go-project/pkg/configpost"
	"path"
	"runtime"
)

var rootCtx context.Context

var rootCmd = &cobra.Command{
	Use:   "server",
	Short: "Go Project Server",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		err := initConfig(cmd, args)
		if err != nil {
			logrus.Fatalf("init config error: %v", err)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		configpost.PostInit(rootCtx)
	},
	Example: "go-project -c config.yaml",
}

var configFile string

const (
	logLevel          = 4
	serverBind        = "0.0.0.0:8080"
	defaultConfigFile = ""
)

func init() {
	// 初始化命令行参数
	rootCmd.Flags().StringVarP(&configFile, "config", "c", defaultConfigFile, "config file")
	rootCmd.Flags().IntP("log-level", "l", logLevel, "log level")
	rootCmd.Flags().String("server-bind", serverBind, "server bind addr")
}

func initConfig(cmd *cobra.Command, args []string) error {
	// 设置日志格式
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			fileName := path.Base(frame.File)
			return frame.Function, fileName
		},
	})

	viper.SetConfigName("config")
	// 配置文件和命令行参数都不指定时的默认配置
	// viper.SetDefault("conn_timeout", 10)

	// 读取配置文件
	if len(configFile) > 0 {
		viper.SetConfigFile(configFile)
		err := viper.ReadInConfig()
		if err != nil {
			return err
		}
	}

	// 绑定命令行参数到配置项
	// 配置项优先级：命令行参数 > 配置文件 > 默认命令行参数
	_ = viper.BindPFlag("log_level", cmd.Flags().Lookup("log-level"))
	_ = viper.BindPFlag("server_bind", cmd.Flags().Lookup("server-bind"))

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
	logrus.Debugf("config init completed: %+v", config.Conf)
	return nil
}

func Execute(ctx context.Context) {
	rootCtx = ctx
	if err := rootCmd.Execute(); err != nil {
		logrus.Fatalf("cmd execute error: %v", err)
	}
}
