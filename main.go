package main

import (
	_ "embed"
	"flag"
	"fmt"
	"github.com/leungyauming/api/app"
	"github.com/leungyauming/api/app/config"
	"github.com/leungyauming/api/common"
	"github.com/leungyauming/api/services/rest"
	"log"
	"os"
	"os/signal"
	"runtime"
	"runtime/pprof"
)

var (
	shouldInitDb  bool
	shouldInitCfg bool

	cpuProfileFilename string
	memProfileFilename string
)

func init() {
	flag.BoolVar(&shouldInitDb, "init-db", false, "database initialization trigger")
	flag.BoolVar(&shouldInitCfg, "only-cfg", false, "only creates config trigger")

	flag.StringVar(&cpuProfileFilename, "pprof-cpu", "", "cpu profile filename")
	flag.StringVar(&memProfileFilename, "pprof-mem", "", "mem profile filename")
}

func Main() int {
	flag.Parse()

	fmt.Println(app.VersionStatement())

	logger := common.NewLogger("main")
	timer := common.NewTimer()

	if cpuProfileFilename != "" {
		logger.Println("creating CPU profile file")
		timer.Start()
		cpuProfileFile, err := os.Create(cpuProfileFilename)
		if err != nil {
			logger.Printf("failed to create CPU profile file -> %v", err)
			return 1
		}
		defer func() {
			if err := cpuProfileFile.Close(); err != nil {
				logger.Printf("failed to close CPU profile file -> %v", err)
			}
		}()
		timer.Stop()
		logger.Printf("created CPU profile file (%d ms)", timer.Duration().Milliseconds())

		logger.Println("starting CPU profiling")
		timer.Start()
		err = pprof.StartCPUProfile(cpuProfileFile)
		if err != nil {
			logger.Printf("failed to start CPU profiling -> %v", err)
			return 1
		}
		defer pprof.StopCPUProfile()
		timer.Stop()
		logger.Printf("started CPU profiling (%d ms)", timer.Duration().Milliseconds())
	}

	logger.Print("initializing config file")
	timer.Start()
	err := initConfigFile()
	if err != nil {
		logger.Printf("failed to save default config -> %v", err)
		return 1
	}
	timer.Stop()
	logger.Printf("initialized config file (%d ms)", timer.Duration().Milliseconds())
	if shouldInitCfg {
		return 0
	}

	logger.Print("parsing config")
	timer.Start()
	cfg, err := config.ParseFile(configPath)
	if err != nil {
		logger.Printf("failed to load config -> %v", err)
		return 1
	}
	timer.Stop()
	logger.Printf("parsed config (%d ms)", timer.Duration().Milliseconds())

	logger.Print("initializing database connection pool")
	timer.Start()
	dbPool, err := initDbPool(cfg.App.DB)
	if err != nil {
		logger.Printf("failed to initialize database connection pool -> %v", err)
		return 1
	}
	timer.Stop()
	logger.Printf("initialized database connection pool (%d ms)", timer.Duration().Milliseconds())

	if shouldInitDb {
		logger.Println("initializing database")
		timer.Start()
		err := initDb(dbPool)
		if err != nil {
			logger.Printf("failed to initialize database -> %v", err)
			return 1
		}
		timer.Stop()
		logger.Printf("initialized database (%d ms)", timer.Duration().Milliseconds())
	}

	deps := &app.Deps{
		Database: dbPool,
		Config:   cfg,
	}

	app_ := app.New(deps)
	app_.RegisterService(rest.New(deps))

	logger.Println("starting services")
	timer.Start()
	err = app_.Start()
	if err != nil {
		logger.Printf("failed to start app -> %v", err)
		return 1
	}
	timer.Stop()
	logger.Printf("started all services (%d ms)", timer.Duration().Milliseconds())

	{
		shutdownCh := make(chan os.Signal)
		signal.Notify(shutdownCh, os.Interrupt)
		<-shutdownCh
	}

	logger.Println("shutting down services")
	timer.Start()
	errs := app_.Shutdown()
	timer.Stop()
	if len(errs) != 0 {
		for err := range errs {
			logger.Printf("failed to shutdown services -> %v", err)
		}
		return 1
	}
	logger.Printf("shut down all services (%d ms)", timer.Duration().Milliseconds())

	if memProfileFilename != "" {
		logger.Println("creating memory profile file")
		timer.Start()
		memProfileFile, err := os.Create(memProfileFilename)
		if err != nil {
			log.Printf("failed to create memory profile file -> %v", err)
			return 1
		}
		defer func() {
			if err := memProfileFile.Close(); err != nil {
				logger.Printf("failed to close memory profile file -> %v", err)
			}
		}()
		timer.Stop()
		logger.Printf("created memory profile file (%d ms)", timer.Duration().Milliseconds())

		runtime.GC()

		logger.Println("writing memory profile to file")
		timer.Start()
		err = pprof.WriteHeapProfile(memProfileFile)
		if err != nil {
			log.Printf("failed to write memory profile to file -> %v", err)
			return 1
		}
		timer.Stop()
		logger.Printf("wrote memory profile to file (%d ms)", timer.Duration().Milliseconds())
	}

	return 0
}

func main() {
	os.Exit(Main())
}
