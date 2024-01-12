// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

package main

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"os/user"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/GuanceCloud/cliutils/logger"
	"github.com/kardianos/service"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/cmd/installer/installer"
	upgrader2 "gitlab.jiagouyun.com/cloudcare-tools/datakit/cmd/installer/upgrader"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/cmd/upgrader/upgrader"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/cmds"
	cp "gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/colorprint"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/config"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/datakit"
	dl "gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/downloader"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/git"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/httpcli"
	dkservice "gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/service"
)

var (
	oldInstallDir      = "/usr/local/cloudcare/dataflux/datakit"
	oldInstallDirWin   = `C:\Program Files\dataflux\datakit`
	oldInstallDirWin32 = `C:\Program Files (x86)\dataflux\datakit`

	DataKitBaseURL = ""
	DataKitVersion = ""
	dataURL        = "https://" + path.Join(DataKitBaseURL, "data.tar.gz")
	dkUpgraderURL  = "https://" + path.Join(DataKitBaseURL,
		fmt.Sprintf("%s-%s-%s.tar.gz", upgrader.BuildBinName, runtime.GOOS, runtime.GOARCH))
	datakitURL = "https://" + path.Join(DataKitBaseURL,
		fmt.Sprintf("datakit-%s-%s-%s.tar.gz",
			runtime.GOOS,
			runtime.GOARCH,
			DataKitVersion))
	datakitLiteURL = "https://" + path.Join(DataKitBaseURL,
		fmt.Sprintf("datakit_lite-%s-%s-%s.tar.gz",
			runtime.GOOS,
			runtime.GOARCH,
			DataKitVersion))
	InstallerBaseURL = ""
	l                = logger.DefaultSLogger("installer")

	isLite  = false
	liteReg = regexp.MustCompile(`Build Tag:\s*lite`)
)

// Installer flags.
var (
	flagDKUpgrade,
	flagOffline,
	flagDownloadOnly,
	flagInfo bool

	flagUpgradeServIPWhiteList,
	flagUserName,
	flagInstallLog,
	flagLite,
	flagSrc string

	flagUpgradeManagerService,
	flagInstallOnly int
)

const (
	datakitBin = "datakit"
)

