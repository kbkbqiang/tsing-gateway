package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strconv"

	"github.com/rs/zerolog/log"
	"golang.org/x/net/http2"
)

func main() {
	var err error

	// 加载配置文件
	if err = loadConfigFile(); err != nil {
		panic(err.Error())
	}

	// 设置logger
	if err = setLogger(); err != nil {
		panic(err.Error())
	}

	// 设置etcd client
	if err = setEtcdCli(); err != nil {
		log.Fatal().Caller().Msg(err.Error())
	}

	// 启动服务
	start()
}

func start() {
	var httpServer *http.Server
	var httpsServer *http.Server
	var err error

	// 创建一个代理引擎实例
	proxy := newProxy()

	// 启动HTTP服务
	if localConfig.Listener.HTTP.Port > 0 {
		httpServer = &http.Server{
			Addr:              localConfig.Listener.IP + ":" + strconv.Itoa(localConfig.Listener.HTTP.Port),
			Handler:           proxy,
			ReadTimeout:       localConfig.Listener.ReadTimeout,       // 读取超时
			WriteTimeout:      localConfig.Listener.WriteTimeout,      // 响应超时
			IdleTimeout:       localConfig.Listener.IdleTimeout,       // 连接空闲超时
			ReadHeaderTimeout: localConfig.Listener.ReadHeaderTimeout, // header读取超时
		}
		go func() {
			log.Info().Msg("启动 HTTP 代理服务 :8000")
			if err = httpServer.ListenAndServe(); err != nil {
				if err == http.ErrServerClosed {
					log.Info().Msg("HTTP 代理服务已关闭")
					return
				}
				log.Fatal().Caller().Msg(err.Error())
			}
		}()
	}

	// 启动HTTPS服务
	if localConfig.Listener.HTTPS.Port > 0 {
		httpsServer = &http.Server{
			Addr:              localConfig.Listener.IP + ":" + strconv.Itoa(localConfig.Listener.HTTPS.Port),
			Handler:           proxy,
			ReadTimeout:       localConfig.Listener.ReadTimeout,       // 读取超时
			WriteTimeout:      localConfig.Listener.WriteTimeout,      // 响应超时
			IdleTimeout:       localConfig.Listener.IdleTimeout,       // 连接空闲超时
			ReadHeaderTimeout: localConfig.Listener.ReadHeaderTimeout, // header读取超时
		}
		go func() {
			log.Info().Msg("启动 HTTPS 代理服务 :8443")
			if localConfig.Listener.HTTPS.HTTP2 {
				log.Info().Msg("启用 HTTP/2 支持在 HTTPS 代理服务")
				if err = http2.ConfigureServer(httpsServer, &http2.Server{}); err != nil {
					log.Fatal().Caller().Msg(err.Error())
				}
			}
			if err = httpsServer.ListenAndServeTLS("server.cert", "server.key"); err != nil {
				if err == http.ErrServerClosed {
					log.Info().Msg("HTTPS 代理服务已关闭")
					return
				}
				log.Fatal().Caller().Msg(err.Error())
			}
		}()
	}

	// 退出进程时等待
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// 设置退出等待超时
	ctx, cancel := context.WithTimeout(context.Background(), localConfig.Listener.QuitWaitTimeout)
	defer cancel()

	// 退出HTTP服务
	if localConfig.Listener.HTTP.Port > 0 {
		if err := httpServer.Shutdown(ctx); err != nil {
			log.Fatal().Caller().Msg(err.Error())
		}
	}
	// 退出HTTPS服务
	if localConfig.Listener.HTTPS.Port > 0 {
		if err := httpsServer.Shutdown(ctx); err != nil {
			log.Fatal().Caller().Msg(err.Error())
		}
	}
}
