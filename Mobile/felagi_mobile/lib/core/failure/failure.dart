 import 'package:equatable/equatable.dart';

class Failure extends Equatable{
  final String message;
  const Failure(this.message);

  List<Object> get props => [message];
}

class ServerFailure extends Failure{
  const ServerFailure(String message) : super(message);
}

class ConnectionFailure extends Failure{
  const ConnectionFailure(String message) : super(message);
}

class DatabaseFailure extends Failure{
  const DatabaseFailure(String message) : super(message);
}


class CacheFailure extends Failure {
  const CacheFailure(super.message);
}

class NetworkFailure extends Failure {
  const NetworkFailure([super.message = 'No internet connection']);
}


