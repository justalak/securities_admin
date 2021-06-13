package main

import config2 "securities_admin/config"

var (
	Config *config2.ApplicationConfig
)

func LoadConfig() {
	Config = config2.LoadConfig()
}