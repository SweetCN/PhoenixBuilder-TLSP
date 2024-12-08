package core

import (
	"phoenixbuilder/fastbuilder/args"
	"phoenixbuilder/fastbuilder/environment"
	"phoenixbuilder/fastbuilder/function"
	fbauth "phoenixbuilder/fastbuilder/mv4"
	fbtask "phoenixbuilder/fastbuilder/task"
	"phoenixbuilder/minecraft"
	"phoenixbuilder/mirror/io/global"
	"phoenixbuilder/mirror/io/lru"
)

func create_environment() *environment.PBEnvironment {
	clientOption := fbauth.MakeDefaultClientOptions()
	clientOption.AuthServer = args.AuthServer
	clientOption.RespondUserOverride = args.CustomGameName
	env := &environment.PBEnvironment{
		ClientOptions: clientOption,
	}
	env.Resources = nil
	env.TaskHolder = fbtask.NewTaskHolder()
	functionHolder := function.NewFunctionHolder(env)
	env.FunctionHolder = functionHolder
	env.LRUMemoryChunkCacher = lru.NewLRUMemoryChunkCacher(12, false)
	env.ChunkFeeder = global.NewChunkFeeder()
	return env
}

// Shouldn't be called when running a debug client
func ConfigRealEnvironment() *environment.PBEnvironment {
	env := create_environment()

	env.LoginInfo = environment.LoginInfo{
		Token:          "",
		ServerCode:     "",
		ServerPasscode: "",
		Username:       "",
		Password:       "",
	}
	env.FBAuthClient = fbauth.CreateClient(env.ClientOptions)
	return env
}

func ConfigDebugEnvironment() *environment.PBEnvironment {
	env := create_environment()
	env.IsDebug = true
	env.LoginInfo = environment.LoginInfo{
		ServerCode: "[DEBUG]",
	}
	return env
}

func DestroyEnvironment(env *environment.PBEnvironment) {
	env.Connection.(*minecraft.Conn).Close()
}