//nolint:gochecknoinits,lll
func init() {
	flag.BoolVar(&flagDKUpgrade, "upgrade", false, "")
	flag.IntVar(&flagUpgradeManagerService, "upgrade-manager", 0, "whether we should upgrade the Datakit upgrade service")
	flag.StringVar(&flagUpgradeServIPWhiteList, "upgrade-ip-whitelist", "", "set datakit upgrade http service allowed request client ip, split by ','")
	flag.StringVar(&flagInstallLog, "install-log", "install.log", "install log")
	flag.StringVar(&flagSrc, "srcs", fmt.Sprintf("./datakit-%s-%s-%s.tar.gz,./data.tar.gz", runtime.GOOS, runtime.GOARCH, DataKitVersion), `local path of install files`)
	flag.IntVar(&flagInstallOnly, "install-only", 0, "install only, not start")
	flag.BoolVar(&flagInfo, "info", false, "show installer info")
	flag.BoolVar(&flagOffline, "offline", false, "-offline option removed")
	flag.BoolVar(&flagDownloadOnly, "download-only", false, "only download install packages")
	flag.StringVar(&InstallerBaseURL, "installer_base_url", "", "install datakit and data BaseUrl")
	flag.StringVar(&flagUserName, "user-name", "root", "install log") // user & group.
	flag.StringVar(&flagLite, "lite", "", "install datakit lite")

	flag.StringVar(&installer.Dataway, "dataway", "", "DataWay host(https://guance.openway.com?token=xxx)")
	flag.StringVar(&installer.Proxy, "proxy", "", "http proxy http://ip:port for datakit")
	flag.StringVar(&installer.DatakitName, "name", "", "specify DataKit name, example: prod-env-datakit")
	flag.StringVar(&installer.EnableInputs, "enable-inputs", "", "default enable inputs(comma splited, example:cpu,mem,disk)")
	flag.StringVar(&installer.HTTPPublicAPIs, "http-public-apis", "", "set which apis can be access by remote, split by comma.")
	flag.StringVar(&installer.HTTPDisabledAPIs, "http-disabled-apis", "", "(Deprecated) set which apis are disallowed access by remote, split by comma.")
	flag.IntVar(&installer.InstallRUMSymbolTools, "install-rum-symbol-tools", 0, "whether to install RUM source map tools")
	flag.BoolVar(&installer.OTA, "ota", false, "auto update")
	flag.StringVar(&installer.InstallExternals, "install-externals", "", "install some external inputs")

	// DCA flags
	flag.StringVar(&installer.DCAEnable, "dca-enable", "", "enable DCA")
	flag.StringVar(&installer.DCAListen, "dca-listen", "0.0.0.0:9531", "DCA listen address and port")
	flag.StringVar(&installer.DCAWhiteList, "dca-white-list", "", "DCA white list")

	// global-host-tags flags
	flag.StringVar(&installer.GlobalHostTags, "global-tags", "", "Deprecated, use global-host-tag")
	flag.StringVar(&installer.GlobalHostTags, "global-host-tags", "", "enable global host tags, example: host= __datakit_hostname,ip= __datakit_ip")

	// election flags
	flag.StringVar(&installer.GlobalElectionTags, "global-election-tags", "", "enable global environment tags, example: project=my-project,cluster=my-cluster")
	flag.StringVar(&installer.GlobalElectionTags, "global-env-tags", "", "Deprecated, use --global-election-tags")
	flag.StringVar(&installer.EnableElection, "enable-election", "", "datakit election")
	flag.StringVar(&installer.ElectionNamespace, "namespace", "", "datakit namespace")

	// datakit HTTP flags
	flag.IntVar(&installer.HTTPPort, "port", 9529, "datakit HTTP port")
	flag.StringVar(&installer.HTTPListen, "listen", "localhost", "datakit HTTP listen")
	flag.StringVar(&installer.HostName, "env_hostname", "", "host name")
	flag.StringVar(&installer.IPDBType, "ipdb-type", "", "ipdb type")
	flag.StringVar(&installer.CloudProvider, "cloud-provider", "", "specify cloud provider(accept aliyun/tencent/aws)")

	// confd flags
	flag.StringVar(&installer.ConfdBackend, "confd-backend", "", "backend kind")
	flag.StringVar(&installer.ConfdBasicAuth, "confd-basic-auth", "", "if backend need auth")
	flag.StringVar(&installer.ConfdClientCaKeys, "confd-client-ca-keys", "", "backend ca key")
	flag.StringVar(&installer.ConfdClientCert, "confd-client-cert", "", "backend cert key")
	flag.StringVar(&installer.ConfdClientKey, "confd-client-key", "", "backend cert key id")
	flag.StringVar(&installer.ConfdBackendNodes, "confd-backend-nodes", "", "backend nodes ip")
	flag.StringVar(&installer.ConfdPassword, "confd-password", "", "backend login password")
	flag.StringVar(&installer.ConfdScheme, "confd-scheme", "", "backend scheme")
	flag.StringVar(&installer.ConfdSeparator, "confd-separator", "", "backend separator")
	flag.StringVar(&installer.ConfdUsername, "confd-username", "", "backend login username")
	flag.StringVar(&installer.ConfdAccessKey, "confd-access-key", "", "backend access key id")
	flag.StringVar(&installer.ConfdSecretKey, "confd-secret-key", "", "backend secret key")
	flag.StringVar(&installer.ConfdConfdNamespace, "confd-confd-namespace", "", "confd config namespace id")
	flag.StringVar(&installer.ConfdPipelineNamespace, "confd-pipeline-namespace", "", "pipeline config namespace id")
	flag.StringVar(&installer.ConfdRegion, "confd-region", "", "aws region")
	flag.IntVar(&installer.ConfdCircleInterval, "confd-circle-interval", 60, "backend loop search interval second")

	// gitrepo flags
	flag.StringVar(&installer.GitURL, "git-url", "", "git repository url")
	flag.StringVar(&installer.GitKeyPath, "git-key-path", "", "git repository access private key path")
	flag.StringVar(&installer.GitKeyPW, "git-key-pw", "", "git repository access private use password")
	flag.StringVar(&installer.GitBranch, "git-branch", "", "git repository branch name")
	flag.StringVar(&installer.GitPullInterval, "git-pull-interval", "", "git repository pull interval")

	// rum flags
	flag.StringVar(&installer.RumOriginIPHeader, "rum-origin-ip-header", "", "rum only")
	flag.StringVar(&installer.RumDisable404Page, "disable-404page", "", "datakit rum 404 page")

	// log flags
	flag.StringVar(&installer.LogLevel, "log-level", "", "log level setting")
	flag.StringVar(&installer.Log, "log", "", "log setting")
	flag.StringVar(&installer.GinLog, "gin-log", "", "gin log setting")

	// pprof flags
	flag.StringVar(&installer.PProfListen, "pprof-listen", "", "pprof listen")

	// sinker flags
	flag.StringVar(&installer.EnableSinker, "enable-dataway-sinker", "", "enable dataway sinker")
	flag.StringVar(&installer.SinkerGlobalCustomerKeys, "sinker-global-customer-keys", "", "sinker configures")

	// resource limit flags
	flag.IntVar(&installer.LimitDisabled, "limit-disabled", 0, "enable disable resource limits for CPU and memory in linux and windows")
	flag.Float64Var(&installer.LimitCPUMax, "limit-cpumax", 30.0, "CPU max usage")
	flag.Float64Var(&installer.LimitCPUMin, "limit-cpumin", 5.0, "CPU min usage")
	flag.Int64Var(&installer.LimitMemMax, "limit-memmax", 4096, "memory limit")
}

