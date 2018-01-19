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
            "description": "Results contains the result of the query.",
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
            "name": "results",
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
            "sq"
        ],
        "create": true,
        "delete": false,
        "description": "TimeSeriesData is a generic API to retrieve time series data stored by the Aporeto system. The API allows different types of queries that are all protected within the namespace of the user.",
        "entity_name": "StatsQuery",
        "extends": [],
        "get": false,
        "package": "Squall API",
        "resource_name": "statsqueries",
        "rest_name": "statsquery",
        "root": null,
        "update": false
    }
}