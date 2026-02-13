# Changelog

## [2.3.1](https://github.com/Filip7/workcd-go/compare/v2.3.0...v2.3.1) (2026-02-13)


### Bug Fixes

* support preview of non md README files ([#17](https://github.com/Filip7/workcd-go/issues/17)) ([d0f8ab1](https://github.com/Filip7/workcd-go/commit/d0f8ab1e25dbe99469b48e85c29f0bf28e7b2a89))

## [2.3.0](https://github.com/Filip7/workcd-go/compare/v2.2.0...v2.3.0) (2025-10-11)


### Features

* support custom name of the function (default is still `wd`) ([13013f3](https://github.com/Filip7/workcd-go/commit/13013f391de958c2492ede3ee84b434f1bee7ba2))

## [2.2.0](https://github.com/Filip7/workcd-go/compare/v2.1.1...v2.2.0) (2025-10-08)


### Features

* shell integration baked in - no more copy pasting whole functions to .zshrc ([#13](https://github.com/Filip7/workcd-go/issues/13)) ([5c7eccf](https://github.com/Filip7/workcd-go/commit/5c7eccf13e38f4612ca403607e0fb6e72c7115a9))

## [2.1.1](https://github.com/Filip7/workcd-go/compare/v2.1.0...v2.1.1) (2025-10-08)


### Bug Fixes

* print config that is currently used - takes into the accout both flags and config ([86c36c9](https://github.com/Filip7/workcd-go/commit/86c36c95a081933b70f725b90ce58daf90bbf42e))

## [2.1.0](https://github.com/Filip7/workcd-go/compare/v2.0.0...v2.1.0) (2025-10-07)


### Features

* print config using --print-config ([#9](https://github.com/Filip7/workcd-go/issues/9)) ([b391d6c](https://github.com/Filip7/workcd-go/commit/b391d6c6566e89f547299423e9cf68a9d56bfa6e))


### Bug Fixes

* print config as yaml ([626c850](https://github.com/Filip7/workcd-go/commit/626c850839ccff529bbfe5f63f6acb5380253234))

## [2.0.0](https://github.com/Filip7/workcd-go/compare/v1.3.0...v2.0.0) (2025-09-15)


### ⚠ BREAKING CHANGES

* move config and flag functions to their files

### Features

* switch to go 1.25 and use experimental garbage collector ([22417fb](https://github.com/Filip7/workcd-go/commit/22417fb2952e98e07df588f42c0f90e94791928a))


### Code Refactoring

* move config and flag functions to their files ([e6b5372](https://github.com/Filip7/workcd-go/commit/e6b53728d7f88ff6891043b0178a863d8c0a9397))

## [1.3.0](https://github.com/Filip7/workcd-go/compare/v1.2.1...v1.3.0) (2025-09-05)


### Features

* show preview of README files in the direcotries wile using fzf ([9741eae](https://github.com/Filip7/workcd-go/commit/9741eaec7f8c9d022fe6cdc6a74b3f1abbf5599d))

## [1.2.1](https://github.com/Filip7/workcd-go/compare/v1.2.0...v1.2.1) (2025-09-03)


### Bug Fixes

* when no directory name is passed to the command, just cd to the basedir ([dc9bb57](https://github.com/Filip7/workcd-go/commit/dc9bb578ee46f64bab72cab23151f1fc3443e969))

## [1.2.0](https://github.com/Filip7/workcd-go/compare/v1.1.0...v1.2.0) (2025-09-03)


### Features

* add flag to set work dir ([ee1c533](https://github.com/Filip7/workcd-go/commit/ee1c5338d60074dd07ae17d4b6608d639a0853a8))
* update shell wrapper to support error handling the correct way ([8c00e71](https://github.com/Filip7/workcd-go/commit/8c00e718b09ba3b688c8e5939b25dc69471aff5a))

## [1.1.0](https://github.com/Filip7/workcd-go/compare/v1.0.0...v1.1.0) (2025-09-02)


### Features

* better error handling ([9702b4f](https://github.com/Filip7/workcd-go/commit/9702b4f9da9237015e60176d6a8d80f9e7fa2305))
* handle opening and querying subdirectories ([dfb15f0](https://github.com/Filip7/workcd-go/commit/dfb15f039a110455dafe17967b35e226024b710d))

## 1.0.0 (2025-09-02)


### Features

* initial version - full functionallity of zsh script ([09ed1b5](https://github.com/Filip7/workcd-go/commit/09ed1b590c513fbe656f6213889a613eb54db814))
* open with editor passed in as --editor, editor must support ([09ed1b5](https://github.com/Filip7/workcd-go/commit/09ed1b590c513fbe656f6213889a613eb54db814))
* support config for setting base workspace ([09ed1b5](https://github.com/Filip7/workcd-go/commit/09ed1b590c513fbe656f6213889a613eb54db814))
