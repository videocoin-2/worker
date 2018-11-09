# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: camera_cloud_internal.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from github.com.gogo.protobuf.gogoproto import gogo_pb2 as github_dot_com_dot_gogo_dot_protobuf_dot_gogoproto_dot_gogo__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='camera_cloud_internal.proto',
  package='proto',
  syntax='proto3',
  serialized_pb=_b('\n\x1b\x63\x61mera_cloud_internal.proto\x12\x05proto\x1a-github.com/gogo/protobuf/gogoproto/gogo.proto\"J\n\x15InternalCameraRequest\x12\x12\n\x02id\x18\x01 \x01(\tB\x06\xe2\xde\x1f\x02ID\x12\x1d\n\x08owner_id\x18\x02 \x01(\rB\x0b\xe2\xde\x1f\x07OwnerID\"M\n\x16InternalCameraResponse\x12\x12\n\x02id\x18\x01 \x01(\tB\x06\xe2\xde\x1f\x02ID\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x11\n\tis_on_air\x18\x03 \x01(\x08\":\n\x19InternalCameraListRequest\x12\x1d\n\x08owner_id\x18\x01 \x01(\rB\x0b\xe2\xde\x1f\x07OwnerID\"J\n\x1aInternalCameraListResponse\x12,\n\x05items\x18\x01 \x03(\x0b\x32\x1d.proto.InternalCameraResponse\"M\n\x16InternalStreamResponse\x12\x12\n\x02id\x18\x01 \x01(\rB\x06\xe2\xde\x1f\x02ID\x12\x1f\n\tcamera_id\x18\x02 \x01(\tB\x0c\xe2\xde\x1f\x08\x43\x61meraID\"[\n\x19InternalStreamListRequest\x12\x1d\n\x08owner_id\x18\x01 \x01(\rB\x0b\xe2\xde\x1f\x07OwnerID\x12\x1f\n\tcamera_id\x18\x02 \x01(\tB\x0c\xe2\xde\x1f\x08\x43\x61meraID\"J\n\x1aInternalStreamListResponse\x12,\n\x05items\x18\x01 \x03(\x0b\x32\x1d.proto.InternalStreamResponse\"^\n\x1cInternalLiveStreamingRequest\x12\x1d\n\x08owner_id\x18\x01 \x01(\rB\x0b\xe2\xde\x1f\x07OwnerID\x12\x1f\n\tcamera_id\x18\x02 \x01(\tB\x0c\xe2\xde\x1f\x08\x43\x61meraID\"\x1f\n\x1dInternalLiveStreamingResponse2\x86\x05\n\x1a\x43\x61meraCloudInternalService\x12J\n\tGetCamera\x12\x1c.proto.InternalCameraRequest\x1a\x1d.proto.InternalCameraResponse\"\x00\x12S\n\nGetCameras\x12 .proto.InternalCameraListRequest\x1a!.proto.InternalCameraListResponse\"\x00\x12Y\n\x10GetActiveStreams\x12 .proto.InternalStreamListRequest\x1a!.proto.InternalStreamListResponse\"\x00\x12\x61\n\x12StartLiveStreaming\x12#.proto.InternalLiveStreamingRequest\x1a$.proto.InternalLiveStreamingResponse\"\x00\x12`\n\x11StopLiveStreaming\x12#.proto.InternalLiveStreamingRequest\x1a$.proto.InternalLiveStreamingResponse\"\x00\x12R\n\x11MarkCameraAsOnAir\x12\x1c.proto.InternalCameraRequest\x1a\x1d.proto.InternalCameraResponse\"\x00\x12S\n\x12MarkCameraAsOffAir\x12\x1c.proto.InternalCameraRequest\x1a\x1d.proto.InternalCameraResponse\"\x00\x42\tH\x01Z\x05protob\x06proto3')
  ,
  dependencies=[github_dot_com_dot_gogo_dot_protobuf_dot_gogoproto_dot_gogo__pb2.DESCRIPTOR,])




