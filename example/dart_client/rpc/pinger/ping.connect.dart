//
// Code generated by protoc-gen-connect-dart. DO NOT EDIT.
// Generated from: pinger/ping.proto
//
// @dart = 2.17
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name
import 'dart:async';
import 'package:ska_connect/ska_connect.dart';
import 'dart:convert';
import './ping.pb.dart';

abstract class Pinger {
	Future<ConnectResponse<PingResponse>>ping(PingRequest pingRequest);
}

class PingerClient implements Pinger {
	final ConnectClient connectClient;

  PingerClient(this.connectClient) {
	}

	@override
	Future<ConnectResponse<PingResponse>> ping(PingRequest pingRequest) async {
    return connectClient.performRequest(
      '/pinger.Pinger/Ping', pingRequest, (b) => PingResponse.fromBuffer(b));
	}
}