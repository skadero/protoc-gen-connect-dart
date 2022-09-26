[![license](https://img.shields.io/badge/license-MIT-purple.svg)](https://opensource.org/licenses/MIT)

A protoc plugin for generating a [Connect](https://connect.build) client for Dart. The generated client only supports unary requests and is suitable for any supported Dart platform.

This plugin is based off of [apptreesoftware/protoc-gen-twirp_dart](https://github.com/apptreesoftware/protoc-gen-twirp_dart) and modifications from [marwan-at-work/protoc-gen-twirp_dart](https://github.com/marwan-at-work/protoc-gen-twirp_dart). The code was modified to remove all of the JSON type generation, and the actual RPC calls are made using the [ska_connect](https://pub.dev/packages/ska_connect) supporting package.

## Installation

The protobuf v3 compiler is required. You can get the latest precompiled binary for your system [here](https://github.com/google/protobuf/releases). In addition Go 1.18 or higher is required. The executable can be installed with the following command:

    go install github.com/skadero/protoc-gen-connect-dart

### Example Go Server (optional)

While not required, the example directory contains a Go based server implementation that requires:

    go install github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go@latest
    go install github.com/golang/protobuf/protoc-gen-go

## Using

The generated code requires the following Dart packages in your `pubspec.yaml`:

    ska_connect: ^1.0.1

An example command for compiling protoc files:

    protoc --proto_path=./example/proto  --dart_out=./example/dart_client/rpc --connect-dart_out=./example/dart_client/rpc ./example/proto/haberdasher/haberdasher.proto ./example/proto/pinger/ping.proto

Which outputs the generated code to `./example/dart_client/rpc`. Example of Connect calls from generated code:

```dart
import 'dart:async';

import "package:ska_connect/ska_connect.dart";
import 'package:http/http.dart' as http;
import 'rpc/haberdasher/haberdasher.connect.dart' as ha;
import 'rpc/haberdasher/haberdasher.pb.dart';
import 'rpc/pinger/ping.pb.dart';

Future main(List<String> args) async {
  final hClient = http.Client();
  final client = HttpConnectClient('http://localhost:9080', hClient);

  final hService = ha.HaberdasherClient(client);

  final hatResponse = await hService.makeHat(Size(inches: 6));
  print(hatResponse.response);

  final invResponse = await hService.invalidArg(PingRequest(name: "example"));
  if (invResponse.isError) {
    print(invResponse.error!.code);
  }
  hClient.close();
}
```
