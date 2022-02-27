// THIS IS AN AUTOGENERATED FILE. DO NOT EDIT THIS FILE DIRECTLY.
// Source: protos/common/v1/geopoint.proto


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
  
  
export interface GeoPoint {
lon: number;
lat: number;
}


  //========================================//
  //        Protobuf Encode / Decode        //
  //========================================//
  
  
export const GeoPoint = {        /**
         * Serializes a GeoPoint to protobuf.
         */
        encode: function(geoPoint: Partial<GeoPoint>): Uint8Array {
          return GeoPoint._writeMessage(geoPoint, new BinaryWriter()).getResultBuffer();
        },
        

        /**
         * Deserializes a GeoPoint from protobuf.
         */
        decode: function(bytes: ByteSource): GeoPoint {
          return GeoPoint._readMessage(GeoPoint.initialize(), new BinaryReader(bytes));
        },
        

        /**
         * Serializes a GeoPoint to JSON.
         */
        encodeJSON: function(geoPoint: Partial<GeoPoint>): string {
          return JSON.stringify(GeoPoint._writeMessageJSON(geoPoint));
        },
        

        /**
         * Deserializes a GeoPoint from JSON.
         */
        decodeJSON: function(json: string): GeoPoint {
          return GeoPoint._readMessageJSON(GeoPoint.initialize(), JSON.parse(json));
        },
        

        /**
         * Initializes a GeoPoint with all fields set to their default value.
         */
        initialize: function(): GeoPoint {
          return {
            lon: 0,lat: 0,
          };
        },

        /**
         * @private
         */
        _writeMessage: function(msg : Partial<GeoPoint>, writer: BinaryWriter): BinaryWriter {
          if (msg.lon) {writer.writeDouble(1, msg.lon);}
if (msg.lat) {writer.writeDouble(2, msg.lat);}
            return writer;
        },

        /**
         * @private
         */
        _writeMessageJSON: function(msg : Partial<GeoPoint>): Record<string, unknown> {
          const json: Record<string, unknown> = {};
          if (msg.lon) {json.lon = msg.lon;}
if (msg.lat) {json.lat = msg.lat;}
          return json;
        },
        
        /**
         * @private
         */
        _readMessage: function(msg: GeoPoint, reader: BinaryReader): GeoPoint {
          while (reader.nextField()) {
            const field = reader.getFieldNumber();
            switch (field) {
              case 1: {msg.lon = reader.readDouble();break;
}
case 2: {msg.lat = reader.readDouble();break;
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
        _readMessageJSON: function(msg: GeoPoint, json: any): GeoPoint {
          const lon = json.lon ?? json.lon;if (lon) {msg.lon = lon;}
const lat = json.lat ?? json.lat;if (lat) {msg.lat = lat;}
          return msg;
        },

      };


