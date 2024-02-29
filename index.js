import createPlugin from '@extism/extism';

const plugin = await createPlugin(
	'https://github.com/extism/plugins/releases/latest/download/count_vowels.wasm',
	{ useWasi: true },
);

const datePlugin = await createPlugin(
	'https://cdn.modsurfer.dylibso.com/api/v1/module/15edf0aaa4d9bc0fa6498ecdf8e59c2666b2497b690beb282ef4b133bda99edc.wasm',
	{ useWasi: true },
);

const input = 'Hello chat!';
let out = await plugin.call('count_vowels', input);
console.log(out.text());

const dateInput =
	'{"time": "2024-03-05T06:10:27.471Z", "format": "Jan 2, 3:04pm MST"}';
const output = await datePlugin.call('format', dateInput);

console.log(output.text());
