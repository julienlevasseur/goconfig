package powerlineshell

import (
	"log"

	homedir "github.com/julienlevasseur/go-homedir"
	"github.com/julienlevasseur/goconfig/pkg/apt"
	"github.com/julienlevasseur/goconfig/pkg/command"
	"github.com/julienlevasseur/goconfig/pkg/file"
	"github.com/julienlevasseur/goconfig/pkg/notif"
)

const name = "powerline_shell"

func Install(notIf *bool) {
	if notIf != nil && !*notIf {
		log.Printf("[%v][Install]\n", name)
		// fmt.Printf("[%v][Install]\n", name)

		powerlineFontsPkg := apt.IPackage{
			Name: "fonts-powerline",
		}

		err := powerlineFontsPkg.Install(file.Exists("/usr/share/fonts/opentype/PowerlineSymbols.otf"))
		if err != nil {
			// fmt.Println(err)
			log.Fatal(err)
		}

		err = file.Download(
			"https://github.com/julienlevasseur/powerline-go/releases/download/v1.26/powerline-go-linux-amd64",
			"/usr/local/bin/powerline-go",
		)
		if err != nil {
			// fmt.Println(err)
			log.Fatal(err)
		}

		err = command.Exec(
			"sudo",
			[]string{
				"chmod",
				"+x",
				"/usr/local/bin/powerline-go",
			},
			nil,
		)
		if err != nil {
			// fmt.Println(err)
			log.Fatal(err)
		}

		homePath, err := homedir.Dir()
		if err != nil {
			// fmt.Println(err)
			log.Fatal(err)
		}

		bashrcPath := homePath + "/.bashrc"

		updatePS1FuncLine, err := file.LineIsPresent(
			bashrcPath,
			"function _update_ps1() {",
		)
		if err != nil {
			// fmt.Println(err)
			log.Fatal(err)
		}

		file.Append(
			bashrcPath,
			`
function _update_ps1() {
  PS1="$(/usr/local/bin/powerline-go -error $? -jobs $(jobs -p|wc -l) -theme default -mode compatible -modules time,host,cwd,git,venv,terraform-workspace,kube,profiler,root,newline)"
}`,
			&updatePS1FuncLine,
		)

		promptCmdLine, err := file.LineIsPresent(
			bashrcPath,
			"PROMPT_COMMAND",
		)
		if err != nil {
			// fmt.Println(err)
			log.Fatal(err)
		}

		file.Append(
			bashrcPath,
			`
if [ "$TERM" != "linux" ] && [ -f "/usr/local/bin/powerline-go" ]; then
    PROMPT_COMMAND="_update_ps1; $PROMPT_COMMAND"
fi`,
			&promptCmdLine,
		)

	} else {
		notif.IgnoreDueToNotIf(name, "Install")
	}
}
