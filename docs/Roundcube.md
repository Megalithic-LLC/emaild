# How to Best Support Roundcube

Here is my current thinking on how to best support [Roundcube Webmail](https://github.com/roundcube/roundcubemail) from a self-contained Go binary whose configuration is sealed by release versioning.

## 1. Embed a PHP+HTTP server for Go

Embed a PHP server for Go such as [RoadRunner](https://github.com/spiral/roadrunner).

## 2. Expose the HTTP endpoint

Expose the RoadRunner (or equivalent) HTTP server via a public endpoint, so that end users can access the email server with a browser. Remember that the email server does not automatically have an http endpoint, since it is only ever configured indirectly through the cloud administration portal where configuration release versioning can be applied.

## 3. Allow a site operator to deploy it

Allow a site operator to use the cloud administration portal to perform a download of a given version of RoadRunner. The email server would then see a new versioned release of the site configuration, and download it.