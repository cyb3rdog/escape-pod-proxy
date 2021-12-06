/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

goog.provide('proto.cybervector.ProxyMessaage');

goog.require('jspb.BinaryReader');
goog.require('jspb.BinaryWriter');
goog.require('jspb.Map');
goog.require('jspb.Message');

goog.forwardDeclare('proto.cybervector.MessageType');

/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.cybervector.ProxyMessaage = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.cybervector.ProxyMessaage, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.cybervector.ProxyMessaage.displayName = 'proto.cybervector.ProxyMessaage';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.cybervector.ProxyMessaage.prototype.toObject = function(opt_includeInstance) {
  return proto.cybervector.ProxyMessaage.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.cybervector.ProxyMessaage} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.cybervector.ProxyMessaage.toObject = function(includeInstance, msg) {
  var f, obj = {
    messageType: jspb.Message.getFieldWithDefault(msg, 1, 0),
    messageData: jspb.Message.getFieldWithDefault(msg, 2, ""),
    intentName: jspb.Message.getFieldWithDefault(msg, 3, ""),
    parametersMap: (f = msg.getParametersMap()) ? f.toObject(includeInstance, undefined) : []
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.cybervector.ProxyMessaage}
 */
proto.cybervector.ProxyMessaage.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.cybervector.ProxyMessaage;
  return proto.cybervector.ProxyMessaage.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.cybervector.ProxyMessaage} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.cybervector.ProxyMessaage}
 */
proto.cybervector.ProxyMessaage.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!proto.cybervector.MessageType} */ (reader.readEnum());
      msg.setMessageType(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setMessageData(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setIntentName(value);
      break;
    case 4:
      var value = msg.getParametersMap();
      reader.readMessage(value, function(message, reader) {
        jspb.Map.deserializeBinary(message, reader, jspb.BinaryReader.prototype.readString, jspb.BinaryReader.prototype.readString, null, "");
         });
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.cybervector.ProxyMessaage.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.cybervector.ProxyMessaage.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.cybervector.ProxyMessaage} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.cybervector.ProxyMessaage.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getMessageType();
  if (f !== 0.0) {
    writer.writeEnum(
      1,
      f
    );
  }
  f = message.getMessageData();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getIntentName();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getParametersMap(true);
  if (f && f.getLength() > 0) {
    f.serializeBinary(4, writer, jspb.BinaryWriter.prototype.writeString, jspb.BinaryWriter.prototype.writeString);
  }
};


/**
 * optional MessageType message_type = 1;
 * @return {!proto.cybervector.MessageType}
 */
proto.cybervector.ProxyMessaage.prototype.getMessageType = function() {
  return /** @type {!proto.cybervector.MessageType} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {!proto.cybervector.MessageType} value */
proto.cybervector.ProxyMessaage.prototype.setMessageType = function(value) {
  jspb.Message.setProto3EnumField(this, 1, value);
};


/**
 * optional string message_data = 2;
 * @return {string}
 */
proto.cybervector.ProxyMessaage.prototype.getMessageData = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.cybervector.ProxyMessaage.prototype.setMessageData = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string intent_name = 3;
 * @return {string}
 */
proto.cybervector.ProxyMessaage.prototype.getIntentName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/** @param {string} value */
proto.cybervector.ProxyMessaage.prototype.setIntentName = function(value) {
  jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * map<string, string> parameters = 4;
 * @param {boolean=} opt_noLazyCreate Do not create the map if
 * empty, instead returning `undefined`
 * @return {!jspb.Map<string,string>}
 */
proto.cybervector.ProxyMessaage.prototype.getParametersMap = function(opt_noLazyCreate) {
  return /** @type {!jspb.Map<string,string>} */ (
      jspb.Message.getMapField(this, 4, opt_noLazyCreate,
      null));
};


proto.cybervector.ProxyMessaage.prototype.clearParametersMap = function() {
  this.getParametersMap().clear();
};

