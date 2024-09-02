import 'package:felagi_mobile/service_locator.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

 void main() async {

  WidgetsFlutterBinding.ensureInitialized();
  await init();

  runApp(const MyApp());
}
class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MultiBlocProvider(
      providers: [
        // BlocProvider(
        //   create: (context) => SubjectBloc(),
        // ),
       
      ],
      child: MaterialApp(
        initialRoute: '/',
        routes: {
          // Add routes here

        },
      ),
    );
  }
}