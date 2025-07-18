import base64
import urllib.parse
from google.protobuf import descriptor_pb2, message_factory
from google.protobuf.internal import decoder

# This part defines Google's OTP Migration payload
# Proto file structure (manually defined for decoding)
# Based on: https://github.com/google/google-authenticator-export/blob/master/src/main/proto/migration_payload.proto

from google.protobuf import descriptor_pool
from google.protobuf.message_factory import MessageFactory

# Define the OTP parameters message
pool = descriptor_pool.Default()
file_desc_proto = descriptor_pb2.FileDescriptorProto()
file_desc_proto.name = "otp.proto"
file_desc_proto.package = "otp"

file_desc_proto.message_type.add(name="MigrationPayload")
migration_payload = file_desc_proto.message_type[0]

migration_payload.field.add(
    name="otp_parameters", number=1, label=3, type=11, type_name="otp.MigrationPayload.OtpParameters"
)
migration_payload.field.add(name="version", number=2, label=1, type=5)
migration_payload.field.add(name="batch_size", number=3, label=1, type=5)
migration_payload.field.add(name="batch_index", number=4, label=1, type=5)
migration_payload.field.add(name="batch_id", number=5, label=1, type=12)

otp_params = migration_payload.nested_type.add(name="OtpParameters")
otp_params.field.add(name="secret", number=1, label=1, type=12)
otp_params.field.add(name="name", number=2, label=1, type=9)
otp_params.field.add(name="issuer", number=3, label=1, type=9)
otp_params.field.add(name="algorithm", number=4, label=1, type=14)
otp_params.field.add(name="digits", number=5, label=1, type=14)
otp_params.field.add(name="type", number=6, label=1, type=14)
otp_params.field.add(name="counter", number=7, label=1, type=4)

for enum_name in ["Algorithm", "DigitCount", "OtpType"]:
    enum = otp_params.enum_type.add(name=enum_name)
    for i, label in enumerate(["ALGO_SHA1", "ALGO_SHA256", "ALGO_SHA512", "ALGO_MD5"] if enum_name == "Algorithm" else ["DIGIT6", "DIGIT8"] if enum_name == "DigitCount" else ["TOTP", "HOTP"]):
        enum.value.add(name=label, number=i)

desc = pool.Add(file_desc_proto)
msg_class = MessageFactory(pool).GetPrototype(desc.message_types_by_name["MigrationPayload"])

def parse_migration_uri(uri: str):
    parsed = urllib.parse.urlparse(uri)
    if not parsed.scheme.startswith("otpauth-migration"):
        raise ValueError("Invalid scheme")
    data = urllib.parse.parse_qs(parsed.query)["data"][0]
    raw = base64.urlsafe_b64decode(data + '==')  # Ensure correct padding
    message = msg_class()
    message.ParseFromString(raw)

    for param in message.otp_parameters:
        name = param.name
        issuer = param.issuer
        secret = base64.b32encode(param.secret).decode("utf-8").replace("=", "")
        algo = param.algorithm
        digits = [6, 8][param.digits] if param.digits < 2 else 6
        type_ = ["totp", "hotp"][param.type] if param.type < 2 else "totp"
        print(f"Account: {issuer + ' - ' if issuer else ''}{name}")
        print(f"Secret: {secret}")
        print(f"Type: {type_.upper()}, Digits: {digits}, Algorithm: {algo}")
        print(f"otpauth://totp/{urllib.parse.quote(issuer + ':' + name)}?secret={secret}&issuer={urllib.parse.quote(issuer)}")
        print("-" * 40)

# Replace this with your own QR data:
MIGRATION_URI = "otpauth-migration://offline?data=CpYBClBUROIRNUCnbRbNKo2KTm2skvvbBAs2wtI7OE%2FDHQncNQzzpSQAn6h0Zdre32vz127PnG%2B4Yvf8lTuzAnkYyN3csUsn65A1GE7lRCWWYK0H2RIab3BlcmF0aW9uc0BnbG9iYWxwcmltZXguZXUaC09wZW5QYXlkIFVLIAEoATACQhNjNzZjNTIxNzMwOTQ1MjM2NDc0EAIYASAA"
parse_migration_uri(MIGRATION_URI)