func setDatakitLite() {
	if len(flagLite) > 0 {
		v, err := strconv.ParseBool(flagLite)
		if err != nil {
			l.Warnf("parse flag 'lite' error: %s", err.Error())
		} else {
			isLite = v
		}
	} else if flagDKUpgrade { // only for upgrading datakit
		cmd := exec.Command(dkservice.Executable, "version") //nolint:gosec
		res, err := cmd.CombinedOutput()
		if err != nil {
			l.Warnf("check version failed: %s", err.Error())
		} else {
			isLite = liteReg.Match(res)
		}
	}
}

func downloadFiles(to string) error {
	dl.CurDownloading = "datakit"

	cliopt := &httpcli.Options{
		// ignore SSL error
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //nolint
	}

	if installer.Proxy != "" {
		u, err := url.Parse(installer.Proxy)
		if err != nil {
			return err
		}
		cliopt.ProxyURL = u
		l.Infof("set proxy to %s ok", installer.Proxy)
	}

	cli := httpcli.Cli(cliopt)

	dkURL := datakitURL
	if isLite {
		dkURL = datakitLiteURL
	}

	if err := dl.Download(cli, dkURL, to, true, flagDownloadOnly); err != nil {
		return err
	}

	fmt.Printf("\n")

	dl.CurDownloading = "data"
	if err := dl.Download(cli, dataURL, to, true, flagDownloadOnly); err != nil {
		return err
	}

	// We will not upgrade dk-upgrader default when upgrading Datakit except for setting flagUpgradeManagerService flag
	if !flagDKUpgrade || (flagDKUpgrade && flagUpgradeManagerService == 1) || flagDownloadOnly {
		if !flagDownloadOnly {
			to = upgrader.InstallDir
		}
		dl.CurDownloading = upgrader.BuildBinName
		if err := dl.Download(cli, dkUpgraderURL, to, true, flagDownloadOnly); err != nil {
			l.Warnf("unable to download %s from [%s]: %s", upgrader.BuildBinName, dkUpgraderURL, err)
		}
	}

	if installer.IPDBType != "" {
		fmt.Printf("\n")
		baseURL := "https://" + DataKitBaseURL

		l.Debugf("get ipdb from %s", baseURL)
		if _, err := cmds.InstallIPDB(baseURL, installer.IPDBType); err != nil {
			l.Warnf("install IPDB %s failed error: %s, please try later.", installer.IPDBType, err.Error())
			time.Sleep(1 * time.Second)
		} else {
			config.Cfg.Pipeline.IPdbType = installer.IPDBType
		}
	}

	fmt.Printf("\n")
	return nil
}

