// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

package cmds

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/spf13/pflag"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit"
	cp "gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/colorprint"
)

var (

	//
	// doc related flags.
	//
	fsDocName         = "doc"
	fsDoc             = pflag.NewFlagSet(fsDocName, pflag.ContinueOnError)
	flagDocExportDocs = fsDoc.String("export-docs", "", "export all inputs and related docs to specified path")
	flagDocIgnore     = fsDoc.String("ignore", "", "disable list, i.e., --ignore nginx,redis,mem")
	flagDocLogPath    = fsDoc.String("log", commonLogFlag(), "log path")
	flagDocTODO       = fsDoc.String("TODO", "TODO", "set TODO placeholder")
	flagDocVersion    = fsDoc.String("version", datakit.Version, "specify version string in document's header")
	fsDocUsage        = func() {
		fmt.Printf("usage: datakit doc [options]\n\n")
		fmt.Printf("Doc used to manage all documents related to DataKit. Available options:\n\n")
		fmt.Println(fsDoc.FlagUsagesWrapped(0))
	}

	//
	// DQL related flags.
	//
	fsDQLName  = "dql"
	fsDQL      = pflag.NewFlagSet(fsDQLName, pflag.ContinueOnError)
	fsDQLUsage = func() {
		fmt.Printf("usage: datakit dql [options]\n\n")
		fmt.Printf("DQL used to query data. If no option specified, query interactively. Other available options:\n\n")
		fmt.Println(fsDQL.FlagUsagesWrapped(0))
	}

	flagDQLJSON        = fsDQL.BoolP("json", "J", false, "output in json format")
	flagDQLAutoJSON    = fsDQL.Bool("auto-json", false, "pretty output string if field/tag value is JSON")
	flagDQLVerbose     = fsDQL.BoolP("verbose", "V", false, "verbosity mode")
	flagDQLString      = fsDQL.StringP("run", "R", "", "run single DQL")
	flagDQLToken       = fsDQL.StringP("token", "T", "", "run query for specific token(workspace)")
	flagDQLCSV         = fsDQL.String("csv", "", "Specify the directory")
	flagDQLForce       = fsDQL.BoolP("force", "F", false, "overwrite csv if file exists")
	flagDQLDataKitHost = fsDQL.StringP("host", "H", "", "specify datakit host to query")
	flagDQLLogPath     = fsDQL.String("log", commonLogFlag(), "log path")

	//
	// running mode. (not used).
	//
	fsRunName          = "run"
	fsRun              = pflag.NewFlagSet(fsRunName, pflag.ContinueOnError)
	FlagRunInContainer = fsRun.BoolP("container", "C", false, "running in container mode")
	fsRunUsage         = func() {
		fmt.Printf("usage: datakit run [options]\n\n")
		fmt.Printf("Run used to select different datakit running mode.\n\n")
		fmt.Println(fsRun.FlagUsagesWrapped(0))
	}

	//
	// pipeline related flags.
	//
	fsPLName       = "pipeline"
	fsPL           = pflag.NewFlagSet(fsPLName, pflag.ContinueOnError)
	flagPLCategory = fsPL.StringP("category", "C", "logging", "data category (logging, metric, ...)")
	flagPLNS       = fsPL.StringP("namespace", "N", "default", "namespace (default, gitrepo, remote)")
	flagPLName     = fsPL.StringP("name", "P", "", "pipeline name")
	flagPLLogPath  = fsPL.String("log", commonLogFlag(), "log path")
	flagPLTxtData  = fsPL.StringP("txt", "T", "", "text string for the pipeline or grok(json or raw text)")
	flagPLTxtFile  = fsPL.StringP("file", "F", "", "text file path for the pipeline or grok(json or raw text)")
	flagPLTable    = fsPL.Bool("tab", false, "output result in table format")
	flagPLDate     = fsPL.Bool("date", false, "append date display(according to local timezone) on timestamp")
	fsPLUsage      = func() {
		fmt.Printf("usage: datakit pipeline -P [pipeline-script-name.p] -T [text] [other-options...]\n\n")
		fmt.Printf("Pipeline used to debug exists pipeline script.\n\n")
		fmt.Println(fsPL.FlagUsagesWrapped(0))
	}

	//
	// version related flags.
	//
	fsVersionName                 = "version"
	fsVersion                     = pflag.NewFlagSet(fsVersionName, pflag.ContinueOnError)
	flagVersionLogPath            = fsVersion.String("log", commonLogFlag(), "log path")
	flagVersionDisableUpgradeInfo = fsVersion.Bool("upgrade-info-off", false, "do not show upgrade info")
	fsVersionUsage                = func() {
		fmt.Printf("usage: datakit version [options]\n\n")
		fmt.Printf("Version used to handle version related functions.\n\n")
		fmt.Println(fsVersion.FlagUsagesWrapped(0))
	}

	//
	// service management related flags.
	//
	fsServiceName        = "service"
	fsService            = pflag.NewFlagSet(fsServiceName, pflag.ContinueOnError)
	flagServiceLogPath   = fsService.String("log", commonLogFlag(), "log path")
	flagServiceRestart   = fsService.BoolP("restart", "R", false, "restart datakit service")
	flagServiceStop      = fsService.BoolP("stop", "T", false, "stop datakit service")
	flagServiceStart     = fsService.BoolP("start", "S", false, "start datakit service")
	flagServiceUninstall = fsService.BoolP("uninstall", "U", false, "uninstall datakit service")
	flagServiceReinstall = fsService.BoolP("reinstall", "I", false, "reinstall datakit service")
	fsServiceUsage       = func() {
		fmt.Printf("usage: datakit service [options]\n\n")
		fmt.Printf("Service used to manage datakit service\n\n")
		fmt.Println(fsService.FlagUsagesWrapped(0))
	}

	//
	// monitor related flags.
	//
	fsMonitorName              = "monitor"
	fsMonitor                  = pflag.NewFlagSet(fsMonitorName, pflag.ContinueOnError)
	flagMonitorTo              = fsMonitor.String("to", "localhost:9529", "specify the DataKit(IP:Port) to show its statistics")
	flagMonitorMaxTableWidth   = fsMonitor.IntP("max-table-width", "W", 128, "set max table cell width")
	flagMonitorLogPath         = fsMonitor.String("log", commonLogFlag(), "log path")
	flagMonitorRefreshInterval = fsMonitor.DurationP("refresh", "R", 5*time.Second, "refresh interval")
	flagMonitorVerbose         = fsMonitor.BoolP("verbose", "V", false, "show all statistics info, default not show goroutine and inputs config info")
	flagMonitorModule          = fsMonitor.StringP("module", "M", "", "show only specified module stats, seprated by ',', i.e., -M filter,inputs")
	flagMonitorOnlyInputs      = fsMonitor.StringP("input", "I", "", "show only specified inputs stats, seprated by ',', i.e., -I cpu,mem")
	fsMonitorUsage             = func() {
		fmt.Printf("usage: datakit monitor [options]\n\n")
		fmt.Printf("Monitor used to show datakit running statistics\n\n")
		fmt.Println(fsMonitor.FlagUsagesWrapped(0))
	}

	//
	// install related flags.
	//
	fsInstallName         = "install"
	fsInstall             = pflag.NewFlagSet(fsInstallName, pflag.ContinueOnError)
	flagInstallLogPath    = fsInstall.String("log", commonLogFlag(), "log path")
	flagInstallTelegraf   = fsInstall.Bool("telegraf", false, "install Telegraf")
	flagInstallScheck     = fsInstall.Bool("scheck", false, "install SCheck")
	flagInstallIPDB       = fsInstall.String("ipdb", "", "install IP database")
	flagInstallSymbolTool = fsInstall.Bool("symbol-tools", false,
		"install tools for symbolizing crash backtrace address, including Android command line tools, ProGuard, Android-NDK, atosl, etc ...")
	fsInstallUsage = func() {
		fmt.Printf("usage: datakit install [options]\n\n")
		fmt.Printf("Install used to install DataKit related packages and plugins\n\n")
		fmt.Println(fsInstall.FlagUsagesWrapped(0))
	}

	//
	// checking/testing related flags.
	//
	fsCheckName      = "check"
	fsCheck          = pflag.NewFlagSet(fsCheckName, pflag.ContinueOnError)
	flagCheckLogPath = fsCheck.String("log", commonLogFlag(), "log path")

	flagCheckConfig    = fsCheck.Bool("config", false, "check inputs configures and datait.conf")
	flagCheckConfigDir = fsCheck.String("config-dir", "", "check configures under specified path")
	flagCheckSample    = fsCheck.Bool("sample", false,
		"check all inputs config sample, to ensure all sample are valid TOML")
	fsCheckUsage = func() {
		fmt.Printf("usage: datakit check [options]\n\n")
		fmt.Printf("Various check tools for DataKit\n\n")
		fmt.Println(fsCheck.FlagUsagesWrapped(0))
	}

	//
	// debug/trouble-shooting related flags.
	//
	fsDebugName       = "debug"
	fsDebug           = pflag.NewFlagSet(fsDebugName, pflag.ContinueOnError)
	flagDebugLogPath  = fsDebug.String("log", commonLogFlag(), "log path")
	flagDebugLoadLog  = fsDebug.Bool("upload-log", false, "upload log")
	flagDebugGlobConf = fsDebug.String("glob-conf", "",
		"find the glob path and print it, provide a configuration file that contains glob statements written on separate lines.")
	flagDebugRegexConf = fsDebug.String("regex-conf", "",
		"export regex match results, provide a configuration file where the first line is a regular expression and the rest of the file is text.")
	flagDebugPromConf  = fsDebug.String("prom-conf", "", "specify the prom input conf to debug")
	flagDebugBugReport = fsDebug.Bool("bug-report", false, "export DataKit running information for troubleshooting")
	flagDebugTestInput = fsDebug.String("test-input", "", "specify input's config file to test")

	fsDebugUsage = func() {
		fmt.Printf("usage: datakit debug [options]\n\n")
		fmt.Printf("Various debug options for DataKit\n\n")
		fmt.Println(fsDebug.FlagUsagesWrapped(0))
	}

	//
	// tools related flags.
	//
	fsToolName = "tool"
	fsTool     = pflag.NewFlagSet(fsToolName, pflag.ContinueOnError)

	flagToolGrokQ = fsTool.Bool("grokq", false, "query groks interactively")

	flagToolLogPath   = fsTool.String("log", commonLogFlag(), "log path")
	flagToolCloudInfo = fsTool.Bool("show-cloud-info", false,
		"show current host's cloud info(currently support aliyun/tencent/aws/hwcloud/azure)")
	flagToolIPInfo        = fsTool.String("ipinfo", "", "show IP geo info")
	flagToolWorkspaceInfo = fsTool.Bool("workspace-info", false, "show workspace info")

	flagToolDumpSamples       = fsTool.String("dump-samples", "", "dump all inputs samples")
	flagToolDefaultMainConfig = fsTool.Bool("default-main-conf", false, "print default datakit.conf")

	flagToolSetupCompleterScripts = fsTool.Bool("setup-completer-script", false, "auto generate auto completion script(Linux only)")
	flagToolCompleterScripts      = fsTool.Bool("completer-script", false, "show completion script(Linux only)")

	flagToolParseLineProtocol = fsTool.String("parse-lp", "", "parse line-protocol file")
	flagToolJSON              = fsTool.Bool("json", false, "output in JSON format(partially supported)")
	flagToolUpdateIPDB        = fsTool.Bool("update-ipdb", false, "update local IPDB")

	fsToolUsage = func() {
		fmt.Printf("usage: datakit tool [options]\n\n")
		fmt.Printf("Various tools for DataKit\n\n")
		fmt.Println(fsTool.FlagUsagesWrapped(0))
	}
)

