// THIS IS AN AUTOGENERATED FILE. DO NOT EDIT THIS FILE DIRECTLY.
// Source: protos/leads/v1/input.proto


import type {
    ByteSource,

    } from 'twirpscript';
import {
  BinaryReader,
BinaryWriter,

  } from 'twirpscript';
  







  //========================================//
  //                 Types                  //
  //========================================//
  
  
    /**
     * InputName represent the name of a input, it allow use to determine the i18n for the label, helper text...
 * The type of the in
     */
      export type InputName = typeof InputName[keyof typeof InputName];

    /**
     * InputType represent the type of an input
     */
      export type InputType = typeof InputType[keyof typeof InputType];

export interface InputDescriptor {
inputName: InputName;
inputType: InputType;
options: string[];
}

export interface InputValue {
inputName: InputName;
inputValueNumber?: bigint| null | undefined;
inputValueString?: string| null | undefined;
}


  //========================================//
  //        Protobuf Encode / Decode        //
  //========================================//
  
  
export const InputName = {INPUT_NAME_UNSPECIFIED: 0,
INPUT_NAME_AMOUNT: 1,
INPUT_NAME_ACCOUNT_TYPE: 2,
} as const;

export const InputType = {INPUT_TYPE_UNSPECIFIED: 0,
INPUT_TYPE_CURRENCY: 1,
INPUT_TYPE_TEXT: 2,
INPUT_TYPE_TEXT_AREA: 3,
INPUT_TYPE_SELECT_OPTION: 4,
INPUT_TYPE_REGION_OPTION: 5,
INPUT_TYPE_PHONE: 6,
} as const;

export const InputDescriptor = {        /**
         * Serializes a InputDescriptor to protobuf.
         */
        encode: function(inputDescriptor: Partial<InputDescriptor>): Uint8Array {
          return InputDescriptor._writeMessage(inputDescriptor, new BinaryWriter()).getResultBuffer();
        },
        

        /**
         * Deserializes a InputDescriptor from protobuf.
         */
        decode: function(bytes: ByteSource): InputDescriptor {
          return InputDescriptor._readMessage(InputDescriptor.initialize(), new BinaryReader(bytes));
        },
        

        /**
         * Serializes a InputDescriptor to JSON.
         */
        encodeJSON: function(inputDescriptor: Partial<InputDescriptor>): string {
          return JSON.stringify(InputDescriptor._writeMessageJSON(inputDescriptor));
        },
        

        /**
         * Deserializes a InputDescriptor from JSON.
         */
        decodeJSON: function(json: string): InputDescriptor {
          return InputDescriptor._readMessageJSON(InputDescriptor.initialize(), JSON.parse(json));
        },
        

        /**
         * Initializes a InputDescriptor with all fields set to their default value.
         */
        initialize: function(): InputDescriptor {
          return {
            inputName: 0,inputType: 0,options: [],
          };
        },

        /**
         * @private
         */
        _writeMessage: function(msg : Partial<InputDescriptor>, writer: BinaryWriter): BinaryWriter {
          if (msg.inputName) {writer.writeEnum(1, msg.inputName);}
if (msg.inputType) {writer.writeEnum(2, msg.inputType);}
if (msg.options?.length) {writer.writeRepeatedString(3, msg.options);}
            return writer;
        },

        /**
         * @private
         */
        _writeMessageJSON: function(msg : Partial<InputDescriptor>): Record<string, unknown> {
          const json: Record<string, unknown> = {};
          if (msg.inputName) {json.inputName = msg.inputName;}
if (msg.inputType) {json.inputType = msg.inputType;}
if (msg.options?.length) {json.options = msg.options;}
          return json;
        },
        
        /**
         * @private
         */
        _readMessage: function(msg: InputDescriptor, reader: BinaryReader): InputDescriptor {
          while (reader.nextField()) {
            const field = reader.getFieldNumber();
            switch (field) {
              case 1: {msg.inputName = reader.readEnum() as InputName;break;
}
case 2: {msg.inputType = reader.readEnum() as InputType;break;
}
case 3: {msg.options.push(reader.readString());break;
}
              default: {
                reader.skipField();
                break;
              }
            }
          }
          return msg;
        },

        /**
         * @private
         */
        _readMessageJSON: function(msg: InputDescriptor, json: any): InputDescriptor {
          const inputName = json.inputName ?? json.input_name;if (inputName) {msg.inputName = inputName;}
const inputType = json.inputType ?? json.input_type;if (inputType) {msg.inputType = inputType;}
const options = json.options ?? json.options;if (options) {msg.options = options;}
          return msg;
        },

      };

export const InputValue = {        /**
         * Serializes a InputValue to protobuf.
         */
        encode: function(inputValue: Partial<InputValue>): Uint8Array {
          return InputValue._writeMessage(inputValue, new BinaryWriter()).getResultBuffer();
        },
        

        /**
         * Deserializes a InputValue from protobuf.
         */
        decode: function(bytes: ByteSource): InputValue {
          return InputValue._readMessage(InputValue.initialize(), new BinaryReader(bytes));
        },
        

        /**
         * Serializes a InputValue to JSON.
         */
        encodeJSON: function(inputValue: Partial<InputValue>): string {
          return JSON.stringify(InputValue._writeMessageJSON(inputValue));
        },
        

        /**
         * Deserializes a InputValue from JSON.
         */
        decodeJSON: function(json: string): InputValue {
          return InputValue._readMessageJSON(InputValue.initialize(), JSON.parse(json));
        },
        

        /**
         * Initializes a InputValue with all fields set to their default value.
         */
        initialize: function(): InputValue {
          return {
            inputName: 0,
          };
        },

        /**
         * @private
         */
        _writeMessage: function(msg : Partial<InputValue>, writer: BinaryWriter): BinaryWriter {
          if (msg.inputName) {writer.writeEnum(1, msg.inputName);}
if (msg.inputValueNumber != undefined) {writer.writeInt64String(100, msg.inputValueNumber.toString());}
if (msg.inputValueString != undefined) {writer.writeString(101, msg.inputValueString);}
            return writer;
        },

        /**
         * @private
         */
        _writeMessageJSON: function(msg : Partial<InputValue>): Record<string, unknown> {
          const json: Record<string, unknown> = {};
          if (msg.inputName) {json.inputName = msg.inputName;}
if (msg.inputValueNumber != undefined) {json.inputValueNumber = msg.inputValueNumber.toString();}
if (msg.inputValueString != undefined) {json.inputValueString = msg.inputValueString;}
          return json;
        },
        
        /**
         * @private
         */
        _readMessage: function(msg: InputValue, reader: BinaryReader): InputValue {
          while (reader.nextField()) {
            const field = reader.getFieldNumber();
            switch (field) {
              case 1: {msg.inputName = reader.readEnum() as InputName;break;
}
case 100: {msg.inputValueNumber = BigInt(reader.readInt64String());break;
}
case 101: {msg.inputValueString = reader.readString();break;
}
              default: {
                reader.skipField();
                break;
              }
            }
          }
          return msg;
        },

        /**
         * @private
         */
        _readMessageJSON: function(msg: InputValue, json: any): InputValue {
          const inputName = json.inputName ?? json.input_name;if (inputName) {msg.inputName = inputName;}
const inputValueNumber = json.inputValueNumber ?? json.input_value_number;if (inputValueNumber) {msg.inputValueNumber = BigInt(inputValueNumber);}
const inputValueString = json.inputValueString ?? json.input_value_string;if (inputValueString) {msg.inputValueString = inputValueString;}
          return msg;
        },

      };


