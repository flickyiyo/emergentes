// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: app/src/main/proto/imgstream/imgstream.proto

package com.proto.imgstream;

/**
 * Protobuf type {@code imgstream.CloseStreamResponse}
 */
public  final class CloseStreamResponse extends
    com.google.protobuf.GeneratedMessageLite<
        CloseStreamResponse, CloseStreamResponse.Builder> implements
    // @@protoc_insertion_point(message_implements:imgstream.CloseStreamResponse)
    CloseStreamResponseOrBuilder {
  private CloseStreamResponse() {
    msg_ = "";
  }
  public static final int MSG_FIELD_NUMBER = 1;
  private java.lang.String msg_;
  /**
   * <code>string msg = 1;</code>
   */
  @java.lang.Override
  public java.lang.String getMsg() {
    return msg_;
  }
  /**
   * <code>string msg = 1;</code>
   */
  @java.lang.Override
  public com.google.protobuf.ByteString
      getMsgBytes() {
    return com.google.protobuf.ByteString.copyFromUtf8(msg_);
  }
  /**
   * <code>string msg = 1;</code>
   */
  private void setMsg(
      java.lang.String value) {
    if (value == null) {
    throw new NullPointerException();
  }
  
    msg_ = value;
  }
  /**
   * <code>string msg = 1;</code>
   */
  private void clearMsg() {
    
    msg_ = getDefaultInstance().getMsg();
  }
  /**
   * <code>string msg = 1;</code>
   */
  private void setMsgBytes(
      com.google.protobuf.ByteString value) {
    if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
    
    msg_ = value.toStringUtf8();
  }

  public static com.proto.imgstream.CloseStreamResponse parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return com.google.protobuf.GeneratedMessageLite.parseFrom(
        DEFAULT_INSTANCE, data);
  }
  public static com.proto.imgstream.CloseStreamResponse parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return com.google.protobuf.GeneratedMessageLite.parseFrom(
        DEFAULT_INSTANCE, data, extensionRegistry);
  }
  public static com.proto.imgstream.CloseStreamResponse parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return com.google.protobuf.GeneratedMessageLite.parseFrom(
        DEFAULT_INSTANCE, data);
  }
  public static com.proto.imgstream.CloseStreamResponse parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return com.google.protobuf.GeneratedMessageLite.parseFrom(
        DEFAULT_INSTANCE, data, extensionRegistry);
  }
  public static com.proto.imgstream.CloseStreamResponse parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return com.google.protobuf.GeneratedMessageLite.parseFrom(
        DEFAULT_INSTANCE, data);
  }
  public static com.proto.imgstream.CloseStreamResponse parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return com.google.protobuf.GeneratedMessageLite.parseFrom(
        DEFAULT_INSTANCE, data, extensionRegistry);
  }
  public static com.proto.imgstream.CloseStreamResponse parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageLite.parseFrom(
        DEFAULT_INSTANCE, input);
  }
  public static com.proto.imgstream.CloseStreamResponse parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageLite.parseFrom(
        DEFAULT_INSTANCE, input, extensionRegistry);
  }
  public static com.proto.imgstream.CloseStreamResponse parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return parseDelimitedFrom(DEFAULT_INSTANCE, input);
  }
  public static com.proto.imgstream.CloseStreamResponse parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return parseDelimitedFrom(DEFAULT_INSTANCE, input, extensionRegistry);
  }
  public static com.proto.imgstream.CloseStreamResponse parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageLite.parseFrom(
        DEFAULT_INSTANCE, input);
  }
  public static com.proto.imgstream.CloseStreamResponse parseFrom(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageLite.parseFrom(
        DEFAULT_INSTANCE, input, extensionRegistry);
  }

  public static Builder newBuilder() {
    return (Builder) DEFAULT_INSTANCE.createBuilder();
  }
  public static Builder newBuilder(com.proto.imgstream.CloseStreamResponse prototype) {
    return (Builder) DEFAULT_INSTANCE.createBuilder(prototype);
  }

  /**
   * Protobuf type {@code imgstream.CloseStreamResponse}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageLite.Builder<
        com.proto.imgstream.CloseStreamResponse, Builder> implements
      // @@protoc_insertion_point(builder_implements:imgstream.CloseStreamResponse)
      com.proto.imgstream.CloseStreamResponseOrBuilder {
    // Construct using com.proto.imgstream.CloseStreamResponse.newBuilder()
    private Builder() {
      super(DEFAULT_INSTANCE);
    }


    /**
     * <code>string msg = 1;</code>
     */
    @java.lang.Override
    public java.lang.String getMsg() {
      return instance.getMsg();
    }
    /**
     * <code>string msg = 1;</code>
     */
    @java.lang.Override
    public com.google.protobuf.ByteString
        getMsgBytes() {
      return instance.getMsgBytes();
    }
    /**
     * <code>string msg = 1;</code>
     */
    public Builder setMsg(
        java.lang.String value) {
      copyOnWrite();
      instance.setMsg(value);
      return this;
    }
    /**
     * <code>string msg = 1;</code>
     */
    public Builder clearMsg() {
      copyOnWrite();
      instance.clearMsg();
      return this;
    }
    /**
     * <code>string msg = 1;</code>
     */
    public Builder setMsgBytes(
        com.google.protobuf.ByteString value) {
      copyOnWrite();
      instance.setMsgBytes(value);
      return this;
    }

    // @@protoc_insertion_point(builder_scope:imgstream.CloseStreamResponse)
  }
  @java.lang.Override
  @java.lang.SuppressWarnings({"unchecked", "fallthrough"})
  protected final java.lang.Object dynamicMethod(
      com.google.protobuf.GeneratedMessageLite.MethodToInvoke method,
      java.lang.Object arg0, java.lang.Object arg1) {
    switch (method) {
      case NEW_MUTABLE_INSTANCE: {
        return new com.proto.imgstream.CloseStreamResponse();
      }
      case NEW_BUILDER: {
        return new Builder();
      }
      case BUILD_MESSAGE_INFO: {
          java.lang.Object[] objects = new java.lang.Object[] {
            "msg_",
          };
          java.lang.String info =
              "\u0000\u0001\u0000\u0000\u0001\u0001\u0001\u0000\u0000\u0000\u0001\u0208";
          return newMessageInfo(DEFAULT_INSTANCE, info, objects);
      }
      // fall through
      case GET_DEFAULT_INSTANCE: {
        return DEFAULT_INSTANCE;
      }
      case GET_PARSER: {
        com.google.protobuf.Parser<com.proto.imgstream.CloseStreamResponse> parser = PARSER;
        if (parser == null) {
          synchronized (com.proto.imgstream.CloseStreamResponse.class) {
            parser = PARSER;
            if (parser == null) {
              parser =
                  new DefaultInstanceBasedParser<com.proto.imgstream.CloseStreamResponse>(
                      DEFAULT_INSTANCE);
              PARSER = parser;
            }
          }
        }
        return parser;
    }
    case GET_MEMOIZED_IS_INITIALIZED: {
      return (byte) 1;
    }
    case SET_MEMOIZED_IS_INITIALIZED: {
      return null;
    }
    }
    throw new UnsupportedOperationException();
  }


  // @@protoc_insertion_point(class_scope:imgstream.CloseStreamResponse)
  private static final com.proto.imgstream.CloseStreamResponse DEFAULT_INSTANCE;
  static {
    CloseStreamResponse defaultInstance = new CloseStreamResponse();
    // New instances are implicitly immutable so no need to make
    // immutable.
    DEFAULT_INSTANCE = defaultInstance;
    com.google.protobuf.GeneratedMessageLite.registerDefaultInstance(
      CloseStreamResponse.class, defaultInstance);
  }

  public static com.proto.imgstream.CloseStreamResponse getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static volatile com.google.protobuf.Parser<CloseStreamResponse> PARSER;

  public static com.google.protobuf.Parser<CloseStreamResponse> parser() {
    return DEFAULT_INSTANCE.getParserForType();
  }
}