func commonLogFlag() string {
	if runtime.GOOS == datakit.OSWindows {
		return "nul" // under windows, nul is /dev/null
	}
	return "/dev/null"
}

//nolint:lll
const datakitIntro = `DataKit is an open source, integrated data collection agent, which provides full
platform (Linux/Windows/macOS) support and has comprehensive data collection capability,
covering various scenarios such as host, container, middleware, tracing, logging and
security inspection.`

func printHelp() {
	fmt.Fprintf(os.Stderr, "%s\n", datakitIntro)
	fmt.Fprintf(os.Stderr, "\nUsage:\n\n")

	fmt.Fprintf(os.Stderr, "\tdatakit <command> [arguments]\n\n")

	fmt.Fprintf(os.Stderr, "The commands are:\n\n")

	fmt.Fprintf(os.Stderr, "\tservice    manage datakit service\n")
	fmt.Fprintf(os.Stderr, "\tdql        query DQL for various usage\n")
	fmt.Fprintf(os.Stderr, "\trun        select DataKit running mode(defaul running as service)\n")
	fmt.Fprintf(os.Stderr, "\tpipeline   debug pipeline\n")
	fmt.Fprintf(os.Stderr, "\tmonitor    show datakit running statistics\n")
	fmt.Fprintf(os.Stderr, "\tinstall    install DataKit related packages and plugins\n")
	fmt.Fprintf(os.Stderr, "\tcheck      methods of all check tools within DataKit\n")
	fmt.Fprintf(os.Stderr, "\tdebug      methods of all debug tools within DataKit\n")
	fmt.Fprintf(os.Stderr, "\ttool       methods of all tools within DataKit\n")
	fmt.Fprintf(os.Stderr, "\tdoc        manage all documents for DataKit\n")

	// TODO: add more commands...

	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "Use 'datakit help <command>' for more information about a command.\n\n")
}

