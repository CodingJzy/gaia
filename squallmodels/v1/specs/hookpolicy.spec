{
    "attributes": [
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": null,
            "channel": null,
            "creation_only": null,
            "default_order": null,
            "default_value": null,
            "deprecated": null,
            "description": "CertificateAuthority contains the pem block of the certificate authority used by the remote endpoint.",
            "exposed": true,
            "filterable": false,
            "foreign_key": null,
            "format": "free",
            "getter": null,
            "identifier": null,
            "index": null,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "certificateAuthority",
            "orderable": true,
            "primary_key": null,
            "read_only": null,
            "required": true,
            "secret": null,
            "setter": null,
            "stored": true,
            "subtype": null,
            "transient": null,
            "type": "string",
            "unique": null,
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
            "description": "ClientCertificate contains the client certificate that will be used to connect to the remote endoint.",
            "exposed": true,
            "filterable": false,
            "foreign_key": null,
            "format": "free",
            "getter": null,
            "identifier": null,
            "index": null,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "clientCertificate",
            "orderable": true,
            "primary_key": null,
            "read_only": null,
            "required": true,
            "secret": null,
            "setter": null,
            "stored": true,
            "subtype": null,
            "transient": null,
            "type": "string",
            "unique": null,
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
            "description": "ClientCertificateKey contains the key associated to the clientCertificate.",
            "exposed": true,
            "filterable": false,
            "foreign_key": null,
            "format": "free",
            "getter": null,
            "identifier": null,
            "index": null,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "clientCertificateKey",
            "orderable": true,
            "primary_key": null,
            "read_only": null,
            "required": true,
            "secret": true,
            "setter": null,
            "stored": true,
            "subtype": null,
            "transient": null,
            "type": "string",
            "unique": null,
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
            "description": "Endpoint contains the full address of the remote processor endoint.",
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
            "name": "endpoint",
            "orderable": true,
            "primary_key": null,
            "read_only": false,
            "required": true,
            "secret": null,
            "setter": null,
            "stored": true,
            "subtype": null,
            "transient": null,
            "type": "string",
            "unique": null,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": [
                "Both",
                "Post",
                "Pre"
            ],
            "autogenerated": null,
            "channel": null,
            "creation_only": null,
            "default_order": null,
            "default_value": "Post",
            "deprecated": null,
            "description": "Mode define the type of the hook.",
            "exposed": true,
            "filterable": true,
            "foreign_key": null,
            "format": null,
            "getter": null,
            "identifier": null,
            "index": null,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "mode",
            "orderable": true,
            "primary_key": null,
            "read_only": null,
            "required": false,
            "secret": null,
            "setter": null,
            "stored": true,
            "subtype": null,
            "transient": null,
            "type": "enum",
            "unique": null,
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
            "description": "Subject contains the tag expression that an object must match in order to trigger the hook.",
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
            "name": "subject",
            "orderable": true,
            "primary_key": null,
            "read_only": null,
            "required": null,
            "secret": null,
            "setter": null,
            "stored": true,
            "subtype": null,
            "transient": null,
            "type": "string",
            "unique": null,
            "uniqueScope": null
        }
    ],
    "children": [],
    "model": {
        "aliases": [
            "hookpols",
            "hookpol",
            "hooks",
            "hook"
        ],
        "create": null,
        "delete": true,
        "description": "Hook allows to to define hooks to the write operations in squall. Hooks are sent to an external Rufus server that will do the processing and eventually return a modified version of the object before we save it.",
        "entity_name": "HookPolicy",
        "extends": [
            "@base",
            "@described",
            "@disabled",
            "@identifiable-nopk-nostored",
            "@metadatable",
            "@named",
            "@propagated"
        ],
        "get": true,
        "package": "Policies",
        "resource_name": "hookpolicies",
        "rest_name": "hookpolicy",
        "root": null,
        "update": true
    }
}