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
