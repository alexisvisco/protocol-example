import { createDummyService, DummyService, LogRequest, LogResponse } from './gen/ts/protos/dummy/v1/service_dummy.pb'
// import fetch from 'node-fetch';
import { client, createTwirpServer, TwirpError } from 'twirpscript'
import { nodeHttpTransport } from 'twirpscript/dist/node'
import { createServer } from 'http'
import { HealthRequest, HealthResponse } from './gen/ts/protos/common/v1/health.pb'

client.rpcTransport = nodeHttpTransport


const service: DummyService<{}> = {
  Health(healthRequest: HealthRequest, context: {}): Promise<HealthResponse> | HealthResponse {
    return {
      status: 0
    }
  }, Log(logRequest: LogRequest, context: {}): Promise<LogResponse> | LogResponse {
    throw new TwirpError({
      msg: 'unable to do that', code: 'canceled', meta: {
        'hi': 'salut'
      }
    })
    return {
      response: logRequest.name
    }
  }

}

const services = [ createDummyService(service) ]
const app = createTwirpServer<{}, typeof services>(services, {
  debug: false,
})


createServer(app).listen('8080', () =>
  console.log(`Server listening on port ${8081}`)
)