func applyFlags() {
	var err error

	// setup logging
	if flagInstallLog == "stdout" {
		l.Infof("Set log file to stdout")

		if err = logger.InitRoot(
			&logger.Option{
				Level: logger.DEBUG,
				Flags: logger.OPT_DEFAULT | logger.OPT_STDOUT,
			}); err != nil {
			l.Errorf("Set root log faile: %s", err.Error())
		}
	} else {
		l.Infof("Set log file to %s", flagInstallLog)
		if err = logger.InitRoot(&logger.Option{
			Path:  flagInstallLog,
			Level: logger.DEBUG,
			Flags: logger.OPT_DEFAULT,
		}); err != nil {
			l.Errorf("Set root log faile: %s", err.Error())
		}
	}

	config.SetLog()
	installer.SetLog()
	l = logger.SLogger("installer")

	installer.DataKitVersion = DataKitVersion

	if flagDownloadOnly {
		if err = downloadFiles(""); err != nil { // download 过程直接覆盖已有安装
			l.Fatalf("download failed: %s", err.Error())
		}
		os.Exit(0)
	}

	if flagSrc != "" && flagOffline {
		for _, f := range strings.Split(flagSrc, ",") {
			fd, err := os.Open(filepath.Clean(f))
			if err != nil {
				l.Fatalf("Open: %s", err)
			}

			if err := dl.Extract(fd, datakit.InstallDir); err != nil {
				l.Fatalf("Extract: %s", err)
			} else if err := fd.Close(); err != nil {
				l.Warnf("Close: %s, ignored", err)
			}
		}

		// NOTE: continue to install/upgrade
	}

	if installer.Proxy != "" {
		if !strings.HasPrefix(installer.Proxy, "http") {
			installer.Proxy = "http://" + installer.Proxy
		}

		if _, err = url.Parse(installer.Proxy); err != nil {
			l.Warnf("bad proxy config expect http://ip:port given %s", installer.Proxy)
		} else {
			l.Infof("set proxy to %s", installer.Proxy)
		}
	}

	if InstallerBaseURL != "" {
		_, err := url.Parse(InstallerBaseURL)
		if err != nil {
			l.Errorf("ENV:$DK_INSTALLER_BASE_URL can not parse to URL, err=%v", err)
			os.Exit(0)
		}

		InstallerBaseURL = cmds.CanonicalInstallBaseURL(InstallerBaseURL)

		l.Infof("Set installer base URL to %s", InstallerBaseURL)
		dataURL = InstallerBaseURL + "data.tar.gz"

		datakitURL = InstallerBaseURL + fmt.Sprintf("datakit-%s-%s-%s.tar.gz",
			runtime.GOOS,
			runtime.GOARCH,
			DataKitVersion)

		dkUpgraderURL = InstallerBaseURL + fmt.Sprintf("%s-%s-%s.tar.gz", upgrader.BuildBinName, runtime.GOOS, runtime.GOARCH)
	}
}

