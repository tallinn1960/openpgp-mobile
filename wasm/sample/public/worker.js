self.importScripts('wasm_exec.js');

let loaded = false;
load = () => {
    const go = new Go();
    let mod, inst;
    return WebAssembly.instantiateStreaming(fetch("openpgp.wasm"), go.importObject).then(
        async result => {
            mod = result.module;
            inst = result.instance;
            loaded = true
            const run = async () => {
                try {
                    await go.run(inst);
                } catch (e) {
                    console.warn(e)
                    loaded = false
                    await load()
                }
            }
            run()
        }
    );
}

onmessage = async (e) => {
    if (!loaded) {
        await load();
    }
    const response = {
        id: e.data.id,
        data: 'test'
    }
    postMessage(response);
};