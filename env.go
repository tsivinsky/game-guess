package main

type env struct {
	RAWG_API_KEY string `env:"RAWG_API_KEY,required"`
}

var Env = new(env)