func main() {
	flag.Parse()

	if flagInfo {
		fmt.Printf(`
Version        : %s
Build At       : %s
Golang Version : %s
BaseUrl        : %s
Data           : %s
`, DataKitVersion, git.BuildAt, git.Golang, datakitURL, dataURL)
		os.Exit(0)
	}

	var err error

	dkservice.Executable = filepath.Join(datakit.InstallDir, datakitBin)
	if runtime.GOOS == datakit.OSWindows {
		dkservice.Executable += ".exe"
	}

	// fix user name.
	var userName string
	if runtime.GOOS == datakit.OSLinux && len(flagUserName) > 0 && flagUserName != "root" {
		userName = flagUserName
		groupName := flagUserName

		if _, err := user.LookupGroup(groupName); err != nil {
			l.Errorf("Group %s not existed! Please create it first.", groupName)
			return
		}

		if _, err := user.Lookup(userName); err != nil {
			l.Errorf("User %s not existed! Please create it first.", userName)
			return
		}

		l.Infof("datakit service run as user: %q", userName)
	}

	svc, err := dkservice.NewService(userName)
	if err != nil {
		l.Errorf("new %s service failed: %s", runtime.GOOS, err.Error())
		return
	}

	svcStatus, err := svc.Status()
	if err != nil {
		if errors.Is(err, service.ErrNotInstalled) {
			l.Infof("datakit service not installed before")
		} else {
			l.Warnf("svc.Status: %s, ignored", err.Error())
		}
	} else {
		switch svcStatus {
		case service.StatusUnknown: // not installed
			l.Infof("DataKit service maybe not installed")
		case service.StatusStopped: // pass
			l.Infof("DataKit service stopped")
		case service.StatusRunning:
			l.Infof("Stopping running DataKit...")
			if err = service.Control(svc, "stop"); err != nil {
				l.Warnf("stop service failed %s, ignored", err.Error())
			}
		}
	}

	if !flagDKUpgrade || flagUpgradeManagerService == 1 {
		upgrader2.StopUpgradeService(userName)
	}

	applyFlags()
	setDatakitLite()

	// 迁移老版本 datakit 数据目录
	mvOldDatakit(svc)

	if !flagOffline {
		dlRetry := 5

		l.Infof("Download installer...\n")

		for i := 0; i < dlRetry; i++ {
			if err = downloadFiles(datakit.InstallDir); err != nil { // download 过程直接覆盖已有安装
				l.Warnf("[%d] download failed: %s, retry...", i, err.Error())
				continue
			}

			goto __downloadOK
		}

		l.Errorf("Download failed, please check your network settings.")
		return
	}

__downloadOK:
	datakit.InitDirs()

	upgrader2.InstallUpgradeService(userName,
		flagDKUpgrade,
		flagInstallOnly,
		flagUpgradeManagerService,
		flagUpgradeServIPWhiteList,
		InstallerBaseURL)

	if flagDKUpgrade { // upgrade new version
		l.Infof("Upgrading to version %s...", DataKitVersion)
		if err = installer.Upgrade(); err != nil {
			l.Warnf("upgrade datakit failed: %s, ignored", err.Error())
		}
	} else { // install new datakit
		l.Infof("Installing version %s...", DataKitVersion)
		installer.Install(svc, userName)
	}

	setupUserGroup(userName, userName)

	if flagInstallOnly != 0 {
		l.Warnf("Only install service %q, NOT started", dkservice.Name)
	} else {
		if err = service.Control(svc, "start"); err != nil {
			l.Warnf("Start service %q failed: %s", dkservice.Name, err.Error())
		} else {
			l.Infof("Starting service %q ok", dkservice.Name)
		}
	}

	if err := config.CreateSymlinks(); err != nil {
		l.Errorf("CreateSymlinks failed: %s", err.Error())
		l.Infof("Your may need to run datakit command under install path %q", datakit.InstallDir)
	} else {
		l.Infof("Create symlinks ok")
	}

	if flagInstallOnly == 0 {
		if err := checkIsNewVersion("http://"+config.Cfg.HTTPAPI.Listen, DataKitVersion); err != nil {
			l.Warnf("Check current datakit version failed(expect version %q), we can ignore the message and go on", DataKitVersion)
			l.Infof("Please visit https://docs.guance.com/datakit/datakit-update/#version-check-failed to see more info about the version checking.")
		} else {
			l.Infof("Current datakit version is %q", DataKitVersion)
		}
	} else {
		l.Infof("Current datakit version is %q(NOT running)", DataKitVersion)
	}

	if flagDKUpgrade {
		l.Infof("Upgrade OK.")
	} else {
		l.Infof("Install OK.")
	}
	promptReferences()
}

