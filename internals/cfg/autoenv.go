package cfg

import (
	"fmt"

	"github.com/ARC5RF/fsup"
	"github.com/joho/godotenv"
)

func init() {
	env, err := fsup.FromWD(".env")
	if err != nil {
		panic(err)
	}
	fmt.Println("using .env ", env)

	err = godotenv.Load(env)
	if err != nil {
		panic(err)
	}
}
