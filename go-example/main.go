package main

import (
	"context"
	"fmt"
	"os"

	extism "github.com/extism/go-sdk"
)

func main() {
	manifest := extism.Manifest{
		Wasm: []extism.Wasm{
			extism.WasmUrl{
				Url: "https://github.com/extism/plugins/releases/latest/download/count_vowels.wasm",
			},
		},
	}

	manifest2 := extism.Manifest{
		Wasm: []extism.Wasm{
			extism.WasmUrl{
				Url: "https://cdn.modsurfer.dylibso.com/api/v1/module/15edf0aaa4d9bc0fa6498ecdf8e59c2666b2497b690beb282ef4b133bda99edc.wasm",
			},
		},
	}

	ctx := context.Background()
	config := extism.PluginConfig{
		EnableWasi: true,
	}
	plugin, err := extism.NewPlugin(ctx, manifest, config, []extism.HostFunction{})

	plugin2, err2 := extism.NewPlugin(ctx, manifest2, config, []extism.HostFunction{})

	if err != nil {
		fmt.Printf("Failed to initialize plugin: %v\n", err)
		os.Exit(1)
	}

	if err2 != nil {
		fmt.Printf("Failed to initialize plugin: %v\n", err2)
		os.Exit(1)
	}

	data := []byte("Hello chat!")
	exit, out, err := plugin.Call("count_vowels", data)
	if err != nil {
		fmt.Println(err)
		os.Exit(int(exit))
	}

	response := string(out)
	fmt.Println(response)

	data2 := []byte(`{"time": "2024-02-29T06:10:27.471Z", "format": "January 2nd @ 3:04pm MST"}`)

	exit2, out2, err3 := plugin2.Call("format", data2)

	if err3 != nil {
		fmt.Println(err3)
		os.Exit(int(exit2))
	}

	response2 := string(out2)
	fmt.Println(response2)
}
