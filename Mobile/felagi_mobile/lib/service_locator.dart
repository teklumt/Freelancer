import 'package:felagi_mobile/core/network/network_info.dart';
import 'package:get_it/get_it.dart';
import 'package:internet_connection_checker/internet_connection_checker.dart';
import 'package:shared_preferences/shared_preferences.dart';
import 'package:http/http.dart' as http;

final getIt = GetIt.instance;

Future<void> init()async {

     var sharedPreferences = await SharedPreferences.getInstance();
  var client = http.Client();
  var internetChecker = InternetConnectionChecker();
  getIt.registerFactory<NetworkInfo>(() => NetworkInfoImpl(internetChecker));


}