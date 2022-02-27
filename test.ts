import {Log} from "./gen/ts/protos/dummy/v1/service_dummy.pb";

async function test() {
    const logResponse = await Log({name: "yeah", level: 0}, {
        baseURL: "http://localhost:8080",
    });

    console.log(logResponse)
}

test()