func runHelpFlags() {
	switch len(os.Args) {
	case 2: // only 'datakit help'
		printHelp()
	case 3: // need help for various commands
		switch os.Args[2] {
		case fsDocName:
			fsDocUsage()

		case fsPLName:
			fsPLUsage()

		case fsDQLName:
			fsDQLUsage()

		case fsRunName:
			fsRunUsage()

		case fsVersionName:
			fsVersionUsage()

		case fsServiceName:
			fsServiceUsage()

		case fsMonitorName:
			fsMonitorUsage()

		case fsInstallName:
			fsInstallUsage()

		case fsToolName:
			fsToolUsage()

		case fsDebugName:
			fsDebugUsage()

		case fsCheckName:
			fsCheckUsage()

		default: // add more
			cp.Errorf("[E] flag provided but not defined: `%s'\n\n", os.Args[2])
			printHelp()
			os.Exit(-1)
		}
	}
}

// nolint:funlen
func doParseAndRunFlags() {
	pflag.Usage = printHelp
	pflag.ErrHelp = errors.New("")

	if len(os.Args) > 1 {
		if os.Args[1] == "help" {
			runHelpFlags()
			os.Exit(0)
		}

		switch os.Args[1] {
		case fsRunName:

			if len(os.Args) < 3 {
				fsRunUsage()
				os.Exit(-1)
			}

			if err := fsRun.Parse(os.Args[2:]); err != nil {
				cp.Errorf("Parse: %s\n", err)
				fsRunUsage()
				os.Exit(-1)
			}

			return

		case fsCheckName:

			if len(os.Args) < 3 {
				fsCheckUsage()
				os.Exit(-1)
			}

			setCmdRootLog(*flagCheckLogPath)

			if err := fsCheck.Parse(os.Args[2:]); err != nil {
				cp.Errorf("Parse: %s\n", err)
				fsCheckUsage()
				os.Exit(-1)
			}

			if err := runCheckFlags(); err != nil {
				cp.Errorf("%s\n", err)
				os.Exit(-1)
			}

		case fsDebugName:

			if len(os.Args) < 3 {
				fsDebugUsage()
				os.Exit(-1)
			}

			setCmdRootLog(*flagDebugLogPath)

			if err := fsDebug.Parse(os.Args[2:]); err != nil {
				cp.Errorf("Parse: %s\n", err)
				fsDebugUsage()
				os.Exit(-1)
			}

			if err := runDebugFlags(); err != nil {
				cp.Errorf("%s\n", err)
				os.Exit(-1)
			}

		case fsDocName:

			if len(os.Args) < 3 {
				fsDocUsage()
				os.Exit(-1)
			}

			setCmdRootLog(*flagDocLogPath)
			if err := fsDoc.Parse(os.Args[2:]); err != nil {
				cp.Errorf("Parse: %s\n", err)
				fsDocUsage()
				os.Exit(-1)
			}

			if err := runDocFlags(); err != nil {
				cp.Errorf("%s\n", err)
				os.Exit(-1)
			}
			os.Exit(0)

		case fsDQLName:

			if err := fsDQL.Parse(os.Args[2:]); err != nil {
				cp.Errorf("Parse: %s\n", err)
				fsDQLUsage()
				os.Exit(-1)
			}

			setCmdRootLog(*flagDQLLogPath)

			tryLoadMainCfg()

			if err := runDQLFlags(); err != nil {
				cp.Errorf("%s\n", err)
				os.Exit(-1)
			}

			os.Exit(0)

		case fsPLName:

			if len(os.Args) < 6 {
				fsPLUsage()
				os.Exit(-1)
			}

			setCmdRootLog(*flagPLLogPath)
			tryLoadMainCfg()

			if err := fsPL.Parse(os.Args[2:]); err != nil {
				cp.Errorf("[E] Parse: %s\n", err)
				fsPLUsage()
				os.Exit(-1)
			}

			if err := runPLFlags(); err != nil {
				cp.Errorf("[E] %s\n", err)
				os.Exit(-1)
			}

			os.Exit(0)

		case fsVersionName:

			if err := fsVersion.Parse(os.Args[2:]); err != nil {
				cp.Errorf("[E] parse: %s\n", err)
				fsVersionUsage()
				os.Exit(-1)
			}

			setCmdRootLog(*flagVersionLogPath)
			tryLoadMainCfg()

			if err := runVersionFlags(*flagVersionDisableUpgradeInfo); err != nil {
				cp.Errorf("[E] %s\n", err)
				os.Exit(-1)
			}

			os.Exit(0)

		case fsServiceName:

			if len(os.Args) < 3 {
				fsServiceUsage()
				os.Exit(-1)
			}

			if err := fsService.Parse(os.Args[2:]); err != nil {
				cp.Errorf("Parse: %s\n", err)
				fsServiceUsage()
				os.Exit(-1)
			}

			setCmdRootLog(*flagServiceLogPath)
			if err := runServiceFlags(); err != nil {
				cp.Errorf("%s\n", err)
				os.Exit(-1)
			}

			os.Exit(0)

		case fsMonitorName:

			if err := fsMonitor.Parse(os.Args[2:]); err != nil {
				cp.Errorf("Parse: %s\n", err)
				fsMonitorUsage()
				os.Exit(-1)
			}

			if *flagMonitorModule != "" {
				nomodule := existsModule(strings.Split(*flagMonitorModule, ","))
				if len(nomodule) != 0 {
					*flagMonitorVerbose = false
					cp.Errorf("has no module:%+v,check please!\n", nomodule)
					os.Exit(-1)
				}
			}

			setCmdRootLog(*flagMonitorLogPath)
			if err := runMonitorFlags(); err != nil {
				cp.Errorf("%s\n", err)
				os.Exit(-1)
			}

			os.Exit(0)

		case fsInstallName:

			if len(os.Args) < 3 {
				fsInstallUsage()
				os.Exit(-1)
			}

			if err := fsInstall.Parse(os.Args[2:]); err != nil {
				cp.Errorf("Parse: %s\n", err)
				fsInstallUsage()
				os.Exit(-1)
			}

			setCmdRootLog(*flagInstallLogPath)
			if err := installPlugins(); err != nil {
				cp.Errorf("%s\n", err)
				os.Exit(-1)
			}
			os.Exit(0)

		case fsToolName:

			if len(os.Args) < 3 {
				fsToolUsage()
				os.Exit(-1)
			}

			if err := fsTool.Parse(os.Args[2:]); err != nil {
				cp.Errorf("Parse: %s\n", err)
				fsToolUsage()
				os.Exit(-1)
			}

			setCmdRootLog(*flagToolLogPath)
			err := runToolFlags()
			if err != nil {
				cp.Errorf("%s\n", err)
				os.Exit(-1)
			}

			// NOTE: Do not exit here, you should exit in sub-tool's command if need

		default:
			cp.Errorf("unknown command `%s'\n", os.Args[1])
			printHelp()
		}
	}
}

func ParseFlags() {
	doParseAndRunFlags()
}
