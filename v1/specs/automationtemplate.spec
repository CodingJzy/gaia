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
            "description": "Kind represents the kind of template.",
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
            "name": "kind",
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
            "autotmpl"
        ],
        "create": null,
        "delete": false,
        "description": "Templates that ca be used in automations",
        "entity_name": "AutomationTemplate",
        "exposed": true,
        "extends": [
            "@described",
            "@named"
        ],
        "get": true,
        "package": "sephiroth",
        "private": null,
        "resource_name": "automationtemplates",
        "rest_name": "automationtemplate",
        "root": null,
        "update": false
    }
}