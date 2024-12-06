package main

import (
	"fmt"

	"github.com/julienlevasseur/goconfig/pkg/apt"
	"github.com/julienlevasseur/goconfig/pkg/file"
	"github.com/julienlevasseur/goconfig/pkg/softwares/helm"
	"github.com/julienlevasseur/goconfig/pkg/softwares/minikube"
)

func main() {
	// Install CNI plugins
	//cni.Install("v1.1.1", "amd64")

	//	// Install Traefik
	//	traefikVersion := "2.9.5"
	//	traefik.Install(
	//		fmt.Sprintf("v%s", traefikVersion),
	//		"amd64",
	//		"linux",
	//		conditions.NotIf(
	//			fmt.Sprintf("traefik version|head -1|grep %s", traefikVersion),
	//		),
	//	)
	//
	//	// traefik config template
	//	traefikCfg := `api:
	//  dashboard: {{ .Dashboard}}
	//  insecure: {{ .Insecure}}
	//`
	//
	//	type traefikVars struct {
	//		Dashboard string
	//		Insecure  string
	//	}
	//
	//	tVars := traefikVars{
	//		Dashboard: "true",
	//		Insecure:  "true",
	//	}
	//
	//	file.Template(
	//		"traefik.yml",
	//		traefikCfg,
	//		tVars,
	//	)

	//	err := crontab.Cron(
	//		"0",
	//		"3",
	//		"*",
	//		"*",
	//		"0",
	//		"systemctl stop traefik && rm /var/log/traefik.log && systemctl start traefik",
	//	)
	//	if err != nil {
	//		panic(err)
	//	}

	//consulVersion := "1.14.1"
	//consul.Install(
	//	fmt.Sprintf("%s", consulVersion),
	//	"amd64",
	//	"linux",
	//	conditions.NotIf(
	//		fmt.Sprintf("consul version|head -1|grep %s", consulVersion),
	//	),
	//)

	// 	folder.Create("/etc/consul")

	// 	consulCfg := `datacenter = "{{ .Datacenter}"
	// data_dir = "{{ .DataDir}}"
	// encrypt = "{{ .Encrypt}}"
	// log_level = "{{ .LogLevel}}"
	// node_name = "{{ .NodeName}}"
	// server = {{ .Server}}
	// bootstrap_expect = {{ .BootstrapExpect}}
	// bind_addr = "{{ .BindAddr}}"
	// client_addr = "{{ .ClientAddr}}"
	// acl {
	//   enabled = {{ .ACL.Enabled}}
	//   default_policy = "{{ .ACL.DefaultPolicy}}"
	//   down_policy = "{{ .ACL.DownPolicy}}"
	//   enable_token_persistence = {{ .ACL.EnableTokenPersistence}}
	//   enable_token_replication = {{ .ACL.EnableTokenReplication}}
	//   tokens {
	// 	master = "{{ .ACL.Tokens.Master}}"
	//   }
	// }
	// 	`

	// 	type ConsulTokens struct {
	// 		Master string
	// 	}

	// 	type consulACL struct {
	// 		Enabled                bool
	// 		DefaultPolicy          string
	// 		DownPolicy             string
	// 		EnableTokenPersistence bool
	// 		EnableTokenReplication bool
	// 		Tokens                 ConsulTokens
	// 	}

	// 	type consulVars struct {
	// 		Datacenter      string
	// 		DataDir         string
	// 		Encrypt         string
	// 		LogLevel        string
	// 		NodeName        string
	// 		Server          bool
	// 		BootstrapExpect int
	// 		BindAddr        string
	// 		ClientAddr      string
	// 		ACL             consulACL
	// 	}
	// 	cVars := consulVars{
	// 		Datacenter:      "aws",
	// 		DataDir:         "/opt/consul",
	// 		Encrypt:         "u00sHTLcDsjucyWN8Jfr2g==",
	// 		LogLevel:        "INFO",
	// 		NodeName:        "Traefik",
	// 		Server:          true,
	// 		BootstrapExpect: 1,
	// 		BindAddr:        "{{ GetInterfaceIP \"ens5\" }}",
	// 		ClientAddr:      "0.0.0.0",
	// 		ACL: consulACL{
	// 			Enabled:                true,
	// 			DefaultPolicy:          "deny",
	// 			DownPolicy:             "extend-cache",
	// 			EnableTokenPersistence: true,
	// 			EnableTokenReplication: true,
	// 			Tokens: ConsulTokens{
	// 				Master: "1d7b246f-8000-e312-fe59-c9a57190119f",
	// 			},
	// 		},
	// 	}

	// 	fmt.Println("create template consul.hcl")
	// 	file.Template(
	// 		"consul.hcl",
	// 		consulCfg,
	// 		cVars,
	// 	)

	// nomadConfig := nomad.NomadConfig{
	// 	BindAddr: "127.0.0.1",
	// }

	// nomad.ConfigFile(nomadConfig, "/tmp/noamd.hcl")

	// az_cli.Install(false)

	apt.Update()
	yamllintPkg := apt.IPackage{
		Name: "yamllint",
	}
	err := yamllintPkg.Install(file.Exists("/usr/bin/yamllint"))
	if err != nil {
		fmt.Println(err)
	}

	// Docker install
	apt.Update()
	dockerPkg := apt.IPackage{
		Name: "docker.io",
	}
	err = dockerPkg.Install(file.Exists("/usr/bin/docker"))
	if err != nil {
		fmt.Println(err)
	}

	// Minikube install
	minikube.Install("amd64", "linux", file.Exists("/usr/local/bin/minikube"))

	helm.Install(false)

}
