package gin

import (
	"context"
	"gateway/infrastructure/util/logging"
	"gateway/launch/gin/router"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func RunGin() {
	gin.SetMode(gin.ReleaseMode)
	r := router.Register()
	srv := &http.Server{
		Addr:    ":8000",
		Handler: r,
	}
	logging.New().Info("项目启动成功", srv.Addr, viper.GetString("app.env"))

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logging.New().Info("Shutdown Server ...", srv.Addr, nil)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		if err == context.DeadlineExceeded {
			logging.New().Info("Server Shutdown: timeout of 3 seconds.", srv.Addr, nil)
		} else {
			logging.New().ErrorL("Server Shutdown Error:", srv.Addr, err)
		}
	}
	logging.New().Info("Server exited", srv.Addr, nil)
}
