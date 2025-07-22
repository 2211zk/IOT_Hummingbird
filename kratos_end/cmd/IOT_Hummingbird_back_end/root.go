package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"

	"IOT_Hummingbird_back_end/internal/conf"

	"path/filepath"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	reportBroker   string
	reportTopic    string
	reportPayload  string
	reportInterval int
)

var rootCmd = &cobra.Command{
	Use:   "hummingbird",
	Short: "Hummingbird 脚本中心",
	Long:  `Hummingbird 脚本中心，支持服务启动、数据迁移、批量任务等自定义命令。`,
}

func newZapLogger() *zap.Logger {
	logDir := "logs"
	_ = os.MkdirAll(logDir, 0755)
	logFile := filepath.Join(logDir, time.Now().Format("2006-01-02")+".log")
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    100, // MB
		MaxBackups: 30,
		MaxAge:     30, // days
		Compress:   true,
	})
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderCfg), w, zap.InfoLevel)
	return zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
}

func newZapWriter() zapcore.WriteSyncer {
	logDir := "logs"
	_ = os.MkdirAll(logDir, 0755)
	logFile := filepath.Join(logDir, time.Now().Format("2006-01-02")+".log")
	return zapcore.AddSync(&lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    100,
		MaxBackups: 30,
		MaxAge:     30,
		Compress:   true,
	})
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "启动服务",
	Run: func(cmd *cobra.Command, args []string) {
		writer := newZapWriter()
		logger := log.With(log.NewStdLogger(writer),
			"ts", log.DefaultTimestamp,
			"caller", log.DefaultCaller,
			"service.id", id,
			"service.name", Name,
			"service.version", Version,
			"trace.id", tracing.TraceID(),
			"span.id", tracing.SpanID(),
		)
		flag.Parse()
		c := config.New(
			config.WithSource(
				file.NewSource(flagconf),
			),
		)
		defer c.Close()

		if err := c.Load(); err != nil {
			panic(err)
		}

		var bc conf.Bootstrap
		if err := c.Scan(&bc); err != nil {
			panic(err)
		}

		app, cleanup, err := wireApp(bc.Server, bc.Data, logger)
		if err != nil {
			panic(err)
		}
		defer cleanup()

		// start and wait for stop signal
		if err := app.Run(); err != nil {
			panic(err)
		}
	},
}

var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "模拟数据上报（MQTT）",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("开始模拟数据上报...")
		opts := mqtt.NewClientOptions().AddBroker(reportBroker)
		client := mqtt.NewClient(opts)
		if token := client.Connect(); token.Wait() && token.Error() != nil {
			fmt.Println("MQTT 连接失败:", token.Error())
			return
		}
		fmt.Println("MQTT 连接成功，开始上报...")
		ticker := time.NewTicker(time.Duration(reportInterval) * time.Second)
		defer ticker.Stop()
		for i := 0; i < 10; i++ { // 默认上报10次，可根据需要调整
			<-ticker.C
			payload := fmt.Sprintf("%s_%d", reportPayload, i+1)
			token := client.Publish(reportTopic, 0, false, payload)
			token.Wait()
			fmt.Printf("已上报: topic=%s payload=%s\n", reportTopic, payload)
		}
		client.Disconnect(250)
		fmt.Println("模拟数据上报完成！")
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	rootCmd.AddCommand(reportCmd)
	reportCmd.Flags().StringVarP(&reportBroker, "broker", "b", "tcp://127.0.0.1:1883", "MQTT Broker 地址")
	reportCmd.Flags().StringVarP(&reportTopic, "topic", "t", "test/topic", "MQTT Topic")
	reportCmd.Flags().StringVarP(&reportPayload, "payload", "p", "hello_mqtt", "上报内容")
	reportCmd.Flags().IntVarP(&reportInterval, "interval", "i", 2, "上报间隔(秒)")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
