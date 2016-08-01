{
    "attributes": [
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": true,
            "channel": null,
            "creation_only": false,
            "default_order": false,
            "default_value": null,
            "deprecated": false,
            "description": "List of namespaces that an user is authorized to use",
            "exposed": true,
            "filterable": true,
            "foreign_key": false,
            "format": null,
            "getter": false,
            "identifier": false,
            "index": false,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "namespaces",
            "orderable": true,
            "primary_key": false,
            "read_only": true,
            "required": false,
            "setter": false,
            "stored": false,
            "subtype": "namespaces_list",
            "transient": false,
            "type": "external",
            "unique": false,
            "uniqueScope": null
        }
    ],
    "children": [],
    "model": {
        "create": false,
        "delete": false,
        "description": "mynamespaces returns the user's authorized namespaces. ",
        "entity_name": "MyNamespace",
        "extends": [
            "@identifiable-nopk-nostored"
        ],
        "get": false,
        "package": "Infrastructure",
        "resource_name": "mynamespaces",
        "rest_name": "mynamespace",
        "root": false,
        "update": false
    }
}
