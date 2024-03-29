## Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [Released]
## [0.3.0] - 2022-04-26
### Added
- Method `RevokePasswordResetToken` to handle the exchange with the 
security microservice to revoke a password reset token to prevent the reset
of a user's password if the user did not authorise the reset.

### Changed
- Moved the `models.go` to `type.go` as it better represents what it is 
in Go.

## [0.2.0] - 2022-04-16
### Added
- `ResetPassword` HTTP exchange function to POST the new user password to
the security microservice.

## [0.1.0] - 2022-04-15
### Added
- `PasswordResetToken` HTTP exchange function to get the password reset
token from the security microservice.
- `payload` file which contains all the payload structure to the security
microservice.

### Changed
- updated the `dutil` package to `0.1.0`.

## [0.0.0] - 2022-04-15
### Added
- Authentication functions such as `Login` and `Logout`.
- Service module to be become a mock and testable service integration.
- Models to be made available to the API Gateways.
- 