_INTERNALCAMERAREQUEST = _descriptor.Descriptor(
  name='InternalCameraRequest',
  full_name='proto.InternalCameraRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='proto.InternalCameraRequest.id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=_descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\342\336\037\002ID'))),
    _descriptor.FieldDescriptor(
      name='owner_id', full_name='proto.InternalCameraRequest.owner_id', index=1,
      number=2, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=_descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\342\336\037\007OwnerID'))),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=85,
  serialized_end=159,
)


_INTERNALCAMERARESPONSE = _descriptor.Descriptor(
  name='InternalCameraResponse',
  full_name='proto.InternalCameraResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='proto.InternalCameraResponse.id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=_descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\342\336\037\002ID'))),
    _descriptor.FieldDescriptor(
      name='name', full_name='proto.InternalCameraResponse.name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='is_on_air', full_name='proto.InternalCameraResponse.is_on_air', index=2,
      number=3, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=161,
  serialized_end=238,
)


_INTERNALCAMERALISTREQUEST = _descriptor.Descriptor(
  name='InternalCameraListRequest',
  full_name='proto.InternalCameraListRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='owner_id', full_name='proto.InternalCameraListRequest.owner_id', index=0,
      number=1, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=_descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\342\336\037\007OwnerID'))),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=240,
  serialized_end=298,
)


_INTERNALCAMERALISTRESPONSE = _descriptor.Descriptor(
  name='InternalCameraListResponse',
  full_name='proto.InternalCameraListResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='items', full_name='proto.InternalCameraListResponse.items', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=300,
  serialized_end=374,
)


_INTERNALSTREAMRESPONSE = _descriptor.Descriptor(
  name='InternalStreamResponse',
  full_name='proto.InternalStreamResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='proto.InternalStreamResponse.id', index=0,
      number=1, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=_descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\342\336\037\002ID'))),
    _descriptor.FieldDescriptor(
      name='camera_id', full_name='proto.InternalStreamResponse.camera_id', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=_descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\342\336\037\010CameraID'))),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=376,
  serialized_end=453,
)


_INTERNALSTREAMLISTREQUEST = _descriptor.Descriptor(
  name='InternalStreamListRequest',
  full_name='proto.InternalStreamListRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='owner_id', full_name='proto.InternalStreamListRequest.owner_id', index=0,
      number=1, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=_descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\342\336\037\007OwnerID'))),
    _descriptor.FieldDescriptor(
      name='camera_id', full_name='proto.InternalStreamListRequest.camera_id', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=_descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\342\336\037\010CameraID'))),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=455,
  serialized_end=546,
)


_INTERNALSTREAMLISTRESPONSE = _descriptor.Descriptor(
  name='InternalStreamListResponse',
  full_name='proto.InternalStreamListResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='items', full_name='proto.InternalStreamListResponse.items', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=548,
  serialized_end=622,
)


_INTERNALLIVESTREAMINGREQUEST = _descriptor.Descriptor(
  name='InternalLiveStreamingRequest',
  full_name='proto.InternalLiveStreamingRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='owner_id', full_name='proto.InternalLiveStreamingRequest.owner_id', index=0,
      number=1, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=_descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\342\336\037\007OwnerID'))),
    _descriptor.FieldDescriptor(
      name='camera_id', full_name='proto.InternalLiveStreamingRequest.camera_id', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=_descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\342\336\037\010CameraID'))),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=624,
  serialized_end=718,
)


_INTERNALLIVESTREAMINGRESPONSE = _descriptor.Descriptor(
  name='InternalLiveStreamingResponse',
  full_name='proto.InternalLiveStreamingResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=720,
  serialized_end=751,
)

_INTERNALCAMERALISTRESPONSE.fields_by_name['items'].message_type = _INTERNALCAMERARESPONSE
_INTERNALSTREAMLISTRESPONSE.fields_by_name['items'].message_type = _INTERNALSTREAMRESPONSE
DESCRIPTOR.message_types_by_name['InternalCameraRequest'] = _INTERNALCAMERAREQUEST
DESCRIPTOR.message_types_by_name['InternalCameraResponse'] = _INTERNALCAMERARESPONSE
DESCRIPTOR.message_types_by_name['InternalCameraListRequest'] = _INTERNALCAMERALISTREQUEST
DESCRIPTOR.message_types_by_name['InternalCameraListResponse'] = _INTERNALCAMERALISTRESPONSE
DESCRIPTOR.message_types_by_name['InternalStreamResponse'] = _INTERNALSTREAMRESPONSE
DESCRIPTOR.message_types_by_name['InternalStreamListRequest'] = _INTERNALSTREAMLISTREQUEST
DESCRIPTOR.message_types_by_name['InternalStreamListResponse'] = _INTERNALSTREAMLISTRESPONSE
DESCRIPTOR.message_types_by_name['InternalLiveStreamingRequest'] = _INTERNALLIVESTREAMINGREQUEST
DESCRIPTOR.message_types_by_name['InternalLiveStreamingResponse'] = _INTERNALLIVESTREAMINGRESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

InternalCameraRequest = _reflection.GeneratedProtocolMessageType('InternalCameraRequest', (_message.Message,), dict(
  DESCRIPTOR = _INTERNALCAMERAREQUEST,
  __module__ = 'camera_cloud_internal_pb2'
  # @@protoc_insertion_point(class_scope:proto.InternalCameraRequest)
  ))
_sym_db.RegisterMessage(InternalCameraRequest)

InternalCameraResponse = _reflection.GeneratedProtocolMessageType('InternalCameraResponse', (_message.Message,), dict(
  DESCRIPTOR = _INTERNALCAMERARESPONSE,
  __module__ = 'camera_cloud_internal_pb2'
  # @@protoc_insertion_point(class_scope:proto.InternalCameraResponse)
  ))
_sym_db.RegisterMessage(InternalCameraResponse)

InternalCameraListRequest = _reflection.GeneratedProtocolMessageType('InternalCameraListRequest', (_message.Message,), dict(
  DESCRIPTOR = _INTERNALCAMERALISTREQUEST,
  __module__ = 'camera_cloud_internal_pb2'
  # @@protoc_insertion_point(class_scope:proto.InternalCameraListRequest)
  ))
_sym_db.RegisterMessage(InternalCameraListRequest)

InternalCameraListResponse = _reflection.GeneratedProtocolMessageType('InternalCameraListResponse', (_message.Message,), dict(
  DESCRIPTOR = _INTERNALCAMERALISTRESPONSE,
  __module__ = 'camera_cloud_internal_pb2'
  # @@protoc_insertion_point(class_scope:proto.InternalCameraListResponse)
  ))
_sym_db.RegisterMessage(InternalCameraListResponse)

InternalStreamResponse = _reflection.GeneratedProtocolMessageType('InternalStreamResponse', (_message.Message,), dict(
  DESCRIPTOR = _INTERNALSTREAMRESPONSE,
  __module__ = 'camera_cloud_internal_pb2'
  # @@protoc_insertion_point(class_scope:proto.InternalStreamResponse)
  ))
_sym_db.RegisterMessage(InternalStreamResponse)

InternalStreamListRequest = _reflection.GeneratedProtocolMessageType('InternalStreamListRequest', (_message.Message,), dict(
  DESCRIPTOR = _INTERNALSTREAMLISTREQUEST,
  __module__ = 'camera_cloud_internal_pb2'
  # @@protoc_insertion_point(class_scope:proto.InternalStreamListRequest)
  ))
_sym_db.RegisterMessage(InternalStreamListRequest)

InternalStreamListResponse = _reflection.GeneratedProtocolMessageType('InternalStreamListResponse', (_message.Message,), dict(
  DESCRIPTOR = _INTERNALSTREAMLISTRESPONSE,
  __module__ = 'camera_cloud_internal_pb2'
  # @@protoc_insertion_point(class_scope:proto.InternalStreamListResponse)
  ))
_sym_db.RegisterMessage(InternalStreamListResponse)

InternalLiveStreamingRequest = _reflection.GeneratedProtocolMessageType('InternalLiveStreamingRequest', (_message.Message,), dict(
  DESCRIPTOR = _INTERNALLIVESTREAMINGREQUEST,
  __module__ = 'camera_cloud_internal_pb2'
  # @@protoc_insertion_point(class_scope:proto.InternalLiveStreamingRequest)
  ))
_sym_db.RegisterMessage(InternalLiveStreamingRequest)

InternalLiveStreamingResponse = _reflection.GeneratedProtocolMessageType('InternalLiveStreamingResponse', (_message.Message,), dict(
  DESCRIPTOR = _INTERNALLIVESTREAMINGRESPONSE,
  __module__ = 'camera_cloud_internal_pb2'
  # @@protoc_insertion_point(class_scope:proto.InternalLiveStreamingResponse)
  ))
_sym_db.RegisterMessage(InternalLiveStreamingResponse)


DESCRIPTOR.has_options = True
DESCRIPTOR._options = _descriptor._ParseOptions(descriptor_pb2.FileOptions(), _b('H\001Z\005proto'))
_INTERNALCAMERAREQUEST.fields_by_name['id'].has_options = True
_INTERNALCAMERAREQUEST.fields_by_name['id']._options = _descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\342\336\037\002ID'))
_INTERNALCAMERAREQUEST.fields_by_name['owner_id'].has_options = True
_INTERNALCAMERAREQUEST.fields_by_name['owner_id']._options = _descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\342\336\037\007OwnerID'))
_INTERNALCAMERARESPONSE.fields_by_name['id'].has_options = True
_INTERNALCAMERARESPONSE.fields_by_name['id']._options = _descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\342\336\037\002ID'))
_INTERNALCAMERALISTREQUEST.fields_by_name['owner_id'].has_options = True
_INTERNALCAMERALISTREQUEST.fields_by_name['owner_id']._options = _descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\342\336\037\007OwnerID'))
_INTERNALSTREAMRESPONSE.fields_by_name['id'].has_options = True
_INTERNALSTREAMRESPONSE.fields_by_name['id']._options = _descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\342\336\037\002ID'))
_INTERNALSTREAMRESPONSE.fields_by_name['camera_id'].has_options = True
_INTERNALSTREAMRESPONSE.fields_by_name['camera_id']._options = _descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\342\336\037\010CameraID'))
_INTERNALSTREAMLISTREQUEST.fields_by_name['owner_id'].has_options = True
_INTERNALSTREAMLISTREQUEST.fields_by_name['owner_id']._options = _descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\342\336\037\007OwnerID'))
_INTERNALSTREAMLISTREQUEST.fields_by_name['camera_id'].has_options = True
_INTERNALSTREAMLISTREQUEST.fields_by_name['camera_id']._options = _descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\342\336\037\010CameraID'))
_INTERNALLIVESTREAMINGREQUEST.fields_by_name['owner_id'].has_options = True
_INTERNALLIVESTREAMINGREQUEST.fields_by_name['owner_id']._options = _descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\342\336\037\007OwnerID'))
_INTERNALLIVESTREAMINGREQUEST.fields_by_name['camera_id'].has_options = True
_INTERNALLIVESTREAMINGREQUEST.fields_by_name['camera_id']._options = _descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\342\336\037\010CameraID'))
try:
  # THESE ELEMENTS WILL BE DEPRECATED.
  # Please use the generated *_pb2_grpc.py files instead.
  import grpc
  from grpc.beta import implementations as beta_implementations
  from grpc.beta import interfaces as beta_interfaces
  from grpc.framework.common import cardinality
  from grpc.framework.interfaces.face import utilities as face_utilities


  class CameraCloudInternalServiceStub(object):
    # missing associated documentation comment in .proto file
    pass

    def __init__(self, channel):
      """Constructor.

      Args:
        channel: A grpc.Channel.
      """
      self.GetCamera = channel.unary_unary(
          '/proto.CameraCloudInternalService/GetCamera',
          request_serializer=InternalCameraRequest.SerializeToString,
          response_deserializer=InternalCameraResponse.FromString,
          )
      self.GetCameras = channel.unary_unary(
          '/proto.CameraCloudInternalService/GetCameras',
          request_serializer=InternalCameraListRequest.SerializeToString,
          response_deserializer=InternalCameraListResponse.FromString,
          )
      self.GetActiveStreams = channel.unary_unary(
          '/proto.CameraCloudInternalService/GetActiveStreams',
          request_serializer=InternalStreamListRequest.SerializeToString,
          response_deserializer=InternalStreamListResponse.FromString,
          )
      self.StartLiveStreaming = channel.unary_unary(
          '/proto.CameraCloudInternalService/StartLiveStreaming',
          request_serializer=InternalLiveStreamingRequest.SerializeToString,
          response_deserializer=InternalLiveStreamingResponse.FromString,
          )
      self.StopLiveStreaming = channel.unary_unary(
          '/proto.CameraCloudInternalService/StopLiveStreaming',
          request_serializer=InternalLiveStreamingRequest.SerializeToString,
          response_deserializer=InternalLiveStreamingResponse.FromString,
          )
      self.MarkCameraAsOnAir = channel.unary_unary(
          '/proto.CameraCloudInternalService/MarkCameraAsOnAir',
          request_serializer=InternalCameraRequest.SerializeToString,
          response_deserializer=InternalCameraResponse.FromString,
          )
      self.MarkCameraAsOffAir = channel.unary_unary(
          '/proto.CameraCloudInternalService/MarkCameraAsOffAir',
          request_serializer=InternalCameraRequest.SerializeToString,
          response_deserializer=InternalCameraResponse.FromString,
          )


  class CameraCloudInternalServiceServicer(object):
    # missing associated documentation comment in .proto file
    pass

    def GetCamera(self, request, context):
      # missing associated documentation comment in .proto file
      pass
      context.set_code(grpc.StatusCode.UNIMPLEMENTED)
      context.set_details('Method not implemented!')
      raise NotImplementedError('Method not implemented!')

    def GetCameras(self, request, context):
      # missing associated documentation comment in .proto file
      pass
      context.set_code(grpc.StatusCode.UNIMPLEMENTED)
      context.set_details('Method not implemented!')
      raise NotImplementedError('Method not implemented!')

    def GetActiveStreams(self, request, context):
      # missing associated documentation comment in .proto file
      pass
      context.set_code(grpc.StatusCode.UNIMPLEMENTED)
      context.set_details('Method not implemented!')
      raise NotImplementedError('Method not implemented!')

    def StartLiveStreaming(self, request, context):
      # missing associated documentation comment in .proto file
      pass
      context.set_code(grpc.StatusCode.UNIMPLEMENTED)
      context.set_details('Method not implemented!')
      raise NotImplementedError('Method not implemented!')

    def StopLiveStreaming(self, request, context):
      # missing associated documentation comment in .proto file
      pass
      context.set_code(grpc.StatusCode.UNIMPLEMENTED)
      context.set_details('Method not implemented!')
      raise NotImplementedError('Method not implemented!')

    def MarkCameraAsOnAir(self, request, context):
      # missing associated documentation comment in .proto file
      pass
      context.set_code(grpc.StatusCode.UNIMPLEMENTED)
      context.set_details('Method not implemented!')
      raise NotImplementedError('Method not implemented!')

    def MarkCameraAsOffAir(self, request, context):
      # missing associated documentation comment in .proto file
      pass
      context.set_code(grpc.StatusCode.UNIMPLEMENTED)
      context.set_details('Method not implemented!')
      raise NotImplementedError('Method not implemented!')


  def add_CameraCloudInternalServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
        'GetCamera': grpc.unary_unary_rpc_method_handler(
            servicer.GetCamera,
            request_deserializer=InternalCameraRequest.FromString,
            response_serializer=InternalCameraResponse.SerializeToString,
        ),
        'GetCameras': grpc.unary_unary_rpc_method_handler(
            servicer.GetCameras,
            request_deserializer=InternalCameraListRequest.FromString,
            response_serializer=InternalCameraListResponse.SerializeToString,
        ),
        'GetActiveStreams': grpc.unary_unary_rpc_method_handler(
            servicer.GetActiveStreams,
            request_deserializer=InternalStreamListRequest.FromString,
            response_serializer=InternalStreamListResponse.SerializeToString,
        ),
        'StartLiveStreaming': grpc.unary_unary_rpc_method_handler(
            servicer.StartLiveStreaming,
            request_deserializer=InternalLiveStreamingRequest.FromString,
            response_serializer=InternalLiveStreamingResponse.SerializeToString,
        ),
        'StopLiveStreaming': grpc.unary_unary_rpc_method_handler(
            servicer.StopLiveStreaming,
            request_deserializer=InternalLiveStreamingRequest.FromString,
            response_serializer=InternalLiveStreamingResponse.SerializeToString,
        ),
        'MarkCameraAsOnAir': grpc.unary_unary_rpc_method_handler(
            servicer.MarkCameraAsOnAir,
            request_deserializer=InternalCameraRequest.FromString,
            response_serializer=InternalCameraResponse.SerializeToString,
        ),
        'MarkCameraAsOffAir': grpc.unary_unary_rpc_method_handler(
            servicer.MarkCameraAsOffAir,
            request_deserializer=InternalCameraRequest.FromString,
            response_serializer=InternalCameraResponse.SerializeToString,
        ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
        'proto.CameraCloudInternalService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


  class BetaCameraCloudInternalServiceServicer(object):
    """The Beta API is deprecated for 0.15.0 and later.

    It is recommended to use the GA API (classes and functions in this
    file not marked beta) for all further purposes. This class was generated
    only to ease transition from grpcio<0.15.0 to grpcio>=0.15.0."""
    # missing associated documentation comment in .proto file
    pass
    def GetCamera(self, request, context):
      # missing associated documentation comment in .proto file
      pass
      context.code(beta_interfaces.StatusCode.UNIMPLEMENTED)
    def GetCameras(self, request, context):
      # missing associated documentation comment in .proto file
      pass
      context.code(beta_interfaces.StatusCode.UNIMPLEMENTED)
    def GetActiveStreams(self, request, context):
      # missing associated documentation comment in .proto file
      pass
      context.code(beta_interfaces.StatusCode.UNIMPLEMENTED)
    def StartLiveStreaming(self, request, context):
      # missing associated documentation comment in .proto file
      pass
      context.code(beta_interfaces.StatusCode.UNIMPLEMENTED)
    def StopLiveStreaming(self, request, context):
      # missing associated documentation comment in .proto file
      pass
      context.code(beta_interfaces.StatusCode.UNIMPLEMENTED)
    def MarkCameraAsOnAir(self, request, context):
      # missing associated documentation comment in .proto file
      pass
      context.code(beta_interfaces.StatusCode.UNIMPLEMENTED)
    def MarkCameraAsOffAir(self, request, context):
      # missing associated documentation comment in .proto file
      pass
      context.code(beta_interfaces.StatusCode.UNIMPLEMENTED)


  class BetaCameraCloudInternalServiceStub(object):
    """The Beta API is deprecated for 0.15.0 and later.

    It is recommended to use the GA API (classes and functions in this
    file not marked beta) for all further purposes. This class was generated
    only to ease transition from grpcio<0.15.0 to grpcio>=0.15.0."""
    # missing associated documentation comment in .proto file
    pass
    def GetCamera(self, request, timeout, metadata=None, with_call=False, protocol_options=None):
      # missing associated documentation comment in .proto file
      pass
      raise NotImplementedError()
    GetCamera.future = None
    def GetCameras(self, request, timeout, metadata=None, with_call=False, protocol_options=None):
      # missing associated documentation comment in .proto file
      pass
      raise NotImplementedError()
    GetCameras.future = None
    def GetActiveStreams(self, request, timeout, metadata=None, with_call=False, protocol_options=None):
      # missing associated documentation comment in .proto file
      pass
      raise NotImplementedError()
    GetActiveStreams.future = None
    def StartLiveStreaming(self, request, timeout, metadata=None, with_call=False, protocol_options=None):
      # missing associated documentation comment in .proto file
      pass
      raise NotImplementedError()
    StartLiveStreaming.future = None
    def StopLiveStreaming(self, request, timeout, metadata=None, with_call=False, protocol_options=None):
      # missing associated documentation comment in .proto file
      pass
      raise NotImplementedError()
    StopLiveStreaming.future = None
    def MarkCameraAsOnAir(self, request, timeout, metadata=None, with_call=False, protocol_options=None):
      # missing associated documentation comment in .proto file
      pass
      raise NotImplementedError()
    MarkCameraAsOnAir.future = None
    def MarkCameraAsOffAir(self, request, timeout, metadata=None, with_call=False, protocol_options=None):
      # missing associated documentation comment in .proto file
      pass
      raise NotImplementedError()
    MarkCameraAsOffAir.future = None


  def beta_create_CameraCloudInternalService_server(servicer, pool=None, pool_size=None, default_timeout=None, maximum_timeout=None):
    """The Beta API is deprecated for 0.15.0 and later.

    It is recommended to use the GA API (classes and functions in this
    file not marked beta) for all further purposes. This function was
    generated only to ease transition from grpcio<0.15.0 to grpcio>=0.15.0"""
    request_deserializers = {
      ('proto.CameraCloudInternalService', 'GetActiveStreams'): InternalStreamListRequest.FromString,
      ('proto.CameraCloudInternalService', 'GetCamera'): InternalCameraRequest.FromString,
      ('proto.CameraCloudInternalService', 'GetCameras'): InternalCameraListRequest.FromString,
      ('proto.CameraCloudInternalService', 'MarkCameraAsOffAir'): InternalCameraRequest.FromString,
      ('proto.CameraCloudInternalService', 'MarkCameraAsOnAir'): InternalCameraRequest.FromString,
      ('proto.CameraCloudInternalService', 'StartLiveStreaming'): InternalLiveStreamingRequest.FromString,
      ('proto.CameraCloudInternalService', 'StopLiveStreaming'): InternalLiveStreamingRequest.FromString,
    }
    response_serializers = {
      ('proto.CameraCloudInternalService', 'GetActiveStreams'): InternalStreamListResponse.SerializeToString,
      ('proto.CameraCloudInternalService', 'GetCamera'): InternalCameraResponse.SerializeToString,
      ('proto.CameraCloudInternalService', 'GetCameras'): InternalCameraListResponse.SerializeToString,
      ('proto.CameraCloudInternalService', 'MarkCameraAsOffAir'): InternalCameraResponse.SerializeToString,
      ('proto.CameraCloudInternalService', 'MarkCameraAsOnAir'): InternalCameraResponse.SerializeToString,
      ('proto.CameraCloudInternalService', 'StartLiveStreaming'): InternalLiveStreamingResponse.SerializeToString,
      ('proto.CameraCloudInternalService', 'StopLiveStreaming'): InternalLiveStreamingResponse.SerializeToString,
    }
    method_implementations = {
      ('proto.CameraCloudInternalService', 'GetActiveStreams'): face_utilities.unary_unary_inline(servicer.GetActiveStreams),
      ('proto.CameraCloudInternalService', 'GetCamera'): face_utilities.unary_unary_inline(servicer.GetCamera),
      ('proto.CameraCloudInternalService', 'GetCameras'): face_utilities.unary_unary_inline(servicer.GetCameras),
      ('proto.CameraCloudInternalService', 'MarkCameraAsOffAir'): face_utilities.unary_unary_inline(servicer.MarkCameraAsOffAir),
      ('proto.CameraCloudInternalService', 'MarkCameraAsOnAir'): face_utilities.unary_unary_inline(servicer.MarkCameraAsOnAir),
      ('proto.CameraCloudInternalService', 'StartLiveStreaming'): face_utilities.unary_unary_inline(servicer.StartLiveStreaming),
      ('proto.CameraCloudInternalService', 'StopLiveStreaming'): face_utilities.unary_unary_inline(servicer.StopLiveStreaming),
    }
    server_options = beta_implementations.server_options(request_deserializers=request_deserializers, response_serializers=response_serializers, thread_pool=pool, thread_pool_size=pool_size, default_timeout=default_timeout, maximum_timeout=maximum_timeout)
    return beta_implementations.server(method_implementations, options=server_options)


  def beta_create_CameraCloudInternalService_stub(channel, host=None, metadata_transformer=None, pool=None, pool_size=None):
    """The Beta API is deprecated for 0.15.0 and later.

    It is recommended to use the GA API (classes and functions in this
    file not marked beta) for all further purposes. This function was
    generated only to ease transition from grpcio<0.15.0 to grpcio>=0.15.0"""
    request_serializers = {
      ('proto.CameraCloudInternalService', 'GetActiveStreams'): InternalStreamListRequest.SerializeToString,
      ('proto.CameraCloudInternalService', 'GetCamera'): InternalCameraRequest.SerializeToString,
      ('proto.CameraCloudInternalService', 'GetCameras'): InternalCameraListRequest.SerializeToString,
      ('proto.CameraCloudInternalService', 'MarkCameraAsOffAir'): InternalCameraRequest.SerializeToString,
      ('proto.CameraCloudInternalService', 'MarkCameraAsOnAir'): InternalCameraRequest.SerializeToString,
      ('proto.CameraCloudInternalService', 'StartLiveStreaming'): InternalLiveStreamingRequest.SerializeToString,
      ('proto.CameraCloudInternalService', 'StopLiveStreaming'): InternalLiveStreamingRequest.SerializeToString,
    }
    response_deserializers = {
      ('proto.CameraCloudInternalService', 'GetActiveStreams'): InternalStreamListResponse.FromString,
      ('proto.CameraCloudInternalService', 'GetCamera'): InternalCameraResponse.FromString,
      ('proto.CameraCloudInternalService', 'GetCameras'): InternalCameraListResponse.FromString,
      ('proto.CameraCloudInternalService', 'MarkCameraAsOffAir'): InternalCameraResponse.FromString,
      ('proto.CameraCloudInternalService', 'MarkCameraAsOnAir'): InternalCameraResponse.FromString,
      ('proto.CameraCloudInternalService', 'StartLiveStreaming'): InternalLiveStreamingResponse.FromString,
      ('proto.CameraCloudInternalService', 'StopLiveStreaming'): InternalLiveStreamingResponse.FromString,
    }
    cardinalities = {
      'GetActiveStreams': cardinality.Cardinality.UNARY_UNARY,
      'GetCamera': cardinality.Cardinality.UNARY_UNARY,
      'GetCameras': cardinality.Cardinality.UNARY_UNARY,
      'MarkCameraAsOffAir': cardinality.Cardinality.UNARY_UNARY,
      'MarkCameraAsOnAir': cardinality.Cardinality.UNARY_UNARY,
      'StartLiveStreaming': cardinality.Cardinality.UNARY_UNARY,
      'StopLiveStreaming': cardinality.Cardinality.UNARY_UNARY,
    }
    stub_options = beta_implementations.stub_options(host=host, metadata_transformer=metadata_transformer, request_serializers=request_serializers, response_deserializers=response_deserializers, thread_pool=pool, thread_pool_size=pool_size)
    return beta_implementations.dynamic_stub(channel, 'proto.CameraCloudInternalService', cardinalities, options=stub_options)
except ImportError:
  pass
# @@protoc_insertion_point(module_scope)
