{
    "attributes": [
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": false,
            "channel": null,
            "creation_only": false,
            "default_order": false,
            "default_value": null,
            "deprecated": false,
            "description": "e-mail address of the user",
            "exposed": true,
            "filterable": true,
            "foreign_key": false,
            "format": "email",
            "getter": null,
            "identifier": false,
            "index": false,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "email",
            "orderable": true,
            "primary_key": true,
            "read_only": false,
            "required": false,
            "setter": null,
            "stored": true,
            "subtype": null,
            "transient": false,
            "type": "string",
            "unique": false,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": false,
            "channel": null,
            "creation_only": true,
            "default_order": false,
            "default_value": null,
            "deprecated": false,
            "description": "ParentAuthenticator is an Internal attribute that points to the parent authenticator.",
            "exposed": false,
            "filterable": false,
            "foreign_key": false,
            "format": "free",
            "getter": null,
            "identifier": false,
            "index": false,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "parentAuthenticator",
            "orderable": false,
            "primary_key": true,
            "read_only": false,
            "required": false,
            "setter": null,
            "stored": true,
            "subtype": null,
            "transient": false,
            "type": "string",
            "unique": false,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": null,
            "channel": null,
            "creation_only": null,
            "default_order": null,
            "default_value": null,
            "deprecated": null,
            "description": "ParentID is an internal reference to the parent object that is an authenticator. ",
            "exposed": true,
            "filterable": true,
            "foreign_key": null,
            "format": "free",
            "getter": null,
            "identifier": null,
            "index": null,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "parentID",
            "orderable": true,
            "primary_key": null,
            "read_only": false,
            "required": null,
            "setter": null,
            "stored": true,
            "subtype": null,
            "transient": false,
            "type": "string",
            "unique": null,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": false,
            "channel": null,
            "creation_only": false,
            "default_order": false,
            "default_value": null,
            "deprecated": false,
            "description": "OU attribute for the generated certificates",
            "exposed": true,
            "filterable": true,
            "foreign_key": false,
            "format": null,
            "getter": null,
            "identifier": false,
            "index": false,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "subOrganizations",
            "orderable": false,
            "primary_key": false,
            "read_only": false,
            "required": false,
            "setter": null,
            "stored": true,
            "subtype": "string",
            "transient": false,
            "type": "list",
            "unique": false,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": false,
            "channel": null,
            "creation_only": true,
            "default_order": false,
            "default_value": null,
            "deprecated": false,
            "description": "CommonName (CN) for the user certificate",
            "exposed": true,
            "filterable": true,
            "foreign_key": false,
            "format": "free",
            "getter": false,
            "identifier": false,
            "index": false,
            "max_length": 64,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "userName",
            "orderable": true,
            "primary_key": false,
            "read_only": false,
            "required": false,
            "setter": false,
            "stored": true,
            "subtype": null,
            "transient": false,
            "type": "string",
            "unique": false,
            "uniqueScope": null
        }
    ],
    "children": [
        {
            "bulk_create": false,
            "bulk_delete": false,
            "bulk_update": false,
            "create": true,
            "delete": false,
            "deprecated": null,
            "get": true,
            "relationship": "child",
            "rest_name": "certificate",
            "update": false
        }
    ],
    "model": {
        "create": null,
        "delete": true,
        "description": null,
        "entity_name": "User",
        "extends": [
            "@base"
        ],
        "get": true,
        "package": "User",
        "resource_name": "users",
        "rest_name": "user",
        "root": null,
        "update": true
    }
}