// test if installed/upgraded to expected version.
func checkIsNewVersion(host, version string) error {
	x := struct {
		Content map[string]string `json:"content"`
	}{}

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * time.Duration(i+1))

		resp, err := http.Get(host + "/v1/ping")
		if err != nil {
			l.Errorf("http.Get: %s", err)
			continue
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			l.Errorf("io.ReadAll: %s", err)
		}

		resp.Body.Close() //nolint:errcheck,gosec

		if err := json.Unmarshal(body, &x); err != nil {
			l.Errorf("json.Unmarshal: %s", err)
		}

		if x.Content["version"] != version {
			l.Warnf("current version: %s, expect %s", x.Content["version"], version)
		} else {
			return nil
		}
	}

	return fmt.Errorf("check version failed")
}

func promptReferences() {
	cp.Infof("\nVisit https://docs.guance.com/datakit/changelog/ to see DataKit change logs.\n")
	cp.Infof("Use `datakit monitor` to see DataKit running status.\n")
}

func mvOldDatakit(svc service.Service) {
	olddir := oldInstallDir
	switch runtime.GOOS + "/" + runtime.GOARCH {
	case datakit.OSArchWinAmd64:
		olddir = oldInstallDirWin
	case datakit.OSArchWin386:
		olddir = oldInstallDirWin32
	}

	if _, err := os.Stat(olddir); err != nil {
		l.Infof("deprecated install path %s not found\n", olddir)
		return
	}

	if err := service.Control(svc, "uninstall"); err != nil {
		l.Warnf("uninstall service failed: %s", err.Error())
	}

	if err := os.Rename(olddir, datakit.InstallDir); err != nil {
		l.Fatalf("move %s -> %s failed: %s", olddir, datakit.InstallDir, err.Error())
	}
}

func setupUserGroup(userName, groupName string) {
	l.Info("setupUserGroup entry")

	if len(userName) == 0 || userName == "root" || runtime.GOOS != datakit.OSLinux {
		return
	}

	l.Info("setupUserGroup start")

	l.Infof("Install dir: %s", datakit.InstallDir)

	logDir := filepath.Dir(config.Cfg.Logging.Log)
	l.Infof("logDir = %s", logDir)

	// make dirs.
	if err := mkdir(datakit.InstallDir, os.ModePerm); err != nil {
		l.Errorf("make installDir failed: %v", err)
	}

	if err := mkdir(logDir, os.ModePerm); err != nil {
		l.Errorf("make defaultLogDir failed: %v", err)
	}

	// chown.
	if err := executeCmd("chown", "-R", fmt.Sprintf("%s:%s", userName, groupName), datakit.InstallDir, logDir); err != nil {
		l.Errorf("chown failed: %v", err)
		return
	}

	// chmod.
	if err := executeCmd("chmod", "-R", "755", datakit.InstallDir, logDir); err != nil {
		l.Errorf("chmod failed: %v", err)
		return
	}

	// chown.
	if err := executeCmd("chown", "-R", fmt.Sprintf("%s:%s", userName, groupName), upgrader.InstallDir, upgrader.DefaultLogDir); err != nil {
		l.Errorf("chown failed: %v", err)
		return
	}
	// chmod.
	if err := executeCmd("chmod", "-R", "755", upgrader.InstallDir, upgrader.DefaultLogDir); err != nil {
		l.Errorf("chmod failed: %v", err)
		return
	}
}

func mkdir(path string, perm os.FileMode) error {
	l.Infof("MkdirAll %s => %s", path, perm.String())
	return os.MkdirAll(path, perm)
}

func executeCmd(name string, arg ...string) error {
	l.Infof("executing %s %v", name, arg)
	cmd := exec.Command(name, arg...)
	return cmd.Run()
}
