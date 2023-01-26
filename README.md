mackerel-plugin-mysql
=====================

MySQL custom metrics plugin for mackerel.io agent.

## Synopsis

```shell
mackerel-plugin-mysql [-host=<host>] [-port=<port>] [-username=<username>] [-password=<password>] [-tempfile=<tempfile>] [-disable_innodb=true] [-metric-key-prefix=<prefix>] [-enable_extended=true] [-debug=true]
```

## Example of mackerel-agent.conf

```
[plugin.metrics.mysql]
command = "/path/to/mackerel-plugin-mysql"
```

## Supported MySQL version

- `v1.1.0 >=` mysql 5.7, 8.0 and above
- `v1.0.0` mysql 5.0, 5.1, 5.5, 5.6, 5.7, 8.0
