# Project Goals

Existing open source email servers have a heavy operational footprint due to configuration files, file permissions, and external backing services such as attached remote storage and directory services, which creates a combinatorial explosion of things that can go wrong, while making the possibility of internally consistent backup/restore and live snapshotting on an ordinary operating system difficult to impossible.

Which sets the stage for the following project goals:

## 1. On-premise, cloud-managed

Use a cloud administration portal for configuration, so that on-premise email servers are completely operator-less and configuration-free. The cloud provides a place to park snapshots and backups, enabling use cases such as launching a new site from a snapshot. The cloud also provides a place to perform integrations with Managed DNS providers.

Use a single self-contained Go binary for minimal footprint. This means no use of Docker, no lift-and-shift of existing software that doesn't conform to the project goals or [crash-only design](https://en.wikipedia.org/wiki/Crash-only_software), and no use of external dependencies like Perl or Python. Think in terms of being an in-place migration option for a 20 year old cPanel-managed Linux mail host.

As with all other email servers, it is expected to be fully operational while it has lost its uplink to the cloud. Although with this design, it will be locked into its current configuration until its cloud uplink is restored.

## 2. Operator-less storage and live snapshots

Store all configuration and mailboxes within a single transactional database that is able to be live-snapshotted. Use embedded database engines that are statically linked into the single binary footprint:

* [BoltDB](https://github.com/etcd-io/bbolt) for read-optimized sites where bullet-proof minimalism is the highest goal
* [BadgerDB](https://github.com/dgraph-io/badger) for write-optimized sites

For cloud and on-prem neutrality, use a [cloud-native](https://12factor.net) architecture to ensure all code is stateless, and able to use either embedded services or managed backing services such as SQL and NoSQL for future cloud hosting potential.

## 3. Effortless migration

A site administrator should be able to do any of the following with single-click ease:

* start an email site on an on-prem server
* transfer the site to a cloud of your choice
* transfer the site back to an on-prem server
