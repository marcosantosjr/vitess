# Changelog of Vitess v17.0.3

### Bug fixes 
#### Build/CI
 * [release-17.0] Make `Static Code Checks Etc` fail if the `./changelog` folder is out-of-date (#14003) [#14006](https://github.com/vitessio/vitess/pull/14006)
 * [release-17.0] Enable failures in `tools/e2e_test_race.sh` and fix races (#13654) [#14011](https://github.com/vitessio/vitess/pull/14011) 
#### CLI
 * [release-17.0] backport mysqlctl CLI compatibility fix to 17.0 [#14082](https://github.com/vitessio/vitess/pull/14082) 
#### Cluster management
 * [release-17.0] Fix `BackupShard` to get its options from its own flags (#13813) [#13820](https://github.com/vitessio/vitess/pull/13820) 
#### Evalengine
 * [release-17.0] evalengine: Mark UUID() function as non-constant (#14051) [#14057](https://github.com/vitessio/vitess/pull/14057) 
#### Online DDL
 * [release-17.0] OnlineDDL: fix nil 'completed_timestamp' for cancelled migrations (#13928) [#13937](https://github.com/vitessio/vitess/pull/13937)
 * [release-17.0] OnlineDDL: cleanup cancelled migration artifacts; support `--retain-artifacts=<duration>` DDL strategy flag (#14029) [#14037](https://github.com/vitessio/vitess/pull/14037)
 * [release-17.0] TableGC: support DROP VIEW (#14020) [#14045](https://github.com/vitessio/vitess/pull/14045)
 * [release-17.0] bugfix: change column name and type to json (#14093) [#14117](https://github.com/vitessio/vitess/pull/14117) 
#### Query Serving
 * [release-17.0] Rewrite `USING` to `ON` condition for joins (#13931) [#13941](https://github.com/vitessio/vitess/pull/13941)
 * [release-17.0] handle large number of predicates without timing out (#13979) [#13982](https://github.com/vitessio/vitess/pull/13982)
 * [release-17.0] fix data race in join engine primitive olap streaming mode execution (#14012) [#14016](https://github.com/vitessio/vitess/pull/14016)
 * [release-17.0] fix: cost to include subshard opcode (#14023) [#14027](https://github.com/vitessio/vitess/pull/14027)
 * [release-17.0] Add session flag for stream execute grpc api (#14046) [#14053](https://github.com/vitessio/vitess/pull/14053) 
#### Throttler
 * [release-17.0] Tablet throttler: empty list of probes on non-leader (#13926) [#13952](https://github.com/vitessio/vitess/pull/13952) 
#### VReplication
 * [release-17.0] Flakes: skip flaky check that ETA for a VReplication VDiff2 Progress command is in the future. (#13804) [#13817](https://github.com/vitessio/vitess/pull/13817)
 * [release-17.0] VReplication: Handle SQL NULL and JSON 'null' correctly for JSON columns (#13944) [#13947](https://github.com/vitessio/vitess/pull/13947)
 * [release-17.0] copy over existing vreplication rows copied to local counter if resuming from another tablet (#13949) [#13963](https://github.com/vitessio/vitess/pull/13963)
 * [release-17.0] VDiff: correct handling of default source and target cells (#13969) [#13984](https://github.com/vitessio/vitess/pull/13984)
 * Backport [release-17.0]  Flakes: Add recently added 'select rows_copied' query to ignore list #13993 [#14039](https://github.com/vitessio/vitess/pull/14039)
 * [release-17.0] json: Fix quoting JSON keys (#14066) [#14068](https://github.com/vitessio/vitess/pull/14068)
 * [release-17.0] VDiff: properly split cell values in record when using TabletPicker (#14099) [#14103](https://github.com/vitessio/vitess/pull/14103)
 * [release-17.0] VDiff: Cleanup the controller for a VDiff before deleting it (#14107) [#14125](https://github.com/vitessio/vitess/pull/14125)
### CI/Build 
#### Documentation
 * [release-17.0] update docgen to embed commit ID in autogenerated doc frontmatter (#14056) [#14074](https://github.com/vitessio/vitess/pull/14074) 
#### General
 * [release-17.0] Upgrade the Golang version to `go1.20.8` [#13934](https://github.com/vitessio/vitess/pull/13934) 
#### VTorc
 * [release-17.0] docker: add dedicated vtorc container (#14126) [#14147](https://github.com/vitessio/vitess/pull/14147)
### Documentation 
#### Documentation
 * [release-17.0] anonymize homedirs in generated docs (#14101) [#14106](https://github.com/vitessio/vitess/pull/14106)
### Enhancement 
#### Build/CI
 * [release-17.0] Bump upgrade tests to `go1.21.0` [#13855](https://github.com/vitessio/vitess/pull/13855)
### Internal Cleanup 
#### Build/CI
 * [release-17.0] Use Debian Bullseye in Bootstrap [#13757](https://github.com/vitessio/vitess/pull/13757) 
#### Query Serving
 * [release-17.0] moved timeout test to different package (#14028) [#14032](https://github.com/vitessio/vitess/pull/14032)
### Release 
#### General
 * Code freeze of release-17.0 [#14138](https://github.com/vitessio/vitess/pull/14138)
### Testing 
#### Build/CI
 * [release-17.0] Remove FOSSA Test from CI until we can do it in a secure way (#14119) [#14122](https://github.com/vitessio/vitess/pull/14122)
 * [release-17.0] Fix upgrade tests [#14143](https://github.com/vitessio/vitess/pull/14143) 
#### VReplication
 * [release-17.0] Flakes: empty vtdataroot before starting a new vreplication e2e test (#13803) [#13822](https://github.com/vitessio/vitess/pull/13822)

