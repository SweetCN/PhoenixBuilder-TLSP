package core

import (
	"fmt"
	"os"
	"os/exec"
	"phoenixbuilder/fastbuilder/args"
	I18n "phoenixbuilder/fastbuilder/i18n"
	"phoenixbuilder/fastbuilder/utils"

	"github.com/pterm/pterm"

	"phoenixbuilder/fastbuilder/readline"
)

func setup() {
	pterm.Error.Prefix = pterm.Prefix{
		Text:  "ERROR",
		Style: pterm.NewStyle(pterm.BgBlack, pterm.FgRed),
	}

	I18n.Init()
	// iSH.app specific, for foreground ability
	if _, err := os.Stat("/dev/location"); err == nil {
		// Call location service
		pterm.Println(pterm.Yellow(I18n.T(I18n.Notice_iSH_Location_Service)))
		cmd := exec.Command("ash", "-c", "cat /dev/location > /dev/null &")
		err := cmd.Start()
		if err != nil {
			fmt.Println(err)
		}
	}

	if !args.NoReadline {
		readline.InitReadline()
	}
	if args.SkipMCPCheckChallenges {
		pterm.Warning.Println("Will login to the rental server without passing MCPCheckChallenges.")
		pterm.Info.Println("The implementation of command execution has been replaced with sending Netease ai commands.")
		pterm.Info.Println("When a command is executed as `CommandOriginPlayer` role, the response is always received. The OutputType corresponding to the response will reset to `CommandOutputTypeNone` because we don't know the actual value of `sendcommandfeedback` game rule.")
		pterm.Info.Println("Gamerule `sendcommandfeedback` will be updated to false(if we can) in order to reduce screen brushing.")
	}
}

func display_info() {
	pterm.DefaultBox.Println(pterm.LightCyan("https://github.com/SweetCN/PhoenixBuilder-TLSP"))
	pterm.Println(pterm.Yellow("PhoenixBuilder " + args.FBVersion))
	if I18n.ShouldDisplaySpecial() {
		fmt.Printf("%s", I18n.T(I18n.Special_Startup))
	}
}

func check_update() {
	fmt.Printf(I18n.T(I18n.Notice_CheckUpdate))
	hasUpdate, latestVersion := utils.CheckUpdate(args.FBPlainVersion)
	fmt.Printf(I18n.T(I18n.Notice_OK))
	if hasUpdate {
		fmt.Printf(I18n.T(I18n.Notice_UpdateAvailable), latestVersion)
		fmt.Printf(I18n.T(I18n.Notice_UpdateNotice))
		// To ensure user won't ignore it directly, can be suppressed by command line argument.
		os.Exit(0)
	}
}

func Bootstrap() {
	setup()
	display_info()
	defer Fatal()
	if args.DebugMode {
		init_and_run_debug_client()
		return
	}
	runInteractiveClient()
}

func runInteractiveClient() {
	env := ConfigRealEnvironment()
	EstablishConnectionAndInitEnv(env)
	go EnterReadlineThread(env, nil)
	defer DestroyEnvironment(env)
	EnterWorkerThread(env, nil)
}

func init_and_run_debug_client() {
	env := ConfigDebugEnvironment()
	EstablishConnectionAndInitEnv(env)
	go EnterReadlineThread(env, nil)
	defer DestroyEnvironment(env)
	EnterWorkerThread(env, nil)
}
