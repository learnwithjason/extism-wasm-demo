<?php

use Extism\Plugin;
use Extism\Manifest;
use Extism\UrlWasmSource;

$wasm = new UrlWasmSource("https://github.com/extism/plugins/releases/latest/download/count_vowels.wasm");
$manifest = new Manifest($wasm);

$plugin = new Plugin($manifest, true);

$output = $plugin->call("count_vowels", "Hello chat!");

echo $output;
