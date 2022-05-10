package logger

import (
	"gin_zap/setting"
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strings"
)

var Logger *zap.Logger

// 日志初始化
func Init() error {
	// 生成定制化的zap日志对象，由3部分组成
	// 1，encoder 编码
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "time"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder   //配置时间格式
	encoderCfg.EncodeLevel = zapcore.CapitalLevelEncoder //全部大写
	encoder := zapcore.NewJSONEncoder(encoderCfg)        // 返回json数据
	//encoder := zapcore.NewConsoleEncoder(encoderCfg)	// 返回 字符串和空格拼接的数据

	// 2.1 往app.log里面写的core , 3.1 定义info级别
	//logFile1,_ := os.OpenFile("./app.logger", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	logPath := setting.Conf.LogConfig.Filename
	lumberJackLogger := &lumberjack.Logger{ // 切割日志的配置
		Filename:   logPath,
		MaxSize:    setting.Conf.LogConfig.MaxSize,
		MaxBackups: setting.Conf.LogConfig.MaxBackups,
		MaxAge:     setting.Conf.LogConfig.MaxAge, // 保存天数
		Compress:   false,
	}
	writeSyncer1 := zapcore.AddSync(lumberJackLogger)

	//consileWs1 := zapcore.AddSync(os.Stdout) // TODO 输出到终端，如果不输出终端注释掉

	// 获取全量日志级别
	level, err := zapcore.ParseLevel(setting.Conf.LogConfig.Level)
	if err != nil {
		level = zapcore.InfoLevel
	}
	core1 := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(writeSyncer1), level) // TODO 如果不输出终端删掉consileWs1

	// 2.2 往 app.err.log写日志的core , 3.2 定义error级别
	//logFile2, err := os.OpenFile("./app.err.logger", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	errLogPath := strings.Replace(setting.Conf.LogConfig.Filename, ".log", ".err.log", 1)
	lumberJackLogger2 := &lumberjack.Logger{
		Filename:   errLogPath,
		MaxSize:    setting.Conf.LogConfig.MaxSize,
		MaxBackups: setting.Conf.LogConfig.MaxBackups,
		MaxAge:     setting.Conf.LogConfig.MaxAge,
		Compress:   false,
	}

	writeSyncer2 := zapcore.AddSync(lumberJackLogger2)
	core2 := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(writeSyncer2), zapcore.ErrorLevel)

	// 当输出的日志，目的地不统，级别不统，需要新建core，zapcore.NewTee() 把多个core生成1个core
	newCore := zapcore.NewTee(core1, core2)

	// 通过合并的core 生成的新core的配置生成logger对象
	Logger = zap.New(newCore, zap.AddCaller())

	// 把Logger变成全局对象，之后其他包 import "go.uber.org/zap"之后，调用 zap.L() 就是Logger对象
	zap.ReplaceGlobals(Logger)

	fmt.Printf("\n 全量日志级别%v，全量日志路径%v，错误日志路径%v \n", level, logPath, errLogPath)
	return nil
}
