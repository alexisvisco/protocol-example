"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.compile = void 0;
const plugin_pb_1 = require("google-protobuf/google/protobuf/compiler/plugin_pb");
const autogenerate_1 = require("./autogenerate");
const utils_1 = require("./utils");
const deserializeConfig_1 = require("./deserializeConfig");
function compile(input) {
    const request = plugin_pb_1.CodeGeneratorRequest.deserializeBinary(input);
    const options = (0, deserializeConfig_1.deserializeConfig)(request.getParameter() ?? "");
    const isTypescript = options.language === "typescript";
    const response = new plugin_pb_1.CodeGeneratorResponse();
    response.setSupportedFeatures(plugin_pb_1.CodeGeneratorResponse.Feature.FEATURE_PROTO3_OPTIONAL);
    const identifierTable = (0, utils_1.buildIdentifierTable)(request);
    function writeFile(name, content) {
        const file = new plugin_pb_1.CodeGeneratorResponse.File();
        file.setName(name);
        file.setContent(content);
        response.addFile(file);
    }
    request.getProtoFileList().forEach((fileDescriptorProto) => {
        const name = fileDescriptorProto.getName();
        if (!name) {
            return;
        }
        const protobufTs = (0, autogenerate_1.generate)(fileDescriptorProto, identifierTable, options);
        writeFile(isTypescript ? (0, utils_1.getProtobufTSFileName)(name) : (0, utils_1.getProtobufJSFileName)(name), protobufTs);
    });
    return response;
}
exports.compile = compile;
