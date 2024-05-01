package main

import (
	"golang.org/x/exp/slog"
	"showmeyourcode/go/playground/samples"
	"showmeyourcode/go/playground/tasks"
)

func main() {
	slog.Info("========== Go playground ==========\n")
	slog.Info("> Below you can see code samples showing Go features.\n")
	slog.Info("")

	samples.Arrays()
	samples.Errors()
	samples.Concurrency()
	samples.Functions()
	samples.IfElse()
	samples.Generics()
	samples.Loops()
	samples.Maps()
	samples.Pointers()
	samples.Recursions()
	samples.Switches()
	samples.Slices()
	samples.Structs()
	samples.Variables()
	samples.Timers()
	samples.Tickers()
	samples.WorkerPools()
	samples.Defer()
	samples.Recover()
	samples.Regex()
	samples.Jsons()
	samples.HttpClient()
	samples.Reflection()

	slog.Info("")
	slog.Info("========== Tasks ==========")
	slog.Info("> Below you can see coding tasks which help you understand Go.")
	slog.Info("")

	tasks.Trees()
	tasks.Slices()
	tasks.Channels()
	tasks.Defer()
	tasks.Polymorphism()
	tasks.ConcurrentMap()
	tasks.DesignPatters()
	tasks.Types()

	slog.Info("PROGRAM FINISHED")
}
