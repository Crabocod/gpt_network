# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: service.proto
# Protobuf Python Version: 5.28.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    5,
    28,
    1,
    '',
    'service.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\rservice.proto\x12\x05proto\"7\n\x0fGenerateRequest\x12\x10\n\x08question\x18\x01 \x01(\t\x12\x12\n\nmodel_name\x18\x02 \x01(\t\"\"\n\x10GenerateResponse\x12\x0e\n\x06\x61nswer\x18\x01 \x01(\t\"%\n\x0bSaveRequest\x12\x16\n\x0egenerated_text\x18\x01 \x01(\t\"\x1f\n\x0cSaveResponse\x12\x0f\n\x07success\x18\x01 \x01(\x08\x32Q\n\x0eTextGenService\x12?\n\x0cGenerateText\x12\x16.proto.GenerateRequest\x1a\x17.proto.GenerateResponse2O\n\x0fSaveTextService\x12<\n\x11SaveGeneratedText\x12\x12.proto.SaveRequest\x1a\x13.proto.SaveResponseB\x0cZ\n./go/protob\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'service_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\n./go/proto'
  _globals['_GENERATEREQUEST']._serialized_start=24
  _globals['_GENERATEREQUEST']._serialized_end=79
  _globals['_GENERATERESPONSE']._serialized_start=81
  _globals['_GENERATERESPONSE']._serialized_end=115
  _globals['_SAVEREQUEST']._serialized_start=117
  _globals['_SAVEREQUEST']._serialized_end=154
  _globals['_SAVERESPONSE']._serialized_start=156
  _globals['_SAVERESPONSE']._serialized_end=187
  _globals['_TEXTGENSERVICE']._serialized_start=189
  _globals['_TEXTGENSERVICE']._serialized_end=270
  _globals['_SAVETEXTSERVICE']._serialized_start=272
  _globals['_SAVETEXTSERVICE']._serialized_end=351
# @@protoc_insertion_point(module_scope)
