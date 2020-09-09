package main

var pluginJson2 = `{}`
var pluginJson = `
{
	"built_in_apis": [{
	  "name": "",
	  "type": "addition",
	  "methods": [
		"POST"
	  ],
	  "url": "/team/{teamUUID}/plugin/app/upload"
	}, {
	  "name": "",
	  "type": "addition",
	  "methods": [
		"POST"
	  ],
	  "url": "/team/{teamUUID}/plugin/app/delete"
	}],
	"server_update_stamp": 1234567890,
	"plugin_list": [{
	  "service": {
		"name": "plugin_test",
		"id": "a12f1e2c",
		"version": "1.0.0",
		"team_uuid": "BDfDqJU7",
		"orgz_uuid": "75tuQaF7",
		"description": "插件范例"
	  },
	  "package": {
		"file": "plugin_tcp_base.zip",
		"entrance": "loadPlugin_tcp_base"
	  },
	  "apis": [{
		  "name": "addUpdate_webhook",
		  "type": "addition",
		  "methods": [
			"POST"
		  ],
		  "url": "/team/{teamUUID}/webhook/update"
		},
		{
		  "name": "list_webhook_issues",
		  "type": "addition",
		  "methods": [
			"GET"
		  ],
		  "url": "/team/{teamUUID}/webhook/list/issue"
		},
		{
		  "name": "list_webhooks",
		  "type": "addition",
		  "methods": [
			"GET"
		  ],
		  "url": "/team/{teamUUID}/webhook/list"
		},
		{
		  "name": "list_simple_webhook",
		  "type": "addition",
		  "methods": [
			"GET"
		  ],
		  "url": "/team/{teamUUID}/webhook/list/uuid/{uuid}"
		},
		{
		  "name": "delete_webhook",
		  "type": "addition",
		  "methods": [
			"GET"
		  ],
		  "url": "/team/{teamUUID}/webhook/delete/uuid/{uuid}"
		},
		{
		  "name": "test_url",
		  "type": "addition",
		  "methods": [
			"GET"
		  ],
		  "url": "/team/{teamUUID}/webhook/test"
		},
		{
		  "name": "verify_webhook_name",
		  "type": "addition",
		  "methods": [
			"GET"
		  ],
		  "url": "/team/{teamUUID}/webhook/verify/name"
		}
	  ]
	}, {
	  "service": {
		"name": "plugin_test",
		"id": "d57d0608",
		"version": "1.0.0",
		"team_uuid": "BDfDqJU7",
		"orgz_uuid": "75tuQaF7",
		"description": "插件范例"
	  },
	  "package": {
		"file": "plugin_tcp_base.zip",
		"entrance": "loadPlugin_tcp_base"
	  },
	  "apis": [{
		  "name": "addUpdate_webhook",
		  "type": "addition",
		  "methods": [
			"POST"
		  ],
		  "url": "/team/{teamUUID}/webhook/update"
		},
		{
		  "name": "list_webhook_issues",
		  "type": "addition",
		  "methods": [
			"GET"
		  ],
		  "url": "/team/{teamUUID}/webhook/list/issue"
		},
		{
		  "name": "list_webhooks",
		  "type": "addition",
		  "methods": [
			"GET"
		  ],
		  "url": "/team/{teamUUID}/webhook/list"
		},
		{
		  "name": "list_simple_webhook",
		  "type": "addition",
		  "methods": [
			"GET"
		  ],
		  "url": "/team/{teamUUID}/webhook/list/uuid/{uuid}"
		},
		{
		  "name": "delete_webhook",
		  "type": "addition",
		  "methods": [
			"GET"
		  ],
		  "url": "/team/{teamUUID}/webhook/delete/uuid/{uuid}"
		},
		{
		  "name": "test_url",
		  "type": "addition",
		  "methods": [
			"GET"
		  ],
		  "url": "/team/{teamUUID}/webhook/test"
		},
		{
		  "name": "verify_webhook_name",
		  "type": "addition",
		  "methods": [
			"GET"
		  ],
		  "url": "/team/{teamUUID}/webhook/verify/name"
		}
	  ]
	}]
  }`
