# How to Best Support SpamAssassin

It's hard to imagine building the next great open source email server without supporting SpamAssassin as a first-class citizen.

Here is my current thinking on how to best support it from a self-contained Go binary whose configuration is sealed by release versioning.

## 1. Build a Perl VM for Go

Take inspiration from the amazing [GopherLua](https://github.com/yuin/gopher-lua) project, by launching a new "GopherPerl" project on GitHub. The Perl community sorely needs some VM alternatives anyways.

## 2. Embed a MySQL server

Embed a MySQL server such as [go-mysql-server](https://github.com/src-d/go-mysql-server), for SpamAssassin's use. The embedded MySQL server storage adapter would delegate to the sitewide BoltDB or BadgerDB database, so that when a site is snapshotted, the snapshot includes all mailboxes AND all spam-related training data.

## 3. Allow a site operator to deploy it

Allow a site operator to use the cloud configuration portal to initiatite a download of a given version of SpamAssassin. The email server would then see a new versioned release of the site configuration, and download it.

The email server would need to be hard-coded with support for specific versions of SpamAssassin, because supporting a new version would require implementing automatic operator-less migration of SpamAssassin-related stored data. This might even be done in the Drupal manner, whereby an operator could DOWNGRADE to an older version of SpamAssassin as well as UPGRADE to a newer version.

This requirement seems like an important one, if a primary use case for this email server should be the in-place seamless upgrade of a 20-year old Linux email server (well, SpamAssassin is only 18 years old as of this writing, but, close enough). If that old email server is running an old version of SpamAssassin, it would be most practical to have a migration tool perform its migration by creating a snapshot of the existing Dovecot/Sendmail/Postfix/Zimbra site, along with its current configured version of SpamAssassin. Once a new email server site is launched from that snapshot, an operator could then be given options to upgrade SpamAssassin under the control of the usual managed service automation.
